package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcFeopolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFeopolicyFunc,
		ReadContext:   readFeopolicyFunc,
		UpdateContext: updateFeopolicyFunc,
		DeleteContext: deleteFeopolicyFunc,
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

func createFeopolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Get("name").(string)
	feopolicy := feo.Feopolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource("feopolicy", feopolicyName, &feopolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(feopolicyName)

	return readFeopolicyFunc(ctx, d, meta)
}

func readFeopolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading feopolicy state %s", feopolicyName)
	data, err := client.FindResource("feopolicy", feopolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing feopolicy state %s", feopolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateFeopolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Get("name").(string)

	feopolicy := feo.Feopolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for feopolicy %s, starting update", feopolicyName)
		feopolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for feopolicy %s, starting update", feopolicyName)
		feopolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("feopolicy", &feopolicy)
		if err != nil {
			return diag.Errorf("Error updating feopolicy %s", feopolicyName)
		}
	}
	return readFeopolicyFunc(ctx, d, meta)
}

func deleteFeopolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFeopolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	feopolicyName := d.Id()
	err := client.DeleteResource("feopolicy", feopolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
