package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationstorefrontauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationstorefrontauthactionFunc,
		ReadContext:   readAuthenticationstorefrontauthactionFunc,
		UpdateContext: updateAuthenticationstorefrontauthactionFunc,
		DeleteContext: deleteAuthenticationstorefrontauthactionFunc,
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
			"serverurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationstorefrontauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Get("name").(string)
	authenticationstorefrontauthaction := authentication.Authenticationstorefrontauthaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Domain:                     d.Get("domain").(string),
		Name:                       d.Get("name").(string),
		Serverurl:                  d.Get("serverurl").(string),
	}

	_, err := client.AddResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName, &authenticationstorefrontauthaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationstorefrontauthactionName)

	return readAuthenticationstorefrontauthactionFunc(ctx, d, meta)
}

func readAuthenticationstorefrontauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationstorefrontauthaction state %s", authenticationstorefrontauthactionName)
	data, err := client.FindResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationstorefrontauthaction state %s", authenticationstorefrontauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("domain", data["domain"])
	d.Set("name", data["name"])
	d.Set("serverurl", data["serverurl"])

	return nil

}

func updateAuthenticationstorefrontauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Get("name").(string)

	authenticationstorefrontauthaction := authentication.Authenticationstorefrontauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("domain") {
		log.Printf("[DEBUG]  citrixadc-provider: Domain has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Domain = d.Get("domain").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationstorefrontauthaction %s, starting update", authenticationstorefrontauthactionName)
		authenticationstorefrontauthaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName, &authenticationstorefrontauthaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationstorefrontauthaction %s", authenticationstorefrontauthactionName)
		}
	}
	return readAuthenticationstorefrontauthactionFunc(ctx, d, meta)
}

func deleteAuthenticationstorefrontauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationstorefrontauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationstorefrontauthactionName := d.Id()
	err := client.DeleteResource("authenticationstorefrontauthaction", authenticationstorefrontauthactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
