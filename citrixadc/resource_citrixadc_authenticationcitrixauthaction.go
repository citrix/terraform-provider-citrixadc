package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationcitrixauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationcitrixauthactionFunc,
		ReadContext:   readAuthenticationcitrixauthactionFunc,
		UpdateContext: updateAuthenticationcitrixauthactionFunc,
		DeleteContext: deleteAuthenticationcitrixauthactionFunc,
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
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authenticationtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcitrixauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Get("name").(string)
	authenticationcitrixauthaction := authentication.Authenticationcitrixauthaction{
		Authentication:     d.Get("authentication").(string),
		Authenticationtype: d.Get("authenticationtype").(string),
		Name:               d.Get("name").(string),
	}

	_, err := client.AddResource("authenticationcitrixauthaction", authenticationcitrixauthactionName, &authenticationcitrixauthaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationcitrixauthactionName)

	return readAuthenticationcitrixauthactionFunc(ctx, d, meta)
}

func readAuthenticationcitrixauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcitrixauthaction state %s", authenticationcitrixauthactionName)
	data, err := client.FindResource("authenticationcitrixauthaction", authenticationcitrixauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcitrixauthaction state %s", authenticationcitrixauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("authentication", data["authentication"])
	d.Set("authenticationtype", data["authenticationtype"])
	d.Set("name", data["name"])

	return nil

}

func updateAuthenticationcitrixauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Get("name").(string)

	authenticationcitrixauthaction := authentication.Authenticationcitrixauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationcitrixauthaction %s, starting update", authenticationcitrixauthactionName)
		authenticationcitrixauthaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authenticationtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationtype has changed for authenticationcitrixauthaction %s, starting update", authenticationcitrixauthactionName)
		authenticationcitrixauthaction.Authenticationtype = d.Get("authenticationtype").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationcitrixauthaction", authenticationcitrixauthactionName, &authenticationcitrixauthaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationcitrixauthaction %s", authenticationcitrixauthactionName)
		}
	}
	return readAuthenticationcitrixauthactionFunc(ctx, d, meta)
}

func deleteAuthenticationcitrixauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcitrixauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcitrixauthactionName := d.Id()
	err := client.DeleteResource("authenticationcitrixauthaction", authenticationcitrixauthactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
