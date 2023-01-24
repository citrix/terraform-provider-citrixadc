package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationsamlaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationsamlactionFunc,
		Read:          readAuthenticationsamlactionFunc,
		Update:        updateAuthenticationsamlactionFunc,
		Delete:        deleteAuthenticationsamlactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
				ForceNew: true,
			},
			"artifactresolutionserviceurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attributeconsumingserviceindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"audience": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authnctxclassref": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"customauthnctxclassref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digestmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforceusername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forceauthn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnamefield": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logoutbinding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logouturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadatarefreshinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"metadataurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relaystaterule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requestedauthncontext": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlacsindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"samlbinding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlidpcertname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlissuername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlredirecturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlrejectunsignedassertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlsigningcertname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samltwofactor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samluserfield": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendthumbprint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signaturealg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skewtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"storesamlresponse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationsamlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Get("name").(string)
	authenticationsamlaction := authentication.Authenticationsamlaction{
		Artifactresolutionserviceurl:   d.Get("artifactresolutionserviceurl").(string),
		Attribute1:                     d.Get("attribute1").(string),
		Attribute10:                    d.Get("attribute10").(string),
		Attribute11:                    d.Get("attribute11").(string),
		Attribute12:                    d.Get("attribute12").(string),
		Attribute13:                    d.Get("attribute13").(string),
		Attribute14:                    d.Get("attribute14").(string),
		Attribute15:                    d.Get("attribute15").(string),
		Attribute16:                    d.Get("attribute16").(string),
		Attribute2:                     d.Get("attribute2").(string),
		Attribute3:                     d.Get("attribute3").(string),
		Attribute4:                     d.Get("attribute4").(string),
		Attribute5:                     d.Get("attribute5").(string),
		Attribute6:                     d.Get("attribute6").(string),
		Attribute7:                     d.Get("attribute7").(string),
		Attribute8:                     d.Get("attribute8").(string),
		Attribute9:                     d.Get("attribute9").(string),
		Attributeconsumingserviceindex: d.Get("attributeconsumingserviceindex").(int),
		Attributes:                     d.Get("attributes").(string),
		Audience:                       d.Get("audience").(string),
		Authnctxclassref:               toStringList(d.Get("authnctxclassref").([]interface{})),
		Customauthnctxclassref:         d.Get("customauthnctxclassref").(string),
		Defaultauthenticationgroup:     d.Get("defaultauthenticationgroup").(string),
		Digestmethod:                   d.Get("digestmethod").(string),
		Enforceusername:                d.Get("enforceusername").(string),
		Forceauthn:                     d.Get("forceauthn").(string),
		Groupnamefield:                 d.Get("groupnamefield").(string),
		Logoutbinding:                  d.Get("logoutbinding").(string),
		Logouturl:                      d.Get("logouturl").(string),
		Metadatarefreshinterval:        d.Get("metadatarefreshinterval").(int),
		Metadataurl:                    d.Get("metadataurl").(string),
		Name:                           d.Get("name").(string),
		Relaystaterule:                 d.Get("relaystaterule").(string),
		Requestedauthncontext:          d.Get("requestedauthncontext").(string),
		Samlacsindex:                   d.Get("samlacsindex").(int),
		Samlbinding:                    d.Get("samlbinding").(string),
		Samlidpcertname:                d.Get("samlidpcertname").(string),
		Samlissuername:                 d.Get("samlissuername").(string),
		Samlredirecturl:                d.Get("samlredirecturl").(string),
		Samlrejectunsignedassertion:    d.Get("samlrejectunsignedassertion").(string),
		Samlsigningcertname:            d.Get("samlsigningcertname").(string),
		Samltwofactor:                  d.Get("samltwofactor").(string),
		Samluserfield:                  d.Get("samluserfield").(string),
		Sendthumbprint:                 d.Get("sendthumbprint").(string),
		Signaturealg:                   d.Get("signaturealg").(string),
		Skewtime:                       d.Get("skewtime").(int),
		Storesamlresponse:              d.Get("storesamlresponse").(string),
	}

	_, err := client.AddResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName, &authenticationsamlaction)
	if err != nil {
		return err
	}

	d.SetId(authenticationsamlactionName)

	err = readAuthenticationsamlactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationsamlaction but we can't read it ?? %s", authenticationsamlactionName)
		return nil
	}
	return nil
}

func readAuthenticationsamlactionFunc(d *schema.ResourceData, meta interface{}) error {
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
	d.Set("attributeconsumingserviceindex", data["attributeconsumingserviceindex"])
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
	d.Set("metadatarefreshinterval", data["metadatarefreshinterval"])
	d.Set("metadataurl", data["metadataurl"])
	d.Set("name", data["name"])
	d.Set("relaystaterule", data["relaystaterule"])
	d.Set("requestedauthncontext", data["requestedauthncontext"])
	d.Set("samlacsindex", data["samlacsindex"])
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
	d.Set("skewtime", data["skewtime"])
	d.Set("storesamlresponse", data["storesamlresponse"])

	return nil

}

func updateAuthenticationsamlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Get("name").(string)

	authenticationsamlaction := authentication.Authenticationsamlaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
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
		authenticationsamlaction.Attributeconsumingserviceindex = d.Get("attributeconsumingserviceindex").(int)
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
		authenticationsamlaction.Metadatarefreshinterval = d.Get("metadatarefreshinterval").(int)
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
		authenticationsamlaction.Samlacsindex = d.Get("samlacsindex").(int)
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
		authenticationsamlaction.Skewtime = d.Get("skewtime").(int)
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
			return fmt.Errorf("Error updating authenticationsamlaction %s", authenticationsamlactionName)
		}
	}
	return readAuthenticationsamlactionFunc(d, meta)
}

func deleteAuthenticationsamlactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlactionName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlaction.Type(), authenticationsamlactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
