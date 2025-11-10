package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationwebauthpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationwebauthpolicyFunc,
		ReadContext:   readAuthenticationwebauthpolicyFunc,
		UpdateContext: updateAuthenticationwebauthpolicyFunc,
		DeleteContext: deleteAuthenticationwebauthpolicyFunc,
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
		},
	}
}

func createAuthenticationwebauthpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Get("name").(string)
	authenticationwebauthpolicy := authentication.Authenticationwebauthpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationwebauthpolicyName)

	return readAuthenticationwebauthpolicyFunc(ctx, d, meta)
}

func readAuthenticationwebauthpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationwebauthpolicy state %s", authenticationwebauthpolicyName)
	data, err := client.FindResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationwebauthpolicy state %s", authenticationwebauthpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationwebauthpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Get("name").(string)

	authenticationwebauthpolicy := authentication.Authenticationwebauthpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationwebauthpolicy %s, starting update", authenticationwebauthpolicyName)
		authenticationwebauthpolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationwebauthpolicy %s, starting update", authenticationwebauthpolicyName)
		authenticationwebauthpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName, &authenticationwebauthpolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationwebauthpolicy %s", authenticationwebauthpolicyName)
		}
	}
	return readAuthenticationwebauthpolicyFunc(ctx, d, meta)
}

func deleteAuthenticationwebauthpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationwebauthpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationwebauthpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationwebauthpolicy.Type(), authenticationwebauthpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
