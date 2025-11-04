package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationsamlaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationsamlactionFunc,
		ReadContext:   readAuthenticationsamlactionFunc,
		UpdateContext: updateAuthenticationsamlactionFunc,
		DeleteContext: deleteAuthenticationsamlactionFunc,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"statechecks": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferredbindtype": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"artifactresolutionserviceurl": {
				Type:     schema.TypeString,
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
			"attributeconsumingserviceindex": {
				Type:     schema.TypeInt,
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
			"authnctxclassref": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"customauthnctxclassref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digestmethod": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforceusername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forceauthn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnamefield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logoutbinding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logouturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadatarefreshinterval": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metadataurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relaystaterule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestedauthncontext": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlacsindex": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samlbinding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlidpcertname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlissuername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlredirecturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlrejectunsignedassertion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlsigningcertname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samltwofactor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samluserfield": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendthumbprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signaturealg": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skewtime": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storesamlresponse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationsamlactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Get("name").(string)
	authenticationsamlaction := authentication.Authenticationsamlaction{
		Artifactresolutionserviceurl: d.Get("artifactresolutionserviceurl").(string),
		Attribute1:                   d.Get("attribute1").(string),
		Attribute10:                  d.Get("attribute10").(string),
		Attribute11:                  d.Get("attribute11").(string),
		Attribute12:                  d.Get("attribute12").(string),
		Attribute13:                  d.Get("attribute13").(string),
		Attribute14:                  d.Get("attribute14").(string),
		Attribute15:                  d.Get("attribute15").(string),
		Attribute16:                  d.Get("attribute16").(string),
		Attribute2:                   d.Get("attribute2").(string),
		Attribute3:                   d.Get("attribute3").(string),
		Attribute4:                   d.Get("attribute4").(string),
		Attribute5:                   d.Get("attribute5").(string),
		Attribute6:                   d.Get("attribute6").(string),
		Attribute7:                   d.Get("attribute7").(string),
		Attribute8:                   d.Get("attribute8").(string),
		Attribute9:                   d.Get("attribute9").(string),
		Attributes:                   d.Get("attributes").(string),
		Audience:                     d.Get("audience").(string),
		Authnctxclassref:             toStringList(d.Get("authnctxclassref").([]interface{})),
		Preferredbindtype:            toStringList(d.Get("preferredbindtype").([]interface{})),
		Statechecks:                  d.Get("statechecks").(string),
		Customauthnctxclassref:       d.Get("customauthnctxclassref").(string),
		Defaultauthenticationgroup:   d.Get("defaultauthenticationgroup").(string),
		Digestmethod:                 d.Get("digestmethod").(string),
		Enforceusername:              d.Get("enforceusername").(string),
		Forceauthn:                   d.Get("forceauthn").(string),
		Groupnamefield:               d.Get("groupnamefield").(string),
		Logoutbinding:                d.Get("logoutbinding").(string),
		Logouturl:                    d.Get("logouturl").(string),
		Metadataurl:                  d.Get("metadataurl").(string),
		Name:                         d.Get("name").(string),
		Relaystaterule:               d.Get("relaystaterule").(string),
		Requestedauthncontext:        d.Get("requestedauthncontext").(string),
		Samlbinding:                  d.Get("samlbinding").(string),
		Samlidpcertname:              d.Get("samlidpcertname").(string),
		Samlissuername:               d.Get("samlissuername").(string),
		Samlredirecturl:              d.Get("samlredirecturl").(string),
		Samlrejectunsignedassertion:  d.Get("samlrejectunsignedassertion").(string),
		Samlsigningcertname:          d.Get("samlsigningcertname").(string),
		Samltwofactor:                d.Get("samltwofactor").(string),
		Samluserfield:                d.Get("samluserfield").(string),
		Sendthumbprint:               d.Get("sendthumbprint").(string),
		Signaturealg:                 d.Get("signaturealg").(string),
		Storesamlresponse:            d.Get("storesamlresponse").(string),
	}

	if raw := d.GetRawConfig().GetAttr("attributeconsumingserviceindex"); !raw.IsNull() {
		authenticationsamlaction.Attributeconsumingserviceindex = intPtr(d.Get("attributeconsumingserviceindex").(int))
	}
	if raw := d.GetRawConfig().GetAttr("metadatarefreshinterval"); !raw.IsNull() {
		authenticationsamlaction.Metadatarefreshinterval = intPtr(d.Get("metadatarefreshinterval").(int))
	}
	if raw := d.GetRawConfig().GetAttr("samlacsindex"); !raw.IsNull() {
		authenticationsamlaction.Samlacsindex = intPtr(d.Get("samlacsindex").(int))
	}
	if raw := d.GetRawConfig().GetAttr("skewtime"); !raw.IsNull() {
		authenticationsamlaction.Skewtime = intPtr(d.Get("skewtime").(int))
	}

	_, err := client.AddResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName, &authenticationsamlaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationsamlactionName)

	return readAuthenticationsamlactionFunc(ctx, d, meta)
}

func readAuthenticationsamlactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationsamlaction state %s", authenticationsamlactionName)
	data, err := client.FindResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationsamlaction state %s", authenticationsamlactionName)
		d.SetId("")
		return nil
	}
	d.Set("artifactresolutionserviceurl", data["artifactresolutionserviceurl"])
	d.Set("statechecks", data["statechecks"])
	d.Set("preferredbindtype", data["preferredbindtype"])
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
	setToInt("attributeconsumingserviceindex", d, data["attributeconsumingserviceindex"])
	d.Set("attributes", data["attributes"])
	d.Set("audience", data["audience"])
	d.Set("authnctxclassref", data["authnctxclassref"])
	d.Set("customauthnctxclassref", data["customauthnctxclassref"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("digestmethod", data["digestmethod"])
	d.Set("enforceusername", data["enforceusername"])
	d.Set("forceauthn", data["forceauthn"])
	d.Set("groupnamefield", data["groupnamefield"])
	d.Set("logoutbinding", data["logoutbinding"])
	d.Set("logouturl", data["logouturl"])
	setToInt("metadatarefreshinterval", d, data["metadatarefreshinterval"])
	d.Set("metadataurl", data["metadataurl"])
	d.Set("name", data["name"])
	d.Set("relaystaterule", data["relaystaterule"])
	d.Set("requestedauthncontext", data["requestedauthncontext"])
	setToInt("samlacsindex", d, data["samlacsindex"])
	d.Set("samlbinding", data["samlbinding"])
	d.Set("samlidpcertname", data["samlidpcertname"])
	d.Set("samlissuername", data["samlissuername"])
	d.Set("samlredirecturl", data["samlredirecturl"])
	d.Set("samlrejectunsignedassertion", data["samlrejectunsignedassertion"])
	d.Set("samlsigningcertname", data["samlsigningcertname"])
	d.Set("samltwofactor", data["samltwofactor"])
	d.Set("samluserfield", data["samluserfield"])
	d.Set("sendthumbprint", data["sendthumbprint"])
	d.Set("signaturealg", data["signaturealg"])
	setToInt("skewtime", d, data["skewtime"])
	d.Set("storesamlresponse", data["storesamlresponse"])

	return nil

}

func updateAuthenticationsamlactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Get("name").(string)

	authenticationsamlaction := authentication.Authenticationsamlaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("statechecks") {
		log.Printf("[DEBUG]  citrixadc-provider: Statechecks has changed for authenticationsamlaction, starting update")
		authenticationsamlaction.Statechecks = d.Get("statechecks").(string)
		hasChange = true
	}
	if d.HasChange("preferredbindtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Preferredbindtype has changed for authenticationsamlaction, starting update")
		authenticationsamlaction.Preferredbindtype = toStringList(d.Get("preferredbindtype").([]interface{}))
		hasChange = true
	}
	if d.HasChange("artifactresolutionserviceurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Artifactresolutionserviceurl has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Artifactresolutionserviceurl = d.Get("artifactresolutionserviceurl").(string)
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attributeconsumingserviceindex") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributeconsumingserviceindex has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attributeconsumingserviceindex = intPtr(d.Get("attributeconsumingserviceindex").(int))
		hasChange = true
	}
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("audience") {
		log.Printf("[DEBUG]  citrixadc-provider: Audience has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Audience = d.Get("audience").(string)
		hasChange = true
	}
	if d.HasChange("authnctxclassref") {
		log.Printf("[DEBUG]  citrixadc-provider: Authnctxclassref has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Authnctxclassref = toStringList(d.Get("authnctxclassref").([]interface{}))
		hasChange = true
	}
	if d.HasChange("customauthnctxclassref") {
		log.Printf("[DEBUG]  citrixadc-provider: Customauthnctxclassref has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Customauthnctxclassref = d.Get("customauthnctxclassref").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("digestmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Digestmethod has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Digestmethod = d.Get("digestmethod").(string)
		hasChange = true
	}
	if d.HasChange("enforceusername") {
		log.Printf("[DEBUG]  citrixadc-provider: Enforceusername has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Enforceusername = d.Get("enforceusername").(string)
		hasChange = true
	}
	if d.HasChange("forceauthn") {
		log.Printf("[DEBUG]  citrixadc-provider: Forceauthn has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Forceauthn = d.Get("forceauthn").(string)
		hasChange = true
	}
	if d.HasChange("groupnamefield") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnamefield has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Groupnamefield = d.Get("groupnamefield").(string)
		hasChange = true
	}
	if d.HasChange("logoutbinding") {
		log.Printf("[DEBUG]  citrixadc-provider: Logoutbinding has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Logoutbinding = d.Get("logoutbinding").(string)
		hasChange = true
	}
	if d.HasChange("logouturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Logouturl has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Logouturl = d.Get("logouturl").(string)
		hasChange = true
	}
	if d.HasChange("metadatarefreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Metadatarefreshinterval has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Metadatarefreshinterval = intPtr(d.Get("metadatarefreshinterval").(int))
		hasChange = true
	}
	if d.HasChange("metadataurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Metadataurl has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Metadataurl = d.Get("metadataurl").(string)
		hasChange = true
	}
	if d.HasChange("relaystaterule") {
		log.Printf("[DEBUG]  citrixadc-provider: Relaystaterule has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Relaystaterule = d.Get("relaystaterule").(string)
		hasChange = true
	}
	if d.HasChange("requestedauthncontext") {
		log.Printf("[DEBUG]  citrixadc-provider: Requestedauthncontext has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Requestedauthncontext = d.Get("requestedauthncontext").(string)
		hasChange = true
	}
	if d.HasChange("samlacsindex") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlacsindex has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlacsindex = intPtr(d.Get("samlacsindex").(int))
		hasChange = true
	}
	if d.HasChange("samlbinding") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlbinding has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlbinding = d.Get("samlbinding").(string)
		hasChange = true
	}
	if d.HasChange("samlidpcertname") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlidpcertname has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlidpcertname = d.Get("samlidpcertname").(string)
		hasChange = true
	}
	if d.HasChange("samlissuername") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlissuername has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlissuername = d.Get("samlissuername").(string)
		hasChange = true
	}
	if d.HasChange("samlredirecturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlredirecturl has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlredirecturl = d.Get("samlredirecturl").(string)
		hasChange = true
	}
	if d.HasChange("samlrejectunsignedassertion") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlrejectunsignedassertion has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlrejectunsignedassertion = d.Get("samlrejectunsignedassertion").(string)
		hasChange = true
	}
	if d.HasChange("samlsigningcertname") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlsigningcertname has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samlsigningcertname = d.Get("samlsigningcertname").(string)
		hasChange = true
	}
	if d.HasChange("samltwofactor") {
		log.Printf("[DEBUG]  citrixadc-provider: Samltwofactor has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samltwofactor = d.Get("samltwofactor").(string)
		hasChange = true
	}
	if d.HasChange("samluserfield") {
		log.Printf("[DEBUG]  citrixadc-provider: Samluserfield has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Samluserfield = d.Get("samluserfield").(string)
		hasChange = true
	}
	if d.HasChange("sendthumbprint") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendthumbprint has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Sendthumbprint = d.Get("sendthumbprint").(string)
		hasChange = true
	}
	if d.HasChange("signaturealg") {
		log.Printf("[DEBUG]  citrixadc-provider: Signaturealg has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Signaturealg = d.Get("signaturealg").(string)
		hasChange = true
	}
	if d.HasChange("skewtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Skewtime has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Skewtime = intPtr(d.Get("skewtime").(int))
		hasChange = true
	}
	if d.HasChange("storesamlresponse") {
		log.Printf("[DEBUG]  citrixadc-provider: Storesamlresponse has changed for authenticationsamlaction %s, starting update", authenticationsamlactionName)
		authenticationsamlaction.Storesamlresponse = d.Get("storesamlresponse").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName, &authenticationsamlaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationsamlaction %s", authenticationsamlactionName)
		}
	}
	return readAuthenticationsamlactionFunc(ctx, d, meta)
}

func deleteAuthenticationsamlactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
