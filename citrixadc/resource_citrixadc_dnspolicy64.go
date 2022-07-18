package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/dns"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createDnspolicy64Func,
		Read:          readDnspolicy64Func,
		Update:        updateDnspolicy64Func,
		Delete:        deleteDnspolicy64Func,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createDnspolicy64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Get("name").(string)
	dnspolicy64 := dns.Dnspolicy64{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
	if err != nil {
		return err
	}

	d.SetId(dnspolicy64Name)

	err = readDnspolicy64Func(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this dnspolicy64 but we can't read it ?? %s", dnspolicy64Name)
		return nil
	}
	return nil
}

func readDnspolicy64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading dnspolicy64 state %s", dnspolicy64Name)
	data, err := client.FindResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing dnspolicy64 state %s", dnspolicy64Name)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateDnspolicy64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Get("name").(string)

	dnspolicy64 := dns.Dnspolicy64{
		Name: dnspolicy64Name,
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for dnspolicy64 %s, starting update", dnspolicy64Name)
		dnspolicy64.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for dnspolicy64 %s, starting update", dnspolicy64Name)
		dnspolicy64.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Dnspolicy64.Type(), dnspolicy64Name, &dnspolicy64)
		if err != nil {
			return fmt.Errorf("Error updating dnspolicy64 %s", dnspolicy64Name)
		}
	}
	return readDnspolicy64Func(d, meta)
}

func deleteDnspolicy64Func(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteDnspolicy64Func")
	client := meta.(*NetScalerNitroClient).client
	dnspolicy64Name := d.Id()
	err := client.DeleteResource(service.Dnspolicy64.Type(), dnspolicy64Name)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
