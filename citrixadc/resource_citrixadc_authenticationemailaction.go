package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationemailaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationemailactionFunc,
		ReadContext:   readAuthenticationemailactionFunc,
		UpdateContext: updateAuthenticationemailactionFunc,
		DeleteContext: deleteAuthenticationemailactionFunc,
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
			"username": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"serverurl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"content": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailaddress": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationemailactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Get("name").(string)
	authenticationemailaction := authentication.Authenticationemailaction{
		Content:                    d.Get("content").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Emailaddress:               d.Get("emailaddress").(string),
		Name:                       d.Get("name").(string),
		Password:                   d.Get("password").(string),
		Serverurl:                  d.Get("serverurl").(string),
		Timeout:                    d.Get("timeout").(int),
		Type:                       d.Get("type").(string),
		Username:                   d.Get("username").(string),
	}

	_, err := client.AddResource("authenticationemailaction", authenticationemailactionName, &authenticationemailaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationemailactionName)

	return readAuthenticationemailactionFunc(ctx, d, meta)
}

func readAuthenticationemailactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationemailaction state %s", authenticationemailactionName)
	data, err := client.FindResource("authenticationemailaction", authenticationemailactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationemailaction state %s", authenticationemailactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("content", data["content"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("emailaddress", data["emailaddress"])
	d.Set("name", data["name"])
	//d.Set("password", data["password"]) encrypted value recieved from ADC
	d.Set("serverurl", data["serverurl"])
	//setToInt("timeout", d, data["timeout"]) not recieved from ADC
	d.Set("type", data["type"])
	d.Set("username", data["username"])

	return nil

}

func updateAuthenticationemailactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Get("name").(string)

	authenticationemailaction := authentication.Authenticationemailaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("content") {
		log.Printf("[DEBUG]  citrixadc-provider: Content has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Content = d.Get("content").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("emailaddress") {
		log.Printf("[DEBUG]  citrixadc-provider: Emailaddress has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Emailaddress = d.Get("emailaddress").(string)
		hasChange = true
	}
	if d.HasChange("password") {
		log.Printf("[DEBUG]  citrixadc-provider: Password has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Password = d.Get("password").(string)
		hasChange = true
	}
	if d.HasChange("serverurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverurl has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Serverurl = d.Get("serverurl").(string)
		hasChange = true
	}
	if d.HasChange("timeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Timeout has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Timeout = d.Get("timeout").(int)
		hasChange = true
	}
	if d.HasChange("type") {
		log.Printf("[DEBUG]  citrixadc-provider: Type has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Type = d.Get("type").(string)
		hasChange = true
	}
	if d.HasChange("username") {
		log.Printf("[DEBUG]  citrixadc-provider: Username has changed for authenticationemailaction %s, starting update", authenticationemailactionName)
		authenticationemailaction.Username = d.Get("username").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationemailaction", authenticationemailactionName, &authenticationemailaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationemailaction %s", authenticationemailactionName)
		}
	}
	return readAuthenticationemailactionFunc(ctx, d, meta)
}

func deleteAuthenticationemailactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationemailactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationemailactionName := d.Id()
	err := client.DeleteResource("authenticationemailaction", authenticationemailactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
