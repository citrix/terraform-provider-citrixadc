package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationnoauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationnoauthactionFunc,
		ReadContext:   readAuthenticationnoauthactionFunc,
		UpdateContext: updateAuthenticationnoauthactionFunc,
		DeleteContext: deleteAuthenticationnoauthactionFunc,
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
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationnoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Get("name").(string)
	authenticationnoauthaction := authentication.Authenticationnoauthaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Name:                       d.Get("name").(string),
	}

	_, err := client.AddResource("authenticationnoauthaction", authenticationnoauthactionName, &authenticationnoauthaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationnoauthactionName)

	return readAuthenticationnoauthactionFunc(ctx, d, meta)
}

func readAuthenticationnoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationnoauthaction state %s", authenticationnoauthactionName)
	data, err := client.FindResource("authenticationnoauthaction", authenticationnoauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationnoauthaction state %s", authenticationnoauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("name", data["name"])

	return nil

}

func updateAuthenticationnoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Get("name").(string)

	authenticationnoauthaction := authentication.Authenticationnoauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationnoauthaction %s, starting update", authenticationnoauthactionName)
		authenticationnoauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationnoauthaction", authenticationnoauthactionName, &authenticationnoauthaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationnoauthaction %s", authenticationnoauthactionName)
		}
	}
	return readAuthenticationnoauthactionFunc(ctx, d, meta)
}

func deleteAuthenticationnoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationnoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationnoauthactionName := d.Id()
	err := client.DeleteResource("authenticationnoauthaction", authenticationnoauthactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
