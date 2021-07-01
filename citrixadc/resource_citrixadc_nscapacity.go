package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
	"time"
)

func resourceCitrixAdcNscapacity() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNscapacityFunc,
		Read:          readNscapacityFunc,
		Update:        schema.Noop,
		Delete:        deleteNscapacityFunc,
		Schema: map[string]*schema.Schema{
			"bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"edition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vcpu": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"reboot_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10m",
			},
			"poll_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			"poll_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "60s",
			},
			"poll_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10s",
			},
		},
	}
}

func createNscapacityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	nscapacityId := resource.PrefixedUniqueId("tf-nscapacity-")
	nscapacity := ns.Nscapacity{
		Bandwidth: uint32(d.Get("bandwidth").(int)),
		Edition:   d.Get("edition").(string),
		Nodeid:    uint32(d.Get("nodeid").(int)),
		Platform:  d.Get("platform").(string),
		Unit:      d.Get("unit").(string),
		Vcpu:      d.Get("vcpu").(bool),
	}

	err := client.UpdateUnnamedResource("nscapacity", &nscapacity)
	if err != nil {
		return err
	}

	d.SetId(nscapacityId)

	// Read also powercycles the ADC
	err = readNscapacityFunc(d, meta)

	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this nscapacity but we can't read it ?? %s", nscapacityId)
		return err
	}
	return nil
}

func deleteNscapacityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	type nscapacityRemove struct {
		Bandwidth bool `json:"bandwidth,omitempty"`
		Platform  bool `json:"platform,omitempty"`
		Vcpu      bool `json:"vcpu,omitempty"`
	}
	nscapacity := nscapacityRemove{}
	log.Printf("nscapacitydelete struct %v", nscapacity)

	if _, ok := d.GetOk("bandwidth"); ok {
		nscapacity.Bandwidth = true
	}

	if _, ok := d.GetOk("platform"); ok {
		nscapacity.Platform = true
	}

	if _, ok := d.GetOk("vcpu"); ok {
		nscapacity.Vcpu = true
	}

	err := client.ActOnResource("nscapacity", &nscapacity, "unset")
	if err != nil {
		return err
	}

	var timeout time.Duration
	if timeout, err = time.ParseDuration(d.Get("reboot_timeout").(string)); err != nil {
		return err
	}

	err = powerCycleAndWait(d, meta, timeout)
	if err != nil {
		return fmt.Errorf("Error power cycling ADC. %s", err.Error())
	}

	d.SetId("")

	return nil
}

func readNscapacityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	var err error

	var timeout time.Duration
	if timeout, err = time.ParseDuration(d.Get("reboot_timeout").(string)); err != nil {
		return err
	}

	err = powerCycleAndWait(d, meta, timeout)
	if err != nil {
		return fmt.Errorf("Error power cycling ADC. %s", err.Error())
	}

	log.Printf("[DEBUG] citrixadc-provider: Reading nscapacity state")
	data, err := client.FindResource("nscapacity", "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing nscapacity state")
		d.SetId("")
		return nil
	}

	// CICO
	if value, ok := data["platform"]; ok {
		d.Set("platform", value)
	} else {
		d.Set("platform", "")
	}

	// VCPU
	if _, ok := data["vcpucount"]; ok {
		d.Set("vcpu", true)
	} else {
		d.Set("vcpu", false)
	}

	// Pooled
	if value, ok := data["bandwidth"]; ok {
		d.Set("bandwidth", value)
		d.Set("edition", data["edition"])
		d.Set("unit", data["unit"])
	} else {
		d.Set("bandwidth", 0)
		d.Set("edition", "")
		d.Set("unit", "")
	}

	return nil

}
