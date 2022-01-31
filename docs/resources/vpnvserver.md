---
subcategory: "Vpn"
---

# Resource: vpnvserver

The `vpnvserver` resource is used to create SSL VPN virtual server.

## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name           = "tf.citrix.example.com"
  servicetype    = "SSL"
  ipv46          = "3.3.3.3"
  port           = 443
  ipset          = citrixadc_ipset.tf_ipset.name
  dtls           = "OFF"
  downstateflush = "DISABLED"
  listenpolicy   = "NONE"
  tcpprofilename = "nstcp_default_XA_XD_profile"
}
```

## Argument Reference

* `name` - (Required) Name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my server" or 'my server').
* `servicetype` - (Required) Protocol used by the Citrix Gateway virtual server.
* `advancedepa` - (Optional) This option tells whether advanced EPA is enabled on this virtual server
* `appflowlog` - (Optional) Log AppFlow records that contain standard NetFlow or IPFIX information, such as time stamps for the beginning and end of a flow, packet count, and byte count. Also log records that contain application-level information, such as HTTP web addresses, HTTP request methods and response status codes, server response time, and latency.
* `authentication` - (Optional) Require authentication for users connecting to Citrix Gateway.
* `authnprofile` - (Optional) Authentication Profile entity on virtual server. This entity can be used to offload authentication to AAA vserver for multi-factor(nFactor) authentication
* `certkeynames` - (Optional) Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate
* `cginfrahomepageredirect` - (Optional) When client requests ShareFile resources and Citrix Gateway detects that the user is unauthenticated or the user session has expired, disabling this option takes the user to the originally requested ShareFile resource after authentication (instead of taking the user to the default VPN home page)
* `comment` - (Optional) Any comments associated with the virtual server.
* `deploymenttype` - (Optional) 0
* `devicecert` - (Optional) Indicates whether device certificate check as a part of EPA is on or off.
* `doublehop` - (Optional) Use the Citrix Gateway appliance in a double-hop configuration. A double-hop deployment provides an extra layer of security for the internal network by using three firewalls to divide the DMZ into two stages. Such a deployment can have one appliance in the DMZ and one appliance in the secure network.
* `downstateflush` - (Optional) Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN.  Do not enable DOWN state flush on servers that must complete their transactions.
* `dtls` - (Optional) This option starts/stops the turn service on the vserver
* `failedlogintimeout` - (Optional) Number of minutes an account will be locked if user exceeds maximum permissible attempts
* `httpprofilename` - (Optional) Name of the HTTP profile to assign to this virtual server.
* `icaonly` - (Optional) - When set to ON, it implies Basic mode where the user can log on using either Citrix Receiver or a browser and get access to the published apps configured at the XenApp/XenDEsktop environment pointed out by the WIHome parameter. Users are not allowed to connect using the Citrix Gateway Plug-in and end point scans cannot be configured. Number of users that can log in and access the apps are not limited by the license in this mode.   - When set to OFF, it implies Smart Access mode where the user can log on using either Citrix Receiver or a browser or a Citrix Gateway Plug-in. The admin can configure end point scans to be run on the client systems and then use the results to control access to the published apps. In this mode, the client can connect to the gateway in other client modes namely VPN and CVPN. Number of users that can log in and access the resources are limited by the CCU licenses in this mode.
* `icaproxysessionmigration` - (Optional) This option determines if an existing ICA Proxy session is transferred when the user logs on from another device.
* `icmpvsrresponse` - (Optional) Criterion for responding to PING requests sent to this virtual server. If this parameter is set to ACTIVE, respond only if the virtual server is available. With the PASSIVE setting, respond even if the virtual server is not available.
* `ipset` - (Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current vpn vserver
* `ipv46` - (Optional) IPv4 or IPv6 address of the Citrix Gateway virtual server. Usually a public IP address. User devices send connection requests to this IP address.
* `l2conn` - (Optional) Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to coexist on the Citrix ADC.
* `linuxepapluginupgrade` - (Optional) Option to set plugin upgrade behaviour for Linux
* `listenpolicy` - (Optional) String specifying the listen policy for the Citrix Gateway virtual server. Can be either a named expression or an expression. The Citrix Gateway virtual server processes only the traffic for which the expression evaluates to true.
* `listenpriority` - (Optional) Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
* `loginonce` - (Optional) This option enables/disables seamless SSO for this Vserver.
* `logoutonsmartcardremoval` - (Optional) Option to VPN plugin behavior when smartcard or its reader is removed
* `macepapluginupgrade` - (Optional) Option to set plugin upgrade behaviour for Mac
* `maxaaausers` - (Optional)
* `maxloginattempts` - (Optional)
* `netprofile` - (Optional) The name of the network profile.
* `pcoipvserverprofilename` - (Optional) Name of the PCoIP vserver profile associated with the vserver.
* `port` - (Optional) TCP port on which the virtual server listens.
* `range` - (Optional) Range of Citrix Gateway virtual server IP addresses. The consecutively numbered range of IP addresses begins with the address specified by the IP Address parameter.  In the configuration utility, select Network VServer to enter a range.
* `rdpserverprofilename` - (Optional) Name of the RDP server profile associated with the vserver.
* `rhistate` - (Optional) A host route is injected according to the setting on the virtual servers.             * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.             * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.             * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance injects even if one virtual server set to ACTIVE is UP.
* `samesite` - (Optional) SameSite attribute value for Cookies generated in VPN context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
* `state` - (Optional) State of the virtual server. If the virtual server is disabled, requests are not processed.
* `tcpprofilename` - (Optional) Name of the TCP profile to assign to this virtual server.
* `userdomains` - (Optional) List of user domains specified as comma seperated value
* `vserverfqdn` - (Optional) Fully qualified domain name for a VPN virtual server. This is used during StoreFront configuration generation.
* `windowsepapluginupgrade` - (Optional) Option to set plugin upgrade behaviour for Win

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `vpnvserver`. It has the same value as the `name` attribute.

## Import

A vpnvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnvserver.tf_vpnserver tf.citrix.example.com
```
