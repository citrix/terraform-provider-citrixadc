package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceCitrixAdcVpnurlpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnurlpolicyFunc,
		ReadContext:   readVpnurlpolicyFunc,
		UpdateContext: updateVpnurlpolicyFunc,
		DeleteContext: deleteVpnurlpolicyFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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

func createVpnurlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnurlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlpolicyName := d.Get("name").(string)
	vpnurlpolicy := vpn.Vpnurlpolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource("vpnurlpolicy", vpnurlpolicyName, &vpnurlpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(vpnurlpolicyName)

	return readVpnurlpolicyFunc(ctx, d, meta)
}

func readVpnurlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnurlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnurlpolicy state %s", vpnurlpolicyName)
	data, err := client.FindResource("vpnurlpolicy", vpnurlpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnurlpolicy state %s", vpnurlpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("action", data["action"])
	d.Set("comment", data["comment"])
	d.Set("logaction", data["logaction"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateVpnurlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnurlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlpolicyName := d.Get("name").(string)

	vpnurlpolicy := vpn.Vpnurlpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for vpnurlpolicy %s, starting update", vpnurlpolicyName)
		vpnurlpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for vpnurlpolicy %s, starting update", vpnurlpolicyName)
		vpnurlpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for vpnurlpolicy %s, starting update", vpnurlpolicyName)
		vpnurlpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for vpnurlpolicy %s, starting update", vpnurlpolicyName)
		vpnurlpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("vpnurlpolicy", vpnurlpolicyName, &vpnurlpolicy)
		if err != nil {
			return diag.Errorf("Error updating vpnurlpolicy %s", vpnurlpolicyName)
		}
	}
	return readVpnurlpolicyFunc(ctx, d, meta)
}

func deleteVpnurlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnurlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnurlpolicyName := d.Id()
	err := client.DeleteResource("vpnurlpolicy", vpnurlpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
