package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationsamlpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationsamlpolicyFunc,
		ReadContext:   readAuthenticationsamlpolicyFunc,
		UpdateContext: updateAuthenticationsamlpolicyFunc,
		DeleteContext: deleteAuthenticationsamlpolicyFunc,
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
			"reqaction": {
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

func createAuthenticationsamlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Get("name").(string)
	authenticationsamlpolicy := authentication.Authenticationsamlpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Rule:      d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationsamlpolicyName)

	return readAuthenticationsamlpolicyFunc(ctx, d, meta)
}

func readAuthenticationsamlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationsamlpolicy state %s", authenticationsamlpolicyName)
	data, err := client.FindResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationsamlpolicy state %s", authenticationsamlpolicyName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("reqaction", data["reqaction"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationsamlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Get("name").(string)

	authenticationsamlpolicy := authentication.Authenticationsamlpolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("reqaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Reqaction has changed for authenticationsamlpolicy %s, starting update", authenticationsamlpolicyName)
		authenticationsamlpolicy.Reqaction = d.Get("reqaction").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationsamlpolicy %s, starting update", authenticationsamlpolicyName)
		authenticationsamlpolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName, &authenticationsamlpolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationsamlpolicy %s", authenticationsamlpolicyName)
		}
	}
	return readAuthenticationsamlpolicyFunc(ctx, d, meta)
}

func deleteAuthenticationsamlpolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlpolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlpolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlpolicy.Type(), authenticationsamlpolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
