package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationradiuspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationradiuspolicyFunc,
		ReadContext:   readAuthenticationradiuspolicyFunc,
		UpdateContext: updateAuthenticationradiuspolicyFunc,
		DeleteContext: deleteAuthenticationradiuspolicyFunc,
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

func createAuthenticationradiuspolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Get("name").(string)
	authenticationradiuspolicy := authentication.Authenticationradiuspolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationradiuspolicyName)

	return readAuthenticationradiuspolicyFunc(ctx, d, meta)
}

func readAuthenticationradiuspolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationradiuspolicy state %s", authenticationradiuspolicyName)
	data, err := client.FindResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationradiuspolicy state %s", authenticationradiuspolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationradiuspolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Get("name").(string)

	authenticationradiuspolicy := authentication.Authenticationradiuspolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationradiuspolicy %s, starting update", authenticationradiuspolicyName)
		authenticationradiuspolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationradiuspolicy %s, starting update", authenticationradiuspolicyName)
		authenticationradiuspolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName, &authenticationradiuspolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationradiuspolicy %s", authenticationradiuspolicyName)
		}
	}
	return readAuthenticationradiuspolicyFunc(ctx, d, meta)
}

func deleteAuthenticationradiuspolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationradiuspolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationradiuspolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationradiuspolicy.Type(), authenticationradiuspolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
