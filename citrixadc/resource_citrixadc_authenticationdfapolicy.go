package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationdfapolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationdfapolicyFunc,
		ReadContext:   readAuthenticationdfapolicyFunc,
		UpdateContext: updateAuthenticationdfapolicyFunc,
		DeleteContext: deleteAuthenticationdfapolicyFunc,
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
			"action": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
		},
	}
}

func createAuthenticationdfapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationdfapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfapolicyName := d.Get("name").(string)
	authenticationdfapolicy := authentication.Authenticationdfapolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	_, err := client.AddResource(service.Authenticationdfapolicy.Type(), authenticationdfapolicyName, &authenticationdfapolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationdfapolicyName)

	return readAuthenticationdfapolicyFunc(ctx, d, meta)
}

func readAuthenticationdfapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationdfapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfapolicyName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationdfapolicy state %s", authenticationdfapolicyName)
	data, err := client.FindResource(service.Authenticationdfapolicy.Type(), authenticationdfapolicyName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationdfapolicy state %s", authenticationdfapolicyName)
		d.SetId("")
		return nil
	}
	d.Set("action", data["action"])
	d.Set("name", data["name"])
	d.Set("rule", data["rule"])

	return nil

}

func updateAuthenticationdfapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationdfapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfapolicyName := d.Get("name").(string)

	authenticationdfapolicy := authentication.Authenticationdfapolicy{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("action") {
		log.Printf("[DEBUG]  citrixadc-provider: Action has changed for authenticationdfapolicy %s, starting update", authenticationdfapolicyName)
		authenticationdfapolicy.Action = d.Get("action").(string)
		hasChange = true
	}
	if d.HasChange("rule") {
		log.Printf("[DEBUG]  citrixadc-provider: Rule has changed for authenticationdfapolicy %s, starting update", authenticationdfapolicyName)
		authenticationdfapolicy.Rule = d.Get("rule").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationdfapolicy.Type(), authenticationdfapolicyName, &authenticationdfapolicy)
		if err != nil {
			return diag.Errorf("Error updating authenticationdfapolicy %s", authenticationdfapolicyName)
		}
	}
	return readAuthenticationdfapolicyFunc(ctx, d, meta)
}

func deleteAuthenticationdfapolicyFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationdfapolicyFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationdfapolicyName := d.Id()
	err := client.DeleteResource(service.Authenticationdfapolicy.Type(), authenticationdfapolicyName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
