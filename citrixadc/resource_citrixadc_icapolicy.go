package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcIcapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createIcapolicyFunc,
		ReadContext:   readIcapolicyFunc,
		UpdateContext: updateIcapolicyFunc,
		DeleteContext: deleteIcapolicyFunc,
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
		},
	}
}

func createIcapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Get("name").(string)
	icapolicy := ica.Icapolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource("icapolicy", icapolicyName, &icapolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(icapolicyName)

	return readIcapolicyFunc(ctx, d, meta)
}

func readIcapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading icapolicy state %s", icapolicyName)
	data, err := client.FindResource("icapolicy", icapolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing icapolicy state %s", icapolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateIcapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Get("name").(string)

	icapolicy := ica.Icapolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for icapolicy %s, starting update", icapolicyName)
		icapolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource("icapolicy", &icapolicy)
		if err != nil {
			return diag.Errorf("Error updating icapolicy %s", icapolicyName)
		}
	}
	return readIcapolicyFunc(ctx, d, meta)
}

func deleteIcapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteIcapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	icapolicyName := d.Id()
	err := client.DeleteResource("icapolicy", icapolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
