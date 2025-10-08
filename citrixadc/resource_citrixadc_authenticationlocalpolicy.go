package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationlocalpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationlocalpolicyFunc,
		ReadContext:   readAuthenticationlocalpolicyFunc,
		UpdateContext: updateAuthenticationlocalpolicyFunc,
		DeleteContext: deleteAuthenticationlocalpolicyFunc,
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
		},
	}
}

func createAuthenticationlocalpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Get("name").(string)
	authenticationlocalpolicy := authentication.Authenticationlocalpolicy{
		Name: d.Get("name").(string),
		Rule: d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName, &authenticationlocalpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationlocalpolicyName)

	return readAuthenticationlocalpolicyFunc(ctx, d, meta)
}

func readAuthenticationlocalpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationlocalpolicy state %s", authenticationlocalpolicyName)
	data, err := client.FindResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationlocalpolicy state %s", authenticationlocalpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationlocalpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Get("name").(string)

	authenticationlocalpolicy := authentication.Authenticationlocalpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationlocalpolicy %s, starting update", authenticationlocalpolicyName)
		authenticationlocalpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName, &authenticationlocalpolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationlocalpolicy %s", authenticationlocalpolicyName)
		}
	}
	return readAuthenticationlocalpolicyFunc(ctx, d, meta)
}

func deleteAuthenticationlocalpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationlocalpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationlocalpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationlocalpolicy.Type(), authenticationlocalpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
