package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcCrpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createCrpolicyFunc,
		Read:          readCrpolicyFunc,
		Update:        updateCrpolicyFunc,
		Delete:        deleteCrpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCrpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createCrpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	crpolicyName := d.Get("policyname").(string)
	crpolicy := cr.Crpolicy{
		Action:     d.Get("action").(string),
		Logaction:  d.Get("logaction").(string),
		Policyname: crpolicyName,
		Rule:       d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Crpolicy.Type(), crpolicyName, &crpolicy)
	if err != nil {
		return err
	}

	d.SetId(crpolicyName)

	err = readCrpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this crpolicy but we can't read it ?? %s", crpolicyName)
		return nil
	}
	return nil
}

func readCrpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readCrpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	crpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading crpolicy state %s", crpolicyName)
	data, err := client.FindResource(service.Crpolicy.Type(), crpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing crpolicy state %s", crpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("policyname", data["policyname"])
	d.Set("action", data["action"])
	d.Set("logaction", data["logaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateCrpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateCrpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	crpolicyName := d.Get("policyname").(string)

	crpolicy := cr.Crpolicy{
		Policyname: crpolicyName,
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for crpolicy %s, starting update", crpolicyName)
		crpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for crpolicy %s, starting update", crpolicyName)
		crpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for crpolicy %s, starting update", crpolicyName)
		crpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Crpolicy.Type(), crpolicyName, &crpolicy)
		if err != nil {
			return fmt.Errorf("Error updating crpolicy %s", crpolicyName)
		}
	}
	return readCrpolicyFunc(d, meta)
}

func deleteCrpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	crpolicyName := d.Id()
	err := client.DeleteResource(service.Crpolicy.Type(), crpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
