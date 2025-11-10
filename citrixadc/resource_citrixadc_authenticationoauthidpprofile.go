package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationoauthidpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationoauthidpprofileFunc,
		ReadContext:   readAuthenticationoauthidpprofileFunc,
		UpdateContext: updateAuthenticationoauthidpprofileFunc,
		DeleteContext: deleteAuthenticationoauthidpprofileFunc,
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
			"clientid": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"clientsecret": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"redirecturl": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"audience": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypttoken": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refreshinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"relyingpartymetadataurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendpassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signaturealg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatureservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skewtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationoauthidpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationoauthidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidpprofileName := d.Get("name").(string)
	authenticationoauthidpprofile := authentication.Authenticationoauthidpprofile{
		Attributes:                 d.Get("attributes").(string),
		Audience:                   d.Get("audience").(string),
		Clientid:                   d.Get("clientid").(string),
		Clientsecret:               d.Get("clientsecret").(string),
		Configservice:              d.Get("configservice").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Encrypttoken:               d.Get("encrypttoken").(string),
		Issuer:                     d.Get("issuer").(string),
		Name:                       d.Get("name").(string),
		Redirecturl:                d.Get("redirecturl").(string),
		Relyingpartymetadataurl:    d.Get("relyingpartymetadataurl").(string),
		Sendpassword:               d.Get("sendpassword").(string),
		Signaturealg:               d.Get("signaturealg").(string),
		Signatureservice:           d.Get("signatureservice").(string),
	}

	if raw := d.GetRawConfig().GetAttr("refreshinterval"); !raw.IsNull() {
		authenticationoauthidpprofile.Refreshinterval = intPtr(d.Get("refreshinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("skewtime"); !raw.IsNull() {
		authenticationoauthidpprofile.Skewtime = intPtr(d.Get("skewtime").(int))
	}

	_, err := client.AddResource("authenticationoauthidpprofile", authenticationoauthidpprofileName, &authenticationoauthidpprofile)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationoauthidpprofileName)

	return readAuthenticationoauthidpprofileFunc(ctx, d, meta)
}

func readAuthenticationoauthidpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationoauthidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidpprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationoauthidpprofile state %s", authenticationoauthidpprofileName)
	data, err := client.FindResource("authenticationoauthidpprofile", authenticationoauthidpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationoauthidpprofile state %s", authenticationoauthidpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("attributes", data["attributes"])
	d.Set("audience", data["audience"])
	d.Set("clientid", data["clientid"])
	// d.Set("clientsecret", data["clientsecret"]) Every time it changes when recieved from ADC
	d.Set("configservice", data["configservice"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("encrypttoken", data["encrypttoken"])
	d.Set("issuer", data["issuer"])
	d.Set("name", data["name"])
	d.Set("redirecturl", data["redirecturl"])
	setToInt("refreshinterval", d, data["refreshinterval"])
	d.Set("relyingpartymetadataurl", data["relyingpartymetadataurl"])
	d.Set("sendpassword", data["sendpassword"])
	d.Set("signaturealg", data["signaturealg"])
	d.Set("signatureservice", data["signatureservice"])
	setToInt("skewtime", d, data["skewtime"])

	return nil

}

func updateAuthenticationoauthidpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationoauthidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidpprofileName := d.Get("name").(string)

	authenticationoauthidpprofile := authentication.Authenticationoauthidpprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("audience") {
		log.Printf("[DEBUG]  citrixadc-provider: Audience has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Audience = d.Get("audience").(string)
		hasChange = true
	}
	if d.HasChange("clientid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientid has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Clientid = d.Get("clientid").(string)
		hasChange = true
	}
	if d.HasChange("clientsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecret has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Clientsecret = d.Get("clientsecret").(string)
		hasChange = true
	}
	if d.HasChange("configservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Configservice has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Configservice = d.Get("configservice").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("encrypttoken") {
		log.Printf("[DEBUG]  citrixadc-provider: Encrypttoken has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Encrypttoken = d.Get("encrypttoken").(string)
		hasChange = true
	}
	if d.HasChange("issuer") {
		log.Printf("[DEBUG]  citrixadc-provider: Issuer has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Issuer = d.Get("issuer").(string)
		hasChange = true
	}
	if d.HasChange("redirecturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirecturl has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Redirecturl = d.Get("redirecturl").(string)
		hasChange = true
	}
	if d.HasChange("refreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Refreshinterval has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Refreshinterval = intPtr(d.Get("refreshinterval").(int))
		hasChange = true
	}
	if d.HasChange("relyingpartymetadataurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Relyingpartymetadataurl has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Relyingpartymetadataurl = d.Get("relyingpartymetadataurl").(string)
		hasChange = true
	}
	if d.HasChange("sendpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendpassword has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Sendpassword = d.Get("sendpassword").(string)
		hasChange = true
	}
	if d.HasChange("signaturealg") {
		log.Printf("[DEBUG]  citrixadc-provider: Signaturealg has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Signaturealg = d.Get("signaturealg").(string)
		hasChange = true
	}
	if d.HasChange("signatureservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureservice has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Signatureservice = d.Get("signatureservice").(string)
		hasChange = true
	}
	if d.HasChange("skewtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Skewtime has changed for authenticationoauthidpprofile %s, starting update", authenticationoauthidpprofileName)
		authenticationoauthidpprofile.Skewtime = intPtr(d.Get("skewtime").(int))
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource("authenticationoauthidpprofile", authenticationoauthidpprofileName, &authenticationoauthidpprofile)
		if err != nil {
			return diag.Errorf("Error updating authenticationoauthidpprofile %s", authenticationoauthidpprofileName)
		}
	}
	return readAuthenticationoauthidpprofileFunc(ctx, d, meta)
}

func deleteAuthenticationoauthidpprofileFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationoauthidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthidpprofileName := d.Id()
	err := client.DeleteResource("authenticationoauthidpprofile", authenticationoauthidpprofileName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
