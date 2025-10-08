package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/spillover"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSpilloverpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSpilloverpolicyFunc,
		ReadContext:   readSpilloverpolicyFunc,
		UpdateContext: updateSpilloverpolicyFunc,
		DeleteContext: deleteSpilloverpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
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
				Required: true,
				Computed: false,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"newname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSpilloverpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Get("name").(string)
	spilloverpolicy := spillover.Spilloverpolicy{
		Action:  d.Get("action").(string),
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Newname: d.Get("newname").(string),
		Rule:    d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Spilloverpolicy.Type(), spilloverpolicyName, &spilloverpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(spilloverpolicyName)

	return readSpilloverpolicyFunc(ctx, d, meta)
}

func readSpilloverpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading spilloverpolicy state %s", spilloverpolicyName)
	data, err := client.FindResource(service.Spilloverpolicy.Type(), spilloverpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing spilloverpolicy state %s", spilloverpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("newname", data["newname"])
	d.Set("rule", data["rule"])

	return nil

}

func updateSpilloverpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Get("name").(string)

	spilloverpolicy := spillover.Spilloverpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("newname") {
		log.Printf("[DEBUG]  citrixadc-provider: Newname has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Newname = d.Get("newname").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for spilloverpolicy %s, starting update", spilloverpolicyName)
		spilloverpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Spilloverpolicy.Type(), spilloverpolicyName, &spilloverpolicy)
		if err != nil {
			return diag.Errorf("Error updating spilloverpolicy %s", spilloverpolicyName)
		}
	}
	return readSpilloverpolicyFunc(ctx, d, meta)
}

func deleteSpilloverpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSpilloverpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	spilloverpolicyName := d.Id()
	err := client.DeleteResource(service.Spilloverpolicy.Type(), spilloverpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
