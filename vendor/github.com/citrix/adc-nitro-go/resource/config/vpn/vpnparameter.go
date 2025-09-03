/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package vpn

/**
* Configuration for VPN parameter resource.
*/
type Vpnparameter struct {
	/**
	* Destination port numbers other than port 80, added as a comma-separated list. Traffic to these ports is processed as HTTP traffic, which allows functionality, such as HTTP authorization and single sign-on to a web application to work.
	*/
	Httpport []int `json:"httpport,omitempty"`
	/**
	* WINS server IP address to add to Citrix Gateway for name resolution.
	*/
	Winsip string `json:"winsip,omitempty"`
	/**
	* Name of the DNS virtual server for the user session.
	*/
	Dnsvservername string `json:"dnsvservername,omitempty"`
	/**
	* Route the DNS requests to the local DNS server configured on the user device, or Citrix Gateway (remote), or both.
	*/
	Splitdns string `json:"splitdns,omitempty"`
	/**
	* The name of the radiusPolicy to use for RADIUS user accounting info on the session.
	*/
	Icauseraccounting string `json:"icauseraccounting,omitempty"`
	/**
	* Number of minutes after which the session times out.
	*/
	Sesstimeout int `json:"sesstimeout,omitempty"`
	/**
	* Specify the client security check for the user device to permit a Citrix Gateway session. The web address or IP address is not included in the expression for the client security check.
	*/
	Clientsecurity string `json:"clientsecurity,omitempty"`
	/**
	* The client security group that will be assigned on failure of the client security check. Users can in general be organized into Groups. In this case, the Client Security Group may have a more restrictive security policy.
	*/
	Clientsecuritygroup string `json:"clientsecuritygroup,omitempty"`
	/**
	* The client security message that will be displayed on failure of the client security check.
	*/
	Clientsecuritymessage string `json:"clientsecuritymessage,omitempty"`
	/**
	* Specifies whether or not to display all failed Client Security scans to the end user
	*/
	Clientsecuritylog string `json:"clientsecuritylog,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Smartgroup string `json:"smartgroup,omitempty"`
	/**
	* Send, through the tunnel, traffic only for intranet applications that are defined in Citrix Gateway. Route all other traffic directly to the Internet. The OFF setting routes all traffic through Citrix Gateway. With the REVERSE setting, intranet applications define the network traffic that is not intercepted. All network traffic directed to internal IP addresses bypasses the VPN tunnel, while other traffic goes through Citrix Gateway. Reverse split tunneling can be used to log all non-local LAN traffic. For example, if users have a home network and are logged on through the Citrix Gateway Plug-in, network traffic destined to a printer or another device within the home network is not intercepted.
	*/
	Splittunnel string `json:"splittunnel,omitempty"`
	/**
	* Set local LAN access. If split tunneling is OFF, and you set local LAN access to ON, the local client can route traffic to its local interface. When the local area network switch is specified, this combination of switches is useful. The client can allow local LAN access to devices that commonly have non-routable addresses, such as local printers or local file servers.
	*/
	Locallanaccess string `json:"locallanaccess,omitempty"`
	/**
	* As defined in the local area network, allow only the following local area network addresses to bypass the VPN tunnel when the local LAN access feature is enabled:
		* 10.*.*.*,
		* 172.16.*.*,
		* 192.168.*.*
	*/
	Rfc1918 string `json:"rfc1918,omitempty"`
	/**
	* Indicate whether or not the application requires IP spoofing, which routes the connection to the intranet application through the virtual adapter.
	*/
	Spoofiip string `json:"spoofiip,omitempty"`
	/**
	* Specify whether the Citrix Gateway Plug-in should disconnect all preexisting connections, such as the connections existing before the user logged on to Citrix Gateway, and prevent new incoming connections on the Citrix Gateway Plug-in for Windows and MAC when the user is connected to Citrix Gateway and split tunneling is disabled.
	*/
	Killconnections string `json:"killconnections,omitempty"`
	/**
	* Allow access to network resources by using a single IP address and subnet mask or a range of IP addresses. The OFF setting sets the mode to proxy, in which you configure destination and source IP addresses and port numbers. If you are using the Citrix Gateway Plug-in for Windows, set this parameter to ON, in which the mode is set to transparent. If you are using the Citrix Gateway Plug-in for Java, set this parameter to OFF.
	*/
	Transparentinterception string `json:"transparentinterception,omitempty"`
	/**
	* The Windows client type. Choose between two types of Windows Client\
		a) Application Agent - which always runs in the task bar as a standalone application and also has a supporting service which runs permanently when installed\
		b) Activex Control - ActiveX control run by Microsoft Internet Explorer.
	*/
	Windowsclienttype string `json:"windowsclienttype,omitempty"`
	/**
	* Specify the network resources that users have access to when they log on to the internal network. The default setting for authorization is to deny access to all network resources. Citrix recommends using the default global setting and then creating authorization policies to define the network resources users can access. If you set the default authorization policy to DENY, you must explicitly authorize access to any network resource, which improves security.
	*/
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	/**
	* Comma-separated list of groups in which the user is placed when none of the groups that the user is a part of is configured on Citrix Gateway. The authorization policy can be bound to these groups to control access to the resources.
	*/
	Authorizationgroup string `json:"authorizationgroup,omitempty"`
	/**
	* Time, in minutes, after which to time out the user session if Citrix Gateway does not detect mouse or keyboard activity.
	*/
	Clientidletimeout int `json:"clientidletimeout,omitempty"`
	/**
	* Set options to apply proxy for accessing the internal resources. Available settings function as follows:
		* BROWSER - Proxy settings are configured only in Internet Explorer and Firefox browsers.
		* NS - Proxy settings are configured on the Citrix ADC.
		* OFF - Proxy settings are not configured.
	*/
	Proxy string `json:"proxy,omitempty"`
	/**
	* IP address of the proxy server to use for all protocols supported by Citrix Gateway.
	*/
	Allprotocolproxy string `json:"allprotocolproxy,omitempty"`
	/**
	* IP address of the proxy server to be used for HTTP access for all subsequent connections to the internal network.
	*/
	Httpproxy string `json:"httpproxy,omitempty"`
	/**
	* IP address of the proxy server to be used for FTP access for all subsequent connections to the internal network.
	*/
	Ftpproxy string `json:"ftpproxy,omitempty"`
	/**
	* IP address of the proxy server to be used for SOCKS access for all subsequent connections to the internal network.
	*/
	Socksproxy string `json:"socksproxy,omitempty"`
	/**
	* IP address of the proxy server to be used for GOPHER access for all subsequent connections to the internal network.
	*/
	Gopherproxy string `json:"gopherproxy,omitempty"`
	/**
	* IP address of the proxy server to be used for SSL access for all subsequent connections to the internal network.
	*/
	Sslproxy string `json:"sslproxy,omitempty"`
	/**
	* Proxy exception string that will be configured in the browser for bypassing the previously configured proxies. Allowed only if proxy type is Browser.
	*/
	Proxyexception string `json:"proxyexception,omitempty"`
	/**
	* Bypass proxy server for local addresses option in Internet Explorer and Firefox proxy server settings.
	*/
	Proxylocalbypass string `json:"proxylocalbypass,omitempty"`
	/**
	* Prompt for client-side cache clean-up when a client-initiated session closes.
	*/
	Clientcleanupprompt string `json:"clientcleanupprompt,omitempty"`
	/**
	* Force cache clean-up when the user closes a session. You can specify all, none, or any combination of the client-side items.
	*/
	Forcecleanup []string `json:"forcecleanup,omitempty"`
	/**
	* Display only the configured menu options when you select the "Configure Citrix Gateway" option in the Citrix Gateway Plug-in's system tray icon for Windows.
	*/
	Clientoptions []string `json:"clientoptions,omitempty"`
	/**
	* Allow users to change client Debug logging level in Configuration tab of the Citrix Gateway Plug-in for Windows.
	*/
	Clientconfiguration []string `json:"clientconfiguration,omitempty"`
	/**
	* Set single sign-on (SSO) for the session. When the user accesses a server, the user's logon credentials are passed to the server for authentication.
		NOTE : This configuration does not honor the following authentication types for security reason. BASIC, DIGEST, and NTLM (without Negotiate NTLM2 Key or Negotiate Sign Flag). Use VPN TrafficAction to configure SSO for these authentication types.
	*/
	Sso string `json:"sso,omitempty"`
	/**
	* Specify whether to use the primary or secondary authentication credentials for single sign-on to the server.
	*/
	Ssocredential string `json:"ssocredential,omitempty"`
	/**
	* Enable or disable the Windows Auto Logon for the session. If a VPN session is established after this setting is enabled, the user is automatically logged on by using Windows credentials after the system is restarted.
	*/
	Windowsautologon string `json:"windowsautologon,omitempty"`
	/**
	* Enable or disable the use of a unique IP address alias, or a mapped IP address, as the client IP address for each client session. Allow Citrix Gateway to use the mapped IP address as an intranet IP address when all other IP addresses are not available.
		When IP pooling is configured and the mapped IP is used as an intranet IP address, the mapped IP address is used when an intranet IP address cannot be assigned.
	*/
	Usemip string `json:"usemip,omitempty"`
	/**
	* Define IP address pool options. Available settings function as follows:
		* SPILLOVER - When an address pool is configured and the mapped IP is used as an intranet IP address, the mapped IP address is used when an intranet IP address cannot be assigned.
		* NOSPILLOVER - When intranet IP addresses are enabled and the mapped IP address is not used, the Transfer Login page appears for users who have used all available intranet IP addresses.
		* OFF - Address pool is not configured.
	*/
	Useiip string `json:"useiip,omitempty"`
	/**
	* Set the trace level on Citrix Gateway. Technical support technicians use these debug logs for in-depth debugging and troubleshooting purposes. Available settings function as follows:
		* DEBUG - Detailed debug messages are collected and written into the specified file.
		* STATS - Application audit level error messages and debug statistic counters are written into the specified file.
		* EVENTS - Application audit-level error messages are written into the specified file.
		* OFF - Only critical events are logged into the Windows Application Log.
	*/
	Clientdebug string `json:"clientdebug,omitempty"`
	/**
	* Path to the logon script that is run when a session is established. Separate multiple scripts by using comma. A "$" in the path signifies that the word following the "$" is an environment variable.
	*/
	Loginscript string `json:"loginscript,omitempty"`
	/**
	* Path to the logout script. Separate multiple scripts by using comma. A "$" in the path signifies that the word following the "$" is an environment variable.
	*/
	Logoutscript string `json:"logoutscript,omitempty"`
	/**
	* Web address of the home page that appears when users log on. Otherwise, users receive the default home page for Citrix Gateway, which is the Access Interface.
	*/
	Homepage string `json:"homepage,omitempty"`
	/**
	* Enable ICA proxy to configure secure Internet access to servers running Citrix XenApp or XenDesktop by using Citrix Receiver instead of the Citrix Gateway Plug-in.
	*/
	Icaproxy string `json:"icaproxy,omitempty"`
	/**
	* Web address of the Web Interface server, such as http://<ipAddress>/Citrix/XenApp, or Receiver for Web, which enumerates the virtualized resources, such as XenApp, XenDesktop, and cloud applications. This web address is used as the home page in ICA proxy mode.
		If Client Choices is ON, you must configure this setting. Because the user can choose between FullClient and ICAProxy, the user may see a different home page. An Internet web site may appear if the user gets the FullClient option, or a Web Interface site if the user gets the ICAProxy option. If the setting is not configured, the XenApp option does not appear as a client choice.
	*/
	Wihome string `json:"wihome,omitempty"`
	/**
	* Type of the wihome address(IPV4/V6)
	*/
	Wihomeaddresstype string `json:"wihomeaddresstype,omitempty"`
	/**
	* Web address for the Citrix Receiver home page. Configure Citrix Gateway so that when users log on to the appliance, the Citrix Gateway Plug-in opens a web browser that allows single sign-on to the Citrix Receiver home page.
	*/
	Citrixreceiverhome string `json:"citrixreceiverhome,omitempty"`
	/**
	* Layout on the Access Interface. The COMPACT value indicates the use of small icons.
	*/
	Wiportalmode string `json:"wiportalmode,omitempty"`
	/**
	* Provide users with multiple logon options. With client choices, users have the option of logging on by using the Citrix Gateway Plug-in for Windows, Citrix Gateway Plug-in for Java, the Web Interface, or clientless access from one location. Depending on how Citrix Gateway is configured, users are presented with up to three icons for logon choices. The most common are the Citrix Gateway Plug-in for Windows, Web Interface, and clientless access.
	*/
	Clientchoices string `json:"clientchoices,omitempty"`
	/**
	* Choose between two types of End point Windows Client
		a) Application Agent - which always runs in the task bar as a standalone application and also has a supporting service which runs permanently when installed
		b) Activex Control - ActiveX control run by Microsoft Internet Explorer.
	*/
	Epaclienttype string `json:"epaclienttype,omitempty"`
	/**
	* An intranet IP DNS suffix. When a user logs on to Citrix Gateway and is assigned an IP address, a DNS record for the user name and IP address combination is added to the Citrix Gateway DNS cache. You can configure a DNS suffix to append to the user name when the DNS record is added to the cache. You can reach to the host from where the user is logged on by using the user's name, which can be easier to remember than an IP address. When the user logs off from Citrix Gateway, the record is removed from the DNS cache.
	*/
	Iipdnssuffix string `json:"iipdnssuffix,omitempty"`
	/**
	* Force a disconnection from the Citrix Gateway Plug-in with Citrix Gateway after a specified number of minutes. If the session closes, the user must log on again.
	*/
	Forcedtimeout int `json:"forcedtimeout,omitempty"`
	/**
	* Number of minutes to warn a user before the user session is disconnected.
	*/
	Forcedtimeoutwarning int `json:"forcedtimeoutwarning,omitempty"`
	/**
	* Single sign-on domain to use for single sign-on to applications in the internal network. This setting can be overwritten by the domain that users specify at the time of logon or by the domain that the authentication server returns.
	*/
	Ntdomain string `json:"ntdomain,omitempty"`
	/**
	* Enable clientless access for web, XenApp or XenDesktop, and FileShare resources without installing the Citrix Gateway Plug-in. Available settings function as follows:
		* ON - Allow only clientless access.
		* OFF - Allow clientless access after users log on with the Citrix Gateway Plug-in.
		* DISABLED - Do not allow clientless access.
	*/
	Clientlessvpnmode string `json:"clientlessvpnmode,omitempty"`
	/**
	* When clientless access is enabled, you can choose to encode the addresses of internal web applications or to leave the address as clear text. Available settings function as follows:
		* OPAQUE - Use standard encoding mechanisms to make the domain and protocol part of the resource unclear to users.
		* TRANSPARENT - Do not encode the web address and make it visible to users.
		* ENCRYPT - Allow the domain and protocol to be encrypted using a session key. When the web address is encrypted, the URL is different for each user session for the same web resource. If users bookmark the encoded web address, save it in the web browser and then log off, they cannot connect to the web address when they log on and use the bookmark. If users save the encrypted bookmark in the Access Interface during their session, the bookmark works each time the user logs on.
	*/
	Clientlessmodeurlencoding string `json:"clientlessmodeurlencoding,omitempty"`
	/**
	* State of persistent cookies in clientless access mode. Persistent cookies are required for accessing certain features of SharePoint, such as opening and editing Microsoft Word, Excel, and PowerPoint documents hosted on the SharePoint server. A persistent cookie remains on the user device and is sent with each HTTP request. Citrix Gateway encrypts the persistent cookie before sending it to the plug-in on the user device, and refreshes the cookie periodically as long as the session exists. The cookie becomes stale if the session ends. Available settings function as follows:
		* ALLOW - Enable persistent cookies. Users can open and edit Microsoft documents stored in SharePoint.
		* DENY - Disable persistent cookies. Users cannot open and edit Microsoft documents stored in SharePoint.
		* PROMPT - Prompt users to allow or deny persistent cookies during the session. Persistent cookies are not required for clientless access if users do not connect to SharePoint.
	*/
	Clientlesspersistentcookie string `json:"clientlesspersistentcookie,omitempty"`
	/**
	* Web address for the web-based email, such as Outlook Web Access.
	*/
	Emailhome string `json:"emailhome,omitempty"`
	/**
	* Specify groups that have permission to log on to Citrix Gateway. Users who do not belong to this group or groups are denied access even if they have valid credentials.
	*/
	Allowedlogingroups string `json:"allowedlogingroups,omitempty"`
	/**
	* Enable encryption of client security expressions.
	*/
	Encryptcsecexp string `json:"encryptcsecexp,omitempty"`
	/**
	* The timeout value in seconds for tokens to access XenMobile applications
	*/
	Apptokentimeout int `json:"apptokentimeout,omitempty"`
	/**
	* Validity of MDX Token in minutes. This token is used for mdx services to access backend and valid  HEAD and GET request.
	*/
	Mdxtokentimeout int `json:"mdxtokentimeout,omitempty"`
	/**
	* Set VPN UI Theme to Green-Bubble, Caxton or Custom; default is Caxton.
	*/
	Uitheme string `json:"uitheme,omitempty"`
	/**
	* Allow users to connect through Citrix Gateway to network resources from iOS and Android mobile devices with Citrix Receiver. Users do not need to establish a full VPN tunnel to access resources in the secure network.
	*/
	Securebrowse string `json:"securebrowse,omitempty"`
	/**
	* Web address for StoreFront to be used in this session for enumeration of resources from XenApp or XenDesktop.
	*/
	Storefronturl string `json:"storefronturl,omitempty"`
	/**
	* The KCD account details to be used in SSO
	*/
	Kcdaccount string `json:"kcdaccount,omitempty"`
	/**
	* checkversion api
	*/
	Clientversions string `json:"clientversions,omitempty"`
	/**
	* Name of the RDP profile associated with the vserver.
	*/
	Rdpclientprofilename string `json:"rdpclientprofilename,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Win
	*/
	Windowspluginupgrade string `json:"windowspluginupgrade,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Mac
	*/
	Macpluginupgrade string `json:"macpluginupgrade,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Linux
	*/
	Linuxpluginupgrade string `json:"linuxpluginupgrade,omitempty"`
	/**
	* Option to decide whether to show plugin icon along with receiver icon
	*/
	Iconwithreceiver string `json:"iconwithreceiver,omitempty"`
	/**
	* List of user domains specified as comma seperated value
	*/
	Userdomains string `json:"userdomains,omitempty"`
	/**
	* Enable or disable ica session timeout. If enabled and in case AAA session gets terminated, ICA connections associated with that will also get terminated
	*/
	Icasessiontimeout string `json:"icasessiontimeout,omitempty"`
	/**
	* Enable or disable HTTP tracking for packets proxied via vpn vserver using GSLB connection proxy feature.
	*/
	Httptrackconnproxy string `json:"httptrackconnproxy,omitempty"`
	/**
	* Name of the AlwaysON profile. The builtin profile named none can be used to explicitly disable AlwaysON.
	*/
	Alwaysonprofilename string `json:"alwaysonprofilename,omitempty"`
	/**
	* URL to auto proxy config file
	*/
	Autoproxyurl string `json:"autoproxyurl,omitempty"`
	/**
	* Option to enable/disable Advanced ClientlessVpnMode. Additionaly, it can be set to STRICT to block Classic ClientlessVpnMode while in AdvancedClientlessMode.
	*/
	Advancedclientlessvpnmode string `json:"advancedclientlessvpnmode,omitempty"`
	/**
	* Name of the PCOIP profile.
	*/
	Pcoipprofilename string `json:"pcoipprofilename,omitempty"`
	/**
	* enables sni extension for backend server handshakes
	*/
	Backendserversni string `json:"backendserversni,omitempty"`
	/**
	* enables backend server certificate validation
	*/
	Backendcertvalidation string `json:"backendcertvalidation,omitempty"`
	/**
	* Enables or disables the secure private access configuration.
	*/
	Secureprivateaccess string `json:"secureprivateaccess,omitempty"`
	/**
	* By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.
	*/
	Accessrestrictedpageredirect string `json:"accessrestrictedpageredirect,omitempty"`
	/**
	* Spoofed IP address range that can be used by client for FQDN based split tunneling
	*/
	Fqdnspoofedip string `json:"fqdnspoofedip,omitempty"`
	/**
	* The netmask for the spoofed ip address
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* SameSite attribute value for Cookies generated in VPN context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
	*/
	Samesite string `json:"samesite,omitempty"`
	/**
	* Maximum number of Intranet IP that can be assigned to a user from AAA group, VPN vserver or VPN global pool. This setting is not applicable for AAA user level Intranet IP configuration
	*/
	Maxiipperuser int `json:"maxiipperuser,omitempty"`
	/**
	* Enable device posture
	*/
	Deviceposture string `json:"deviceposture,omitempty"`
	/**
	* Enables DTLS 1.2 for backend server handshakes
	*/
	Backenddtls12 string `json:"backenddtls12,omitempty"`

	//------- Read only Parameter ---------;

	Name string `json:"name,omitempty"`
	Clientidletimeoutwarning string `json:"clientidletimeoutwarning,omitempty"`
	Vpnsessionpolicybindtype string `json:"vpnsessionpolicybindtype,omitempty"`
	Vpnsessionpolicycount string `json:"vpnsessionpolicycount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
