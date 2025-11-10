package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationldappolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationldappolicyFunc,
		ReadContext:   readAuthenticationldappolicyFunc,
		UpdateContext: updateAuthenticationldappolicyFunc,
		DeleteContext: deleteAuthenticationldappolicyFunc,
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

func createAuthenticationldappolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Get("name").(string)
	authenticationldappolicy := authentication.Authenticationldappolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName, &authenticationldappolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationldappolicyName)

	return readAuthenticationldappolicyFunc(ctx, d, meta)
}

func readAuthenticationldappolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationldappolicy state %s", authenticationldappolicyName)
	data, err := client.FindResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationldappolicy state %s", authenticationldappolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationldappolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Get("name").(string)

	authenticationldappolicy := authentication.Authenticationldappolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationldappolicy %s, starting update", authenticationldappolicyName)
		authenticationldappolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationldappolicy %s, starting update", authenticationldappolicyName)
		authenticationldappolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName, &authenticationldappolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationldappolicy %s", authenticationldappolicyName)
		}
	}
	return readAuthenticationldappolicyFunc(ctx, d, meta)
}

func deleteAuthenticationldappolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationldappolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldappolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationldappolicy.Type(), authenticationldappolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
