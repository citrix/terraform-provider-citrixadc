package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationcertaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationcertactionFunc,
		ReadContext:   readAuthenticationcertactionFunc,
		UpdateContext: updateAuthenticationcertactionFunc,
		DeleteContext: deleteAuthenticationcertactionFunc,
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
			"groupnamefield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"twofactor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usernamefield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationcertactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationcertactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertactionName := d.Get("name").(string)
	authenticationcertaction := authentication.Authenticationcertaction{
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Groupnamefield:             d.Get("groupnamefield").(string),
		Name:                       d.Get("name").(string),
		Twofactor:                  d.Get("twofactor").(string),
		Usernamefield:              d.Get("usernamefield").(string),
	}

	_, err := client.AddResource(service.Authenticationcertaction.Type(), authenticationcertactionName, &authenticationcertaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationcertactionName)

	return readAuthenticationcertactionFunc(ctx, d, meta)
}

func readAuthenticationcertactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationcertactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationcertaction state %s", authenticationcertactionName)
	data, err := client.FindResource(service.Authenticationcertaction.Type(), authenticationcertactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationcertaction state %s", authenticationcertactionName)
		d.SetId("")
		return nil
	}
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("groupnamefield", data["groupnamefield"])
	d.Set("name", data["name"])
	d.Set("twofactor", data["twofactor"])
	d.Set("usernamefield", data["usernamefield"])

	return nil

}

func updateAuthenticationcertactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationcertactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertactionName := d.Get("name").(string)

	authenticationcertaction := authentication.Authenticationcertaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationcertaction %s, starting update", authenticationcertactionName)
		authenticationcertaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("groupnamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnamefield has changed for authenticationcertaction %s, starting update", authenticationcertactionName)
		authenticationcertaction.Groupnamefield = d.Get("groupnamefield").(string)
		hasChange = true
	}
	if d.HasChange("twofactor") {
		log.Printf("[DEBUG]  citrixadc-provider: Twofactor has changed for authenticationcertaction %s, starting update", authenticationcertactionName)
		authenticationcertaction.Twofactor = d.Get("twofactor").(string)
		hasChange = true
	}
	if d.HasChange("usernamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Usernamefield has changed for authenticationcertaction %s, starting update", authenticationcertactionName)
		authenticationcertaction.Usernamefield = d.Get("usernamefield").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationcertaction.Type(), authenticationcertactionName, &authenticationcertaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationcertaction %s", authenticationcertactionName)
		}
	}
	return readAuthenticationcertactionFunc(ctx, d, meta)
}

func deleteAuthenticationcertactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationcertactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationcertactionName := d.Id()
	err := client.DeleteResource(service.Authenticationcertaction.Type(), authenticationcertactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
