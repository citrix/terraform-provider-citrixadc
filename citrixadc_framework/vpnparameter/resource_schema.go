package vpnparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnparameterResourceModel describes the resource data model.
type VpnparameterResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Accessrestrictedpageredirect types.String `tfsdk:"accessrestrictedpageredirect"`
	Advancedclientlessvpnmode    types.String `tfsdk:"advancedclientlessvpnmode"`
	Allowedlogingroups           types.String `tfsdk:"allowedlogingroups"`
	Allprotocolproxy             types.String `tfsdk:"allprotocolproxy"`
	Alwaysonprofilename          types.String `tfsdk:"alwaysonprofilename"`
	Apptokentimeout              types.Int64  `tfsdk:"apptokentimeout"`
	Authorizationgroup           types.String `tfsdk:"authorizationgroup"`
	Autoproxyurl                 types.String `tfsdk:"autoproxyurl"`
	Backendcertvalidation        types.String `tfsdk:"backendcertvalidation"`
	Backenddtls12                types.String `tfsdk:"backenddtls12"`
	Backendserversni             types.String `tfsdk:"backendserversni"`
	Citrixreceiverhome           types.String `tfsdk:"citrixreceiverhome"`
	Clientchoices                types.String `tfsdk:"clientchoices"`
	Clientcleanupprompt          types.String `tfsdk:"clientcleanupprompt"`
	Clientconfiguration          types.List   `tfsdk:"clientconfiguration"`
	Clientdebug                  types.String `tfsdk:"clientdebug"`
	Clientidletimeout            types.Int64  `tfsdk:"clientidletimeout"`
	Clientlessmodeurlencoding    types.String `tfsdk:"clientlessmodeurlencoding"`
	Clientlesspersistentcookie   types.String `tfsdk:"clientlesspersistentcookie"`
	Clientlessvpnmode            types.String `tfsdk:"clientlessvpnmode"`
	Clientoptions                types.List   `tfsdk:"clientoptions"`
	Clientsecurity               types.String `tfsdk:"clientsecurity"`
	Clientsecuritygroup          types.String `tfsdk:"clientsecuritygroup"`
	Clientsecuritylog            types.String `tfsdk:"clientsecuritylog"`
	Clientsecuritymessage        types.String `tfsdk:"clientsecuritymessage"`
	Clientversions               types.String `tfsdk:"clientversions"`
	Defaultauthorizationaction   types.String `tfsdk:"defaultauthorizationaction"`
	Deviceposture                types.String `tfsdk:"deviceposture"`
	Dnsvservername               types.String `tfsdk:"dnsvservername"`
	Emailhome                    types.String `tfsdk:"emailhome"`
	Encryptcsecexp               types.String `tfsdk:"encryptcsecexp"`
	Epaclienttype                types.String `tfsdk:"epaclienttype"`
	Forcecleanup                 types.List   `tfsdk:"forcecleanup"`
	Forcedtimeout                types.Int64  `tfsdk:"forcedtimeout"`
	Forcedtimeoutwarning         types.Int64  `tfsdk:"forcedtimeoutwarning"`
	Fqdnspoofedip                types.String `tfsdk:"fqdnspoofedip"`
	Ftpproxy                     types.String `tfsdk:"ftpproxy"`
	Gopherproxy                  types.String `tfsdk:"gopherproxy"`
	Homepage                     types.String `tfsdk:"homepage"`
	Httpport                     types.List   `tfsdk:"httpport"`
	Httpproxy                    types.String `tfsdk:"httpproxy"`
	Httptrackconnproxy           types.String `tfsdk:"httptrackconnproxy"`
	Icaproxy                     types.String `tfsdk:"icaproxy"`
	Icasessiontimeout            types.String `tfsdk:"icasessiontimeout"`
	Icauseraccounting            types.String `tfsdk:"icauseraccounting"`
	Iconwithreceiver             types.String `tfsdk:"iconwithreceiver"`
	Iipdnssuffix                 types.String `tfsdk:"iipdnssuffix"`
	Kcdaccount                   types.String `tfsdk:"kcdaccount"`
	Killconnections              types.String `tfsdk:"killconnections"`
	Linuxpluginupgrade           types.String `tfsdk:"linuxpluginupgrade"`
	Locallanaccess               types.String `tfsdk:"locallanaccess"`
	Loginscript                  types.String `tfsdk:"loginscript"`
	Logoutscript                 types.String `tfsdk:"logoutscript"`
	Macpluginupgrade             types.String `tfsdk:"macpluginupgrade"`
	Maxiipperuser                types.Int64  `tfsdk:"maxiipperuser"`
	Mdxtokentimeout              types.Int64  `tfsdk:"mdxtokentimeout"`
	Netmask                      types.String `tfsdk:"netmask"`
	Ntdomain                     types.String `tfsdk:"ntdomain"`
	Pcoipprofilename             types.String `tfsdk:"pcoipprofilename"`
	Proxy                        types.String `tfsdk:"proxy"`
	Proxyexception               types.String `tfsdk:"proxyexception"`
	Proxylocalbypass             types.String `tfsdk:"proxylocalbypass"`
	Rdpclientprofilename         types.String `tfsdk:"rdpclientprofilename"`
	Rfc1918                      types.String `tfsdk:"rfc1918"`
	Samesite                     types.String `tfsdk:"samesite"`
	Securebrowse                 types.String `tfsdk:"securebrowse"`
	Secureprivateaccess          types.String `tfsdk:"secureprivateaccess"`
	Sesstimeout                  types.Int64  `tfsdk:"sesstimeout"`
	Smartgroup                   types.String `tfsdk:"smartgroup"`
	Socksproxy                   types.String `tfsdk:"socksproxy"`
	Splitdns                     types.String `tfsdk:"splitdns"`
	Splittunnel                  types.String `tfsdk:"splittunnel"`
	Spoofiip                     types.String `tfsdk:"spoofiip"`
	Sslproxy                     types.String `tfsdk:"sslproxy"`
	Sso                          types.String `tfsdk:"sso"`
	Ssocredential                types.String `tfsdk:"ssocredential"`
	Storefronturl                types.String `tfsdk:"storefronturl"`
	Transparentinterception      types.String `tfsdk:"transparentinterception"`
	Uitheme                      types.String `tfsdk:"uitheme"`
	Useiip                       types.String `tfsdk:"useiip"`
	Usemip                       types.String `tfsdk:"usemip"`
	Userdomains                  types.String `tfsdk:"userdomains"`
	Wihome                       types.String `tfsdk:"wihome"`
	Wihomeaddresstype            types.String `tfsdk:"wihomeaddresstype"`
	Windowsautologon             types.String `tfsdk:"windowsautologon"`
	Windowsclienttype            types.String `tfsdk:"windowsclienttype"`
	Windowspluginupgrade         types.String `tfsdk:"windowspluginupgrade"`
	Winsip                       types.String `tfsdk:"winsip"`
	Wiportalmode                 types.String `tfsdk:"wiportalmode"`
}

