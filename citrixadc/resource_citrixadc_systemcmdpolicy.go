package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcSystemcmdpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSystemcmdpolicyFunc,
		Read:          readSystemcmdpolicyFunc,
		Update:        updateSystemcmdpolicyFunc,
		Delete:        deleteSystemcmdpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cmdspec": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createSystemcmdpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Get("policyname").(string)

	systemcmdpolicy := system.Systemcmdpolicy{
		Action:     d.Get("action").(string),
		Cmdspec:    d.Get("cmdspec").(string),
		Policyname: d.Get("policyname").(string),
	}

	_, err := client.AddResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName, &systemcmdpolicy)
	if err != nil {
		return err
	}

	d.SetId(systemcmdpolicyName)

	err = readSystemcmdpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this systemcmdpolicy but we can't read it ?? %s", systemcmdpolicyName)
		return nil
	}
	return nil
}

func readSystemcmdpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading systemcmdpolicy state %s", systemcmdpolicyName)
	data, err := client.FindResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing systemcmdpolicy state %s", systemcmdpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("cmdspec", data["cmdspec"])
	d.Set("policyname", data["policyname"])

	return nil

}

func updateSystemcmdpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Get("policyname").(string)

	systemcmdpolicy := system.Systemcmdpolicy{
		Policyname: d.Get("policyname").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for systemcmdpolicy %s, starting update", systemcmdpolicyName)
		systemcmdpolicy.Action = d.Get("action").(string)
		systemcmdpolicy.Cmdspec = d.Get("cmdspec").(string)
		hasChange = true
	}
	if d.HasChange("cmdspec") {
		log.Printf("[DEBUG]  citrixadc-provider: Cmdspec has changed for systemcmdpolicy %s, starting update", systemcmdpolicyName)
		systemcmdpolicy.Cmdspec = d.Get("cmdspec").(string)
		systemcmdpolicy.Action = d.Get("action").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName, &systemcmdpolicy)
		if err != nil {
			return fmt.Errorf("Error updating systemcmdpolicy %s:%s", systemcmdpolicyName, err.Error())
		}
	}
	return readSystemcmdpolicyFunc(d, meta)
}

func deleteSystemcmdpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSystemcmdpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	systemcmdpolicyName := d.Id()
	err := client.DeleteResource(service.Systemcmdpolicy.Type(), systemcmdpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
