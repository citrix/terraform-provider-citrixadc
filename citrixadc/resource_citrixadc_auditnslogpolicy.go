package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuditnslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuditnslogpolicyFunc,
		Read:          readAuditnslogpolicyFunc,
		Update:        updateAuditnslogpolicyFunc,
		Delete:        deleteAuditnslogpolicyFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": &schema.Schema{
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

func createAuditnslogpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuditnslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogpolicyName := d.Get("name").(string)
	auditnslogpolicy := audit.Auditnslogpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Auditnslogpolicy.Type(), auditnslogpolicyName, &auditnslogpolicy)
	if err != nil {
		return err
	}

	d.SetId(auditnslogpolicyName)

	err = readAuditnslogpolicyFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this auditnslogpolicy but we can't read it ?? %s", auditnslogpolicyName)
		return nil
	}
	return nil
}

func readAuditnslogpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuditnslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading auditnslogpolicy state %s", auditnslogpolicyName)
	data, err := client.FindResource(service.Auditnslogpolicy.Type(), auditnslogpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing auditnslogpolicy state %s", auditnslogpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuditnslogpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuditnslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogpolicyName := d.Get("name").(string)

	auditnslogpolicy := audit.Auditnslogpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for auditnslogpolicy %s, starting update", auditnslogpolicyName)
		auditnslogpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for auditnslogpolicy %s, starting update", auditnslogpolicyName)
		auditnslogpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Auditnslogpolicy.Type(), &auditnslogpolicy)
		if err != nil {
			return fmt.Errorf("Error updating auditnslogpolicy %s", auditnslogpolicyName)
		}
	}
	return readAuditnslogpolicyFunc(d, meta)
}

func deleteAuditnslogpolicyFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditnslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogpolicyName := d.Id()
	err := client.DeleteResource(service.Auditnslogpolicy.Type(), auditnslogpolicyName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
