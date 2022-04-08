package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcAuthenticationsamlidpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createAuthenticationsamlidpprofileFunc,
		Read:          readAuthenticationsamlidpprofileFunc,
		Update:        updateAuthenticationsamlidpprofileFunc,
		Delete:        deleteAuthenticationsamlidpprofileFunc,
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
			"acsurlrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assertionconsumerserviceurl": &schema.Schema{
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
			"attribute10expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute10friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute11friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute12friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute13friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute14friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute15friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute16friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute1expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute1format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute1friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute2friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute3friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute4friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute5friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute6friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute7friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute8friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9expr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attribute9friendlyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"audience": &schema.Schema{
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
			"encryptassertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryptionalgorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keytransportalg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logoutbinding": &schema.Schema{
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
			"nameidexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nameidformat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rejectunsignedrequests": &schema.Schema{
				Type:     schema.TypeString,
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
			"samlsigningcertversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlspcertname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"samlspcertversion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendpassword": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serviceproviderid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signassertion": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signaturealg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signatureservice": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skewtime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"splogouturl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationsamlidpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationsamlidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidpprofileName := d.Get("name").(string)
	authenticationsamlidpprofile := authentication.Authenticationsamlidpprofile{
		Acsurlrule:                  d.Get("acsurlrule").(string),
		Assertionconsumerserviceurl: d.Get("assertionconsumerserviceurl").(string),
		Attribute1:                  d.Get("attribute1").(string),
		Attribute10:                 d.Get("attribute10").(string),
		Attribute10expr:             d.Get("attribute10expr").(string),
		Attribute10format:           d.Get("attribute10format").(string),
		Attribute10friendlyname:     d.Get("attribute10friendlyname").(string),
		Attribute11:                 d.Get("attribute11").(string),
		Attribute11expr:             d.Get("attribute11expr").(string),
		Attribute11format:           d.Get("attribute11format").(string),
		Attribute11friendlyname:     d.Get("attribute11friendlyname").(string),
		Attribute12:                 d.Get("attribute12").(string),
		Attribute12expr:             d.Get("attribute12expr").(string),
		Attribute12format:           d.Get("attribute12format").(string),
		Attribute12friendlyname:     d.Get("attribute12friendlyname").(string),
		Attribute13:                 d.Get("attribute13").(string),
		Attribute13expr:             d.Get("attribute13expr").(string),
		Attribute13format:           d.Get("attribute13format").(string),
		Attribute13friendlyname:     d.Get("attribute13friendlyname").(string),
		Attribute14:                 d.Get("attribute14").(string),
		Attribute14expr:             d.Get("attribute14expr").(string),
		Attribute14format:           d.Get("attribute14format").(string),
		Attribute14friendlyname:     d.Get("attribute14friendlyname").(string),
		Attribute15:                 d.Get("attribute15").(string),
		Attribute15expr:             d.Get("attribute15expr").(string),
		Attribute15format:           d.Get("attribute15format").(string),
		Attribute15friendlyname:     d.Get("attribute15friendlyname").(string),
		Attribute16:                 d.Get("attribute16").(string),
		Attribute16expr:             d.Get("attribute16expr").(string),
		Attribute16format:           d.Get("attribute16format").(string),
		Attribute16friendlyname:     d.Get("attribute16friendlyname").(string),
		Attribute1expr:              d.Get("attribute1expr").(string),
		Attribute1format:            d.Get("attribute1format").(string),
		Attribute1friendlyname:      d.Get("attribute1friendlyname").(string),
		Attribute2:                  d.Get("attribute2").(string),
		Attribute2expr:              d.Get("attribute2expr").(string),
		Attribute2format:            d.Get("attribute2format").(string),
		Attribute2friendlyname:      d.Get("attribute2friendlyname").(string),
		Attribute3:                  d.Get("attribute3").(string),
		Attribute3expr:              d.Get("attribute3expr").(string),
		Attribute3format:            d.Get("attribute3format").(string),
		Attribute3friendlyname:      d.Get("attribute3friendlyname").(string),
		Attribute4:                  d.Get("attribute4").(string),
		Attribute4expr:              d.Get("attribute4expr").(string),
		Attribute4format:            d.Get("attribute4format").(string),
		Attribute4friendlyname:      d.Get("attribute4friendlyname").(string),
		Attribute5:                  d.Get("attribute5").(string),
		Attribute5expr:              d.Get("attribute5expr").(string),
		Attribute5format:            d.Get("attribute5format").(string),
		Attribute5friendlyname:      d.Get("attribute5friendlyname").(string),
		Attribute6:                  d.Get("attribute6").(string),
		Attribute6expr:              d.Get("attribute6expr").(string),
		Attribute6format:            d.Get("attribute6format").(string),
		Attribute6friendlyname:      d.Get("attribute6friendlyname").(string),
		Attribute7:                  d.Get("attribute7").(string),
		Attribute7expr:              d.Get("attribute7expr").(string),
		Attribute7format:            d.Get("attribute7format").(string),
		Attribute7friendlyname:      d.Get("attribute7friendlyname").(string),
		Attribute8:                  d.Get("attribute8").(string),
		Attribute8expr:              d.Get("attribute8expr").(string),
		Attribute8format:            d.Get("attribute8format").(string),
		Attribute8friendlyname:      d.Get("attribute8friendlyname").(string),
		Attribute9:                  d.Get("attribute9").(string),
		Attribute9expr:              d.Get("attribute9expr").(string),
		Attribute9format:            d.Get("attribute9format").(string),
		Attribute9friendlyname:      d.Get("attribute9friendlyname").(string),
		Audience:                    d.Get("audience").(string),
		Defaultauthenticationgroup:  d.Get("defaultauthenticationgroup").(string),
		Digestmethod:                d.Get("digestmethod").(string),
		Encryptassertion:            d.Get("encryptassertion").(string),
		Encryptionalgorithm:         d.Get("encryptionalgorithm").(string),
		Keytransportalg:             d.Get("keytransportalg").(string),
		Logoutbinding:               d.Get("logoutbinding").(string),
		Metadatarefreshinterval:     d.Get("metadatarefreshinterval").(int),
		Metadataurl:                 d.Get("metadataurl").(string),
		Name:                        d.Get("name").(string),
		Nameidexpr:                  d.Get("nameidexpr").(string),
		Nameidformat:                d.Get("nameidformat").(string),
		Rejectunsignedrequests:      d.Get("rejectunsignedrequests").(string),
		Samlbinding:                 d.Get("samlbinding").(string),
		Samlidpcertname:             d.Get("samlidpcertname").(string),
		Samlissuername:              d.Get("samlissuername").(string),
		Samlsigningcertversion:      d.Get("samlsigningcertversion").(string),
		Samlspcertname:              d.Get("samlspcertname").(string),
		Samlspcertversion:           d.Get("samlspcertversion").(string),
		Sendpassword:                d.Get("sendpassword").(string),
		Serviceproviderid:           d.Get("serviceproviderid").(string),
		Signassertion:               d.Get("signassertion").(string),
		Signaturealg:                d.Get("signaturealg").(string),
		Signatureservice:            d.Get("signatureservice").(string),
		Skewtime:                    d.Get("skewtime").(int),
		Splogouturl:                 d.Get("splogouturl").(string),
	}

	_, err := client.AddResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName, &authenticationsamlidpprofile)
	if err != nil {
		return err
	}

	d.SetId(authenticationsamlidpprofileName)

	err = readAuthenticationsamlidpprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this authenticationsamlidpprofile but we can't read it ?? %s", authenticationsamlidpprofileName)
		return nil
	}
	return nil
}

func readAuthenticationsamlidpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationsamlidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidpprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationsamlidpprofile state %s", authenticationsamlidpprofileName)
	data, err := client.FindResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationsamlidpprofile state %s", authenticationsamlidpprofileName)
		d.SetId("")
		return nil
	}
	d.Set("acsurlrule", data["acsurlrule"])
	d.Set("assertionconsumerserviceurl", data["assertionconsumerserviceurl"])
	d.Set("attribute1", data["attribute1"])
	d.Set("attribute10", data["attribute10"])
	d.Set("attribute10expr", data["attribute10expr"])
	d.Set("attribute10format", data["attribute10format"])
	d.Set("attribute10friendlyname", data["attribute10friendlyname"])
	d.Set("attribute11", data["attribute11"])
	d.Set("attribute11expr", data["attribute11expr"])
	d.Set("attribute11format", data["attribute11format"])
	d.Set("attribute11friendlyname", data["attribute11friendlyname"])
	d.Set("attribute12", data["attribute12"])
	d.Set("attribute12expr", data["attribute12expr"])
	d.Set("attribute12format", data["attribute12format"])
	d.Set("attribute12friendlyname", data["attribute12friendlyname"])
	d.Set("attribute13", data["attribute13"])
	d.Set("attribute13expr", data["attribute13expr"])
	d.Set("attribute13format", data["attribute13format"])
	d.Set("attribute13friendlyname", data["attribute13friendlyname"])
	d.Set("attribute14", data["attribute14"])
	d.Set("attribute14expr", data["attribute14expr"])
	d.Set("attribute14format", data["attribute14format"])
	d.Set("attribute14friendlyname", data["attribute14friendlyname"])
	d.Set("attribute15", data["attribute15"])
	d.Set("attribute15expr", data["attribute15expr"])
	d.Set("attribute15format", data["attribute15format"])
	d.Set("attribute15friendlyname", data["attribute15friendlyname"])
	d.Set("attribute16", data["attribute16"])
	d.Set("attribute16expr", data["attribute16expr"])
	d.Set("attribute16format", data["attribute16format"])
	d.Set("attribute16friendlyname", data["attribute16friendlyname"])
	d.Set("attribute1expr", data["attribute1expr"])
	d.Set("attribute1format", data["attribute1format"])
	d.Set("attribute1friendlyname", data["attribute1friendlyname"])
	d.Set("attribute2", data["attribute2"])
	d.Set("attribute2expr", data["attribute2expr"])
	d.Set("attribute2format", data["attribute2format"])
	d.Set("attribute2friendlyname", data["attribute2friendlyname"])
	d.Set("attribute3", data["attribute3"])
	d.Set("attribute3expr", data["attribute3expr"])
	d.Set("attribute3format", data["attribute3format"])
	d.Set("attribute3friendlyname", data["attribute3friendlyname"])
	d.Set("attribute4", data["attribute4"])
	d.Set("attribute4expr", data["attribute4expr"])
	d.Set("attribute4format", data["attribute4format"])
	d.Set("attribute4friendlyname", data["attribute4friendlyname"])
	d.Set("attribute5", data["attribute5"])
	d.Set("attribute5expr", data["attribute5expr"])
	d.Set("attribute5format", data["attribute5format"])
	d.Set("attribute5friendlyname", data["attribute5friendlyname"])
	d.Set("attribute6", data["attribute6"])
	d.Set("attribute6expr", data["attribute6expr"])
	d.Set("attribute6format", data["attribute6format"])
	d.Set("attribute6friendlyname", data["attribute6friendlyname"])
	d.Set("attribute7", data["attribute7"])
	d.Set("attribute7expr", data["attribute7expr"])
	d.Set("attribute7format", data["attribute7format"])
	d.Set("attribute7friendlyname", data["attribute7friendlyname"])
	d.Set("attribute8", data["attribute8"])
	d.Set("attribute8expr", data["attribute8expr"])
	d.Set("attribute8format", data["attribute8format"])
	d.Set("attribute8friendlyname", data["attribute8friendlyname"])
	d.Set("attribute9", data["attribute9"])
	d.Set("attribute9expr", data["attribute9expr"])
	d.Set("attribute9format", data["attribute9format"])
	d.Set("attribute9friendlyname", data["attribute9friendlyname"])
	d.Set("audience", data["audience"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("digestmethod", data["digestmethod"])
	d.Set("encryptassertion", data["encryptassertion"])
	d.Set("encryptionalgorithm", data["encryptionalgorithm"])
	d.Set("keytransportalg", data["keytransportalg"])
	d.Set("logoutbinding", data["logoutbinding"])
	d.Set("metadatarefreshinterval", data["metadatarefreshinterval"])
	d.Set("metadataurl", data["metadataurl"])
	d.Set("name", data["name"])
	d.Set("nameidexpr", data["nameidexpr"])
	d.Set("nameidformat", data["nameidformat"])
	d.Set("rejectunsignedrequests", data["rejectunsignedrequests"])
	d.Set("samlbinding", data["samlbinding"])
	d.Set("samlidpcertname", data["samlidpcertname"])
	d.Set("samlissuername", data["samlissuername"])
	d.Set("samlsigningcertversion", data["samlsigningcertversion"])
	d.Set("samlspcertname", data["samlspcertname"])
	d.Set("samlspcertversion", data["samlspcertversion"])
	// d.Set("sendpassword", data["sendpassword"])
	d.Set("serviceproviderid", data["serviceproviderid"])
	d.Set("signassertion", data["signassertion"])
	d.Set("signaturealg", data["signaturealg"])
	d.Set("signatureservice", data["signatureservice"])
	d.Set("skewtime", data["skewtime"])
	d.Set("splogouturl", data["splogouturl"])

	return nil

}

func updateAuthenticationsamlidpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationsamlidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidpprofileName := d.Get("name").(string)

	authenticationsamlidpprofile := authentication.Authenticationsamlidpprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("acsurlrule") {
		log.Printf("[DEBUG]  citrixadc-provider: Acsurlrule has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Acsurlrule = d.Get("acsurlrule").(string)
		hasChange = true
	}
	if d.HasChange("assertionconsumerserviceurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Assertionconsumerserviceurl has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Assertionconsumerserviceurl = d.Get("assertionconsumerserviceurl").(string)
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute10expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute10expr = d.Get("attribute10expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute10format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute10format = d.Get("attribute10format").(string)
		hasChange = true
	}
	if d.HasChange("attribute10friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute10friendlyname = d.Get("attribute10friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute11expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute11expr = d.Get("attribute11expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute11format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute11format = d.Get("attribute11format").(string)
		hasChange = true
	}
	if d.HasChange("attribute11friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute11friendlyname = d.Get("attribute11friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute12expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute12expr = d.Get("attribute12expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute12format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute12format = d.Get("attribute12format").(string)
		hasChange = true
	}
	if d.HasChange("attribute12friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute12friendlyname = d.Get("attribute12friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute13expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute13expr = d.Get("attribute13expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute13format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute13format = d.Get("attribute13format").(string)
		hasChange = true
	}
	if d.HasChange("attribute13friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute13friendlyname = d.Get("attribute13friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute14expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute14expr = d.Get("attribute14expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute14format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute14format = d.Get("attribute14format").(string)
		hasChange = true
	}
	if d.HasChange("attribute14friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute14friendlyname = d.Get("attribute14friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute15expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute15expr = d.Get("attribute15expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute15format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute15format = d.Get("attribute15format").(string)
		hasChange = true
	}
	if d.HasChange("attribute15friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute15friendlyname = d.Get("attribute15friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute16expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute16expr = d.Get("attribute16expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute16format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute16format = d.Get("attribute16format").(string)
		hasChange = true
	}
	if d.HasChange("attribute16friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute16friendlyname = d.Get("attribute16friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute1expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute1expr = d.Get("attribute1expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute1format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute1format = d.Get("attribute1format").(string)
		hasChange = true
	}
	if d.HasChange("attribute1friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute1friendlyname = d.Get("attribute1friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute2expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute2expr = d.Get("attribute2expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute2format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute2format = d.Get("attribute2format").(string)
		hasChange = true
	}
	if d.HasChange("attribute2friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute2friendlyname = d.Get("attribute2friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute3expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute3expr = d.Get("attribute3expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute3format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute3format = d.Get("attribute3format").(string)
		hasChange = true
	}
	if d.HasChange("attribute3friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute3friendlyname = d.Get("attribute3friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute4expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute4expr = d.Get("attribute4expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute4format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute4format = d.Get("attribute4format").(string)
		hasChange = true
	}
	if d.HasChange("attribute4friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute4friendlyname = d.Get("attribute4friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute5expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute5expr = d.Get("attribute5expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute5format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute5format = d.Get("attribute5format").(string)
		hasChange = true
	}
	if d.HasChange("attribute5friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute5friendlyname = d.Get("attribute5friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute6expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute6expr = d.Get("attribute6expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute6format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute6format = d.Get("attribute6format").(string)
		hasChange = true
	}
	if d.HasChange("attribute6friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute6friendlyname = d.Get("attribute6friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute7expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute7expr = d.Get("attribute7expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute7format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute7format = d.Get("attribute7format").(string)
		hasChange = true
	}
	if d.HasChange("attribute7friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute7friendlyname = d.Get("attribute7friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute8expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute8expr = d.Get("attribute8expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute8format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute8format = d.Get("attribute8format").(string)
		hasChange = true
	}
	if d.HasChange("attribute8friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute8friendlyname = d.Get("attribute8friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attribute9expr") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9expr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute9expr = d.Get("attribute9expr").(string)
		hasChange = true
	}
	if d.HasChange("attribute9format") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9format has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute9format = d.Get("attribute9format").(string)
		hasChange = true
	}
	if d.HasChange("attribute9friendlyname") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9friendlyname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Attribute9friendlyname = d.Get("attribute9friendlyname").(string)
		hasChange = true
	}
	if d.HasChange("audience") {
		log.Printf("[DEBUG]  citrixadc-provider: Audience has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Audience = d.Get("audience").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("digestmethod") {
		log.Printf("[DEBUG]  citrixadc-provider: Digestmethod has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Digestmethod = d.Get("digestmethod").(string)
		hasChange = true
	}
	if d.HasChange("encryptassertion") {
		log.Printf("[DEBUG]  citrixadc-provider: Encryptassertion has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Encryptassertion = d.Get("encryptassertion").(string)
		hasChange = true
	}
	if d.HasChange("encryptionalgorithm") {
		log.Printf("[DEBUG]  citrixadc-provider: Encryptionalgorithm has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Encryptionalgorithm = d.Get("encryptionalgorithm").(string)
		hasChange = true
	}
	if d.HasChange("keytransportalg") {
		log.Printf("[DEBUG]  citrixadc-provider: Keytransportalg has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Keytransportalg = d.Get("keytransportalg").(string)
		hasChange = true
	}
	if d.HasChange("logoutbinding") {
		log.Printf("[DEBUG]  citrixadc-provider: Logoutbinding has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Logoutbinding = d.Get("logoutbinding").(string)
		hasChange = true
	}
	if d.HasChange("metadatarefreshinterval") {
		log.Printf("[DEBUG]  citrixadc-provider: Metadatarefreshinterval has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Metadatarefreshinterval = d.Get("metadatarefreshinterval").(int)
		hasChange = true
	}
	if d.HasChange("metadataurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Metadataurl has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Metadataurl = d.Get("metadataurl").(string)
		hasChange = true
	}
	if d.HasChange("nameidexpr") {
		log.Printf("[DEBUG]  citrixadc-provider: Nameidexpr has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Nameidexpr = d.Get("nameidexpr").(string)
		hasChange = true
	}
	if d.HasChange("nameidformat") {
		log.Printf("[DEBUG]  citrixadc-provider: Nameidformat has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Nameidformat = d.Get("nameidformat").(string)
		hasChange = true
	}
	if d.HasChange("rejectunsignedrequests") {
		log.Printf("[DEBUG]  citrixadc-provider: Rejectunsignedrequests has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Rejectunsignedrequests = d.Get("rejectunsignedrequests").(string)
		hasChange = true
	}
	if d.HasChange("samlbinding") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlbinding has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlbinding = d.Get("samlbinding").(string)
		hasChange = true
	}
	if d.HasChange("samlidpcertname") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlidpcertname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlidpcertname = d.Get("samlidpcertname").(string)
		hasChange = true
	}
	if d.HasChange("samlissuername") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlissuername has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlissuername = d.Get("samlissuername").(string)
		hasChange = true
	}
	if d.HasChange("samlsigningcertversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlsigningcertversion has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlsigningcertversion = d.Get("samlsigningcertversion").(string)
		hasChange = true
	}
	if d.HasChange("samlspcertname") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlspcertname has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlspcertname = d.Get("samlspcertname").(string)
		hasChange = true
	}
	if d.HasChange("samlspcertversion") {
		log.Printf("[DEBUG]  citrixadc-provider: Samlspcertversion has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Samlspcertversion = d.Get("samlspcertversion").(string)
		hasChange = true
	}
	if d.HasChange("sendpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendpassword has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Sendpassword = d.Get("sendpassword").(string)
		hasChange = true
	}
	if d.HasChange("serviceproviderid") {
		log.Printf("[DEBUG]  citrixadc-provider: Serviceproviderid has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Serviceproviderid = d.Get("serviceproviderid").(string)
		hasChange = true
	}
	if d.HasChange("signassertion") {
		log.Printf("[DEBUG]  citrixadc-provider: Signassertion has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Signassertion = d.Get("signassertion").(string)
		hasChange = true
	}
	if d.HasChange("signaturealg") {
		log.Printf("[DEBUG]  citrixadc-provider: Signaturealg has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Signaturealg = d.Get("signaturealg").(string)
		hasChange = true
	}
	if d.HasChange("signatureservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Signatureservice has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Signatureservice = d.Get("signatureservice").(string)
		hasChange = true
	}
	if d.HasChange("skewtime") {
		log.Printf("[DEBUG]  citrixadc-provider: Skewtime has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Skewtime = d.Get("skewtime").(int)
		hasChange = true
	}
	if d.HasChange("splogouturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Splogouturl has changed for authenticationsamlidpprofile %s, starting update", authenticationsamlidpprofileName)
		authenticationsamlidpprofile.Splogouturl = d.Get("splogouturl").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName, &authenticationsamlidpprofile)
		if err != nil {
			return fmt.Errorf("Error updating authenticationsamlidpprofile %s", authenticationsamlidpprofileName)
		}
	}
	return readAuthenticationsamlidpprofileFunc(d, meta)
}

func deleteAuthenticationsamlidpprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationsamlidpprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationsamlidpprofileName := d.Id()
	err := client.DeleteResource(service.Authenticationsamlidpprofile.Type(), authenticationsamlidpprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
