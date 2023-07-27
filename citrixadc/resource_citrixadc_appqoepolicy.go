package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/appqoe"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAppqoepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAppqoepolicyFunc,
		Read:          readAppqoepolicyFunc,
		Update:        updateAppqoepolicyFunc,
		Delete:        deleteAppqoepolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createAppqoepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAppqoepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoepolicyName := d.Get("name").(string)
	appqoepolicy := appqoe.Appqoepolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Appqoepolicy.Type(), appqoepolicyName, &appqoepolicy)
	if err != nil {
		return err
	}

	d.SetId(appqoepolicyName)

	err = readAppqoepolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this appqoepolicy but we can't read it ?? %s", appqoepolicyName)
		return nil
	}
	return nil
}

func readAppqoepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAppqoepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading appqoepolicy state %s", appqoepolicyName)
	data, err := client.FindResource(service.Appqoepolicy.Type(), appqoepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing appqoepolicy state %s", appqoepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAppqoepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAppqoepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoepolicyName := d.Get("name").(string)

	appqoepolicy := appqoe.Appqoepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for appqoepolicy %s, starting update", appqoepolicyName)
		appqoepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for appqoepolicy %s, starting update", appqoepolicyName)
		appqoepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Appqoepolicy.Type(), &appqoepolicy)
		if err != nil {
			return fmt.Errorf("Error updating appqoepolicy %s", appqoepolicyName)
		}
	}
	return readAppqoepolicyFunc(d, meta)
}

func deleteAppqoepolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAppqoepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	appqoepolicyName := d.Id()
	err := client.DeleteResource(service.Appqoepolicy.Type(), appqoepolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
