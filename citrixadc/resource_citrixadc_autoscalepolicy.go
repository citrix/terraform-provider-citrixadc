package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAutoscalepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAutoscalepolicyFunc,
		ReadContext:   readAutoscalepolicyFunc,
		UpdateContext: updateAutoscalepolicyFunc,
		DeleteContext: deleteAutoscalepolicyFunc,
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
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logaction": {
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

func createAutoscalepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Get("name").(string)
	autoscalepolicy := autoscale.Autoscalepolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Autoscalepolicy.Type(), autoscalepolicyName, &autoscalepolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(autoscalepolicyName)

	return readAutoscalepolicyFunc(ctx, d, meta)
}

func readAutoscalepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading autoscalepolicy state %s", autoscalepolicyName)
	data, err := client.FindResource(service.Autoscalepolicy.Type(), autoscalepolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing autoscalepolicy state %s", autoscalepolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAutoscalepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Get("name").(string)

	autoscalepolicy := autoscale.Autoscalepolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for autoscalepolicy %s, starting update", autoscalepolicyName)
		autoscalepolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Autoscalepolicy.Type(), &autoscalepolicy)
		if err != nil {
			return diag.Errorf("Error updating autoscalepolicy %s", autoscalepolicyName)
		}
	}
	return readAutoscalepolicyFunc(ctx, d, meta)
}

func deleteAutoscalepolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAutoscalepolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	autoscalepolicyName := d.Id()
	err := client.DeleteResource(service.Autoscalepolicy.Type(), autoscalepolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
