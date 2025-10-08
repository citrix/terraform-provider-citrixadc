package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTmsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTmsessionpolicyFunc,
		ReadContext:   readTmsessionpolicyFunc,
		UpdateContext: updateTmsessionpolicyFunc,
		DeleteContext: deleteTmsessionpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
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

func createTmsessionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTmsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionpolicyName := d.Get("name").(string)

	tmsessionpolicy := tm.Tmsessionpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName, &tmsessionpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tmsessionpolicyName)

	return readTmsessionpolicyFunc(ctx, d, meta)
}

func readTmsessionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTmsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tmsessionpolicy state %s", tmsessionpolicyName)
	data, err := client.FindResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tmsessionpolicy state %s", tmsessionpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateTmsessionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTmsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionpolicyName := d.Get("name").(string)

	tmsessionpolicy := tm.Tmsessionpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for tmsessionpolicy %s, starting update", tmsessionpolicyName)
		tmsessionpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for tmsessionpolicy %s, starting update", tmsessionpolicyName)
		tmsessionpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tmsessionpolicy.Type(), &tmsessionpolicy)
		if err != nil {
			return diag.Errorf("Error updating tmsessionpolicy %s", tmsessionpolicyName)
		}
	}
	return readTmsessionpolicyFunc(ctx, d, meta)
}

func deleteTmsessionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTmsessionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tmsessionpolicyName := d.Id()
	err := client.DeleteResource(service.Tmsessionpolicy.Type(), tmsessionpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
