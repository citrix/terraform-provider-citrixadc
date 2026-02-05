---
subcategory: "VPN"
---

# Data Source: vpnsessionaction

The vpnsessionaction data source allows you to retrieve information about a VPN session action configuration.

## Example usage

```terraform
data "citrixadc_vpnsessionaction" "foo" {
  name = "newsession"
}

output "sesstimeout" {
  value = data.citrixadc_vpnsessionaction.foo.sesstimeout
}

output "defaultauthorizationaction" {
  value = data.citrixadc_vpnsessionaction.foo.defaultauthorizationaction
}
```

## Argument Reference

* `name` - (Required) Name for the session action

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnsessionaction. It is the same as the `name` attribute.
* `advancedclientlessvpnmode` - Option to enable/disable Advanced ClientlessVpnMode. Additionaly, it can be set to STRICT to block Classic ClientlessVpnMode while in AdvancedClientlessMode.
* `allowedlogingroups` - Specify groups that have permission to log on to Citrix Gateway. Users who do not belong to this group or groups are denied access even if they have valid credentials.
* `allprotocolproxy` - IP address of the proxy server to use for all protocols supported by Citrix Gateway.
* `alwaysonprofilename` - Name of the AlwaysON profile associated with the session action. The builtin profile named none can be used to explicitly disable AlwaysON for the session action.
* `authorizationgroup` - Comma-separated list of groups in which the user is placed when none of the groups that the user is a part of is configured on Citrix Gateway. The authorization policy can be bound to these groups to control access to the resources.
* `autoproxyurl` - URL to auto proxy config file
* `citrixreceiverhome` - Web address for the Citrix Receiver home page. Configure Citrix Gateway so that when users log on to the appliance, the Citrix Gateway Plug-in opens a web browser that allows single sign-on to the Citrix Receiver home page.
* `clientchoices` - Provide users with multiple logon options. With client choices, users have the option of logging on by using the Citrix Gateway Plug-in for Windows, Citrix Gateway Plug-in for Java, the Web Interface, or clientless access from one location. Depending on how Citrix Gateway is configured, users are presented with up to three icons for logon choices. The most common are the Citrix Gateway Plug-in for Windows, Web Interface, and clientless access.
* `clientcleanupprompt` - Prompt for client-side cache clean-up when a client-initiated session closes.
* `clientconfiguration` - Allow users to change client Debug logging level in Configuration tab of the Citrix Gateway Plug-in for Windows.
* `clientdebug` - Set the trace level on Citrix Gateway. Technical support technicians use these debug logs for in-depth debugging and troubleshooting purposes.
* `clientidletimeout` - Time, in minutes, after which to time out the user session if Citrix Gateway does not detect mouse or keyboard activity.
* `clientlessmodeurlencoding` - When clientless access is enabled, you can choose to encode the addresses of internal web applications or to leave the address as clear text.
* `clientlesspersistentcookie` - State of persistent cookies in clientless access mode.
* `clientlessvpnmode` - Enable clientless access for web, XenApp or XenDesktop, and FileShare resources without installing the Citrix Gateway Plug-in.
* `clientoptions` - Display only the configured menu options when you select the "Configure Citrix Gateway" option in the Citrix Gateway Plug-in system tray icon for Windows.
* `clientsecurity` - Specify the client security check for the user device to permit a Citrix Gateway session.
* `clientsecuritygroup` - The client security group that will be assigned on failure of the client security check.
* `clientsecuritylog` - Set the logging of client security checks.
* `clientsecuritymessage` - The client security message that will be displayed on failure of the client security check.
* `defaultauthorizationaction` - Specify the authorization action that is applied by default if no authorization policy matches.
* `dnsvservername` - Name of the DNS virtual server for the user session.
* `emailhome` - Web address for the web-based email, such as Outlook Web Access.
* `epaclienttype` - Choose between two types of End point analysis (EPA) scans. Possible values: [ AGENT, AGENTLESS ]
* `forcecleanup` - Force cache clean-up when the user session closes.
* `forcedtimeout` - Force a disconnection from the Citrix Gateway after a specified number of minutes.
* `forcedtimeoutwarning` - Display a warning when the user session is about to expire.
* `ftpproxy` - Proxy server for FTP traffic.
* `gopherproxy` - Proxy server for Gopher traffic.
* `homepage` - Web address of the home page that appears when users log on.
* `httpport` - Destination port numbers for HTTP traffic.
* `httponlycookie` - Use httponly cookie with the session cookie.
* `icaproxy` - Enable ICA proxy to configure secure Internet access to servers running Citrix XenApp or XenDesktop by using Citrix Receiver instead of the Citrix Gateway Plug-in.
* `iconwithreceiver` - Option to customize icon with the Citrix Receiver.
* `iipdnssuffix` - An intranet IP DNS suffix.
* `kcdaccount` - Kerberos constrained delegation account name
* `linuxpluginupgrade` - Option to set plugin upgrade behaviour for Linux
* `locallanaccess` - Set local LAN access.
* `loginscript` - Path to the logon script that is run when a session is established.
* `logoutscript` - Path to the logout script that runs when the user logs out.
* `macpluginupgrade` - Option to set plugin upgrade behaviour for Mac
* `mdxtokensecret` - The MDX token secret for the session.
* `netmask` - The netmask for the intranet IP addresses.
* `ntdomain` - Domain of the user
* `pcoipprofilename` - Name of the PCoIP profile associated with the session action.
* `proxy` - Proxy server for HTTP/SSL traffic.
* `proxyexception` - Proxy exception string that will be configured in the browser.
* `proxylocalbypass` - Bypass proxy server for local addresses.
* `rdpclientprofilename` - Name of the RDP profile associated with the session action.
* `rfc1918` - Enable users to use RFC 1918 addresses.
* `secondauthorizationaction` - Specify the authorization action that is applied if a second authorization policy matches.
* `sesstimeout` - Number of minutes after which the session times out.
* `smartgroup` - Name of the configured SmartGroup.
* `splitdns` - Route the DNS requests to the local DNS server configured on the user device, or Citrix Gateway (remote), or both.
* `splittunnel` - Send, through the tunnel, traffic only for intranet applications that are defined in Citrix Gateway.
* `sso` - Use single sign-on (SSO) to log users on to all web applications automatically after they authenticate, or pass users to the web application logon page to authenticate to each application individually.
* `ssocredential` - Specify whether to use the primary or secondary credentials as SSO credentials.
* `ssodomain` - Domain to use for single sign-on.
* `storefronturl` - Web address of the StoreFront server that provides resources for this VPN session.
* `transparentinterception` - Allow access to network resources by using a single IP address and subnet mask or a range of IP addresses.
* `useiip` - Define IP address pool options.
* `usemip` - Define IP address pool options.
* `useraccounting` - User account for generating user audit logs
* `wihome` - Web address of the Web Interface server, such as http://citrix-xd.com/Citrix/XenApp, or Receiver for Web, which enumerates the virtualized resources, such as XenApp, XenDesktop, and cloud applications.
* `wihomeaddresstype` - Type of the wihome address (IPV4/IPV6). Possible values: [ IPV4, IPV6 ]
* `windowsautologon` - Enable or disable the Windows Auto Logon for the session.
* `windowsclienttype` - Specify the client type for Windows clients.
* `windowspluginupgrade` - Option to set plugin upgrade behaviour for Win
* `winsip` - WINS server IP address to add to Citrix Gateway for name resolution.
* `wiportalmode` - Layout on the Access Interface. Possible values: [ NORMAL, COMPACT ]