func (r *VpnparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnparameter resource.",
			},
			"accessrestrictedpageredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CDN"),
				Description: "By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.",
			},
			"advancedclientlessvpnmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Option to enable/disable Advanced ClientlessVpnMode. Additionaly, it can be set to STRICT to block Classic ClientlessVpnMode while in AdvancedClientlessMode.",
			},
			"allowedlogingroups": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify groups that have permission to log on to Citrix Gateway. Users who do not belong to this group or groups are denied access even if they have valid credentials.",
			},
			"allprotocolproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to use for all protocols supported by Citrix Gateway.",
			},
			"alwaysonprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AlwaysON profile. The builtin profile named none can be used to explicitly disable AlwaysON.",
			},
			"apptokentimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100),
				Description: "The timeout value in seconds for tokens to access XenMobile applications",
			},
			"authorizationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of groups in which the user is placed when none of the groups that the user is a part of is configured on Citrix Gateway. The authorization policy can be bound to these groups to control access to the resources.",
			},
			"autoproxyurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to auto proxy config file",
			},
			"backendcertvalidation": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "enables backend server certificate validation",
			},
			"backenddtls12": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enables DTLS 1.2 for backend server handshakes",
			},
			"backendserversni": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "enables sni extension for backend server handshakes",
			},
			"citrixreceiverhome": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address for the Citrix Receiver home page. Configure Citrix Gateway so that when users log on to the appliance, the Citrix Gateway Plug-in opens a web browser that allows single sign-on to the Citrix Receiver home page.",
			},
			"clientchoices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Provide users with multiple logon options. With client choices, users have the option of logging on by using the Citrix Gateway Plug-in for Windows, Citrix Gateway Plug-in for Java, the Web Interface, or clientless access from one location. Depending on how Citrix Gateway is configured, users are presented with up to three icons for logon choices. The most common are the Citrix Gateway Plug-in for Windows, Web Interface, and clientless access.",
			},
			"clientcleanupprompt": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Prompt for client-side cache clean-up when a client-initiated session closes.",
			},
			"clientconfiguration": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Allow users to change client Debug logging level in Configuration tab of the Citrix Gateway Plug-in for Windows.",
			},
			"clientdebug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the trace level on Citrix Gateway. Technical support technicians use these debug logs for in-depth debugging and troubleshooting purposes. Available settings function as follows:\n* DEBUG - Detailed debug messages are collected and written into the specified file.\n* STATS - Application audit level error messages and debug statistic counters are written into the specified file.\n* EVENTS - Application audit-level error messages are written into the specified file.\n* OFF - Only critical events are logged into the Windows Application Log.",
			},
			"clientidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in minutes, after which to time out the user session if Citrix Gateway does not detect mouse or keyboard activity.",
			},
			"clientlessmodeurlencoding": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("OPAQUE"),
				Description: "When clientless access is enabled, you can choose to encode the addresses of internal web applications or to leave the address as clear text. Available settings function as follows:\n* OPAQUE - Use standard encoding mechanisms to make the domain and protocol part of the resource unclear to users.\n* TRANSPARENT - Do not encode the web address and make it visible to users.\n* ENCRYPT - Allow the domain and protocol to be encrypted using a session key. When the web address is encrypted, the URL is different for each user session for the same web resource. If users bookmark the encoded web address, save it in the web browser and then log off, they cannot connect to the web address when they log on and use the bookmark. If users save the encrypted bookmark in the Access Interface during their session, the bookmark works each time the user logs on.",
			},
			"clientlesspersistentcookie": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DENY"),
				Description: "State of persistent cookies in clientless access mode. Persistent cookies are required for accessing certain features of SharePoint, such as opening and editing Microsoft Word, Excel, and PowerPoint documents hosted on the SharePoint server. A persistent cookie remains on the user device and is sent with each HTTP request. Citrix Gateway encrypts the persistent cookie before sending it to the plug-in on the user device, and refreshes the cookie periodically as long as the session exists. The cookie becomes stale if the session ends. Available settings function as follows:\n* ALLOW - Enable persistent cookies. Users can open and edit Microsoft documents stored in SharePoint.\n* DENY - Disable persistent cookies. Users cannot open and edit Microsoft documents stored in SharePoint.\n* PROMPT - Prompt users to allow or deny persistent cookies during the session. Persistent cookies are not required for clientless access if users do not connect to SharePoint.",
			},
			"clientlessvpnmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable clientless access for web, XenApp or XenDesktop, and FileShare resources without installing the Citrix Gateway Plug-in. Available settings function as follows:\n* ON - Allow only clientless access.\n* OFF - Allow clientless access after users log on with the Citrix Gateway Plug-in.\n* DISABLED - Do not allow clientless access.",
			},
			"clientoptions": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Display only the configured menu options when you select the \"Configure Citrix Gateway\" option in the Citrix Gateway Plug-in's system tray icon for Windows.",
			},
			"clientsecurity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the client security check for the user device to permit a Citrix Gateway session. The web address or IP address is not included in the expression for the client security check.",
			},
			"clientsecuritygroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The client security group that will be assigned on failure of the client security check. Users can in general be organized into Groups. In this case, the Client Security Group may have a more restrictive security policy.",
			},
			"clientsecuritylog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether or not to display all failed Client Security scans to the end user",
			},
			"clientsecuritymessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The client security message that will be displayed on failure of the client security check.",
			},
			"clientversions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "checkversion api",
			},
			"defaultauthorizationaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DENY"),
				Description: "Specify the network resources that users have access to when they log on to the internal network. The default setting for authorization is to deny access to all network resources. Citrix recommends using the default global setting and then creating authorization policies to define the network resources users can access. If you set the default authorization policy to DENY, you must explicitly authorize access to any network resource, which improves security.",
			},
			"deviceposture": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable device posture",
			},
			"dnsvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS virtual server for the user session.",
			},
			"emailhome": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address for the web-based email, such as Outlook Web Access.",
			},
			"encryptcsecexp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable encryption of client security expressions.",
			},
			"epaclienttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Choose between two types of End point Windows Client\na) Application Agent - which always runs in the task bar as a standalone application and also has a supporting service which runs permanently when installed\nb) Activex Control - ActiveX control run by Microsoft Internet Explorer.",
			},
			"forcecleanup": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Force cache clean-up when the user closes a session. You can specify all, none, or any combination of the client-side items.",
			},
			"forcedtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Force a disconnection from the Citrix Gateway Plug-in with Citrix Gateway after a specified number of minutes. If the session closes, the user must log on again.",
			},
			"forcedtimeoutwarning": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes to warn a user before the user session is disconnected.",
			},
			"fqdnspoofedip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Spoofed IP address range that can be used by client for FQDN based split tunneling",
			},
			"ftpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to be used for FTP access for all subsequent connections to the internal network.",
			},
			"gopherproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to be used for GOPHER access for all subsequent connections to the internal network.",
			},
			"homepage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address of the home page that appears when users log on. Otherwise, users receive the default home page for Citrix Gateway, which is the Access Interface.",
			},
			"httpport": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Destination port numbers other than port 80, added as a comma-separated list. Traffic to these ports is processed as HTTP traffic, which allows functionality, such as HTTP authorization and single sign-on to a web application to work.",
			},
			"httpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to be used for HTTP access for all subsequent connections to the internal network.",
			},
			"httptrackconnproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable HTTP tracking for packets proxied via vpn vserver using GSLB connection proxy feature.",
			},
			"icaproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable ICA proxy to configure secure Internet access to servers running Citrix XenApp or XenDesktop by using Citrix Receiver instead of the Citrix Gateway Plug-in.",
			},
			"icasessiontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable ica session timeout. If enabled and in case AAA session gets terminated, ICA connections associated with that will also get terminated",
			},
			"icauseraccounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the radiusPolicy to use for RADIUS user accounting info on the session.",
			},
			"iconwithreceiver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to decide whether to show plugin icon along with receiver icon",
			},
			"iipdnssuffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An intranet IP DNS suffix. When a user logs on to Citrix Gateway and is assigned an IP address, a DNS record for the user name and IP address combination is added to the Citrix Gateway DNS cache. You can configure a DNS suffix to append to the user name when the DNS record is added to the cache. You can reach to the host from where the user is logged on by using the user's name, which can be easier to remember than an IP address. When the user logs off from Citrix Gateway, the record is removed from the DNS cache.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The KCD account details to be used in SSO",
			},
			"killconnections": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether the Citrix Gateway Plug-in should disconnect all preexisting connections, such as the connections existing before the user logged on to Citrix Gateway, and prevent new incoming connections on the Citrix Gateway Plug-in for Windows and MAC when the user is connected to Citrix Gateway and split tunneling is disabled.",
			},
			"linuxpluginupgrade": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Always"),
				Description: "Option to set plugin upgrade behaviour for Linux",
			},
			"locallanaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set local LAN access. If split tunneling is OFF, and you set local LAN access to ON, the local client can route traffic to its local interface. When the local area network switch is specified, this combination of switches is useful. The client can allow local LAN access to devices that commonly have non-routable addresses, such as local printers or local file servers.",
			},
			"loginscript": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the logon script that is run when a session is established. Separate multiple scripts by using comma. A \"$\" in the path signifies that the word following the \"$\" is an environment variable.",
			},
			"logoutscript": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path to the logout script. Separate multiple scripts by using comma. A \"$\" in the path signifies that the word following the \"$\" is an environment variable.",
			},
			"macpluginupgrade": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Always"),
				Description: "Option to set plugin upgrade behaviour for Mac",
			},
			"maxiipperuser": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Maximum number of Intranet IP that can be assigned to a user from AAA group, VPN vserver or VPN global pool. This setting is not applicable for AAA user level Intranet IP configuration",
			},
			"mdxtokentimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Validity of MDX Token in minutes. This token is used for mdx services to access backend and valid  HEAD and GET request.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask for the spoofed ip address",
			},
			"ntdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Single sign-on domain to use for single sign-on to applications in the internal network. This setting can be overwritten by the domain that users specify at the time of logon or by the domain that the authentication server returns.",
			},
			"pcoipprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the PCOIP profile.",
			},
			"proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set options to apply proxy for accessing the internal resources. Available settings function as follows:\n* BROWSER - Proxy settings are configured only in Internet Explorer and Firefox browsers.\n* NS - Proxy settings are configured on the Citrix ADC.\n* OFF - Proxy settings are not configured.",
			},
			"proxyexception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy exception string that will be configured in the browser for bypassing the previously configured proxies. Allowed only if proxy type is Browser.",
			},
			"proxylocalbypass": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Bypass proxy server for local addresses option in Internet Explorer and Firefox proxy server settings.",
			},
			"rdpclientprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the RDP profile associated with the vserver.",
			},
			"rfc1918": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "As defined in the local area network, allow only the following local area network addresses to bypass the VPN tunnel when the local LAN access feature is enabled:\n* 10.*.*.*,\n* 172.16.*.*,\n* 192.168.*.*",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in VPN context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"securebrowse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow users to connect through Citrix Gateway to network resources from iOS and Android mobile devices with Citrix Receiver. Users do not need to establish a full VPN tunnel to access resources in the secure network.",
			},
			"secureprivateaccess": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enables or disables the secure private access configuration.",
			},
			"sesstimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Number of minutes after which the session times out.",
			},
			"smartgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"socksproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to be used for SOCKS access for all subsequent connections to the internal network.",
			},
			"splitdns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Route the DNS requests to the local DNS server configured on the user device, or Citrix Gateway (remote), or both.",
			},
			"splittunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send, through the tunnel, traffic only for intranet applications that are defined in Citrix Gateway. Route all other traffic directly to the Internet. The OFF setting routes all traffic through Citrix Gateway. With the REVERSE setting, intranet applications define the network traffic that is not intercepted. All network traffic directed to internal IP addresses bypasses the VPN tunnel, while other traffic goes through Citrix Gateway. Reverse split tunneling can be used to log all non-local LAN traffic. For example, if users have a home network and are logged on through the Citrix Gateway Plug-in, network traffic destined to a printer or another device within the home network is not intercepted.",
			},
			"spoofiip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Indicate whether or not the application requires IP spoofing, which routes the connection to the intranet application through the virtual adapter.",
			},
			"sslproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server to be used for SSL access for all subsequent connections to the internal network.",
			},
			"sso": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set single sign-on (SSO) for the session. When the user accesses a server, the user's logon credentials are passed to the server for authentication.\n	    NOTE : This configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use VPN TrafficAction to configure SSO for these authentication types.",
			},
			"ssocredential": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PRIMARY"),
				Description: "Specify whether to use the primary or secondary authentication credentials for single sign-on to the server.",
			},
			"storefronturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address for StoreFront to be used in this session for enumeration of resources from XenApp or XenDesktop.",
			},
			"transparentinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow access to network resources by using a single IP address and subnet mask or a range of IP addresses. The OFF setting sets the mode to proxy, in which you configure destination and source IP addresses and port numbers. If you are using the Citrix Gateway Plug-in for Windows, set this parameter to ON, in which the mode is set to transparent. If you are using the Citrix Gateway Plug-in for Java, set this parameter to OFF.",
			},
			"uitheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set VPN UI Theme to Green-Bubble, Caxton or Custom; default is Caxton.",
			},
			"useiip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOSPILLOVER"),
				Description: "Define IP address pool options. Available settings function as follows:\n* SPILLOVER - When an address pool is configured and the mapped IP is used as an intranet IP address, the mapped IP address is used when an intranet IP address cannot be assigned.\n* NOSPILLOVER - When intranet IP addresses are enabled and the mapped IP address is not used, the Transfer Login page appears for users who have used all available intranet IP addresses.\n* OFF - Address pool is not configured.",
			},
			"usemip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NS"),
				Description: "Enable or disable the use of a unique IP address alias, or a mapped IP address, as the client IP address for each client session. Allow Citrix Gateway to use the mapped IP address as an intranet IP address when all other IP addresses are not available.\nWhen IP pooling is configured and the mapped IP is used as an intranet IP address, the mapped IP address is used when an intranet IP address cannot be assigned.",
			},
			"userdomains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of user domains specified as comma seperated value",
			},
			"wihome": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Web address of the Web Interface server, such as http://<ipAddress>/Citrix/XenApp, or Receiver for Web, which enumerates the virtualized resources, such as XenApp, XenDesktop, and cloud applications. This web address is used as the home page in ICA proxy mode.\nIf Client Choices is ON, you must configure this setting. Because the user can choose between FullClient and ICAProxy, the user may see a different home page. An Internet web site may appear if the user gets the FullClient option, or a Web Interface site if the user gets the ICAProxy option. If the setting is not configured, the XenApp option does not appear as a client choice.",
			},
			"wihomeaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the wihome address(IPV4/V6)",
			},
			"windowsautologon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the Windows Auto Logon for the session. If a VPN session is established after this setting is enabled, the user is automatically logged on by using Windows credentials after the system is restarted.",
			},
			"windowsclienttype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("AGENT"),
				Description: "The Windows client type. Choose between two types of Windows Client\\\na) Application Agent - which always runs in the task bar as a standalone application and also has a supporting service which runs permanently when installed\\\nb) Activex Control - ActiveX control run by Microsoft Internet Explorer.",
			},
			"windowspluginupgrade": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("Always"),
				Description: "Option to set plugin upgrade behaviour for Win",
			},
			"winsip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "WINS server IP address to add to Citrix Gateway for name resolution.",
			},
			"wiportalmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layout on the Access Interface. The COMPACT value indicates the use of small icons.",
			},
		},
	}
}

