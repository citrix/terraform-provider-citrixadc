package citrixadc

import (
	"github.com/chiradeep/go-nitro/config/ssl"

	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCitrixAdcSslprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSslprofileFunc,
		Read:          readSslprofileFunc,
		Update:        updateSslprofileFunc,
		Delete:        deleteSslprofileFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipherpriority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"cipherredirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cipherurl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cleartextport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientauthuseboundcachain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientcert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"commonname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"denysslreneg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dhekeyexchangewithpsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhkeyexpsizelimit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dropreqwithnohostheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encrypttriggerpktcount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ersa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ersacount": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"hsts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"includesubdomains": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"insertionencoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maxage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ocspstapling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prevsessionkeylifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushenctrigger": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pushenctriggertimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"pushflag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"quantumsize": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirectportrewrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sendclosenotify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serverauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionkeylifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessionticket": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionticketkeydata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionticketkeyrefresh": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sessionticketlifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sessreuse": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"skipclientcertpolicycheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snienable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snihttphostmatch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslimaxsessperserver": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sslinterception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssliocspcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslireneg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssllogprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslprofiletype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslredirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssltriggertimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"strictcachecks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strictsigdigestcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls11": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls12": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls13": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls13sessionticketsperauthcontext": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"zerorttearlydata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// sslprofile_ecccurve_binding
			"ecccurvebindings": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			// sslprofile_cipher_binding
			"cipherbindings": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Set:      sslprofileCipherbindingMappingHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ciphername": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							// ForceNew: true,
						},
						"cipherpriority": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							// Computed: true,
						},
					},
				},
			},
		},
	}
}
func createSslprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)

	sslprofile := ssl.Sslprofile{
		Ciphername:                        d.Get("ciphername").(string),
		Cipherpriority:                    d.Get("cipherpriority").(int),
		Cipherredirect:                    d.Get("cipherredirect").(string),
		Cipherurl:                         d.Get("cipherurl").(string),
		Cleartextport:                     d.Get("cleartextport").(int),
		Clientauth:                        d.Get("clientauth").(string),
		Clientauthuseboundcachain:         d.Get("clientauthuseboundcachain").(string),
		Clientcert:                        d.Get("clientcert").(string),
		Commonname:                        d.Get("commonname").(string),
		Denysslreneg:                      d.Get("denysslreneg").(string),
		Dh:                                d.Get("dh").(string),
		Dhcount:                           d.Get("dhcount").(int),
		Dhekeyexchangewithpsk:             d.Get("dhekeyexchangewithpsk").(string),
		Dhfile:                            d.Get("dhfile").(string),
		Dhkeyexpsizelimit:                 d.Get("dhkeyexpsizelimit").(string),
		Dropreqwithnohostheader:           d.Get("dropreqwithnohostheader").(string),
		Encrypttriggerpktcount:            d.Get("encrypttriggerpktcount").(int),
		Ersa:                              d.Get("ersa").(string),
		Ersacount:                         d.Get("ersacount").(int),
		Hsts:                              d.Get("hsts").(string),
		Includesubdomains:                 d.Get("includesubdomains").(string),
		Insertionencoding:                 d.Get("insertionencoding").(string),
		Maxage:                            d.Get("maxage").(int),
		Name:                              d.Get("name").(string),
		Ocspstapling:                      d.Get("ocspstapling").(string),
		Preload:                           d.Get("preload").(string),
		Prevsessionkeylifetime:            d.Get("prevsessionkeylifetime").(int),
		Pushenctrigger:                    d.Get("pushenctrigger").(string),
		Pushenctriggertimeout:             d.Get("pushenctriggertimeout").(int),
		Pushflag:                          d.Get("pushflag").(int),
		Quantumsize:                       d.Get("quantumsize").(string),
		Redirectportrewrite:               d.Get("redirectportrewrite").(string),
		Sendclosenotify:                   d.Get("sendclosenotify").(string),
		Serverauth:                        d.Get("serverauth").(string),
		Sessionkeylifetime:                d.Get("sessionkeylifetime").(int),
		Sessionticket:                     d.Get("sessionticket").(string),
		Sessionticketkeydata:              d.Get("sessionticketkeydata").(string),
		Sessionticketkeyrefresh:           d.Get("sessionticketkeyrefresh").(string),
		Sessionticketlifetime:             d.Get("sessionticketlifetime").(int),
		Sessreuse:                         d.Get("sessreuse").(string),
		Sesstimeout:                       d.Get("sesstimeout").(int),
		Skipclientcertpolicycheck:         d.Get("skipclientcertpolicycheck").(string),
		Snienable:                         d.Get("snienable").(string),
		Snihttphostmatch:                  d.Get("snihttphostmatch").(string),
		Ssl3:                              d.Get("ssl3").(string),
		Sslimaxsessperserver:              d.Get("sslimaxsessperserver").(int),
		Sslinterception:                   d.Get("sslinterception").(string),
		Ssliocspcheck:                     d.Get("ssliocspcheck").(string),
		Sslireneg:                         d.Get("sslireneg").(string),
		Ssllogprofile:                     d.Get("ssllogprofile").(string),
		Sslprofiletype:                    d.Get("sslprofiletype").(string),
		Sslredirect:                       d.Get("sslredirect").(string),
		Ssltriggertimeout:                 d.Get("ssltriggertimeout").(int),
		Strictcachecks:                    d.Get("strictcachecks").(string),
		Strictsigdigestcheck:              d.Get("strictsigdigestcheck").(string),
		Tls1:                              d.Get("tls1").(string),
		Tls11:                             d.Get("tls11").(string),
		Tls12:                             d.Get("tls12").(string),
		Tls13:                             d.Get("tls13").(string),
		Tls13sessionticketsperauthcontext: d.Get("tls13sessionticketsperauthcontext").(int),
		Zerorttearlydata:                  d.Get("zerorttearlydata").(string),
	}

	_, err := client.AddResource(netscaler.Sslprofile.Type(), sslprofileName, &sslprofile)
	if err != nil {
		return err
	}

	d.SetId(sslprofileName)

	err = createSslprofileEcccurveBindings(d, meta)
	if err != nil {
		return err
	}
	err = createSslprofileCipherBindings(d, meta)
	if err != nil {
		return err
	}
	err = readSslprofileFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this sslprofile but we can't read it ?? %s", sslprofileName)
		return nil
	}
	return nil
}

func readSslprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readSslprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading sslprofile state %s", sslprofileName)
	data, err := client.FindResource(netscaler.Sslprofile.Type(), sslprofileName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing sslprofile state %s", sslprofileName)
		d.SetId("")
		return nil
	}

	err = readSslprofileEcccurvebindings(d, meta)
	if err != nil {
		return err
	}
	err = readSslprofileCipherbindings(d, meta)
	if err != nil {
		return err
	}

	d.Set("name", data["name"])
	d.Set("ciphername", data["ciphername"])
	d.Set("cipherpriority", data["cipherpriority"])
	d.Set("cipherredirect", data["cipherredirect"])
	d.Set("cipherurl", data["cipherurl"])
	d.Set("cleartextport", data["cleartextport"])
	d.Set("clientauth", data["clientauth"])
	d.Set("clientauthuseboundcachain", data["clientauthuseboundcachain"])
	d.Set("clientcert", data["clientcert"])
	d.Set("commonname", data["commonname"])
	d.Set("denysslreneg", data["denysslreneg"])
	d.Set("dh", data["dh"])
	d.Set("dhcount", data["dhcount"])
	d.Set("dhekeyexchangewithpsk", data["dhekeyexchangewithpsk"])
	d.Set("dhfile", data["dhfile"])
	d.Set("dhkeyexpsizelimit", data["dhkeyexpsizelimit"])
	d.Set("dropreqwithnohostheader", data["dropreqwithnohostheader"])
	d.Set("encrypttriggerpktcount", data["encrypttriggerpktcount"])
	d.Set("ersa", data["ersa"])
	d.Set("ersacount", data["ersacount"])
	d.Set("hsts", data["hsts"])
	d.Set("includesubdomains", data["includesubdomains"])
	d.Set("insertionencoding", data["insertionencoding"])
	d.Set("maxage", data["maxage"])
	d.Set("name", data["name"])
	d.Set("ocspstapling", data["ocspstapling"])
	d.Set("preload", data["preload"])
	d.Set("prevsessionkeylifetime", data["prevsessionkeylifetime"])
	d.Set("pushenctrigger", data["pushenctrigger"])
	d.Set("pushenctriggertimeout", data["pushenctriggertimeout"])
	d.Set("pushflag", data["pushflag"])
	d.Set("quantumsize", data["quantumsize"])
	d.Set("redirectportrewrite", data["redirectportrewrite"])
	d.Set("sendclosenotify", data["sendclosenotify"])
	d.Set("serverauth", data["serverauth"])
	d.Set("sessionkeylifetime", data["sessionkeylifetime"])
	d.Set("sessionticket", data["sessionticket"])
	d.Set("sessionticketkeydata", data["sessionticketkeydata"])
	d.Set("sessionticketkeyrefresh", data["sessionticketkeyrefresh"])
	d.Set("sessionticketlifetime", data["sessionticketlifetime"])
	d.Set("sessreuse", data["sessreuse"])
	d.Set("sesstimeout", data["sesstimeout"])
	d.Set("skipclientcertpolicycheck", data["skipclientcertpolicycheck"])
	d.Set("snienable", data["snienable"])
	d.Set("snihttphostmatch", data["snihttphostmatch"])
	d.Set("ssl3", data["ssl3"])
	d.Set("sslimaxsessperserver", data["sslimaxsessperserver"])
	d.Set("sslinterception", data["sslinterception"])
	d.Set("ssliocspcheck", data["ssliocspcheck"])
	d.Set("sslireneg", data["sslireneg"])
	d.Set("ssllogprofile", data["ssllogprofile"])
	d.Set("sslprofiletype", data["sslprofiletype"])
	d.Set("sslredirect", data["sslredirect"])
	d.Set("ssltriggertimeout", data["ssltriggertimeout"])
	d.Set("strictcachecks", data["strictcachecks"])
	d.Set("strictsigdigestcheck", data["strictsigdigestcheck"])
	d.Set("tls1", data["tls1"])
	d.Set("tls11", data["tls11"])
	d.Set("tls12", data["tls12"])
	d.Set("tls13", data["tls13"])
	d.Set("tls13sessionticketsperauthcontext", data["tls13sessionticketsperauthcontext"])
	d.Set("zerorttearlydata", data["zerorttearlydata"])

	return nil

}

func updateSslprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)

	sslprofile := ssl.Sslprofile{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("ciphername") {
		log.Printf("[DEBUG]  citrixadc-provider: Ciphername has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ciphername = d.Get("ciphername").(string)
		hasChange = true
	}
	if d.HasChange("cipherpriority") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherpriority has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Cipherpriority = d.Get("cipherpriority").(int)
		hasChange = true
	}
	if d.HasChange("cipherredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherredirect has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Cipherredirect = d.Get("cipherredirect").(string)
		hasChange = true
	}
	if d.HasChange("cipherurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Cipherurl has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Cipherurl = d.Get("cipherurl").(string)
		hasChange = true
	}
	if d.HasChange("cleartextport") {
		log.Printf("[DEBUG]  citrixadc-provider: Cleartextport has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Cleartextport = d.Get("cleartextport").(int)
		hasChange = true
	}
	if d.HasChange("clientauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientauth has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Clientauth = d.Get("clientauth").(string)
		hasChange = true
	}
	if d.HasChange("clientauthuseboundcachain") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientauthuseboundcachain has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Clientauthuseboundcachain = d.Get("clientauthuseboundcachain").(string)
		hasChange = true
	}
	if d.HasChange("clientcert") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcert has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Clientcert = d.Get("clientcert").(string)
		hasChange = true
	}
	if d.HasChange("commonname") {
		log.Printf("[DEBUG]  citrixadc-provider: Commonname has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Commonname = d.Get("commonname").(string)
		hasChange = true
	}
	if d.HasChange("denysslreneg") {
		log.Printf("[DEBUG]  citrixadc-provider: Denysslreneg has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Denysslreneg = d.Get("denysslreneg").(string)
		hasChange = true
	}
	if d.HasChange("dh") {
		log.Printf("[DEBUG]  citrixadc-provider: Dh has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dh = d.Get("dh").(string)
		hasChange = true
	}
	if d.HasChange("dhcount") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhcount has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dhcount = d.Get("dhcount").(int)
		hasChange = true
	}
	if d.HasChange("dhekeyexchangewithpsk") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhekeyexchangewithpsk has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dhekeyexchangewithpsk = d.Get("dhekeyexchangewithpsk").(string)
		hasChange = true
	}
	if d.HasChange("dhfile") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhfile has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dhfile = d.Get("dhfile").(string)
		hasChange = true
	}
	if d.HasChange("dhkeyexpsizelimit") {
		log.Printf("[DEBUG]  citrixadc-provider: Dhkeyexpsizelimit has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dhkeyexpsizelimit = d.Get("dhkeyexpsizelimit").(string)
		hasChange = true
	}
	if d.HasChange("dropreqwithnohostheader") {
		log.Printf("[DEBUG]  citrixadc-provider: Dropreqwithnohostheader has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Dropreqwithnohostheader = d.Get("dropreqwithnohostheader").(string)
		hasChange = true
	}
	if d.HasChange("encrypttriggerpktcount") {
		log.Printf("[DEBUG]  citrixadc-provider: Encrypttriggerpktcount has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Encrypttriggerpktcount = d.Get("encrypttriggerpktcount").(int)
		hasChange = true
	}
	if d.HasChange("ersa") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersa has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ersa = d.Get("ersa").(string)
		hasChange = true
	}
	if d.HasChange("ersacount") {
		log.Printf("[DEBUG]  citrixadc-provider: Ersacount has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ersacount = d.Get("ersacount").(int)
		hasChange = true
	}
	if d.HasChange("hsts") {
		log.Printf("[DEBUG]  citrixadc-provider: Hsts has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Hsts = d.Get("hsts").(string)
		hasChange = true
	}
	if d.HasChange("includesubdomains") {
		log.Printf("[DEBUG]  citrixadc-provider: Includesubdomains has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Includesubdomains = d.Get("includesubdomains").(string)
		hasChange = true
	}
	if d.HasChange("insertionencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Insertionencoding has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Insertionencoding = d.Get("insertionencoding").(string)
		hasChange = true
	}
	if d.HasChange("maxage") {
		log.Printf("[DEBUG]  citrixadc-provider: Maxage has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Maxage = d.Get("maxage").(int)
		hasChange = true
	}
	if d.HasChange("name") {
		log.Printf("[DEBUG]  citrixadc-provider: Name has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Name = d.Get("name").(string)
		hasChange = true
	}
	if d.HasChange("ocspstapling") {
		log.Printf("[DEBUG]  citrixadc-provider: Ocspstapling has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ocspstapling = d.Get("ocspstapling").(string)
		hasChange = true
	}
	if d.HasChange("preload") {
		log.Printf("[DEBUG]  citrixadc-provider: Preload has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Preload = d.Get("preload").(string)
		hasChange = true
	}
	if d.HasChange("prevsessionkeylifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Prevsessionkeylifetime has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Prevsessionkeylifetime = d.Get("prevsessionkeylifetime").(int)
		hasChange = true
	}
	if d.HasChange("pushenctrigger") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctrigger has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Pushenctrigger = d.Get("pushenctrigger").(string)
		hasChange = true
	}
	if d.HasChange("pushenctriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushenctriggertimeout has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Pushenctriggertimeout = d.Get("pushenctriggertimeout").(int)
		hasChange = true
	}
	if d.HasChange("pushflag") {
		log.Printf("[DEBUG]  citrixadc-provider: Pushflag has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Pushflag = d.Get("pushflag").(int)
		hasChange = true
	}
	if d.HasChange("quantumsize") {
		log.Printf("[DEBUG]  citrixadc-provider: Quantumsize has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Quantumsize = d.Get("quantumsize").(string)
		hasChange = true
	}
	if d.HasChange("redirectportrewrite") {
		log.Printf("[DEBUG]  citrixadc-provider: Redirectportrewrite has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Redirectportrewrite = d.Get("redirectportrewrite").(string)
		hasChange = true
	}
	if d.HasChange("sendclosenotify") {
		log.Printf("[DEBUG]  citrixadc-provider: Sendclosenotify has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sendclosenotify = d.Get("sendclosenotify").(string)
		hasChange = true
	}
	if d.HasChange("serverauth") {
		log.Printf("[DEBUG]  citrixadc-provider: Serverauth has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Serverauth = d.Get("serverauth").(string)
		hasChange = true
	}
	if d.HasChange("sessionkeylifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionkeylifetime has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessionkeylifetime = d.Get("sessionkeylifetime").(int)
		hasChange = true
	}
	if d.HasChange("sessionticket") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionticket has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessionticket = d.Get("sessionticket").(string)
		hasChange = true
	}
	if d.HasChange("sessionticketkeydata") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionticketkeydata has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessionticketkeydata = d.Get("sessionticketkeydata").(string)
		hasChange = true
	}
	if d.HasChange("sessionticketkeyrefresh") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionticketkeyrefresh has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessionticketkeyrefresh = d.Get("sessionticketkeyrefresh").(string)
		hasChange = true
	}
	if d.HasChange("sessionticketlifetime") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessionticketlifetime has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessionticketlifetime = d.Get("sessionticketlifetime").(int)
		hasChange = true
	}
	if d.HasChange("sessreuse") {
		log.Printf("[DEBUG]  citrixadc-provider: Sessreuse has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sessreuse = d.Get("sessreuse").(string)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("skipclientcertpolicycheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Skipclientcertpolicycheck has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Skipclientcertpolicycheck = d.Get("skipclientcertpolicycheck").(string)
		hasChange = true
	}
	if d.HasChange("snienable") {
		log.Printf("[DEBUG]  citrixadc-provider: Snienable has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Snienable = d.Get("snienable").(string)
		hasChange = true
	}
	if d.HasChange("snihttphostmatch") {
		log.Printf("[DEBUG]  citrixadc-provider: Snihttphostmatch has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Snihttphostmatch = d.Get("snihttphostmatch").(string)
		hasChange = true
	}
	if d.HasChange("ssl3") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssl3 has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ssl3 = d.Get("ssl3").(string)
		hasChange = true
	}
	if d.HasChange("sslimaxsessperserver") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslimaxsessperserver has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sslimaxsessperserver = d.Get("sslimaxsessperserver").(int)
		hasChange = true
	}
	if d.HasChange("sslinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslinterception has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sslinterception = d.Get("sslinterception").(string)
		hasChange = true
	}
	if d.HasChange("ssliocspcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssliocspcheck has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ssliocspcheck = d.Get("ssliocspcheck").(string)
		hasChange = true
	}
	if d.HasChange("sslireneg") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslireneg has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sslireneg = d.Get("sslireneg").(string)
		hasChange = true
	}
	if d.HasChange("ssllogprofile") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssllogprofile has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ssllogprofile = d.Get("ssllogprofile").(string)
		hasChange = true
	}
	if d.HasChange("sslprofiletype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslprofiletype has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sslprofiletype = d.Get("sslprofiletype").(string)
		hasChange = true
	}
	if d.HasChange("sslredirect") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslredirect has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Sslredirect = d.Get("sslredirect").(string)
		hasChange = true
	}
	if d.HasChange("ssltriggertimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssltriggertimeout has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Ssltriggertimeout = d.Get("ssltriggertimeout").(int)
		hasChange = true
	}
	if d.HasChange("strictcachecks") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictcachecks has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Strictcachecks = d.Get("strictcachecks").(string)
		hasChange = true
	}
	if d.HasChange("strictsigdigestcheck") {
		log.Printf("[DEBUG]  citrixadc-provider: Strictsigdigestcheck has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Strictsigdigestcheck = d.Get("strictsigdigestcheck").(string)
		hasChange = true
	}
	if d.HasChange("tls1") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls1 has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Tls1 = d.Get("tls1").(string)
		hasChange = true
	}
	if d.HasChange("tls11") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls11 has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Tls11 = d.Get("tls11").(string)
		hasChange = true
	}
	if d.HasChange("tls12") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls12 has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Tls12 = d.Get("tls12").(string)
		hasChange = true
	}
	if d.HasChange("tls13") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13 has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Tls13 = d.Get("tls13").(string)
		hasChange = true
	}
	if d.HasChange("tls13sessionticketsperauthcontext") {
		log.Printf("[DEBUG]  citrixadc-provider: Tls13sessionticketsperauthcontext has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Tls13sessionticketsperauthcontext = d.Get("tls13sessionticketsperauthcontext").(int)
		hasChange = true
	}
	if d.HasChange("zerorttearlydata") {
		log.Printf("[DEBUG]  citrixadc-provider: Zerorttearlydata has changed for sslprofile %s, starting update", sslprofileName)
		sslprofile.Zerorttearlydata = d.Get("zerorttearlydata").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(netscaler.Sslprofile.Type(), sslprofileName, &sslprofile)
		if err != nil {
			return fmt.Errorf("Error updating sslprofile %s", sslprofileName)
		}
	}

	if d.HasChange("ecccurvebindings") {
		err := updateSslprofileEcccurveBindings(d, meta)
		if err != nil {
			return err
		}
	}

	if d.HasChange("cipherbindings") {
		err := updateSslprofileCipherBindings(d, meta)
		if err != nil {
			return err
		}
	}
	return readSslprofileFunc(d, meta)
}

func deleteSslprofileFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSslprofileFunc")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Id()
	err := client.DeleteResource(netscaler.Sslprofile.Type(), sslprofileName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteSingleSslprofileEcccurveBinding(d *schema.ResourceData, meta interface{}, ecccurvename string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSslprofileEcccurveBinding")
	client := meta.(*NetScalerNitroClient).client

	sslprofileName := d.Get("name").(string)
	args := make([]string, 0, 1)

	s := fmt.Sprintf("ecccurvename:%s", ecccurvename)
	args = append(args, s)

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs(netscaler.Sslprofile_ecccurve_binding.Type(), sslprofileName, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting EccCurve binding %v\n", sslprofileName)
		return err
	}

	return nil
}

func addSingleSslprofileEcccurveBinding(d *schema.ResourceData, meta interface{}, ecccurvename string) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSslprofileEcccurveBinding")
	client := meta.(*NetScalerNitroClient).client

	bindingStruct := ssl.Sslprofileecccurvebinding{}
	bindingStruct.Name = d.Get("name").(string)
	bindingStruct.Ecccurvename = ecccurvename

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource(netscaler.Sslprofile_ecccurve_binding.Type(), bindingStruct.Name, bindingStruct); err != nil {
		return err
	}
	return nil
}

func getDefaultSslprofileEcccurveBindings(d *schema.ResourceData, meta interface{}) ([]string, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In getDefaultSslprofileEcccurveBindings")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)
	bindings, _ := client.FindResourceArray(netscaler.Sslprofile_ecccurve_binding.Type(), sslprofileName)
	log.Printf("bindings %v\n", bindings)

	defaultSslprofileEcccurveBindings := make([]string, len(bindings))

	for i, val := range bindings {
		defaultSslprofileEcccurveBindings[i] = val["ecccurvename"].(string)
	}

	return defaultSslprofileEcccurveBindings, nil
}

func createSslprofileEcccurveBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofileEcccurveBindings")
	oldSet, newSet := d.GetChange("ecccurvebindings")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))

	// Delete default ECCcurves being bound to SSLprofile
	// If a user explicitly gives these ECCcurves, these will be added in the next step
	// DO NOT catch any errors while deleting. If delete fails here, just continue

	// get default Ecccurve bindings to the created SSLprofile
	defaultEcccurves, err := getDefaultSslprofileEcccurveBindings(d, meta)
	log.Printf("[DEBUG] citrixadc-provider: defaultSslprofileEcccurveBindings: %v", defaultEcccurves)
	if err != nil {
		return err
	}

	for _, ecccurvename := range defaultEcccurves {
		deleteSingleSslprofileEcccurveBinding(d, meta, ecccurvename)
	}

	for _, ecccurvename := range add.List() {
		if err := addSingleSslprofileEcccurveBinding(d, meta, ecccurvename.(string)); err != nil {
			return err
		}
	}
	return nil
}

func updateSslprofileEcccurveBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslprofileEcccurveBindings")
	oldSet, newSet := d.GetChange("ecccurvebindings")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, ecccurvename := range remove.List() {
		if err := deleteSingleSslprofileEcccurveBinding(d, meta, ecccurvename.(string)); err != nil {
			return err
		}
	}

	for _, ecccurvename := range add.List() {
		if err := addSingleSslprofileEcccurveBinding(d, meta, ecccurvename.(string)); err != nil {
			return err
		}
	}
	return nil
}

func readSslprofileEcccurvebindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSslprofileEcccurvebindings")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)
	bindings, _ := client.FindResourceArray(netscaler.Sslprofile_ecccurve_binding.Type(), sslprofileName)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))
	for i, val := range bindings {
		processedBindings[i] = val["ecccurvename"].(string)
	}

	updatedSet := processedBindings
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("ecccurvebindings", updatedSet); err != nil {
		return err
	}
	return nil
}

// Cipher bindings

func deleteSingleSslprofileCipherBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteSingleSslprofileCipherBinding")
	client := meta.(*NetScalerNitroClient).client

	sslprofileName := d.Get("name").(string)
	// construct args from binding data
	args := make([]string, 0, 1)

	if d, ok := binding["ciphername"]; ok {
		s := fmt.Sprintf("ciphername:%v", d.(string))
		args = append(args, s)
	}

	log.Printf("args is %v", args)

	if err := client.DeleteResourceWithArgs(netscaler.Sslprofile_sslcipher_binding.Type(), sslprofileName, args); err != nil {
		log.Printf("[DEBUG]  citrixadc-provider: Error deleting Cipher binding %v\n", binding)
		return err
	}

	return nil
}

