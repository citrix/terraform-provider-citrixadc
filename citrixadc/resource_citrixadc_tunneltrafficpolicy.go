package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tunnel"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcTunneltrafficpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createTunneltrafficpolicyFunc,
		ReadContext:   readTunneltrafficpolicyFunc,
		UpdateContext: updateTunneltrafficpolicyFunc,
		DeleteContext: deleteTunneltrafficpolicyFunc,
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

func createTunneltrafficpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createTunneltrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tunneltrafficpolicyName := d.Get("name").(string)
	tunneltrafficpolicy := tunnel.Tunneltrafficpolicy{
		Action:    d.Get("action").(string),
		Comment:   d.Get("comment").(string),
		Logaction: d.Get("logaction").(string),
		Name:      d.Get("name").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Tunneltrafficpolicy.Type(), tunneltrafficpolicyName, &tunneltrafficpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tunneltrafficpolicyName)

	return readTunneltrafficpolicyFunc(ctx, d, meta)
}

func readTunneltrafficpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readTunneltrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tunneltrafficpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading tunneltrafficpolicy state %s", tunneltrafficpolicyName)
	data, err := client.FindResource(service.Tunneltrafficpolicy.Type(), tunneltrafficpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing tunneltrafficpolicy state %s", tunneltrafficpolicyName)
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

func updateTunneltrafficpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateTunneltrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tunneltrafficpolicyName := d.Get("name").(string)

	tunneltrafficpolicy := tunnel.Tunneltrafficpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for tunneltrafficpolicy %s, starting update", tunneltrafficpolicyName)
		tunneltrafficpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("comment") {
		log.Printf("[DEBUG]  citrixadc-provider: Comment has changed for tunneltrafficpolicy %s, starting update", tunneltrafficpolicyName)
		tunneltrafficpolicy.Comment = d.Get("comment").(string)
		hasChange = true
	}
	if d.HasChange("logaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Logaction has changed for tunneltrafficpolicy %s, starting update", tunneltrafficpolicyName)
		tunneltrafficpolicy.Logaction = d.Get("logaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for tunneltrafficpolicy %s, starting update", tunneltrafficpolicyName)
		tunneltrafficpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Tunneltrafficpolicy.Type(), &tunneltrafficpolicy)
		if err != nil {
			return diag.Errorf("Error updating tunneltrafficpolicy %s", tunneltrafficpolicyName)
		}
	}
	return readTunneltrafficpolicyFunc(ctx, d, meta)
}

func deleteTunneltrafficpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteTunneltrafficpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	tunneltrafficpolicyName := d.Id()
	err := client.DeleteResource(service.Tunneltrafficpolicy.Type(), tunneltrafficpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
