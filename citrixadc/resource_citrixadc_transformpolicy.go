package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/transform"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTransformpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTransformpolicyFunc,
		ReadContext:   readTransformpolicyFunc,
		UpdateContext: updateTransformpolicyFunc,
		DeleteContext: deleteTransformpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"profilename": {
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

func createTransformpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTransformpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicyName := d.Get("name").(string)
	transformpolicy := transform.Transformpolicy{
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Transformpolicy.Type(), transformpolicyName, &transformpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(transformpolicyName)

	return readTransformpolicyFunc(ctx, d, meta)
}

func readTransformpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTransformpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading transformpolicy state %s", transformpolicyName)
	data, err := client.FindResource(service.Transformpolicy.Type(), transformpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing transformpolicy state %s", transformpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("profilename", data["profilename"])
	d.Set("rule", data["rule"])

	return nil

}

func updateTransformpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTransformpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicyName := d.Get("name").(string)

	transformpolicy := transform.Transformpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for transformpolicy %s, starting update", transformpolicyName)
		transformpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for transformpolicy %s, starting update", transformpolicyName)
		transformpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for transformpolicy %s, starting update", transformpolicyName)
		transformpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("profilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Profilename has changed for transformpolicy %s, starting update", transformpolicyName)
		transformpolicy.Profilename = d.Get("profilename").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for transformpolicy %s, starting update", transformpolicyName)
		transformpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Transformpolicy.Type(), transformpolicyName, &transformpolicy)
		if err != nil {
			return diag.Errorf("Error updating transformpolicy %s", transformpolicyName)
		}
	}
	return readTransformpolicyFunc(ctx, d, meta)
}

func deleteTransformpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTransformpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	transformpolicyName := d.Id()
	err := client.DeleteResource(service.Transformpolicy.Type(), transformpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