func vpnparameterGetThePayloadFromtheConfig(ctx context.Context, data *VpnparameterResourceModel) vpn.Vpnparameter {
	tflog.Debug(ctx, "In vpnparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnparameter := vpn.Vpnparameter{}
	if !data.Accessrestrictedpageredirect.IsNull() {
		vpnparameter.Accessrestrictedpageredirect = data.Accessrestrictedpageredirect.ValueString()
	}
	if !data.Advancedclientlessvpnmode.IsNull() {
		vpnparameter.Advancedclientlessvpnmode = data.Advancedclientlessvpnmode.ValueString()
	}
	if !data.Allowedlogingroups.IsNull() {
		vpnparameter.Allowedlogingroups = data.Allowedlogingroups.ValueString()
	}
	if !data.Allprotocolproxy.IsNull() {
		vpnparameter.Allprotocolproxy = data.Allprotocolproxy.ValueString()
	}
	if !data.Alwaysonprofilename.IsNull() {
		vpnparameter.Alwaysonprofilename = data.Alwaysonprofilename.ValueString()
	}
	if !data.Apptokentimeout.IsNull() {
		vpnparameter.Apptokentimeout = utils.IntPtr(int(data.Apptokentimeout.ValueInt64()))
	}
	if !data.Authorizationgroup.IsNull() {
		vpnparameter.Authorizationgroup = data.Authorizationgroup.ValueString()
	}
	if !data.Autoproxyurl.IsNull() {
		vpnparameter.Autoproxyurl = data.Autoproxyurl.ValueString()
	}
	if !data.Backendcertvalidation.IsNull() {
		vpnparameter.Backendcertvalidation = data.Backendcertvalidation.ValueString()
	}
	if !data.Backenddtls12.IsNull() {
		vpnparameter.Backenddtls12 = data.Backenddtls12.ValueString()
	}
	if !data.Backendserversni.IsNull() {
		vpnparameter.Backendserversni = data.Backendserversni.ValueString()
	}
	if !data.Citrixreceiverhome.IsNull() {
		vpnparameter.Citrixreceiverhome = data.Citrixreceiverhome.ValueString()
	}
	if !data.Clientchoices.IsNull() {
		vpnparameter.Clientchoices = data.Clientchoices.ValueString()
	}
	if !data.Clientcleanupprompt.IsNull() {
		vpnparameter.Clientcleanupprompt = data.Clientcleanupprompt.ValueString()
	}
	if !data.Clientdebug.IsNull() {
		vpnparameter.Clientdebug = data.Clientdebug.ValueString()
	}
	if !data.Clientidletimeout.IsNull() {
		vpnparameter.Clientidletimeout = utils.IntPtr(int(data.Clientidletimeout.ValueInt64()))
	}
	if !data.Clientlessmodeurlencoding.IsNull() {
		vpnparameter.Clientlessmodeurlencoding = data.Clientlessmodeurlencoding.ValueString()
	}
	if !data.Clientlesspersistentcookie.IsNull() {
		vpnparameter.Clientlesspersistentcookie = data.Clientlesspersistentcookie.ValueString()
	}
	if !data.Clientlessvpnmode.IsNull() {
		vpnparameter.Clientlessvpnmode = data.Clientlessvpnmode.ValueString()
	}
	if !data.Clientsecurity.IsNull() {
		vpnparameter.Clientsecurity = data.Clientsecurity.ValueString()
	}
	if !data.Clientsecuritygroup.IsNull() {
		vpnparameter.Clientsecuritygroup = data.Clientsecuritygroup.ValueString()
	}
	if !data.Clientsecuritylog.IsNull() {
		vpnparameter.Clientsecuritylog = data.Clientsecuritylog.ValueString()
	}
	if !data.Clientsecuritymessage.IsNull() {
		vpnparameter.Clientsecuritymessage = data.Clientsecuritymessage.ValueString()
	}
	if !data.Clientversions.IsNull() {
		vpnparameter.Clientversions = data.Clientversions.ValueString()
	}
	if !data.Defaultauthorizationaction.IsNull() {
		vpnparameter.Defaultauthorizationaction = data.Defaultauthorizationaction.ValueString()
	}
	if !data.Deviceposture.IsNull() {
		vpnparameter.Deviceposture = data.Deviceposture.ValueString()
	}
	if !data.Dnsvservername.IsNull() {
		vpnparameter.Dnsvservername = data.Dnsvservername.ValueString()
	}
	if !data.Emailhome.IsNull() {
		vpnparameter.Emailhome = data.Emailhome.ValueString()
	}
	if !data.Encryptcsecexp.IsNull() {
		vpnparameter.Encryptcsecexp = data.Encryptcsecexp.ValueString()
	}
	if !data.Epaclienttype.IsNull() {
		vpnparameter.Epaclienttype = data.Epaclienttype.ValueString()
	}
	if !data.Forcedtimeout.IsNull() {
		vpnparameter.Forcedtimeout = utils.IntPtr(int(data.Forcedtimeout.ValueInt64()))
	}
	if !data.Forcedtimeoutwarning.IsNull() {
		vpnparameter.Forcedtimeoutwarning = utils.IntPtr(int(data.Forcedtimeoutwarning.ValueInt64()))
	}
	if !data.Fqdnspoofedip.IsNull() {
		vpnparameter.Fqdnspoofedip = data.Fqdnspoofedip.ValueString()
	}
	if !data.Ftpproxy.IsNull() {
		vpnparameter.Ftpproxy = data.Ftpproxy.ValueString()
	}
	if !data.Gopherproxy.IsNull() {
		vpnparameter.Gopherproxy = data.Gopherproxy.ValueString()
	}
	if !data.Homepage.IsNull() {
		vpnparameter.Homepage = data.Homepage.ValueString()
	}
	if !data.Httpproxy.IsNull() {
		vpnparameter.Httpproxy = data.Httpproxy.ValueString()
	}
	if !data.Httptrackconnproxy.IsNull() {
		vpnparameter.Httptrackconnproxy = data.Httptrackconnproxy.ValueString()
	}
	if !data.Icaproxy.IsNull() {
		vpnparameter.Icaproxy = data.Icaproxy.ValueString()
	}
	if !data.Icasessiontimeout.IsNull() {
		vpnparameter.Icasessiontimeout = data.Icasessiontimeout.ValueString()
	}
	if !data.Icauseraccounting.IsNull() {
		vpnparameter.Icauseraccounting = data.Icauseraccounting.ValueString()
	}
	if !data.Iconwithreceiver.IsNull() {
		vpnparameter.Iconwithreceiver = data.Iconwithreceiver.ValueString()
	}
	if !data.Iipdnssuffix.IsNull() {
		vpnparameter.Iipdnssuffix = data.Iipdnssuffix.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		vpnparameter.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Killconnections.IsNull() {
		vpnparameter.Killconnections = data.Killconnections.ValueString()
	}
	if !data.Linuxpluginupgrade.IsNull() {
		vpnparameter.Linuxpluginupgrade = data.Linuxpluginupgrade.ValueString()
	}
	if !data.Locallanaccess.IsNull() {
		vpnparameter.Locallanaccess = data.Locallanaccess.ValueString()
	}
	if !data.Loginscript.IsNull() {
		vpnparameter.Loginscript = data.Loginscript.ValueString()
	}
	if !data.Logoutscript.IsNull() {
		vpnparameter.Logoutscript = data.Logoutscript.ValueString()
	}
	if !data.Macpluginupgrade.IsNull() {
		vpnparameter.Macpluginupgrade = data.Macpluginupgrade.ValueString()
	}
	if !data.Maxiipperuser.IsNull() {
		vpnparameter.Maxiipperuser = utils.IntPtr(int(data.Maxiipperuser.ValueInt64()))
	}
	if !data.Mdxtokentimeout.IsNull() {
		vpnparameter.Mdxtokentimeout = utils.IntPtr(int(data.Mdxtokentimeout.ValueInt64()))
	}
	if !data.Netmask.IsNull() {
		vpnparameter.Netmask = data.Netmask.ValueString()
	}
	if !data.Ntdomain.IsNull() {
		vpnparameter.Ntdomain = data.Ntdomain.ValueString()
	}
	if !data.Pcoipprofilename.IsNull() {
		vpnparameter.Pcoipprofilename = data.Pcoipprofilename.ValueString()
	}
	if !data.Proxy.IsNull() {
		vpnparameter.Proxy = data.Proxy.ValueString()
	}
	if !data.Proxyexception.IsNull() {
		vpnparameter.Proxyexception = data.Proxyexception.ValueString()
	}
	if !data.Proxylocalbypass.IsNull() {
		vpnparameter.Proxylocalbypass = data.Proxylocalbypass.ValueString()
	}
	if !data.Rdpclientprofilename.IsNull() {
		vpnparameter.Rdpclientprofilename = data.Rdpclientprofilename.ValueString()
	}
	if !data.Rfc1918.IsNull() {
		vpnparameter.Rfc1918 = data.Rfc1918.ValueString()
	}
	if !data.Samesite.IsNull() {
		vpnparameter.Samesite = data.Samesite.ValueString()
	}
	if !data.Securebrowse.IsNull() {
		vpnparameter.Securebrowse = data.Securebrowse.ValueString()
	}
	if !data.Secureprivateaccess.IsNull() {
		vpnparameter.Secureprivateaccess = data.Secureprivateaccess.ValueString()
	}
	if !data.Sesstimeout.IsNull() {
		vpnparameter.Sesstimeout = utils.IntPtr(int(data.Sesstimeout.ValueInt64()))
	}
	if !data.Smartgroup.IsNull() {
		vpnparameter.Smartgroup = data.Smartgroup.ValueString()
	}
	if !data.Socksproxy.IsNull() {
		vpnparameter.Socksproxy = data.Socksproxy.ValueString()
	}
	if !data.Splitdns.IsNull() {
		vpnparameter.Splitdns = data.Splitdns.ValueString()
	}
	if !data.Splittunnel.IsNull() {
		vpnparameter.Splittunnel = data.Splittunnel.ValueString()
	}
	if !data.Spoofiip.IsNull() {
		vpnparameter.Spoofiip = data.Spoofiip.ValueString()
	}
	if !data.Sslproxy.IsNull() {
		vpnparameter.Sslproxy = data.Sslproxy.ValueString()
	}
	if !data.Sso.IsNull() {
		vpnparameter.Sso = data.Sso.ValueString()
	}
	if !data.Ssocredential.IsNull() {
		vpnparameter.Ssocredential = data.Ssocredential.ValueString()
	}
	if !data.Storefronturl.IsNull() {
		vpnparameter.Storefronturl = data.Storefronturl.ValueString()
	}
	if !data.Transparentinterception.IsNull() {
		vpnparameter.Transparentinterception = data.Transparentinterception.ValueString()
	}
	if !data.Uitheme.IsNull() {
		vpnparameter.Uitheme = data.Uitheme.ValueString()
	}
	if !data.Useiip.IsNull() {
		vpnparameter.Useiip = data.Useiip.ValueString()
	}
	if !data.Usemip.IsNull() {
		vpnparameter.Usemip = data.Usemip.ValueString()
	}
	if !data.Userdomains.IsNull() {
		vpnparameter.Userdomains = data.Userdomains.ValueString()
	}
	if !data.Wihome.IsNull() {
		vpnparameter.Wihome = data.Wihome.ValueString()
	}
	if !data.Wihomeaddresstype.IsNull() {
		vpnparameter.Wihomeaddresstype = data.Wihomeaddresstype.ValueString()
	}
	if !data.Windowsautologon.IsNull() {
		vpnparameter.Windowsautologon = data.Windowsautologon.ValueString()
	}
	if !data.Windowsclienttype.IsNull() {
		vpnparameter.Windowsclienttype = data.Windowsclienttype.ValueString()
	}
	if !data.Windowspluginupgrade.IsNull() {
		vpnparameter.Windowspluginupgrade = data.Windowspluginupgrade.ValueString()
	}
	if !data.Winsip.IsNull() {
		vpnparameter.Winsip = data.Winsip.ValueString()
	}
	if !data.Wiportalmode.IsNull() {
		vpnparameter.Wiportalmode = data.Wiportalmode.ValueString()
	}

	return vpnparameter
}

func vpnparameterSetAttrFromGet(ctx context.Context, data *VpnparameterResourceModel, getResponseData map[string]interface{}) *VpnparameterResourceModel {
	tflog.Debug(ctx, "In vpnparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accessrestrictedpageredirect"]; ok && val != nil {
		data.Accessrestrictedpageredirect = types.StringValue(val.(string))
	} else {
		data.Accessrestrictedpageredirect = types.StringNull()
	}
	if val, ok := getResponseData["advancedclientlessvpnmode"]; ok && val != nil {
		data.Advancedclientlessvpnmode = types.StringValue(val.(string))
	} else {
		data.Advancedclientlessvpnmode = types.StringNull()
	}
	if val, ok := getResponseData["allowedlogingroups"]; ok && val != nil {
		data.Allowedlogingroups = types.StringValue(val.(string))
	} else {
		data.Allowedlogingroups = types.StringNull()
	}
	if val, ok := getResponseData["allprotocolproxy"]; ok && val != nil {
		data.Allprotocolproxy = types.StringValue(val.(string))
	} else {
		data.Allprotocolproxy = types.StringNull()
	}
	if val, ok := getResponseData["alwaysonprofilename"]; ok && val != nil {
		data.Alwaysonprofilename = types.StringValue(val.(string))
	} else {
		data.Alwaysonprofilename = types.StringNull()
	}
	if val, ok := getResponseData["apptokentimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Apptokentimeout = types.Int64Value(intVal)
		}
	} else {
		data.Apptokentimeout = types.Int64Null()
	}
	if val, ok := getResponseData["authorizationgroup"]; ok && val != nil {
		data.Authorizationgroup = types.StringValue(val.(string))
	} else {
		data.Authorizationgroup = types.StringNull()
	}
	if val, ok := getResponseData["autoproxyurl"]; ok && val != nil {
		data.Autoproxyurl = types.StringValue(val.(string))
	} else {
		data.Autoproxyurl = types.StringNull()
	}
	if val, ok := getResponseData["backendcertvalidation"]; ok && val != nil {
		data.Backendcertvalidation = types.StringValue(val.(string))
	} else {
		data.Backendcertvalidation = types.StringNull()
	}
	if val, ok := getResponseData["backenddtls12"]; ok && val != nil {
		data.Backenddtls12 = types.StringValue(val.(string))
	} else {
		data.Backenddtls12 = types.StringNull()
	}
	if val, ok := getResponseData["backendserversni"]; ok && val != nil {
		data.Backendserversni = types.StringValue(val.(string))
	} else {
		data.Backendserversni = types.StringNull()
	}
	if val, ok := getResponseData["citrixreceiverhome"]; ok && val != nil {
		data.Citrixreceiverhome = types.StringValue(val.(string))
	} else {
		data.Citrixreceiverhome = types.StringNull()
	}
	if val, ok := getResponseData["clientchoices"]; ok && val != nil {
		data.Clientchoices = types.StringValue(val.(string))
	} else {
		data.Clientchoices = types.StringNull()
	}
	if val, ok := getResponseData["clientcleanupprompt"]; ok && val != nil {
		data.Clientcleanupprompt = types.StringValue(val.(string))
	} else {
		data.Clientcleanupprompt = types.StringNull()
	}
	if val, ok := getResponseData["clientdebug"]; ok && val != nil {
		data.Clientdebug = types.StringValue(val.(string))
	} else {
		data.Clientdebug = types.StringNull()
	}
	if val, ok := getResponseData["clientidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clientidletimeout = types.Int64Value(intVal)
		}
	} else {
		data.Clientidletimeout = types.Int64Null()
	}
	if val, ok := getResponseData["clientlessmodeurlencoding"]; ok && val != nil {
		data.Clientlessmodeurlencoding = types.StringValue(val.(string))
	} else {
		data.Clientlessmodeurlencoding = types.StringNull()
	}
	if val, ok := getResponseData["clientlesspersistentcookie"]; ok && val != nil {
		data.Clientlesspersistentcookie = types.StringValue(val.(string))
	} else {
		data.Clientlesspersistentcookie = types.StringNull()
	}
	if val, ok := getResponseData["clientlessvpnmode"]; ok && val != nil {
		data.Clientlessvpnmode = types.StringValue(val.(string))
	} else {
		data.Clientlessvpnmode = types.StringNull()
	}
	if val, ok := getResponseData["clientsecurity"]; ok && val != nil {
		data.Clientsecurity = types.StringValue(val.(string))
	} else {
		data.Clientsecurity = types.StringNull()
	}
	if val, ok := getResponseData["clientsecuritygroup"]; ok && val != nil {
		data.Clientsecuritygroup = types.StringValue(val.(string))
	} else {
		data.Clientsecuritygroup = types.StringNull()
	}
	if val, ok := getResponseData["clientsecuritylog"]; ok && val != nil {
		data.Clientsecuritylog = types.StringValue(val.(string))
	} else {
		data.Clientsecuritylog = types.StringNull()
	}
	if val, ok := getResponseData["clientsecuritymessage"]; ok && val != nil {
		data.Clientsecuritymessage = types.StringValue(val.(string))
	} else {
		data.Clientsecuritymessage = types.StringNull()
	}
	if val, ok := getResponseData["clientversions"]; ok && val != nil {
		data.Clientversions = types.StringValue(val.(string))
	} else {
		data.Clientversions = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthorizationaction"]; ok && val != nil {
		data.Defaultauthorizationaction = types.StringValue(val.(string))
	} else {
		data.Defaultauthorizationaction = types.StringNull()
	}
	if val, ok := getResponseData["deviceposture"]; ok && val != nil {
		data.Deviceposture = types.StringValue(val.(string))
	} else {
		data.Deviceposture = types.StringNull()
	}
	if val, ok := getResponseData["dnsvservername"]; ok && val != nil {
		data.Dnsvservername = types.StringValue(val.(string))
	} else {
		data.Dnsvservername = types.StringNull()
	}
	if val, ok := getResponseData["emailhome"]; ok && val != nil {
		data.Emailhome = types.StringValue(val.(string))
	} else {
		data.Emailhome = types.StringNull()
	}
	if val, ok := getResponseData["encryptcsecexp"]; ok && val != nil {
		data.Encryptcsecexp = types.StringValue(val.(string))
	} else {
		data.Encryptcsecexp = types.StringNull()
	}
	if val, ok := getResponseData["epaclienttype"]; ok && val != nil {
		data.Epaclienttype = types.StringValue(val.(string))
	} else {
		data.Epaclienttype = types.StringNull()
	}
	if val, ok := getResponseData["forcedtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Forcedtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Forcedtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["forcedtimeoutwarning"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Forcedtimeoutwarning = types.Int64Value(intVal)
		}
	} else {
		data.Forcedtimeoutwarning = types.Int64Null()
	}
	if val, ok := getResponseData["fqdnspoofedip"]; ok && val != nil {
		data.Fqdnspoofedip = types.StringValue(val.(string))
	} else {
		data.Fqdnspoofedip = types.StringNull()
	}
	if val, ok := getResponseData["ftpproxy"]; ok && val != nil {
		data.Ftpproxy = types.StringValue(val.(string))
	} else {
		data.Ftpproxy = types.StringNull()
	}
	if val, ok := getResponseData["gopherproxy"]; ok && val != nil {
		data.Gopherproxy = types.StringValue(val.(string))
	} else {
		data.Gopherproxy = types.StringNull()
	}
	if val, ok := getResponseData["homepage"]; ok && val != nil {
		data.Homepage = types.StringValue(val.(string))
	} else {
		data.Homepage = types.StringNull()
	}
	if val, ok := getResponseData["httpproxy"]; ok && val != nil {
		data.Httpproxy = types.StringValue(val.(string))
	} else {
		data.Httpproxy = types.StringNull()
	}
	if val, ok := getResponseData["httptrackconnproxy"]; ok && val != nil {
		data.Httptrackconnproxy = types.StringValue(val.(string))
	} else {
		data.Httptrackconnproxy = types.StringNull()
	}
	if val, ok := getResponseData["icaproxy"]; ok && val != nil {
		data.Icaproxy = types.StringValue(val.(string))
	} else {
		data.Icaproxy = types.StringNull()
	}
	if val, ok := getResponseData["icasessiontimeout"]; ok && val != nil {
		data.Icasessiontimeout = types.StringValue(val.(string))
	} else {
		data.Icasessiontimeout = types.StringNull()
	}
	if val, ok := getResponseData["icauseraccounting"]; ok && val != nil {
		data.Icauseraccounting = types.StringValue(val.(string))
	} else {
		data.Icauseraccounting = types.StringNull()
	}
	if val, ok := getResponseData["iconwithreceiver"]; ok && val != nil {
		data.Iconwithreceiver = types.StringValue(val.(string))
	} else {
		data.Iconwithreceiver = types.StringNull()
	}
	if val, ok := getResponseData["iipdnssuffix"]; ok && val != nil {
		data.Iipdnssuffix = types.StringValue(val.(string))
	} else {
		data.Iipdnssuffix = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["killconnections"]; ok && val != nil {
		data.Killconnections = types.StringValue(val.(string))
	} else {
		data.Killconnections = types.StringNull()
	}
	if val, ok := getResponseData["linuxpluginupgrade"]; ok && val != nil {
		data.Linuxpluginupgrade = types.StringValue(val.(string))
	} else {
		data.Linuxpluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["locallanaccess"]; ok && val != nil {
		data.Locallanaccess = types.StringValue(val.(string))
	} else {
		data.Locallanaccess = types.StringNull()
	}
	if val, ok := getResponseData["loginscript"]; ok && val != nil {
		data.Loginscript = types.StringValue(val.(string))
	} else {
		data.Loginscript = types.StringNull()
	}
	if val, ok := getResponseData["logoutscript"]; ok && val != nil {
		data.Logoutscript = types.StringValue(val.(string))
	} else {
		data.Logoutscript = types.StringNull()
	}
	if val, ok := getResponseData["macpluginupgrade"]; ok && val != nil {
		data.Macpluginupgrade = types.StringValue(val.(string))
	} else {
		data.Macpluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["maxiipperuser"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxiipperuser = types.Int64Value(intVal)
		}
	} else {
		data.Maxiipperuser = types.Int64Null()
	}
	if val, ok := getResponseData["mdxtokentimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mdxtokentimeout = types.Int64Value(intVal)
		}
	} else {
		data.Mdxtokentimeout = types.Int64Null()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["ntdomain"]; ok && val != nil {
		data.Ntdomain = types.StringValue(val.(string))
	} else {
		data.Ntdomain = types.StringNull()
	}
	if val, ok := getResponseData["pcoipprofilename"]; ok && val != nil {
		data.Pcoipprofilename = types.StringValue(val.(string))
	} else {
		data.Pcoipprofilename = types.StringNull()
	}
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["proxyexception"]; ok && val != nil {
		data.Proxyexception = types.StringValue(val.(string))
	} else {
		data.Proxyexception = types.StringNull()
	}
	if val, ok := getResponseData["proxylocalbypass"]; ok && val != nil {
		data.Proxylocalbypass = types.StringValue(val.(string))
	} else {
		data.Proxylocalbypass = types.StringNull()
	}
	if val, ok := getResponseData["rdpclientprofilename"]; ok && val != nil {
		data.Rdpclientprofilename = types.StringValue(val.(string))
	} else {
		data.Rdpclientprofilename = types.StringNull()
	}
	if val, ok := getResponseData["rfc1918"]; ok && val != nil {
		data.Rfc1918 = types.StringValue(val.(string))
	} else {
		data.Rfc1918 = types.StringNull()
	}
	if val, ok := getResponseData["samesite"]; ok && val != nil {
		data.Samesite = types.StringValue(val.(string))
	} else {
		data.Samesite = types.StringNull()
	}
	if val, ok := getResponseData["securebrowse"]; ok && val != nil {
		data.Securebrowse = types.StringValue(val.(string))
	} else {
		data.Securebrowse = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccess"]; ok && val != nil {
		data.Secureprivateaccess = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccess = types.StringNull()
	}
	if val, ok := getResponseData["sesstimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sesstimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sesstimeout = types.Int64Null()
	}
	if val, ok := getResponseData["smartgroup"]; ok && val != nil {
		data.Smartgroup = types.StringValue(val.(string))
	} else {
		data.Smartgroup = types.StringNull()
	}
	if val, ok := getResponseData["socksproxy"]; ok && val != nil {
		data.Socksproxy = types.StringValue(val.(string))
	} else {
		data.Socksproxy = types.StringNull()
	}
	if val, ok := getResponseData["splitdns"]; ok && val != nil {
		data.Splitdns = types.StringValue(val.(string))
	} else {
		data.Splitdns = types.StringNull()
	}
	if val, ok := getResponseData["splittunnel"]; ok && val != nil {
		data.Splittunnel = types.StringValue(val.(string))
	} else {
		data.Splittunnel = types.StringNull()
	}
	if val, ok := getResponseData["spoofiip"]; ok && val != nil {
		data.Spoofiip = types.StringValue(val.(string))
	} else {
		data.Spoofiip = types.StringNull()
	}
	if val, ok := getResponseData["sslproxy"]; ok && val != nil {
		data.Sslproxy = types.StringValue(val.(string))
	} else {
		data.Sslproxy = types.StringNull()
	}
	if val, ok := getResponseData["sso"]; ok && val != nil {
		data.Sso = types.StringValue(val.(string))
	} else {
		data.Sso = types.StringNull()
	}
	if val, ok := getResponseData["ssocredential"]; ok && val != nil {
		data.Ssocredential = types.StringValue(val.(string))
	} else {
		data.Ssocredential = types.StringNull()
	}
	if val, ok := getResponseData["storefronturl"]; ok && val != nil {
		data.Storefronturl = types.StringValue(val.(string))
	} else {
		data.Storefronturl = types.StringNull()
	}
	if val, ok := getResponseData["transparentinterception"]; ok && val != nil {
		data.Transparentinterception = types.StringValue(val.(string))
	} else {
		data.Transparentinterception = types.StringNull()
	}
	if val, ok := getResponseData["uitheme"]; ok && val != nil {
		data.Uitheme = types.StringValue(val.(string))
	} else {
		data.Uitheme = types.StringNull()
	}
	if val, ok := getResponseData["useiip"]; ok && val != nil {
		data.Useiip = types.StringValue(val.(string))
	} else {
		data.Useiip = types.StringNull()
	}
	if val, ok := getResponseData["usemip"]; ok && val != nil {
		data.Usemip = types.StringValue(val.(string))
	} else {
		data.Usemip = types.StringNull()
	}
	if val, ok := getResponseData["userdomains"]; ok && val != nil {
		data.Userdomains = types.StringValue(val.(string))
	} else {
		data.Userdomains = types.StringNull()
	}
	if val, ok := getResponseData["wihome"]; ok && val != nil {
		data.Wihome = types.StringValue(val.(string))
	} else {
		data.Wihome = types.StringNull()
	}
	if val, ok := getResponseData["wihomeaddresstype"]; ok && val != nil {
		data.Wihomeaddresstype = types.StringValue(val.(string))
	} else {
		data.Wihomeaddresstype = types.StringNull()
	}
	if val, ok := getResponseData["windowsautologon"]; ok && val != nil {
		data.Windowsautologon = types.StringValue(val.(string))
	} else {
		data.Windowsautologon = types.StringNull()
	}
	if val, ok := getResponseData["windowsclienttype"]; ok && val != nil {
		data.Windowsclienttype = types.StringValue(val.(string))
	} else {
		data.Windowsclienttype = types.StringNull()
	}
	if val, ok := getResponseData["windowspluginupgrade"]; ok && val != nil {
		data.Windowspluginupgrade = types.StringValue(val.(string))
	} else {
		data.Windowspluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["winsip"]; ok && val != nil {
		data.Winsip = types.StringValue(val.(string))
	} else {
		data.Winsip = types.StringNull()
	}
	if val, ok := getResponseData["wiportalmode"]; ok && val != nil {
		data.Wiportalmode = types.StringValue(val.(string))
	} else {
		data.Wiportalmode = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("vpnparameter-config")

	return data
}
