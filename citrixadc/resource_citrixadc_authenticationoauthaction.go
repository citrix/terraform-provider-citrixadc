package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationoauthaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationoauthactionFunc,
		ReadContext:   readAuthenticationoauthactionFunc,
		UpdateContext: updateAuthenticationoauthactionFunc,
		DeleteContext: deleteAuthenticationoauthactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"requestattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oauthmiscflags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"intunedeviceidexpression": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"authorizationendpoint": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
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
			"tokenendpoint": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"allowedalgorithms": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"attribute1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certendpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certfilepath": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"granttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graphendpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idtokendecryptendpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"introspecturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadataurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oauthtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pkce": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"refreshinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resourceuri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skewtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"tenantid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tokenendpointauthmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"userinfourl": {
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

func createAuthenticationoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthactionName := d.Get("name").(string)
	authenticationoauthaction := authentication.Authenticationoauthaction{
		Allowedalgorithms:          toStringList(d.Get("allowedalgorithms").([]interface{})),
		Intunedeviceidexpression:   d.Get("intunedeviceidexpression").(string),
		Oauthmiscflags:             toStringList(d.Get("oauthmiscflags").([]interface{})),
		Requestattribute:           d.Get("requestattribute").(string),
		Attribute1:                 d.Get("attribute1").(string),
		Attribute10:                d.Get("attribute10").(string),
		Attribute11:                d.Get("attribute11").(string),
		Attribute12:                d.Get("attribute12").(string),
		Attribute13:                d.Get("attribute13").(string),
		Attribute14:                d.Get("attribute14").(string),
		Attribute15:                d.Get("attribute15").(string),
		Attribute16:                d.Get("attribute16").(string),
		Attribute2:                 d.Get("attribute2").(string),
		Attribute3:                 d.Get("attribute3").(string),
		Attribute4:                 d.Get("attribute4").(string),
		Attribute5:                 d.Get("attribute5").(string),
		Attribute6:                 d.Get("attribute6").(string),
		Attribute7:                 d.Get("attribute7").(string),
		Attribute8:                 d.Get("attribute8").(string),
		Attribute9:                 d.Get("attribute9").(string),
		Attributes:                 d.Get("attributes").(string),
		Audience:                   d.Get("audience").(string),
		Authentication:             d.Get("authentication").(string),
		Authorizationendpoint:      d.Get("authorizationendpoint").(string),
		Certendpoint:               d.Get("certendpoint").(string),
		Certfilepath:               d.Get("certfilepath").(string),
		Clientid:                   d.Get("clientid").(string),
		Clientsecret:               d.Get("clientsecret").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Granttype:                  d.Get("granttype").(string),
		Graphendpoint:              d.Get("graphendpoint").(string),
		Idtokendecryptendpoint:     d.Get("idtokendecryptendpoint").(string),
		Introspecturl:              d.Get("introspecturl").(string),
		Issuer:                     d.Get("issuer").(string),
		Metadataurl:                d.Get("metadataurl").(string),
		Name:                       d.Get("name").(string),
		Oauthtype:                  d.Get("oauthtype").(string),
		Pkce:                       d.Get("pkce").(string),
		Resourceuri:                d.Get("resourceuri").(string),
		Tenantid:                   d.Get("tenantid").(string),
		Tokenendpoint:              d.Get("tokenendpoint").(string),
		Tokenendpointauthmethod:    d.Get("tokenendpointauthmethod").(string),
		Userinfourl:                d.Get("userinfourl").(string),
		Usernamefield:              d.Get("usernamefield").(string),
	}

	if raw := d.GetRawConfig().GetAttr("refreshinterval"); !raw.IsNull() {
		authenticationoauthaction.Refreshinterval = intPtr(d.Get("refreshinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("skewtime"); !raw.IsNull() {
		authenticationoauthaction.Skewtime = intPtr(d.Get("skewtime").(int))
	}

	_, err := client.AddResource(service.Authenticationoauthaction.Type(), authenticationoauthactionName, &authenticationoauthaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationoauthactionName)

	return readAuthenticationoauthactionFunc(ctx, d, meta)
}

func readAuthenticationoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationoauthaction state %s", authenticationoauthactionName)
	data, err := client.FindResource(service.Authenticationoauthaction.Type(), authenticationoauthactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationoauthaction state %s", authenticationoauthactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("requestattribute", data["requestattribute"])
	d.Set("oauthmiscflags", data["oauthmiscflags"])
	d.Set("intunedeviceidexpression", data["intunedeviceidexpression"])
	d.Set("allowedalgorithms", data["allowedalgorithms"])
	d.Set("attribute1", data["attribute1"])
	d.Set("attribute10", data["attribute10"])
	d.Set("attribute11", data["attribute11"])
	d.Set("attribute12", data["attribute12"])
	d.Set("attribute13", data["attribute13"])
	d.Set("attribute14", data["attribute14"])
	d.Set("attribute15", data["attribute15"])
	d.Set("attribute16", data["attribute16"])
	d.Set("attribute2", data["attribute2"])
	d.Set("attribute3", data["attribute3"])
	d.Set("attribute4", data["attribute4"])
	d.Set("attribute5", data["attribute5"])
	d.Set("attribute6", data["attribute6"])
	d.Set("attribute7", data["attribute7"])
	d.Set("attribute8", data["attribute8"])
	d.Set("attribute9", data["attribute9"])
	d.Set("attributes", data["attributes"])
	d.Set("audience", data["audience"])
	d.Set("authentication", data["authentication"])
	d.Set("authorizationendpoint", data["authorizationendpoint"])
	d.Set("certendpoint", data["certendpoint"])
	d.Set("certfilepath", data["certfilepath"])
	d.Set("clientid", data["clientid"])
	//d.Set("clientsecret", data["clientsecret"]) It returns different value
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("granttype", data["granttype"])
	d.Set("graphendpoint", data["graphendpoint"])
	d.Set("idtokendecryptendpoint", data["idtokendecryptendpoint"])
	d.Set("introspecturl", data["introspecturl"])
	d.Set("issuer", data["issuer"])
	d.Set("metadataurl", data["metadataurl"])
	d.Set("name", data["name"])
	d.Set("oauthtype", data["oauthtype"])
	d.Set("pkce", data["pkce"])
	setToInt("refreshinterval", d, data["refreshinterval"])
	d.Set("resourceuri", data["resourceuri"])
	setToInt("skewtime", d, data["skewtime"])
	d.Set("tenantid", data["tenantid"])
	d.Set("tokenendpoint", data["tokenendpoint"])
	d.Set("tokenendpointauthmethod", data["tokenendpointauthmethod"])
	d.Set("userinfourl", data["userinfourl"])
	d.Set("usernamefield", data["usernamefield"])

	return nil

}

func updateAuthenticationoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthactionName := d.Get("name").(string)

	authenticationoauthaction := authentication.Authenticationoauthaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("requestattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Requestattribute has changed for authenticationoauthaction, starting update")
		authenticationoauthaction.Requestattribute = d.Get("requestattribute").(string)
		hasChange = true
	}
	if d.HasChange("oauthmiscflags") {
		log.Printf("[DEBUG]  citrixadc-provider: Oauthmiscflags has changed for authenticationoauthaction, starting update")
		authenticationoauthaction.Oauthmiscflags = toStringList(d.Get("oauthmiscflags").([]interface{}))
		hasChange = true
	}
	if d.HasChange("intunedeviceidexpression") {
		log.Printf("[DEBUG]  citrixadc-provider: Intunedeviceidexpression has changed for authenticationoauthaction, starting update")
		authenticationoauthaction.Intunedeviceidexpression = d.Get("intunedeviceidexpression").(string)
		hasChange = true
	}
	if d.HasChange("allowedalgorithms") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowedalgorithms has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Allowedalgorithms = toStringList(d.Get("allowedalgorithms").([]interface{}))
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("audience") {
		log.Printf("[DEBUG]  citrixadc-provider: Audience has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Audience = d.Get("audience").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authorizationendpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorizationendpoint has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Authorizationendpoint = d.Get("authorizationendpoint").(string)
		hasChange = true
	}
	if d.HasChange("certendpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Certendpoint has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Certendpoint = d.Get("certendpoint").(string)
		hasChange = true
	}
	if d.HasChange("certfilepath") {
		log.Printf("[DEBUG]  citrixadc-provider: Certfilepath has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Certfilepath = d.Get("certfilepath").(string)
		hasChange = true
	}
	if d.HasChange("clientid") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientid has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Clientid = d.Get("clientid").(string)
		hasChange = true
	}
	if d.HasChange("clientsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecret has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Clientsecret = d.Get("clientsecret").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("granttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Granttype has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Granttype = d.Get("granttype").(string)
		hasChange = true
	}
	if d.HasChange("graphendpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Graphendpoint has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Graphendpoint = d.Get("graphendpoint").(string)
		hasChange = true
	}
	if d.HasChange("idtokendecryptendpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Idtokendecryptendpoint has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Idtokendecryptendpoint = d.Get("idtokendecryptendpoint").(string)
		hasChange = true
	}
	if d.HasChange("introspecturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Introspecturl has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Introspecturl = d.Get("introspecturl").(string)
		hasChange = true
	}
	if d.HasChange("issuer") {
		log.Printf("[DEBUG]  citrixadc-provider: Issuer has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Issuer = d.Get("issuer").(string)
		hasChange = true
	}
	if d.HasChange("metadataurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Metadataurl has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Metadataurl = d.Get("metadataurl").(string)
		hasChange = true
	}
	if d.HasChange("oauthtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Oauthtype has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Oauthtype = d.Get("oauthtype").(string)
		hasChange = true
	}
	if d.HasChange("pkce") {
		log.Printf("[DEBUG]  citrixadc-provider: Pkce has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Pkce = d.Get("pkce").(string)
		hasChange = true
	}
	if d.HasChange("refreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Refreshinterval has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Refreshinterval = intPtr(d.Get("refreshinterval").(int))
		hasChange = true
	}
	if d.HasChange("resourceuri") {
		log.Printf("[DEBUG]  citrixadc-provider: Resourceuri has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Resourceuri = d.Get("resourceuri").(string)
		hasChange = true
	}
	if d.HasChange("skewtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Skewtime has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Skewtime = intPtr(d.Get("skewtime").(int))
		hasChange = true
	}
	if d.HasChange("tenantid") {
		log.Printf("[DEBUG]  citrixadc-provider: Tenantid has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Tenantid = d.Get("tenantid").(string)
		hasChange = true
	}
	if d.HasChange("tokenendpoint") {
		log.Printf("[DEBUG]  citrixadc-provider: Tokenendpoint has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Tokenendpoint = d.Get("tokenendpoint").(string)
		hasChange = true
	}
	if d.HasChange("tokenendpointauthmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Tokenendpointauthmethod has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Tokenendpointauthmethod = d.Get("tokenendpointauthmethod").(string)
		hasChange = true
	}
	if d.HasChange("userinfourl") {
		log.Printf("[DEBUG]  citrixadc-provider: Userinfourl has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Userinfourl = d.Get("userinfourl").(string)
		hasChange = true
	}
	if d.HasChange("usernamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Usernamefield has changed for authenticationoauthaction %s, starting update", authenticationoauthactionName)
		authenticationoauthaction.Usernamefield = d.Get("usernamefield").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationoauthaction.Type(), authenticationoauthactionName, &authenticationoauthaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationoauthaction %s", authenticationoauthactionName)
		}
	}
	return readAuthenticationoauthactionFunc(ctx, d, meta)
}

func deleteAuthenticationoauthactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationoauthactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationoauthactionName := d.Id()
	err := client.DeleteResource(service.Authenticationoauthaction.Type(), authenticationoauthactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
