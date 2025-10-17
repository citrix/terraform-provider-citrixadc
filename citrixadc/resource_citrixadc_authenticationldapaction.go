package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"
	"github.com/citrix/adc-nitro-go/service"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcAuthenticationldapaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createAuthenticationldapactionFunc,
		ReadContext:   readAuthenticationldapactionFunc,
		UpdateContext: updateAuthenticationldapactionFunc,
		DeleteContext: deleteAuthenticationldapactionFunc,
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
			"alternateemailattr": {
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
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cloudattributes": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthenticationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"followreferrals": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupattrname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupnameidentifier": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchfilter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"groupsearchsubattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kbattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbase": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldapbinddnpassword": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldaphostname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldaploginname": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxldapreferrals": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maxnestinglevel": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"mssrvrecordlocation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nestedgroupextraction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"otpsecret": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwdchange": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushservice": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"referraldnslookup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"requireuser": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"searchfilter": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sectype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverport": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sshpublickey": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssonameattribute": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subattributename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"svrtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validateservercert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createAuthenticationldapactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Get("name").(string)
	authenticationldapaction := authentication.Authenticationldapaction{
		Alternateemailattr:         d.Get("alternateemailattr").(string),
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
		Authentication:             d.Get("authentication").(string),
		Cloudattributes:            d.Get("cloudattributes").(string),
		Defaultauthenticationgroup: d.Get("defaultauthenticationgroup").(string),
		Email:                      d.Get("email").(string),
		Followreferrals:            d.Get("followreferrals").(string),
		Groupattrname:              d.Get("groupattrname").(string),
		Groupnameidentifier:        d.Get("groupnameidentifier").(string),
		Groupsearchattribute:       d.Get("groupsearchattribute").(string),
		Groupsearchfilter:          d.Get("groupsearchfilter").(string),
		Groupsearchsubattribute:    d.Get("groupsearchsubattribute").(string),
		Kbattribute:                d.Get("kbattribute").(string),
		Ldapbase:                   d.Get("ldapbase").(string),
		Ldapbinddn:                 d.Get("ldapbinddn").(string),
		Ldapbinddnpassword:         d.Get("ldapbinddnpassword").(string),
		Ldaphostname:               d.Get("ldaphostname").(string),
		Ldaploginname:              d.Get("ldaploginname").(string),
		Mssrvrecordlocation:        d.Get("mssrvrecordlocation").(string),
		Name:                       d.Get("name").(string),
		Nestedgroupextraction:      d.Get("nestedgroupextraction").(string),
		Otpsecret:                  d.Get("otpsecret").(string),
		Passwdchange:               d.Get("passwdchange").(string),
		Pushservice:                d.Get("pushservice").(string),
		Referraldnslookup:          d.Get("referraldnslookup").(string),
		Requireuser:                d.Get("requireuser").(string),
		Searchfilter:               d.Get("searchfilter").(string),
		Sectype:                    d.Get("sectype").(string),
		Serverip:                   d.Get("serverip").(string),
		Servername:                 d.Get("servername").(string),
		Sshpublickey:               d.Get("sshpublickey").(string),
		Ssonameattribute:           d.Get("ssonameattribute").(string),
		Subattributename:           d.Get("subattributename").(string),
		Svrtype:                    d.Get("svrtype").(string),
		Validateservercert:         d.Get("validateservercert").(string),
	}

	if raw := d.GetRawConfig().GetAttr("authtimeout"); !raw.IsNull() {
		authenticationldapaction.Authtimeout = intPtr(d.Get("authtimeout").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxldapreferrals"); !raw.IsNull() {
		authenticationldapaction.Maxldapreferrals = intPtr(d.Get("maxldapreferrals").(int))
	}
	if raw := d.GetRawConfig().GetAttr("maxnestinglevel"); !raw.IsNull() {
		authenticationldapaction.Maxnestinglevel = intPtr(d.Get("maxnestinglevel").(int))
	}
	if raw := d.GetRawConfig().GetAttr("serverport"); !raw.IsNull() {
		authenticationldapaction.Serverport = intPtr(d.Get("serverport").(int))
	}

	_, err := client.AddResource(service.Authenticationldapaction.Type(), authenticationldapactionName, &authenticationldapaction)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(authenticationldapactionName)

	return readAuthenticationldapactionFunc(ctx, d, meta)
}

func readAuthenticationldapactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading authenticationldapaction state %s", authenticationldapactionName)
	data, err := client.FindResource(service.Authenticationldapaction.Type(), authenticationldapactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing authenticationldapaction state %s", authenticationldapactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("alternateemailattr", data["alternateemailattr"])
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
	d.Set("authentication", data["authentication"])
	setToInt("authtimeout", d, data["authtimeout"])
	d.Set("cloudattributes", data["cloudattributes"])
	d.Set("defaultauthenticationgroup", data["defaultauthenticationgroup"])
	d.Set("email", data["email"])
	d.Set("followreferrals", data["followreferrals"])
	d.Set("groupattrname", data["groupattrname"])
	d.Set("groupnameidentifier", data["groupnameidentifier"])
	d.Set("groupsearchattribute", data["groupsearchattribute"])
	d.Set("groupsearchfilter", data["groupsearchfilter"])
	d.Set("groupsearchsubattribute", data["groupsearchsubattribute"])
	d.Set("kbattribute", data["kbattribute"])
	d.Set("ldapbase", data["ldapbase"])
	d.Set("ldapbinddn", data["ldapbinddn"])
	// d.Set("ldapbinddnpassword", data["ldapbinddnpassword"]) // We get the hash value from the NetScaler, which creates terraform to update the resource attribute on our next terraform apply command
	d.Set("ldaphostname", data["ldaphostname"])
	d.Set("ldaploginname", data["ldaploginname"])
	setToInt("maxldapreferrals", d, data["maxldapreferrals"])
	setToInt("maxnestinglevel", d, data["maxnestinglevel"])
	d.Set("mssrvrecordlocation", data["mssrvrecordlocation"])
	d.Set("name", data["name"])
	d.Set("nestedgroupextraction", data["nestedgroupextraction"])
	d.Set("otpsecret", data["otpsecret"])
	d.Set("passwdchange", data["passwdchange"])
	d.Set("pushservice", data["pushservice"])
	d.Set("referraldnslookup", data["referraldnslookup"])
	d.Set("requireuser", data["requireuser"])
	d.Set("searchfilter", data["searchfilter"])
	d.Set("sectype", data["sectype"])
	d.Set("serverip", data["serverip"])
	d.Set("servername", data["servername"])
	setToInt("serverport", d, data["serverport"])
	d.Set("sshpublickey", data["sshpublickey"])
	d.Set("ssonameattribute", data["ssonameattribute"])
	d.Set("subattributename", data["subattributename"])
	d.Set("svrtype", data["svrtype"])
	d.Set("validateservercert", data["validateservercert"])

	return nil

}

func updateAuthenticationldapactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Get("name").(string)

	authenticationldapaction := authentication.Authenticationldapaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("alternateemailattr") {
		log.Printf("[DEBUG]  citrixadc-provider: Alternateemailattr has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Alternateemailattr = d.Get("alternateemailattr").(string)
		hasChange = true
	}
	if d.HasChange("attribute1") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute1 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute1 = d.Get("attribute1").(string)
		hasChange = true
	}
	if d.HasChange("attribute10") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute10 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute10 = d.Get("attribute10").(string)
		hasChange = true
	}
	if d.HasChange("attribute11") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute11 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute11 = d.Get("attribute11").(string)
		hasChange = true
	}
	if d.HasChange("attribute12") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute12 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute12 = d.Get("attribute12").(string)
		hasChange = true
	}
	if d.HasChange("attribute13") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute13 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute13 = d.Get("attribute13").(string)
		hasChange = true
	}
	if d.HasChange("attribute14") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute14 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute14 = d.Get("attribute14").(string)
		hasChange = true
	}
	if d.HasChange("attribute15") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute15 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute15 = d.Get("attribute15").(string)
		hasChange = true
	}
	if d.HasChange("attribute16") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute16 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute16 = d.Get("attribute16").(string)
		hasChange = true
	}
	if d.HasChange("attribute2") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute2 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute2 = d.Get("attribute2").(string)
		hasChange = true
	}
	if d.HasChange("attribute3") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute3 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute3 = d.Get("attribute3").(string)
		hasChange = true
	}
	if d.HasChange("attribute4") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute4 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute4 = d.Get("attribute4").(string)
		hasChange = true
	}
	if d.HasChange("attribute5") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute5 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute5 = d.Get("attribute5").(string)
		hasChange = true
	}
	if d.HasChange("attribute6") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute6 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute6 = d.Get("attribute6").(string)
		hasChange = true
	}
	if d.HasChange("attribute7") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute7 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute7 = d.Get("attribute7").(string)
		hasChange = true
	}
	if d.HasChange("attribute8") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute8 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute8 = d.Get("attribute8").(string)
		hasChange = true
	}
	if d.HasChange("attribute9") {
		log.Printf("[DEBUG]  citrixadc-provider: Attribute9 has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attribute9 = d.Get("attribute9").(string)
		hasChange = true
	}
	if d.HasChange("attributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Attributes has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Attributes = d.Get("attributes").(string)
		hasChange = true
	}
	if d.HasChange("authentication") {
		log.Printf("[DEBUG]  citrixadc-provider: Authentication has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Authentication = d.Get("authentication").(string)
		hasChange = true
	}
	if d.HasChange("authtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Authtimeout has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Authtimeout = intPtr(d.Get("authtimeout").(int))
		hasChange = true
	}
	if d.HasChange("cloudattributes") {
		log.Printf("[DEBUG]  citrixadc-provider: Cloudattributes has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Cloudattributes = d.Get("cloudattributes").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthenticationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthenticationgroup has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Defaultauthenticationgroup = d.Get("defaultauthenticationgroup").(string)
		hasChange = true
	}
	if d.HasChange("email") {
		log.Printf("[DEBUG]  citrixadc-provider: Email has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Email = d.Get("email").(string)
		hasChange = true
	}
	if d.HasChange("followreferrals") {
		log.Printf("[DEBUG]  citrixadc-provider: Followreferrals has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Followreferrals = d.Get("followreferrals").(string)
		hasChange = true
	}
	if d.HasChange("groupattrname") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupattrname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupattrname = d.Get("groupattrname").(string)
		hasChange = true
	}
	if d.HasChange("groupnameidentifier") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupnameidentifier has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupnameidentifier = d.Get("groupnameidentifier").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchattribute = d.Get("groupsearchattribute").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchfilter has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchfilter = d.Get("groupsearchfilter").(string)
		hasChange = true
	}
	if d.HasChange("groupsearchsubattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Groupsearchsubattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Groupsearchsubattribute = d.Get("groupsearchsubattribute").(string)
		hasChange = true
	}
	if d.HasChange("kbattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Kbattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Kbattribute = d.Get("kbattribute").(string)
		hasChange = true
	}
	if d.HasChange("ldapbase") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbase has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbase = d.Get("ldapbase").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddn") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddn has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbinddn = d.Get("ldapbinddn").(string)
		hasChange = true
	}
	if d.HasChange("ldapbinddnpassword") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldapbinddnpassword has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldapbinddnpassword = d.Get("ldapbinddnpassword").(string)
		hasChange = true
	}
	if d.HasChange("ldaphostname") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldaphostname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldaphostname = d.Get("ldaphostname").(string)
		hasChange = true
	}
	if d.HasChange("ldaploginname") {
		log.Printf("[DEBUG]  citrixadc-provider: Ldaploginname has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ldaploginname = d.Get("ldaploginname").(string)
		hasChange = true
	}
	if d.HasChange("maxldapreferrals") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxldapreferrals has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Maxldapreferrals = intPtr(d.Get("maxldapreferrals").(int))
		hasChange = true
	}
	if d.HasChange("maxnestinglevel") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxnestinglevel has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Maxnestinglevel = intPtr(d.Get("maxnestinglevel").(int))
		hasChange = true
	}
	if d.HasChange("mssrvrecordlocation") {
		log.Printf("[DEBUG]  citrixadc-provider: Mssrvrecordlocation has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Mssrvrecordlocation = d.Get("mssrvrecordlocation").(string)
		hasChange = true
	}
	if d.HasChange("nestedgroupextraction") {
		log.Printf("[DEBUG]  citrixadc-provider: Nestedgroupextraction has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Nestedgroupextraction = d.Get("nestedgroupextraction").(string)
		hasChange = true
	}
	if d.HasChange("otpsecret") {
		log.Printf("[DEBUG]  citrixadc-provider: Otpsecret has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Otpsecret = d.Get("otpsecret").(string)
		hasChange = true
	}
	if d.HasChange("passwdchange") {
		log.Printf("[DEBUG]  citrixadc-provider: Passwdchange has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Passwdchange = d.Get("passwdchange").(string)
		hasChange = true
	}
	if d.HasChange("pushservice") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushservice has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Pushservice = d.Get("pushservice").(string)
		hasChange = true
	}
	if d.HasChange("referraldnslookup") {
		log.Printf("[DEBUG]  citrixadc-provider: Referraldnslookup has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Referraldnslookup = d.Get("referraldnslookup").(string)
		hasChange = true
	}
	if d.HasChange("requireuser") {
		log.Printf("[DEBUG]  citrixadc-provider: Requireuser has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Requireuser = d.Get("requireuser").(string)
		hasChange = true
	}
	if d.HasChange("searchfilter") {
		log.Printf("[DEBUG]  citrixadc-provider: Searchfilter has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Searchfilter = d.Get("searchfilter").(string)
		hasChange = true
	}
	if d.HasChange("sectype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sectype has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Sectype = d.Get("sectype").(string)
		hasChange = true
	}
	if d.HasChange("serverip") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverip has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Serverip = d.Get("serverip").(string)
		hasChange = true
	}
	if d.HasChange("servername") {
		log.Printf("[DEBUG]  citrixadc-provider: Servername has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Servername = d.Get("servername").(string)
		hasChange = true
	}
	if d.HasChange("serverport") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverport has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Serverport = intPtr(d.Get("serverport").(int))
		hasChange = true
	}
	if d.HasChange("sshpublickey") {
		log.Printf("[DEBUG]  citrixadc-provider: Sshpublickey has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Sshpublickey = d.Get("sshpublickey").(string)
		hasChange = true
	}
	if d.HasChange("ssonameattribute") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssonameattribute has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Ssonameattribute = d.Get("ssonameattribute").(string)
		hasChange = true
	}
	if d.HasChange("subattributename") {
		log.Printf("[DEBUG]  citrixadc-provider: Subattributename has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Subattributename = d.Get("subattributename").(string)
		hasChange = true
	}
	if d.HasChange("svrtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Svrtype has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Svrtype = d.Get("svrtype").(string)
		hasChange = true
	}
	if d.HasChange("validateservercert") {
		log.Printf("[DEBUG]  citrixadc-provider: Validateservercert has changed for authenticationldapaction %s, starting update", authenticationldapactionName)
		authenticationldapaction.Validateservercert = d.Get("validateservercert").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Authenticationldapaction.Type(), authenticationldapactionName, &authenticationldapaction)
		if err != nil {
			return diag.Errorf("Error updating authenticationldapaction %s", authenticationldapactionName)
		}
	}
	return readAuthenticationldapactionFunc(ctx, d, meta)
}

func deleteAuthenticationldapactionFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteAuthenticationldapactionFunc")
	client := meta.(*NetScalerNitroClient).client
	authenticationldapactionName := d.Id()
	err := client.DeleteResource(service.Authenticationldapaction.Type(), authenticationldapactionName)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
