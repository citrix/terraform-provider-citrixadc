package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/filter"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcFilterpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFilterpolicyFunc,
		Read:          readFilterpolicyFunc,
		Update:        updateFilterpolicyFunc,
		Delete:        deleteFilterpolicyFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createFilterpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Get("name").(string)
	filterpolicy := filter.Filterpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(netscaler.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
	if err != nil {
		return err
	}

	d.SetId(filterpolicyName)

	err = readFilterpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this filterpolicy but we can't read it ?? %s", filterpolicyName)
		return nil
	}
	return nil
}

func readFilterpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading filterpolicy state %s", filterpolicyName)
	data, err := client.FindResource(netscaler.Filterpolicy.Type(), filterpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing filterpolicy state %s", filterpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("resaction", data["resaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateFilterpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Get("name").(string)

	filterpolicy := filter.Filterpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for filterpolicy %s, starting update", filterpolicyName)
		filterpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for filterpolicy %s, starting update", filterpolicyName)
		filterpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("resaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Resaction has changed for filterpolicy %s, starting update", filterpolicyName)
		filterpolicy.Resaction = d.Get("resaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for filterpolicy %s, starting update", filterpolicyName)
		filterpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
		if err != nil {
			return fmt.Errorf("Error updating filterpolicy %s", filterpolicyName)
		}
	}
	return readFilterpolicyFunc(d, meta)
}

func deleteFilterpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Id()
	err := client.DeleteResource(netscaler.Filterpolicy.Type(), filterpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
