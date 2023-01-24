package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcNscapacity() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createNscapacityFunc,
		Read:          readNscapacityFunc,
		Update:        createNscapacityFunc,
		Delete:        deleteNscapacityFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"edition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"platform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcpu": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createNscapacityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	nscapacityId := resource.PrefixedUniqueId("tf-nscapacity-")
	nscapacity := ns.Nscapacity{
		Bandwidth: d.Get("bandwidth").(int),
		Edition:   d.Get("edition").(string),
		Nodeid:    d.Get("nodeid").(int),
		Platform:  d.Get("platform").(string),
		Unit:      d.Get("unit").(string),
		Vcpu:      d.Get("vcpu").(bool),
	}

	err := client.UpdateUnnamedResource("nscapacity", &nscapacity)
	if err != nil {
		return err
	}

	d.SetId(nscapacityId)

	warm := true
	if err = rebootNetScaler(d, meta, warm); err != nil {
		return fmt.Errorf("Error warm rebooting ADC. %s", err.Error())
	}
	
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

	d.SetId("")

	return nil
}

func readNscapacityFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readNscapacityFunc")
	client := meta.(*NetScalerNitroClient).client

	var err error

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
		setToInt("bandwidth", d, value)
		d.Set("edition", data["edition"])
		d.Set("unit", data["unit"])
	} else {
		d.Set("bandwidth", 0)
		d.Set("edition", "")
		d.Set("unit", "")
	}

	return nil

}
