package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuditnslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuditnslogpolicyFunc,
		ReadContext:   readAuditnslogpolicyFunc,
		UpdateContext: updateAuditnslogpolicyFunc,
		DeleteContext: deleteAuditnslogpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuditnslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(auditnslogpolicyName)

	return readAuditnslogpolicyFunc(ctx, d, meta)
}

func readAuditnslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func updateAuditnslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating auditnslogpolicy %s", auditnslogpolicyName)
		}
	}
	return readAuditnslogpolicyFunc(ctx, d, meta)
}

func deleteAuditnslogpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuditnslogpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	auditnslogpolicyName := d.Id()
	err := client.DeleteResource(service.Auditnslogpolicy.Type(), auditnslogpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
