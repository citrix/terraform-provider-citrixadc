package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcSslpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createSslpolicyFunc,
		ReadContext:   readSslpolicyFunc,
		UpdateContext: updateSslpolicyFunc,
		DeleteContext: deleteSslpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"reqaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"undefaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createSslpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicyName := d.Get("name").(string)

	sslpolicy := ssl.Sslpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Name:        sslpolicyName,
		Reqaction:   d.Get("reqaction").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	_, err := client.AddResource(service.Sslpolicy.Type(), sslpolicyName, &sslpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(sslpolicyName)

	return readSslpolicyFunc(ctx, d, meta)
}

func readSslpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslpolicy state %s", sslpolicyName)
	data, err := client.FindResource(service.Sslpolicy.Type(), sslpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslpolicy state %s", sslpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])
	d.Set("undefaction", data["undefaction"])

	return nil

}

func updateSslpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicyName := d.Get("name").(string)

	sslpolicy := ssl.Sslpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}
	if d.HasChange("undefaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Undefaction has changed for sslpolicy %s, starting update", sslpolicyName)
		sslpolicy.Undefaction = d.Get("undefaction").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Sslpolicy.Type(), sslpolicyName, &sslpolicy)
		if err != nil {
			return diag.Errorf("Error updating sslpolicy %s", sslpolicyName)
		}
	}
	return readSslpolicyFunc(ctx, d, meta)
}

func deleteSslpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	sslpolicyName := d.Id()
	err := client.DeleteResource(service.Sslpolicy.Type(), sslpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
