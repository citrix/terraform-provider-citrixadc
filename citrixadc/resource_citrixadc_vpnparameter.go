package citrixadc

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCitrixAdcVpnparameter() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		CreateContext: createVpnparameterFunc,
		ReadContext:   readVpnparameterFunc,
		UpdateContext: updateVpnparameterFunc,
		DeleteContext: deleteVpnparameterFunc, // Thought vpnparameter resource donot have DELETE operation, it is required to set ID to "" d.SetID("") to maintain terraform state
		Schema: map[string]*schema.Schema{
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
			"apptokentimeout": {
				Type:     schema.TypeInt,
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
			"backendcertvalidation": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backendserversni": {
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
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"clientversions": {
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
			"encryptcsecexp": {
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
			"icasessiontimeout": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icauseraccounting": {
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
			"mdxtokentimeout": {
				Type:     schema.TypeInt,
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
			"samesite": {
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
			"uitheme": {
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
			"userdomains": {
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

func createVpnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In createVpnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	var vpnparameterName string
	// there is no primary key in VPNPARAMETER resource. Hence generate one for terraform state maintenance
	vpnparameterName = resource.PrefixedUniqueId("tf-vpnparameter-")
	vpnparameter := vpn.Vpnparameter{
		Advancedclientlessvpnmode:  d.Get("advancedclientlessvpnmode").(string),
		Allowedlogingroups:         d.Get("allowedlogingroups").(string),
		Allprotocolproxy:           d.Get("allprotocolproxy").(string),
		Alwaysonprofilename:        d.Get("alwaysonprofilename").(string),
		Apptokentimeout:            d.Get("apptokentimeout").(int),
		Authorizationgroup:         d.Get("authorizationgroup").(string),
		Autoproxyurl:               d.Get("autoproxyurl").(string),
		Backendcertvalidation:      d.Get("backendcertvalidation").(string),
		Backendserversni:           d.Get("backendserversni").(string),
		Citrixreceiverhome:         d.Get("citrixreceiverhome").(string),
		Clientchoices:              d.Get("clientchoices").(string),
		Clientcleanupprompt:        d.Get("clientcleanupprompt").(string),
		Clientconfiguration:        toStringList(d.Get("clientconfiguration").([]interface{})),
		Clientdebug:                d.Get("clientdebug").(string),
		Clientidletimeout:          d.Get("clientidletimeout").(int),
		Clientlessmodeurlencoding:  d.Get("clientlessmodeurlencoding").(string),
		Clientlesspersistentcookie: d.Get("clientlesspersistentcookie").(string),
		Clientlessvpnmode:          d.Get("clientlessvpnmode").(string),
		Clientoptions:              toStringList(d.Get("clientoptions").([]interface{})),
		Clientsecurity:             d.Get("clientsecurity").(string),
		Clientsecuritygroup:        d.Get("clientsecuritygroup").(string),
		Clientsecuritylog:          d.Get("clientsecuritylog").(string),
		Clientsecuritymessage:      d.Get("clientsecuritymessage").(string),
		Clientversions:             d.Get("clientversions").(string),
		Defaultauthorizationaction: d.Get("defaultauthorizationaction").(string),
		Dnsvservername:             d.Get("dnsvservername").(string),
		Emailhome:                  d.Get("emailhome").(string),
		Encryptcsecexp:             d.Get("encryptcsecexp").(string),
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
		Icasessiontimeout:          d.Get("icasessiontimeout").(string),
		Icauseraccounting:          d.Get("icauseraccounting").(string),
		Iconwithreceiver:           d.Get("iconwithreceiver").(string),
		Iipdnssuffix:               d.Get("iipdnssuffix").(string),
		Kcdaccount:                 d.Get("kcdaccount").(string),
		Killconnections:            d.Get("killconnections").(string),
		Linuxpluginupgrade:         d.Get("linuxpluginupgrade").(string),
		Locallanaccess:             d.Get("locallanaccess").(string),
		Loginscript:                d.Get("loginscript").(string),
		Logoutscript:               d.Get("logoutscript").(string),
		Macpluginupgrade:           d.Get("macpluginupgrade").(string),
		Mdxtokentimeout:            d.Get("mdxtokentimeout").(int),
		Netmask:                    d.Get("netmask").(string),
		Ntdomain:                   d.Get("ntdomain").(string),
		Pcoipprofilename:           d.Get("pcoipprofilename").(string),
		Proxy:                      d.Get("proxy").(string),
		Proxyexception:             d.Get("proxyexception").(string),
		Proxylocalbypass:           d.Get("proxylocalbypass").(string),
		Rdpclientprofilename:       d.Get("rdpclientprofilename").(string),
		Rfc1918:                    d.Get("rfc1918").(string),
		Samesite:                   d.Get("samesite").(string),
		Securebrowse:               d.Get("securebrowse").(string),
		Sesstimeout:                d.Get("sesstimeout").(int),
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
		Uitheme:                    d.Get("uitheme").(string),
		Useiip:                     d.Get("useiip").(string),
		Usemip:                     d.Get("usemip").(string),
		Userdomains:                d.Get("userdomains").(string),
		Wihome:                     d.Get("wihome").(string),
		Wihomeaddresstype:          d.Get("wihomeaddresstype").(string),
		Windowsautologon:           d.Get("windowsautologon").(string),
		Windowsclienttype:          d.Get("windowsclienttype").(string),
		Windowspluginupgrade:       d.Get("windowspluginupgrade").(string),
		Winsip:                     d.Get("winsip").(string),
		Wiportalmode:               d.Get("wiportalmode").(string),
	}

	err := client.UpdateUnnamedResource(service.Vpnparameter.Type(), &vpnparameter)
	if err != nil {
		return diag.Errorf("Error updating vpnparameter")
	}

	d.SetId(vpnparameterName)

	return readVpnparameterFunc(ctx, d, meta)
}

func readVpnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] citrixadc-provider:  In readVpnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	log.Printf("[DEBUG] citrixadc-provider: Reading vpnparameter state")
	data, err := client.FindResource(service.Vpnparameter.Type(), "")
	if err != nil {
		log.Printf("[WARN] citrixadc-provider: Clearing vpnparameter state")
		d.SetId("")
		return nil
	}
	d.Set("advancedclientlessvpnmode", data["advancedclientlessvpnmode"])
	d.Set("allowedlogingroups", data["allowedlogingroups"])
	d.Set("allprotocolproxy", data["allprotocolproxy"])
	d.Set("alwaysonprofilename", data["alwaysonprofilename"])
	setToInt("apptokentimeout", d, data["apptokentimeout"])
	d.Set("authorizationgroup", data["authorizationgroup"])
	d.Set("autoproxyurl", data["autoproxyurl"])
	d.Set("backendcertvalidation", data["backendcertvalidation"])
	d.Set("backendserversni", data["backendserversni"])
	d.Set("citrixreceiverhome", data["citrixreceiverhome"])
	d.Set("clientchoices", data["clientchoices"])
	d.Set("clientcleanupprompt", data["clientcleanupprompt"])
	d.Set("clientconfiguration", data["clientconfiguration"])
	d.Set("clientdebug", data["clientdebug"])
	setToInt("clientidletimeout", d, data["clientidletimeout"])
	d.Set("clientlessmodeurlencoding", data["clientlessmodeurlencoding"])
	d.Set("clientlesspersistentcookie", data["clientlesspersistentcookie"])
	d.Set("clientlessvpnmode", data["clientlessvpnmode"])
	d.Set("clientoptions", data["clientoptions"])
	d.Set("clientsecurity", data["clientsecurity"])
	d.Set("clientsecuritygroup", data["clientsecuritygroup"])
	d.Set("clientsecuritylog", data["clientsecuritylog"])
	d.Set("clientsecuritymessage", data["clientsecuritymessage"])
	d.Set("clientversions", data["clientversions"])
	d.Set("defaultauthorizationaction", data["defaultauthorizationaction"])
	d.Set("dnsvservername", data["dnsvservername"])
	d.Set("emailhome", data["emailhome"])
	d.Set("encryptcsecexp", data["encryptcsecexp"])
	d.Set("epaclienttype", data["epaclienttype"])
	d.Set("forcecleanup", data["forcecleanup"])
	setToInt("forcedtimeout", d, data["forcedtimeout"])
	setToInt("forcedtimeoutwarning", d, data["forcedtimeoutwarning"])
	d.Set("fqdnspoofedip", data["fqdnspoofedip"])
	d.Set("ftpproxy", data["ftpproxy"])
	d.Set("gopherproxy", data["gopherproxy"])
	d.Set("homepage", data["homepage"])
	d.Set("httpproxy", data["httpproxy"])
	d.Set("icaproxy", data["icaproxy"])
	d.Set("icasessiontimeout", data["icasessiontimeout"])
	d.Set("icauseraccounting", data["icauseraccounting"])
	d.Set("iconwithreceiver", data["iconwithreceiver"])
	d.Set("iipdnssuffix", data["iipdnssuffix"])
	d.Set("kcdaccount", data["kcdaccount"])
	d.Set("killconnections", data["killconnections"])
	d.Set("linuxpluginupgrade", data["linuxpluginupgrade"])
	d.Set("locallanaccess", data["locallanaccess"])
	d.Set("loginscript", data["loginscript"])
	d.Set("logoutscript", data["logoutscript"])
	d.Set("macpluginupgrade", data["macpluginupgrade"])
	setToInt("mdxtokentimeout", d, data["mdxtokentimeout"])
	d.Set("netmask", data["netmask"])
	d.Set("ntdomain", data["ntdomain"])
	d.Set("pcoipprofilename", data["pcoipprofilename"])
	d.Set("proxy", data["proxy"])
	d.Set("proxyexception", data["proxyexception"])
	d.Set("proxylocalbypass", data["proxylocalbypass"])
	d.Set("rdpclientprofilename", data["rdpclientprofilename"])
	d.Set("rfc1918", data["rfc1918"])
	d.Set("samesite", data["samesite"])
	d.Set("securebrowse", data["securebrowse"])
	setToInt("sesstimeout", d, data["sesstimeout"])
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
	d.Set("uitheme", data["uitheme"])
	d.Set("useiip", data["useiip"])
	d.Set("usemip", data["usemip"])
	d.Set("userdomains", data["userdomains"])
	d.Set("wihome", data["wihome"])
	d.Set("wihomeaddresstype", data["wihomeaddresstype"])
	d.Set("windowsautologon", data["windowsautologon"])
	d.Set("windowsclienttype", data["windowsclienttype"])
	d.Set("windowspluginupgrade", data["windowspluginupgrade"])
	d.Set("winsip", data["winsip"])
	d.Set("wiportalmode", data["wiportalmode"])
	// Convert httpport from []string to []int before setting
	if httpPort, ok := data["httpport"]; ok && httpPort != nil {
		d.Set("httpport", stringListToIntList(httpPort.([]interface{})))
	}

	return nil

}

func updateVpnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In updateVpnparameterFunc")
	client := meta.(*NetScalerNitroClient).client
	vpnparameter := vpn.Vpnparameter{}
	hasChange := false
	if d.HasChange("advancedclientlessvpnmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Advancedclientlessvpnmode has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Advancedclientlessvpnmode = d.Get("advancedclientlessvpnmode").(string)
		hasChange = true
	}
	if d.HasChange("allowedlogingroups") {
		log.Printf("[DEBUG]  citrixadc-provider: Allowedlogingroups has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Allowedlogingroups = d.Get("allowedlogingroups").(string)
		hasChange = true
	}
	if d.HasChange("allprotocolproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Allprotocolproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Allprotocolproxy = d.Get("allprotocolproxy").(string)
		hasChange = true
	}
	if d.HasChange("alwaysonprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Alwaysonprofilename has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Alwaysonprofilename = d.Get("alwaysonprofilename").(string)
		hasChange = true
	}
	if d.HasChange("apptokentimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Apptokentimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Apptokentimeout = d.Get("apptokentimeout").(int)
		hasChange = true
	}
	if d.HasChange("authorizationgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Authorizationgroup has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Authorizationgroup = d.Get("authorizationgroup").(string)
		hasChange = true
	}
	if d.HasChange("autoproxyurl") {
		log.Printf("[DEBUG]  citrixadc-provider: Autoproxyurl has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Autoproxyurl = d.Get("autoproxyurl").(string)
		hasChange = true
	}
	if d.HasChange("backendcertvalidation") {
		log.Printf("[DEBUG]  citrixadc-provider: Backendcertvalidation has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Backendcertvalidation = d.Get("backendcertvalidation").(string)
		hasChange = true
	}
	if d.HasChange("backendserversni") {
		log.Printf("[DEBUG]  citrixadc-provider: Backendserversni has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Backendserversni = d.Get("backendserversni").(string)
		hasChange = true
	}
	if d.HasChange("citrixreceiverhome") {
		log.Printf("[DEBUG]  citrixadc-provider: Citrixreceiverhome has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Citrixreceiverhome = d.Get("citrixreceiverhome").(string)
		hasChange = true
	}
	if d.HasChange("clientchoices") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientchoices has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientchoices = d.Get("clientchoices").(string)
		hasChange = true
	}
	if d.HasChange("clientcleanupprompt") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientcleanupprompt has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientcleanupprompt = d.Get("clientcleanupprompt").(string)
		hasChange = true
	}
	if d.HasChange("clientconfiguration") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientconfiguration has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientconfiguration = toStringList(d.Get("clientconfiguration").([]interface{}))
		hasChange = true
	}
	if d.HasChange("clientdebug") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientdebug has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientdebug = d.Get("clientdebug").(string)
		hasChange = true
	}
	if d.HasChange("clientidletimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientidletimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientidletimeout = d.Get("clientidletimeout").(int)
		hasChange = true
	}
	if d.HasChange("clientlessmodeurlencoding") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessmodeurlencoding has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientlessmodeurlencoding = d.Get("clientlessmodeurlencoding").(string)
		hasChange = true
	}
	if d.HasChange("clientlesspersistentcookie") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlesspersistentcookie has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientlesspersistentcookie = d.Get("clientlesspersistentcookie").(string)
		hasChange = true
	}
	if d.HasChange("clientlessvpnmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientlessvpnmode has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientlessvpnmode = d.Get("clientlessvpnmode").(string)
		hasChange = true
	}
	if d.HasChange("clientoptions") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientoptions has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientoptions = toStringList(d.Get("clientoptions").([]interface{}))
		hasChange = true
	}
	if d.HasChange("clientsecurity") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecurity has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientsecurity = d.Get("clientsecurity").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritygroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritygroup has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientsecuritygroup = d.Get("clientsecuritygroup").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritylog") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritylog has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientsecuritylog = d.Get("clientsecuritylog").(string)
		hasChange = true
	}
	if d.HasChange("clientsecuritymessage") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientsecuritymessage has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientsecuritymessage = d.Get("clientsecuritymessage").(string)
		hasChange = true
	}
	if d.HasChange("clientversions") {
		log.Printf("[DEBUG]  citrixadc-provider: Clientversions has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Clientversions = d.Get("clientversions").(string)
		hasChange = true
	}
	if d.HasChange("defaultauthorizationaction") {
		log.Printf("[DEBUG]  citrixadc-provider: Defaultauthorizationaction has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Defaultauthorizationaction = d.Get("defaultauthorizationaction").(string)
		hasChange = true
	}
	if d.HasChange("dnsvservername") {
		log.Printf("[DEBUG]  citrixadc-provider: Dnsvservername has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Dnsvservername = d.Get("dnsvservername").(string)
		hasChange = true
	}
	if d.HasChange("emailhome") {
		log.Printf("[DEBUG]  citrixadc-provider: Emailhome has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Emailhome = d.Get("emailhome").(string)
		hasChange = true
	}
	if d.HasChange("encryptcsecexp") {
		log.Printf("[DEBUG]  citrixadc-provider: Encryptcsecexp has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Encryptcsecexp = d.Get("encryptcsecexp").(string)
		hasChange = true
	}
	if d.HasChange("epaclienttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Epaclienttype has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Epaclienttype = d.Get("epaclienttype").(string)
		hasChange = true
	}
	if d.HasChange("forcecleanup") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcecleanup has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Forcecleanup = toStringList(d.Get("forcecleanup").([]interface{}))
		hasChange = true
	}
	if d.HasChange("forcedtimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Forcedtimeout = d.Get("forcedtimeout").(int)
		hasChange = true
	}
	if d.HasChange("forcedtimeoutwarning") {
		log.Printf("[DEBUG]  citrixadc-provider: Forcedtimeoutwarning has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Forcedtimeoutwarning = d.Get("forcedtimeoutwarning").(int)
		hasChange = true
	}
	if d.HasChange("fqdnspoofedip") {
		log.Printf("[DEBUG]  citrixadc-provider: Fqdnspoofedip has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Fqdnspoofedip = d.Get("fqdnspoofedip").(string)
		hasChange = true
	}
	if d.HasChange("ftpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Ftpproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Ftpproxy = d.Get("ftpproxy").(string)
		hasChange = true
	}
	if d.HasChange("gopherproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Gopherproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Gopherproxy = d.Get("gopherproxy").(string)
		hasChange = true
	}
	if d.HasChange("homepage") {
		log.Printf("[DEBUG]  citrixadc-provider: Homepage has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Homepage = d.Get("homepage").(string)
		hasChange = true
	}
	if d.HasChange("httpport") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpport has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Httpport = toIntegerList(d.Get("httpport").([]interface{}))
		hasChange = true
	}
	if d.HasChange("httpproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Httpproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Httpproxy = d.Get("httpproxy").(string)
		hasChange = true
	}
	if d.HasChange("icaproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Icaproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Icaproxy = d.Get("icaproxy").(string)
		hasChange = true
	}
	if d.HasChange("icasessiontimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Icasessiontimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Icasessiontimeout = d.Get("icasessiontimeout").(string)
		hasChange = true
	}
	if d.HasChange("icauseraccounting") {
		log.Printf("[DEBUG]  citrixadc-provider: Icauseraccounting has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Icauseraccounting = d.Get("icauseraccounting").(string)
		hasChange = true
	}
	if d.HasChange("iconwithreceiver") {
		log.Printf("[DEBUG]  citrixadc-provider: Iconwithreceiver has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Iconwithreceiver = d.Get("iconwithreceiver").(string)
		hasChange = true
	}
	if d.HasChange("iipdnssuffix") {
		log.Printf("[DEBUG]  citrixadc-provider: Iipdnssuffix has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Iipdnssuffix = d.Get("iipdnssuffix").(string)
		hasChange = true
	}
	if d.HasChange("kcdaccount") {
		log.Printf("[DEBUG]  citrixadc-provider: Kcdaccount has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Kcdaccount = d.Get("kcdaccount").(string)
		hasChange = true
	}
	if d.HasChange("killconnections") {
		log.Printf("[DEBUG]  citrixadc-provider: Killconnections has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Killconnections = d.Get("killconnections").(string)
		hasChange = true
	}
	if d.HasChange("linuxpluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Linuxpluginupgrade has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Linuxpluginupgrade = d.Get("linuxpluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("locallanaccess") {
		log.Printf("[DEBUG]  citrixadc-provider: Locallanaccess has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Locallanaccess = d.Get("locallanaccess").(string)
		hasChange = true
	}
	if d.HasChange("loginscript") {
		log.Printf("[DEBUG]  citrixadc-provider: Loginscript has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Loginscript = d.Get("loginscript").(string)
		hasChange = true
	}
	if d.HasChange("logoutscript") {
		log.Printf("[DEBUG]  citrixadc-provider: Logoutscript has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Logoutscript = d.Get("logoutscript").(string)
		hasChange = true
	}
	if d.HasChange("macpluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Macpluginupgrade has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Macpluginupgrade = d.Get("macpluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("mdxtokentimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Mdxtokentimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Mdxtokentimeout = d.Get("mdxtokentimeout").(int)
		hasChange = true
	}
	if d.HasChange("netmask") {
		log.Printf("[DEBUG]  citrixadc-provider: Netmask has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Netmask = d.Get("netmask").(string)
		hasChange = true
	}
	if d.HasChange("ntdomain") {
		log.Printf("[DEBUG]  citrixadc-provider: Ntdomain has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Ntdomain = d.Get("ntdomain").(string)
		hasChange = true
	}
	if d.HasChange("pcoipprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Pcoipprofilename has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Pcoipprofilename = d.Get("pcoipprofilename").(string)
		hasChange = true
	}
	if d.HasChange("proxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Proxy = d.Get("proxy").(string)
		hasChange = true
	}
	if d.HasChange("proxyexception") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxyexception has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Proxyexception = d.Get("proxyexception").(string)
		hasChange = true
	}
	if d.HasChange("proxylocalbypass") {
		log.Printf("[DEBUG]  citrixadc-provider: Proxylocalbypass has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Proxylocalbypass = d.Get("proxylocalbypass").(string)
		hasChange = true
	}
	if d.HasChange("rdpclientprofilename") {
		log.Printf("[DEBUG]  citrixadc-provider: Rdpclientprofilename has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Rdpclientprofilename = d.Get("rdpclientprofilename").(string)
		hasChange = true
	}
	if d.HasChange("rfc1918") {
		log.Printf("[DEBUG]  citrixadc-provider: Rfc1918 has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Rfc1918 = d.Get("rfc1918").(string)
		hasChange = true
	}
	if d.HasChange("samesite") {
		log.Printf("[DEBUG]  citrixadc-provider: Samesite has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Samesite = d.Get("samesite").(string)
		hasChange = true
	}
	if d.HasChange("securebrowse") {
		log.Printf("[DEBUG]  citrixadc-provider: Securebrowse has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Securebrowse = d.Get("securebrowse").(string)
		hasChange = true
	}
	if d.HasChange("sesstimeout") {
		log.Printf("[DEBUG]  citrixadc-provider: Sesstimeout has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Sesstimeout = d.Get("sesstimeout").(int)
		hasChange = true
	}
	if d.HasChange("smartgroup") {
		log.Printf("[DEBUG]  citrixadc-provider: Smartgroup has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Smartgroup = d.Get("smartgroup").(string)
		hasChange = true
	}
	if d.HasChange("socksproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Socksproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Socksproxy = d.Get("socksproxy").(string)
		hasChange = true
	}
	if d.HasChange("splitdns") {
		log.Printf("[DEBUG]  citrixadc-provider: Splitdns has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Splitdns = d.Get("splitdns").(string)
		hasChange = true
	}
	if d.HasChange("splittunnel") {
		log.Printf("[DEBUG]  citrixadc-provider: Splittunnel has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Splittunnel = d.Get("splittunnel").(string)
		hasChange = true
	}
	if d.HasChange("spoofiip") {
		log.Printf("[DEBUG]  citrixadc-provider: Spoofiip has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Spoofiip = d.Get("spoofiip").(string)
		hasChange = true
	}
	if d.HasChange("sslproxy") {
		log.Printf("[DEBUG]  citrixadc-provider: Sslproxy has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Sslproxy = d.Get("sslproxy").(string)
		hasChange = true
	}
	if d.HasChange("sso") {
		log.Printf("[DEBUG]  citrixadc-provider: Sso has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Sso = d.Get("sso").(string)
		hasChange = true
	}
	if d.HasChange("ssocredential") {
		log.Printf("[DEBUG]  citrixadc-provider: Ssocredential has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Ssocredential = d.Get("ssocredential").(string)
		hasChange = true
	}
	if d.HasChange("storefronturl") {
		log.Printf("[DEBUG]  citrixadc-provider: Storefronturl has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Storefronturl = d.Get("storefronturl").(string)
		hasChange = true
	}
	if d.HasChange("transparentinterception") {
		log.Printf("[DEBUG]  citrixadc-provider: Transparentinterception has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Transparentinterception = d.Get("transparentinterception").(string)
		hasChange = true
	}
	if d.HasChange("uitheme") {
		log.Printf("[DEBUG]  citrixadc-provider: Uitheme has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Uitheme = d.Get("uitheme").(string)
		hasChange = true
	}
	if d.HasChange("useiip") {
		log.Printf("[DEBUG]  citrixadc-provider: Useiip has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Useiip = d.Get("useiip").(string)
		hasChange = true
	}
	if d.HasChange("usemip") {
		log.Printf("[DEBUG]  citrixadc-provider: Usemip has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Usemip = d.Get("usemip").(string)
		hasChange = true
	}
	if d.HasChange("userdomains") {
		log.Printf("[DEBUG]  citrixadc-provider: Userdomains has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Userdomains = d.Get("userdomains").(string)
		hasChange = true
	}
	if d.HasChange("wihome") {
		log.Printf("[DEBUG]  citrixadc-provider: Wihome has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Wihome = d.Get("wihome").(string)
		hasChange = true
	}
	if d.HasChange("wihomeaddresstype") {
		log.Printf("[DEBUG]  citrixadc-provider: Wihomeaddresstype has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Wihomeaddresstype = d.Get("wihomeaddresstype").(string)
		hasChange = true
	}
	if d.HasChange("windowsautologon") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowsautologon has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Windowsautologon = d.Get("windowsautologon").(string)
		hasChange = true
	}
	if d.HasChange("windowsclienttype") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowsclienttype has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Windowsclienttype = d.Get("windowsclienttype").(string)
		hasChange = true
	}
	if d.HasChange("windowspluginupgrade") {
		log.Printf("[DEBUG]  citrixadc-provider: Windowspluginupgrade has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Windowspluginupgrade = d.Get("windowspluginupgrade").(string)
		hasChange = true
	}
	if d.HasChange("winsip") {
		log.Printf("[DEBUG]  citrixadc-provider: Winsip has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Winsip = d.Get("winsip").(string)
		hasChange = true
	}
	if d.HasChange("wiportalmode") {
		log.Printf("[DEBUG]  citrixadc-provider: Wiportalmode has changed for vpnparameter vpnparameter, starting update")
		vpnparameter.Wiportalmode = d.Get("wiportalmode").(string)
		hasChange = true
	}

	if hasChange {
		err := client.UpdateUnnamedResource(service.Vpnparameter.Type(), &vpnparameter)
		if err != nil {
			return diag.Errorf("Error updating vpnparameter %s", err.Error())
		}
	}
	return readVpnparameterFunc(ctx, d, meta)
}

func deleteVpnparameterFunc(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[DEBUG]  citrixadc-provider: In deleteVpnparameterFunc")
	// vpnparameter do not have DELETE operation, but this function is required to set the ID to ""
	d.SetId("")

	return nil
}
