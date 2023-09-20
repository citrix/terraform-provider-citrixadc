package citrixadc

import (
	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"fmt"
	"log"
)

func resourceCitrixAdcVpnsessionaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createVpnsessionactionFunc,
		Read:          readVpnsessionactionFunc,
		Update:        updateVpnsessionactionFunc,
		Delete:        deleteVpnsessionactionFunc,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"advancedclientlessvpnmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowedlogingroups": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allprotocolproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alwaysonprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorizationgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"autoproxyurl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"citrixreceiverhome": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientchoices": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientcleanupprompt": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientconfiguration": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"clientdebug": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientidletimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"clientlessmodeurlencoding": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientlesspersistentcookie": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientlessvpnmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientoptions": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecurity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecuritygroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecuritylog": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clientsecuritymessage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defaultauthorizationaction": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsvservername": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailhome": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"epaclienttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forcecleanup": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
			"forcedtimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"forcedtimeoutwarning": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"fqdnspoofedip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftpproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gopherproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"homepage": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"httpport": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeInt},
				Optional: true,
				Computed: true,
			},
			"httpproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icaproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iconwithreceiver": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iipdnssuffix": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"kcdaccount": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"killconnections": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linuxpluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locallanaccess": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"loginscript": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logoutscript": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macpluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntdomain": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pcoipprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxyexception": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxylocalbypass": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdpclientprofilename": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rfc1918": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"securebrowse": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sesstimeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sfgatewayauthtype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smartgroup": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"socksproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"splitdns": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"splittunnel": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spoofiip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sslproxy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssocredential": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storefronturl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transparentinterception": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useiip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"usemip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"useraccounting": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wihome": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wihomeaddresstype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windowsautologon": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windowsclienttype": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windowspluginupgrade": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"winsip": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wiportalmode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func createVpnsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionactionName := d.Get("name").(string)
	vpnsessionaction := vpn.Vpnsessionaction{
		Advancedclientlessvpnmode:  d.Get("advancedclientlessvpnmode").(string),
		Allowedlogingroups:         d.Get("allowedlogingroups").(string),
		Allprotocolproxy:           d.Get("allprotocolproxy").(string),
		Alwaysonprofilename:        d.Get("alwaysonprofilename").(string),
		Authorizationgroup:         d.Get("authorizationgroup").(string),
		Autoproxyurl:               d.Get("autoproxyurl").(string),
		Citrixreceiverhome:         d.Get("citrixreceiverhome").(string),
		Clientchoices:              d.Get("clientchoices").(string),
		Clientcleanupprompt:        d.Get("clientcleanupprompt").(string),
		Clientconfiguration:        toStringList(d.Get("clientconfiguration").([]interface{})),
		Clientdebug:                d.Get("clientdebug").(string),
		Clientidletimeout:          d.Get("clientidletimeout").(int),
		Clientlessmodeurlencoding:  d.Get("clientlessmodeurlencoding").(string),
		Clientlesspersistentcookie: d.Get("clientlesspersistentcookie").(string),
		Clientlessvpnmode:          d.Get("clientlessvpnmode").(string),
		Clientoptions:              d.Get("clientoptions").(string),
		Clientsecurity:             d.Get("clientsecurity").(string),
		Clientsecuritygroup:        d.Get("clientsecuritygroup").(string),
		Clientsecuritylog:          d.Get("clientsecuritylog").(string),
		Clientsecuritymessage:      d.Get("clientsecuritymessage").(string),
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Dnsvservername:             d.Get("dnsvservername").(string),
		Emailhome:                  d.Get("emailhome").(string),
		Epaclienttype:              d.Get("epaclienttype").(string),
		Forcecleanup:               toStringList(d.Get("forcecleanup").([]interface{})),
		Forcedtimeout:              d.Get("forcedtimeout").(int),
		Forcedtimeoutwarning:       d.Get("forcedtimeoutwarning").(int),
		Fqdnspoofedip:              d.Get("fqdnspoofedip").(string),
		Ftpproxy:                   d.Get("ftpproxy").(string),
		Gopherproxy:                d.Get("gopherproxy").(string),
		Homepage:                   d.Get("homepage").(string),
		Httpport:                   toIntegerList(d.Get("httpport").([]interface{})),
		Httpproxy:                  d.Get("httpproxy").(string),
		Icaproxy:                   d.Get("icaproxy").(string),
		Iconwithreceiver:           d.Get("iconwithreceiver").(string),
		Iipdnssuffix:               d.Get("iipdnssuffix").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Killconnections:            d.Get("killconnections").(string),
		Linuxpluginupgrade:         d.Get("linuxpluginupgrade").(string),
		Locallanaccess:             d.Get("locallanaccess").(string),
		Loginscript:                d.Get("loginscript").(string),
		Logoutscript:               d.Get("logoutscript").(string),
		Macpluginupgrade:           d.Get("macpluginupgrade").(string),
		Name:                       d.Get("name").(string),
		Netmask:                    d.Get("netmask").(string),
		Ntdomain:                   d.Get("ntdomain").(string),
		Pcoipprofilename:           d.Get("pcoipprofilename").(string),
		Proxy:                      d.Get("proxy").(string),
		Proxyexception:             d.Get("proxyexception").(string),
		Proxylocalbypass:           d.Get("proxylocalbypass").(string),
		Rdpclientprofilename:       d.Get("rdpclientprofilename").(string),
		Rfc1918:                    d.Get("rfc1918").(string),
		Securebrowse:               d.Get("securebrowse").(string),
		Sesstimeout:                d.Get("sesstimeout").(int),
		Sfgatewayauthtype:          d.Get("sfgatewayauthtype").(string),
		Smartgroup:                 d.Get("smartgroup").(string),
		Socksproxy:                 d.Get("socksproxy").(string),
		Splitdns:                   d.Get("splitdns").(string),
		Splittunnel:                d.Get("splittunnel").(string),
		Spoofiip:                   d.Get("spoofiip").(string),
		Sslproxy:                   d.Get("sslproxy").(string),
		Sso:                        d.Get("sso").(string),
		Ssocredential:              d.Get("ssocredential").(string),
		Storefronturl:              d.Get("storefronturl").(string),
		Transparentinterception:    d.Get("transparentinterception").(string),
		Useiip:                     d.Get("useiip").(string),
		Usemip:                     d.Get("usemip").(string),
		Useraccounting:             d.Get("useraccounting").(string),
		Wihome:                     d.Get("wihome").(string),
		Wihomeaddresstype:          d.Get("wihomeaddresstype").(string),
		Windowsautologon:           d.Get("windowsautologon").(string),
		Windowsclienttype:          d.Get("windowsclienttype").(string),
		Windowspluginupgrade:       d.Get("windowspluginupgrade").(string),
		Winsip:                     d.Get("winsip").(string),
		Wiportalmode:               d.Get("wiportalmode").(string),
	}

	_, err := client.AddResource(service.Vpnsessionaction.Type(), vpnsessionactionName, &vpnsessionaction)
	if err != nil {
		return err
	}

	d.SetId(vpnsessionactionName)

	err = readVpnsessionactionFunc(d, meta)
	if err != nil {
		log.Printf("[ERROR] netscaler-provider: ?? we just created this vpnsessionaction but we can't read it ?? %s", vpnsessionactionName)
		return nil
	}
	return nil
}

func readVpnsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionactionName := d.Id()
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnsessionaction state %s", vpnsessionactionName)
	data, err := client.FindResource(service.Vpnsessionaction.Type(), vpnsessionactionName)
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnsessionaction state %s", vpnsessionactionName)
		d.SetId("")
		return nil
	}
	d.Set("name", data["name"])
	d.Set("advancedclientlessvpnmode", data["advancedclientlessvpnmode"])
	d.Set("allowedlogingroups", data["allowedlogingroups"])
	d.Set("allprotocolproxy", data["allprotocolproxy"])
	d.Set("alwaysonprofilename", data["alwaysonprofilename"])
	d.Set("authorizationgroup", data["authorizationgroup"])
	d.Set("autoproxyurl", data["autoproxyurl"])
	d.Set("citrixreceiverhome", data["citrixreceiverhome"])
	d.Set("clientchoices", data["clientchoices"])
	d.Set("clientcleanupprompt", data["clientcleanupprompt"])
	d.Set("clientconfiguration", data["clientconfiguration"])
	d.Set("clientdebug", data["clientdebug"])
	d.Set("clientidletimeout", data["clientidletimeout"])
	d.Set("clientlessmodeurlencoding", data["clientlessmodeurlencoding"])
	d.Set("clientlesspersistentcookie", data["clientlesspersistentcookie"])
	d.Set("clientlessvpnmode", data["clientlessvpnmode"])
	d.Set("clientoptions", data["clientoptions"])
	d.Set("clientsecurity", data["clientsecurity"])
	d.Set("clientsecuritygroup", data["clientsecuritygroup"])
	d.Set("clientsecuritylog", data["clientsecuritylog"])
	d.Set("clientsecuritymessage", data["clientsecuritymessage"])
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("emailhome", data["emailhome"])
	d.Set("epaclienttype", data["epaclienttype"])
	d.Set("forcecleanup", data["forcecleanup"])
	d.Set("forcedtimeout", data["forcedtimeout"])
	d.Set("forcedtimeoutwarning", data["forcedtimeoutwarning"])
	d.Set("fqdnspoofedip", data["fqdnspoofedip"])
	d.Set("ftpproxy", data["ftpproxy"])
	d.Set("gopherproxy", data["gopherproxy"])
	d.Set("homepage", data["homepage"])
	d.Set("httpport", data["httpport"])
	d.Set("httpproxy", data["httpproxy"])
	d.Set("icaproxy", data["icaproxy"])
	d.Set("iconwithreceiver", data["iconwithreceiver"])
	d.Set("iipdnssuffix", data["iipdnssuffix"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("killconnections", data["killconnections"])
	d.Set("linuxpluginupgrade", data["linuxpluginupgrade"])
	d.Set("locallanaccess", data["locallanaccess"])
	d.Set("loginscript", data["loginscript"])
	d.Set("logoutscript", data["logoutscript"])
	d.Set("macpluginupgrade", data["macpluginupgrade"])
	d.Set("name", data["name"])
	d.Set("netmask", data["netmask"])
	d.Set("ntdomain", data["ntdomain"])
	d.Set("pcoipprofilename", data["pcoipprofilename"])
	d.Set("proxy", data["proxy"])
	d.Set("proxyexception", data["proxyexception"])
	d.Set("proxylocalbypass", data["proxylocalbypass"])
	d.Set("rdpclientprofilename", data["rdpclientprofilename"])
	d.Set("rfc1918", data["rfc1918"])
	d.Set("securebrowse", data["securebrowse"])
	d.Set("sesstimeout", data["sesstimeout"])
	d.Set("sfgatewayauthtype", data["sfgatewayauthtype"])
	d.Set("smartgroup", data["smartgroup"])
	d.Set("socksproxy", data["socksproxy"])
	d.Set("splitdns", data["splitdns"])
	d.Set("splittunnel", data["splittunnel"])
	d.Set("spoofiip", data["spoofiip"])
	d.Set("sslproxy", data["sslproxy"])
	d.Set("sso", data["sso"])
	d.Set("ssocredential", data["ssocredential"])
	d.Set("storefronturl", data["storefronturl"])
	d.Set("transparentinterception", data["transparentinterception"])
	d.Set("useiip", data["useiip"])
	d.Set("usemip", data["usemip"])
	d.Set("useraccounting", data["useraccounting"])
	d.Set("wihome", data["wihome"])
	d.Set("wihomeaddresstype", data["wihomeaddresstype"])
	d.Set("windowsautologon", data["windowsautologon"])
	d.Set("windowsclienttype", data["windowsclienttype"])
	d.Set("windowspluginupgrade", data["windowspluginupgrade"])
	d.Set("winsip", data["winsip"])
	d.Set("wiportalmode", data["wiportalmode"])

	return nil

}

func updateVpnsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionactionName := d.Get("name").(string)

	vpnsessionaction := vpn.Vpnsessionaction{
		Name: d.Get("name").(string),
	}
	hasChange := false
	if d.HasChange("advancedclientlessvpnmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Advancedclientlessvpnmode has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Advancedclientlessvpnmode = d.Get("advancedclientlessvpnmode").(string)
		hasChange = true
	}
	if d.HasChange("allowedlogingroups") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowedlogingroups has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Allowedlogingroups = d.Get("allowedlogingroups").(string)
		hasChange = true
	}
	if d.HasChange("allprotocolproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Allprotocolproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Allprotocolproxy = d.Get("allprotocolproxy").(string)
		hasChange = true
	}
	if d.HasChange("alwaysonprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Alwaysonprofilename has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Alwaysonprofilename = d.Get("alwaysonprofilename").(string)
		hasChange = true
	}
	if d.HasChange("authorizationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorizationgroup has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Authorizationgroup = d.Get("authorizationgroup").(string)
		hasChange = true
	}
	if d.HasChange("autoproxyurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoproxyurl has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Autoproxyurl = d.Get("autoproxyurl").(string)
		hasChange = true
	}
	if d.HasChange("citrixreceiverhome") {
		log.Printf("[DEBUG]  citrixadc-provider: Citrixreceiverhome has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Citrixreceiverhome = d.Get("citrixreceiverhome").(string)
		hasChange = true
	}
	if d.HasChange("clientchoices") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientchoices has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientchoices = d.Get("clientchoices").(string)
		hasChange = true
	}
	if d.HasChange("clientcleanupprompt") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcleanupprompt has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientcleanupprompt = d.Get("clientcleanupprompt").(string)
		hasChange = true
	}
	if d.HasChange("clientconfiguration") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientconfiguration has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientconfiguration = toStringList(d.Get("clientconfiguration").([]interface{}))
		hasChange = true
	}
	if d.HasChange("clientdebug") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientdebug has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientdebug = d.Get("clientdebug").(string)
		hasChange = true
	}
	if d.HasChange("clientidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientidletimeout has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientidletimeout = d.Get("clientidletimeout").(int)
		hasChange = true
	}
	if d.HasChange("clientlessmodeurlencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessmodeurlencoding has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientlessmodeurlencoding = d.Get("clientlessmodeurlencoding").(string)
		hasChange = true
	}
	if d.HasChange("clientlesspersistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlesspersistentcookie has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientlesspersistentcookie = d.Get("clientlesspersistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("clientlessvpnmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessvpnmode has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientlessvpnmode = d.Get("clientlessvpnmode").(string)
		hasChange = true
	}
	if d.HasChange("clientoptions") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientoptions has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientoptions = d.Get("clientoptions").(string)
		hasChange = true
	}
	if d.HasChange("clientsecurity") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecurity has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientsecurity = d.Get("clientsecurity").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritygroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritygroup has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientsecuritygroup = d.Get("clientsecuritygroup").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritylog") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritylog has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientsecuritylog = d.Get("clientsecuritylog").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritymessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritymessage has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Clientsecuritymessage = d.Get("clientsecuritymessage").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("dnsvservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsvservername has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Dnsvservername = d.Get("dnsvservername").(string)
		hasChange = true
	}
	if d.HasChange("emailhome") {
		log.Printf("[DEBUG]  citrixadc-provider: Emailhome has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Emailhome = d.Get("emailhome").(string)
		hasChange = true
	}
	if d.HasChange("epaclienttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Epaclienttype has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Epaclienttype = d.Get("epaclienttype").(string)
		hasChange = true
	}
	if d.HasChange("forcecleanup") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcecleanup has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Forcecleanup = toStringList(d.Get("forcecleanup").([]interface{}))
		hasChange = true
	}
	if d.HasChange("forcedtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeout has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Forcedtimeout = d.Get("forcedtimeout").(int)
		hasChange = true
	}
	if d.HasChange("forcedtimeoutwarning") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeoutwarning has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Forcedtimeoutwarning = d.Get("forcedtimeoutwarning").(int)
		hasChange = true
	}
	if d.HasChange("fqdnspoofedip") {
		log.Printf("[DEBUG]  citrixadc-provider: Fqdnspoofedip has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Fqdnspoofedip = d.Get("fqdnspoofedip").(string)
		hasChange = true
	}
	if d.HasChange("ftpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftpproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Ftpproxy = d.Get("ftpproxy").(string)
		hasChange = true
	}
	if d.HasChange("gopherproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Gopherproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Gopherproxy = d.Get("gopherproxy").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httpport") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpport has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Httpport = toIntegerList(d.Get("httpport").([]interface{}))
		hasChange = true
	}
	if d.HasChange("httpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Httpproxy = d.Get("httpproxy").(string)
		hasChange = true
	}
	if d.HasChange("icaproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Icaproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Icaproxy = d.Get("icaproxy").(string)
		hasChange = true
	}
	if d.HasChange("iconwithreceiver") {
		log.Printf("[DEBUG]  citrixadc-provider: Iconwithreceiver has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Iconwithreceiver = d.Get("iconwithreceiver").(string)
		hasChange = true
	}
	if d.HasChange("iipdnssuffix") {
		log.Printf("[DEBUG]  citrixadc-provider: Iipdnssuffix has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Iipdnssuffix = d.Get("iipdnssuffix").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("killconnections") {
		log.Printf("[DEBUG]  citrixadc-provider: Killconnections has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Killconnections = d.Get("killconnections").(string)
		hasChange = true
	}
	if d.HasChange("linuxpluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Linuxpluginupgrade has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Linuxpluginupgrade = d.Get("linuxpluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("locallanaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Locallanaccess has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Locallanaccess = d.Get("locallanaccess").(string)
		hasChange = true
	}
	if d.HasChange("loginscript") {
		log.Printf("[DEBUG]  citrixadc-provider: Loginscript has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Loginscript = d.Get("loginscript").(string)
		hasChange = true
	}
	if d.HasChange("logoutscript") {
		log.Printf("[DEBUG]  citrixadc-provider: Logoutscript has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Logoutscript = d.Get("logoutscript").(string)
		hasChange = true
	}
	if d.HasChange("macpluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Macpluginupgrade has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Macpluginupgrade = d.Get("macpluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("ntdomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ntdomain has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Ntdomain = d.Get("ntdomain").(string)
		hasChange = true
	}
	if d.HasChange("pcoipprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcoipprofilename has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Pcoipprofilename = d.Get("pcoipprofilename").(string)
		hasChange = true
	}
	if d.HasChange("proxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Proxy = d.Get("proxy").(string)
		hasChange = true
	}
	if d.HasChange("proxyexception") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyexception has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Proxyexception = d.Get("proxyexception").(string)
		hasChange = true
	}
	if d.HasChange("proxylocalbypass") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxylocalbypass has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Proxylocalbypass = d.Get("proxylocalbypass").(string)
		hasChange = true
	}
	if d.HasChange("rdpclientprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpclientprofilename has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Rdpclientprofilename = d.Get("rdpclientprofilename").(string)
		hasChange = true
	}
	if d.HasChange("rfc1918") {
		log.Printf("[DEBUG]  citrixadc-provider: Rfc1918 has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Rfc1918 = d.Get("rfc1918").(string)
		hasChange = true
	}
	if d.HasChange("securebrowse") {
		log.Printf("[DEBUG]  citrixadc-provider: Securebrowse has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Securebrowse = d.Get("securebrowse").(string)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("sfgatewayauthtype") {
		log.Printf("[DEBUG]  citrixadc-provider: Sfgatewayauthtype has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Sfgatewayauthtype = d.Get("sfgatewayauthtype").(string)
		hasChange = true
	}
	if d.HasChange("smartgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Smartgroup has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Smartgroup = d.Get("smartgroup").(string)
		hasChange = true
	}
	if d.HasChange("socksproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Socksproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Socksproxy = d.Get("socksproxy").(string)
		hasChange = true
	}
	if d.HasChange("splitdns") {
		log.Printf("[DEBUG]  citrixadc-provider: Splitdns has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Splitdns = d.Get("splitdns").(string)
		hasChange = true
	}
	if d.HasChange("splittunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Splittunnel has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Splittunnel = d.Get("splittunnel").(string)
		hasChange = true
	}
	if d.HasChange("spoofiip") {
		log.Printf("[DEBUG]  citrixadc-provider: Spoofiip has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Spoofiip = d.Get("spoofiip").(string)
		hasChange = true
	}
	if d.HasChange("sslproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslproxy has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Sslproxy = d.Get("sslproxy").(string)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("storefronturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Storefronturl has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Storefronturl = d.Get("storefronturl").(string)
		hasChange = true
	}
	if d.HasChange("transparentinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Transparentinterception has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Transparentinterception = d.Get("transparentinterception").(string)
		hasChange = true
	}
	if d.HasChange("useiip") {
		log.Printf("[DEBUG]  citrixadc-provider: Useiip has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Useiip = d.Get("useiip").(string)
		hasChange = true
	}
	if d.HasChange("usemip") {
		log.Printf("[DEBUG]  citrixadc-provider: Usemip has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Usemip = d.Get("usemip").(string)
		hasChange = true
	}
	if d.HasChange("useraccounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Useraccounting has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Useraccounting = d.Get("useraccounting").(string)
		hasChange = true
	}
	if d.HasChange("wihome") {
		log.Printf("[DEBUG]  citrixadc-provider: Wihome has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Wihome = d.Get("wihome").(string)
		hasChange = true
	}
	if d.HasChange("wihomeaddresstype") {
		log.Printf("[DEBUG]  citrixadc-provider: Wihomeaddresstype has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Wihomeaddresstype = d.Get("wihomeaddresstype").(string)
		hasChange = true
	}
	if d.HasChange("windowsautologon") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowsautologon has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Windowsautologon = d.Get("windowsautologon").(string)
		hasChange = true
	}
	if d.HasChange("windowsclienttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowsclienttype has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Windowsclienttype = d.Get("windowsclienttype").(string)
		hasChange = true
	}
	if d.HasChange("windowspluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowspluginupgrade has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Windowspluginupgrade = d.Get("windowspluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("winsip") {
		log.Printf("[DEBUG]  citrixadc-provider: Winsip has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Winsip = d.Get("winsip").(string)
		hasChange = true
	}
	if d.HasChange("wiportalmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Wiportalmode has changed for vpnsessionaction %s, starting update", vpnsessionactionName)
		vpnsessionaction.Wiportalmode = d.Get("wiportalmode").(string)
		hasChange = true
	}

	if hasChange {
		_, err := client.UpdateResource(service.Vpnsessionaction.Type(), vpnsessionactionName, &vpnsessionaction)
		if err != nil {
			return fmt.Errorf("Error updating vpnsessionaction %s", vpnsessionactionName)
		}
	}
	return readVpnsessionactionFunc(d, meta)
}

func deleteVpnsessionactionFunc(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnsessionactionFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnsessionactionName := d.Id()
	err := client.DeleteResource(service.Vpnsessionaction.Type(), vpnsessionactionName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
