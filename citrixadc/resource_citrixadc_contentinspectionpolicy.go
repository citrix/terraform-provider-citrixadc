package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcContentinspectionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createContentinspectionpolicyFunc,
		ReadContext:   readContentinspectionpolicyFunc,
		UpdateContext: updateContentinspectionpolicyFunc,
		DeleteContext: deleteContentinspectionpolicyFunc,
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
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
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
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createContentinspectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createContentinspectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicyName := d.Get("name").(string)
	contentinspectionpolicy := contentinspection.Contentinspectionpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource("contentinspectionpolicy", contentinspectionpolicyName, &contentinspectionpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(contentinspectionpolicyName)

	return readContentinspectionpolicyFunc(ctx, d, meta)
}

func readContentinspectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readContentinspectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading contentinspectionpolicy state %s", contentinspectionpolicyName)
	data, err := client.FindResource("contentinspectionpolicy", contentinspectionpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing contentinspectionpolicy state %s", contentinspectionpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateContentinspectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateContentinspectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicyName := d.Get("name").(string)

	contentinspectionpolicy := contentinspection.Contentinspectionpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for contentinspectionpolicy %s, starting update", contentinspectionpolicyName)
		contentinspectionpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for contentinspectionpolicy %s, starting update", contentinspectionpolicyName)
		contentinspectionpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for contentinspectionpolicy %s, starting update", contentinspectionpolicyName)
		contentinspectionpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for contentinspectionpolicy %s, starting update", contentinspectionpolicyName)
		contentinspectionpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for contentinspectionpolicy %s, starting update", contentinspectionpolicyName)
		contentinspectionpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("contentinspectionpolicy", &contentinspectionpolicy)
		if err != nil {
			return diag.Errorf("Error updating contentinspectionpolicy %s", contentinspectionpolicyName)
		}
	}
	return readContentinspectionpolicyFunc(ctx, d, meta)
}

func deleteContentinspectionpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteContentinspectionpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	contentinspectionpolicyName := d.Id()
	err := client.DeleteResource("contentinspectionpolicy", contentinspectionpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