func addSingleSslprofileCipherBinding(d *schema.ResourceData, meta interface{}, binding map[string]interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In addSingleSslprofileCipherBinding")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("Adding binding %v", binding)

	bindingStruct := ssl.Sslprofilesslcipherbinding{}
	bindingStruct.Name = d.Get("name").(string)

	if d, ok := binding["ciphername"]; ok {
		bindingStruct.Ciphername = d.(string)
	}

	if d, ok := binding["cipherpriority"]; ok {
		bindingStruct.Cipherpriority = d.(int)
	}

	// We need to do a HTTP PUT hence the UpdateResource
	if _, err := client.UpdateResource(netscaler.Sslprofile_sslcipher_binding.Type(), bindingStruct.Name, bindingStruct); err != nil {
		return err
	}
	return nil
}

func getDefaultSslprofileCipherBindings(d *schema.ResourceData, meta interface{}) ([]interface{}, error) {
	log.Printf("[DEBUG]  citrixadc-provider: In getDefaultSslprofileCipherBindings")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)
	bindings, _ := client.FindResourceArray(netscaler.Sslprofile_sslcipher_binding.Type(), sslprofileName)
	log.Printf("bindings %v\n", bindings)

	defaultSslprofileCipherBindings := make([]interface{}, len(bindings))

	for i := range bindings {
		defaultSslprofileCipherBindings[i] = make(map[string]interface{})
		defaultSslprofileCipherBindings[i].(map[string]interface{})["ciphername"] = bindings[i]["cipheraliasname"].(string)
		cipherpriorityString, err := strconv.Atoi(bindings[i]["cipherpriority"].(string))
		if err != nil {
			return nil, err
		}
		defaultSslprofileCipherBindings[i].(map[string]interface{})["cipherpriority"] = cipherpriorityString
	}

	return defaultSslprofileCipherBindings, nil
}

func createSslprofileCipherBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createSslprofileCipherBindings")
	oldSet, newSet := d.GetChange("cipherbindings")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))

	// Delete the default ciphers being bound to SSLprofile by default
	// If a user explicitly gives this cipher, this will be added in the next step
	// DO NOT catch any errors while deleting. If delete fails here, just continue

	// Get the default ciphers from the created SSL Profile and delete them
	defaultCipherBindings, err := getDefaultSslprofileCipherBindings(d, meta)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] citrixadc-provider: defaultSslprofileCipherBindings: %v", defaultCipherBindings)
	for _, binding := range defaultCipherBindings {
		if err := deleteSingleSslprofileCipherBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	for _, binding := range add.List() {
		if err := addSingleSslprofileCipherBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func updateSslprofileCipherBindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateSslprofileCipherBindings")
	oldSet, newSet := d.GetChange("cipherbindings")
	log.Printf("[DEBUG]  citrixadc-provider: oldSet %v\n", oldSet)
	log.Printf("[DEBUG]  citrixadc-provider: newSet %v\n", newSet)
	remove := oldSet.(*schema.Set).Difference(newSet.(*schema.Set))
	add := newSet.(*schema.Set).Difference(oldSet.(*schema.Set))
	for _, binding := range remove.List() {
		if err := deleteSingleSslprofileCipherBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	for _, binding := range add.List() {
		if err := addSingleSslprofileCipherBinding(d, meta, binding.(map[string]interface{})); err != nil {
			return err
		}
	}
	return nil
}

func readSslprofileCipherbindings(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In readSslprofileCipherbindings")
	client := meta.(*NetScalerNitroClient).client
	sslprofileName := d.Get("name").(string)
	bindings, _ := client.FindResourceArray(netscaler.Sslprofile_sslcipher_binding.Type(), sslprofileName)
	log.Printf("bindings %v\n", bindings)

	processedBindings := make([]interface{}, len(bindings))

	for i := range bindings {
		processedBindings[i] = make(map[string]interface{})
		processedBindings[i].(map[string]interface{})["ciphername"] = bindings[i]["cipheraliasname"].(string)
		cipherpriorityString, err := strconv.Atoi(bindings[i]["cipherpriority"].(string))
		if err != nil {
			return err
		}
		processedBindings[i].(map[string]interface{})["cipherpriority"] = cipherpriorityString
	}

	updatedSet := schema.NewSet(sslprofileCipherbindingMappingHash, processedBindings)
	log.Printf("updatedSet %v\n", updatedSet)
	if err := d.Set("cipherbindings", updatedSet); err != nil {
		return err
	}
	return nil
}

func sslprofileCipherbindingMappingHash(v interface{}) int {
	log.Printf("[DEBUG]  citrixadc-provider: In sslprofileCipherbindingMappingHash")
	var buf bytes.Buffer

	m := v.(map[string]interface{})
	if d, ok := m["ciphername"]; ok {
		buf.WriteString(fmt.Sprintf("%s-", d.(string)))
	}
	if d, ok := m["cipherpriority"]; ok {
		buf.WriteString(fmt.Sprintf("%d-", d.(int)))
	}
	return hashcode.String(buf.String())
}
