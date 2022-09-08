package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcTmtrafficpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createTmtrafficpolicyFunc,
		Read:          readTmtrafficpolicyFunc,
		Update:        updateTmtrafficpolicyFunc,
		Delete:        deleteTmtrafficpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createTmtrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmtrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficpolicyName := d.Get("name").(string)
	
	tmtrafficpolicy := tm.Tmtrafficpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName, &tmtrafficpolicy)
	if err != nil {
		return err
	}

	d.SetId(tmtrafficpolicyName)

	err = readTmtrafficpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this tmtrafficpolicy but we can't read it ?? %s", tmtrafficpolicyName)
		return nil
	}
	return nil
}

func readTmtrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmtrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmtrafficpolicy state %s", tmtrafficpolicyName)
	data, err := client.FindResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmtrafficpolicy state %s", tmtrafficpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateTmtrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmtrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficpolicyName := d.Get("name").(string)

	tmtrafficpolicy := tm.Tmtrafficpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for tmtrafficpolicy %s, starting update", tmtrafficpolicyName)
		tmtrafficpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for tmtrafficpolicy %s, starting update", tmtrafficpolicyName)
		tmtrafficpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmtrafficpolicy.Type(), &tmtrafficpolicy)
		if err != nil {
			return fmt.Errorf("Error updating tmtrafficpolicy %s", tmtrafficpolicyName)
		}
	}
	return readTmtrafficpolicyFunc(d, meta)
}

func deleteTmtrafficpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmtrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmtrafficpolicyName := d.Id()
	err := client.DeleteResource(service.Tmtrafficpolicy.Type(), tmtrafficpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
