package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnssoarec() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnssoarecFunc,
		Read:          readDnssoarecFunc,
		Update:        updateDnssoarecFunc,
		Delete:        deleteDnssoarecFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ecssubnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"minimum": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nodeid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"originserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"retry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createDnssoarecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Get("domain").(string)
	dnssoarec := dns.Dnssoarec{
		Contact:      d.Get("contact").(string),
		Domain:       d.Get("domain").(string),
		Ecssubnet:    d.Get("ecssubnet").(string),
		Expire:       uint64(d.Get("expire").(int)),
		Minimum:      uint64(d.Get("minimum").(int)),
		Nodeid:       uint32(d.Get("nodeid").(int)),
		Originserver: d.Get("originserver").(string),
		Refresh:      uint64(d.Get("refresh").(int)),
		Retry:        uint64(d.Get("retry").(int)),
		Serial:       uint32(d.Get("serial").(int)),
		Ttl:          uint64(d.Get("ttl").(int)),
		Type:         d.Get("type").(string),
	}

	_, err := client.AddResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
	if err != nil {
		return err
	}

	d.SetId(dnssoarecId)

	err = readDnssoarecFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnssoarec but we can't read it ?? %s", dnssoarecId)
		return nil
	}
	return nil
}

func readDnssoarecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnssoarec state %s", dnssoarecName)
	data, err := client.FindResource(service.Dnssoarec.Type(), dnssoarecName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnssoarec state %s", dnssoarecName)
		d.SetId("")
		return nil
	}
	d.Set("contact", data["contact"])
	d.Set("domain", data["domain"])
	d.Set("ecssubnet", data["ecssubnet"])
	d.Set("expire", data["expire"])
	d.Set("minimum", data["minimum"])
	d.Set("nodeid", data["nodeid"])
	d.Set("originserver", data["originserver"])
	d.Set("refresh", data["refresh"])
	d.Set("retry", data["retry"])
	d.Set("serial", data["serial"])
	d.Set("ttl", data["ttl"])
	d.Set("type", data["type"])

	return nil

}

func updateDnssoarecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Get("domain").(string)

	dnssoarec := dns.Dnssoarec{
		Domain: dnssoarecId,
	}
	hasChange := false
	if d.HasChange("contact") {
		log.Printf("[DEBUG]  citrixadc-provider: Contact has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Contact = d.Get("contact").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("ecssubnet") {
		log.Printf("[DEBUG]  citrixadc-provider: Ecssubnet has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Ecssubnet = d.Get("ecssubnet").(string)
		hasChange = true
	}
	if d.HasChange("expire") {
		log.Printf("[DEBUG]  citrixadc-provider: Expire has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Expire = uint64(d.Get("expire").(int))
		hasChange = true
	}
	if d.HasChange("minimum") {
		log.Printf("[DEBUG]  citrixadc-provider: Minimum has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Minimum = uint64(d.Get("minimum").(int))
		hasChange = true
	}
	if d.HasChange("nodeid") {
		log.Printf("[DEBUG]  citrixadc-provider: Nodeid has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Nodeid = uint32(d.Get("nodeid").(int))
		hasChange = true
	}
	if d.HasChange("originserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Originserver has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Originserver = d.Get("originserver").(string)
		hasChange = true
	}
	if d.HasChange("refresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Refresh has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Refresh = uint64(d.Get("refresh").(int))
		hasChange = true
	}
	if d.HasChange("retry") {
		log.Printf("[DEBUG]  citrixadc-provider: Retry has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Retry = uint64(d.Get("retry").(int))
		hasChange = true
	}
	if d.HasChange("serial") {
		log.Printf("[DEBUG]  citrixadc-provider: Serial has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Serial = uint32(d.Get("serial").(int))
		hasChange = true
	}
	if d.HasChange("ttl") {
		log.Printf("[DEBUG]  citrixadc-provider: Ttl has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Ttl = uint64(d.Get("ttl").(int))
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for dnssoarec %s, starting update", dnssoarecId)
		dnssoarec.Type = d.Get("type").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnssoarec.Type(), dnssoarecId, &dnssoarec)
		if err != nil {
			return fmt.Errorf("Error updating dnssoarec %s. %s", dnssoarecId, err)
		}
	}
	return readDnssoarecFunc(d, meta)
}

func deleteDnssoarecFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnssoarecFunc")
	client := meta.(*NetScalerNitroClient).client
	dnssoarecId := d.Id()
	err := client.DeleteResource(service.Dnssoarec.Type(), dnssoarecId)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
