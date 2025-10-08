package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/filter"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcFilterpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createFilterpolicyFunc,
		ReadContext:   readFilterpolicyFunc,
		UpdateContext: updateFilterpolicyFunc,
		DeleteContext: deleteFilterpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resaction": {
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

func createFilterpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Get("name").(string)
	filterpolicy := filter.Filterpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(filterpolicyName)

	return readFilterpolicyFunc(ctx, d, meta)
}

func readFilterpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading filterpolicy state %s", filterpolicyName)
	data, err := client.FindResource(service.Filterpolicy.Type(), filterpolicyName)
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

func updateFilterpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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
		_, err := client.UpdateResource(service.Filterpolicy.Type(), filterpolicyName, &filterpolicy)
		if err != nil {
			return diag.Errorf("Error updating filterpolicy %s", filterpolicyName)
		}
	}
	return readFilterpolicyFunc(ctx, d, meta)
}

func deleteFilterpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteFilterpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	filterpolicyName := d.Id()
	err := client.DeleteResource(service.Filterpolicy.Type(), filterpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
