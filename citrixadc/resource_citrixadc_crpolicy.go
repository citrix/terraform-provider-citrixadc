package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcCrpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createCrpolicyFunc,
		ReadContext:   readCrpolicyFunc,
		UpdateContext: updateCrpolicyFunc,
		DeleteContext: deleteCrpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"policyname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createCrpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		return diag.FromErr(err)
	}

	d.SetId(crpolicyName)

	return readCrpolicyFunc(ctx, d, meta)
}

func readCrpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func updateCrpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
			return diag.Errorf("Error updating crpolicy %s", crpolicyName)
		}
	}
	return readCrpolicyFunc(ctx, d, meta)
}

func deleteCrpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteCrpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	crpolicyName := d.Id()
	err := client.DeleteResource(service.Crpolicy.Type(), crpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
