---
subcategory: "VPN"
---

# Data Source: vpnvserver

The vpnvserver data source allows you to retrieve information about a VPN virtual server.

## Example usage

```terraform
data "citrixadc_vpnvserver" "example" {
  name = "tf.citrix.example.com"
}

output "servicetype" {
  value = data.citrixadc_vpnvserver.example.servicetype
}

output "ipv46" {
  value = data.citrixadc_vpnvserver.example.ipv46
}
```

## Argument Reference

* `name` - (Required) Name for the Citrix Gateway virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `accessrestrictedpageredirect` - By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.
* `advancedepa` - This option tells whether advanced EPA is enabled on this virtual server.
* `appflowlog` - Log AppFlow records that contain standard NetFlow or IPFIX information, such as time stamps for the beginning and end of a flow, packet count, and byte count. Also log records that contain application-level information, such as HTTP web addresses, HTTP request methods and response status codes, server response time, and latency.
* `authentication` - Require authentication for users connecting to Citrix Gateway.
* `authnprofile` - Authentication Profile entity on virtual server. This entity can be used to offload authentication to AAA vserver for multi-factor(nFactor) authentication.
* `certkeynames` - Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate.
* `cginfrahomepageredirect` - When client requests ShareFile resources and Citrix Gateway detects that the user is unauthenticated or the user session has expired, disabling this option takes the user to the originally requested ShareFile resource after authentication (instead of taking the user to the default VPN home page).
* `comment` - Any comments associated with the virtual server.
* `deploymenttype` - Indicates deployment type. Possible values are NONE, ICA_WEBINTERFACE, ICA_STOREFRONT, MOBILITY.
* `devicecert` - Indicates whether device certificate check as a part of EPA is on or off.
* `deviceposture` - Enable device posture.
* `doublehop` - Use the Citrix Gateway appliance in a double-hop configuration. A double-hop deployment provides an extra layer of security for the internal network by using three firewalls to divide the DMZ into two stages. Such a deployment can have one appliance in the DMZ and one appliance in the secure network.
* `downstateflush` - Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN. Do not enable DOWN state flush on servers that must complete their transactions.
* `dtls` - This option starts/stops the turn service on the vserver.
* `failedlogintimeout` - Number of minutes an account will be locked if user exceeds maximum permissible attempts.
* `httpprofilename` - Name of the HTTP profile to assign to this virtual server.
* `icaonly` - When set to ON, it implies Basic mode where the user can log on using either Citrix Receiver or a browser and get access to the published apps configured at the XenApp/XenDEsktop environment pointed out by the WIHome parameter. Users are not allowed to connect using the Citrix Gateway Plug-in and end point scans cannot be configured. When set to OFF, it implies Smart Access mode where the user can log on using either Citrix Receiver or a browser or a Citrix Gateway Plug-in. The admin can configure end point scans to be run on the client systems and then use the results to control access to the published apps.
* `icaproxysessionmigration` - This option determines if an existing ICA Proxy session is transferred when the user logs on from another device.
* `icmpvsrresponse` - Criterion for responding to PING requests sent to this virtual server. If this parameter is set to ACTIVE, respond only if the virtual server is available. With the PASSIVE setting, respond even if the virtual server is not available.
* `ipset` - The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current vpn vserver.
* `ipv46` - IPv4 or IPv6 address of the Citrix Gateway virtual server. Usually a public IP address. User devices send connection requests to this IP address.
* `l2conn` - Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to coexist on the Citrix ADC.
* `linuxepapluginupgrade` - Option to set plugin upgrade behaviour for Linux.
* `listenpolicy` - String specifying the listen policy for the Citrix Gateway virtual server. Can be either a named expression or an expression. The Citrix Gateway virtual server processes only the traffic for which the expression evaluates to true.
* `listenpriority` - Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
* `loginonce` - This option enables/disables seamless SSO for this Vserver.
* `logoutonsmartcardremoval` - Option to VPN plugin behavior when smartcard or its reader is removed.
* `macepapluginupgrade` - Option to set plugin upgrade behaviour for Mac.
* `secureprivateaccess` - Enable secure private access for this virtual server.

### Read-only vpnvserver metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnvserver` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `ip` - The Virtual IP address of the VPN virtual server.
* `value` - Indicates whether or not the certificate is bound or if SSL offload is disabled.
* `type` - The type of virtual server; for example, CONTENT based or ADDRESS based.
* `curstate` - The current state of the virtual server, as UP, DOWN, BUSY, and so on.
* `status` - Whether or not this virtual server responds to ARPs and whether or not round-robin selection is temporarily in effect.
* `cachetype` - Virtual server cache type. The options are: TRANSPARENT, REVERSE, and FORWARD.
* `redirect` - The cache redirect policy.
* `precedence` - The type of policy (URL or RULE) that takes precedence on the content switching virtual server.
* `redirecturl` - The URL where traffic is redirected if the virtual server in system becomes unavailable.
* `curaaausers` - The number of current users logged on to this virtual server.
* `curtotalusers` - The total number of current users connected through this virtual server.
* `domain` - The domain name of the server for which a service needs to be added.
* `rule` - The name of the rule, or expression, if any, that policy for the VPN server is to use.
* `servicename` - The name of the service, if any, to which the virtual server policy is bound.
* `weight` - Weight for this service, if any, used when the system performs load balancing.
* `cachevserver` - The name of the default target cache virtual server, if any, to which requests are redirected.
* `backupvserver` - The name of the backup VPN virtual server for this VPN virtual server.
* `clttimeout` - The idle time, if any, in seconds after which the client connection is terminated.
* `somethod` - The method used to determine whether or not a new connection will spill over the allocated block of intranet IP addresses.
* `sothreshold` - The number of client connections after which the mapped IP address is used as the client source IP address.
* `sopersistence` - Whether or not cookie-based site persistance is enabled for this VPN vserver.
* `sopersistencetimeout` - The timeout, if any, for cookie-based site persistance of this VPN vserver.
* `usemip` - Deprecated. See `map`.
* `map` - Whether or not mapped IP addresses are ON or OFF.
* `bindpoint` - Bindpoint to which the policy is bound.
* `disableprimaryondown` - Tells whether traffic will continue reaching backup virtual servers even after the primary virtual server comes UP from DOWN state.
* `secondary` - Whether the authentication policy is bound as the secondary policy in a two-factor configuration.
* `groupextraction` - Whether the authentication policy is bound to a tertiary chain used only for group extraction.
* `epaprofileoptional` - Whether the EPA profile is marked optional for preauthentication EPA profile.
* `ngname` - Node group devno to which this authentication virtual server belongs.
* `csvserver` - Name of the CS vserver to which the VPN vserver is bound.
* `nodefaultbindings` - Whether the configuration will have default ssl CIPHER and ECC curve bindings.
* `response` - Response.
