package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationloginschema() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationloginschemaFunc,
		ReadContext:   readAuthenticationloginschemaFunc,
		UpdateContext: updateAuthenticationloginschemaFunc,
		DeleteContext: deleteAuthenticationloginschemaFunc,
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
			"authenticationschema": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"authenticationstrength": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"passwdexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwordcredentialindex": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ssocredentials": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usercredentialindex": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"userexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationloginschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationloginschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemaName := d.Get("name").(string)
	authenticationloginschema := authentication.Authenticationloginschema{
		Authenticationschema:    d.Get("authenticationschema").(string),
		Authenticationstrength:  d.Get("authenticationstrength").(int),
		Name:                    d.Get("name").(string),
		Passwdexpression:        d.Get("passwdexpression").(string),
		Passwordcredentialindex: d.Get("passwordcredentialindex").(int),
		Ssocredentials:          d.Get("ssocredentials").(string),
		Usercredentialindex:     d.Get("usercredentialindex").(int),
		Userexpression:          d.Get("userexpression").(string),
	}

	_, err := client.AddResource(service.Authenticationloginschema.Type(), authenticationloginschemaName, &authenticationloginschema)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationloginschemaName)

	return readAuthenticationloginschemaFunc(ctx, d, meta)
}

func readAuthenticationloginschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationloginschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemaName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationloginschema state %s", authenticationloginschemaName)
	data, err := client.FindResource(service.Authenticationloginschema.Type(), authenticationloginschemaName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationloginschema state %s", authenticationloginschemaName)
		d.SetId("")
		return nil
	}
	d.Set("authenticationschema", data["authenticationschema"])
	setToInt("authenticationstrength", d, data["authenticationstrength"])
	d.Set("name", data["name"])
	d.Set("passwdexpression", data["passwdexpression"])
	setToInt("passwordcredentialindex", d, data["passwordcredentialindex"])
	d.Set("ssocredentials", data["ssocredentials"])
	setToInt("usercredentialindex", d, data["usercredentialindex"])
	d.Set("userexpression", data["userexpression"])

	return nil

}

func updateAuthenticationloginschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationloginschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemaName := d.Get("name").(string)

	authenticationloginschema := authentication.Authenticationloginschema{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("authenticationschema") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationschema has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Authenticationschema = d.Get("authenticationschema").(string)
		hasChange = true
	}
	if d.HasChange("authenticationstrength") {
		log.Printf("[DEBUG]  citrixadc-provider: Authenticationstrength has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Authenticationstrength = d.Get("authenticationstrength").(int)
		hasChange = true
	}
	if d.HasChange("passwdexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdexpression has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Passwdexpression = d.Get("passwdexpression").(string)
		hasChange = true
	}
	if d.HasChange("passwordcredentialindex") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwordcredentialindex has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Passwordcredentialindex = d.Get("passwordcredentialindex").(int)
		hasChange = true
	}
	if d.HasChange("ssocredentials") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredentials has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Ssocredentials = d.Get("ssocredentials").(string)
		hasChange = true
	}
	if d.HasChange("usercredentialindex") {
		log.Printf("[DEBUG]  citrixadc-provider: Usercredentialindex has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Usercredentialindex = d.Get("usercredentialindex").(int)
		hasChange = true
	}
	if d.HasChange("userexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Userexpression has changed for authenticationloginschema %s, starting update", authenticationloginschemaName)
		authenticationloginschema.Userexpression = d.Get("userexpression").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationloginschema.Type(), authenticationloginschemaName, &authenticationloginschema)
		if err != nil {
			return diag.Errorf("Error updating authenticationloginschema %s", authenticationloginschemaName)
		}
	}
	return readAuthenticationloginschemaFunc(ctx, d, meta)
}

func deleteAuthenticationloginschemaFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationloginschemaFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationloginschemaName := d.Id()
	err := client.DeleteResource(service.Authenticationloginschema.Type(), authenticationloginschemaName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
