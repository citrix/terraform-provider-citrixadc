package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationcertpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationcertpolicyFunc,
		ReadContext:   readAuthenticationcertpolicyFunc,
		UpdateContext: updateAuthenticationcertpolicyFunc,
		DeleteContext: deleteAuthenticationcertpolicyFunc,
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
			"rule": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"reqaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcertpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Get("name").(string)
	authenticationcertpolicy := authentication.Authenticationcertpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationcertpolicyName)

	return readAuthenticationcertpolicyFunc(ctx, d, meta)
}

func readAuthenticationcertpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcertpolicy state %s", authenticationcertpolicyName)
	data, err := client.FindResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcertpolicy state %s", authenticationcertpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationcertpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Get("name").(string)

	authenticationcertpolicy := authentication.Authenticationcertpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationcertpolicy %s, starting update", authenticationcertpolicyName)
		authenticationcertpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationcertpolicy %s, starting update", authenticationcertpolicyName)
		authenticationcertpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName, &authenticationcertpolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationcertpolicy %s", authenticationcertpolicyName)
		}
	}
	return readAuthenticationcertpolicyFunc(ctx, d, meta)
}

func deleteAuthenticationcertpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcertpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationcertpolicy.Type(), authenticationcertpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
