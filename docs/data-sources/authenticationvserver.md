---
subcategory: "Authentication"
---

# Data Source: authenticationvserver

The authenticationvserver data source allows you to retrieve information about an existing authentication virtual server.


## Example usage

```terraform
data "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name = "my_authenticationvserver"
}

output "servicetype" {
  value = data.citrixadc_authenticationvserver.tf_authenticationvserver.servicetype
}

output "comment" {
  value = data.citrixadc_authenticationvserver.tf_authenticationvserver.comment
}
```


## Argument Reference

* `name` - (Required) Name for the new authentication virtual server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the authentication virtual server is added by using the rename authentication vserver command.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver. It has the same value as the `name` attribute.
* `appflowlog` - Log AppFlow flow information.
* `authentication` - Require users to be authenticated before sending traffic through this virtual server.
* `authenticationdomain` - The domain of the authentication cookie set by Authentication vserver.
* `certkeynames` - Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate.
* `comment` - Any comments associated with this virtual server.
* `failedlogintimeout` - Number of minutes an account will be locked if user exceeds maximum permissible attempts.
* `ipv46` - IP address of the authentication virtual server, if a single IP address is assigned to the virtual server.
* `maxloginattempts` - Maximum Number of login Attempts.
* `newname` - New name of the authentication virtual server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `port` - TCP port on which the virtual server accepts connections.
* `range` - If you are creating a series of virtual servers with a range of IP addresses assigned to them, the length of the range. The new range of authentication virtual servers will have IP addresses consecutively numbered, starting with the primary address specified with the IP Address parameter.
* `samesite` - SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite.
* `servicetype` - Protocol type of the authentication virtual server. Always SSL.
* `state` - Initial state of the new virtual server.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

### Read-only authenticationvserver metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_authenticationvserver` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `ip` - The Virtual IP address of the authentication vserver.
* `value` - Indicates whether or not the certificate is bound or if SSL offload is disabled.
* `type` - The type of Virtual Server, e.g. CONTENT based or ADDRESS based.
* `curstate` - The current state of the Virtual server, e.g. UP, DOWN, BUSY, etc.
* `status` - Whether or not this vserver responds to ARPs and whether or not round-robin selection is temporarily in effect.
* `cachetype` - Virtual server's cache type. The options are: TRANSPARENT, REVERSE and FORWARD.
* `redirect` - The cache redirect policy.
* `precedence` - The type of policy (URL or RULE) that takes precedence on the content switching virtual server.
* `redirecturl` - The URL where traffic is redirected if the virtual server becomes unavailable.
* `curaaausers` - The number of current users logged in to this vserver.
* `policy` - The name of the policy, if any, bound to the authentication vserver.
* `servicename` - The name of the service, if any, to which the vserver policy is bound.
* `weight` - Weight for this service, if any, used when the system performs load balancing.
* `cachevserver` - The name of the default target cache virtual server, if any, to which requests are redirected.
* `backupvserver` - The name of the backup vpn virtual server for this vpn virtual server.
* `clttimeout` - The idle time, if any, in seconds after which the client connection is terminated.
* `somethod` - The method used to determine whether or not a new connection will spillover the allocated block of Intranet IP addresses.
* `sothreshold` - The number of client connections after which the Mapped IP address is used as the client source IP address.
* `sopersistence` - Whether or not cookie-based site persistance is enabled for this VPN vserver.
* `sopersistencetimeout` - The timeout, if any, for cookie-based site persistance of this VPN vserver.
* `priority` - The priority, if any, of the vpn vserver policy.
* `downstateflush` - Perform delayed clean up of connections on this vserver.
* `bindpoint` - Bindpoint to which the policy is bound.
* `disableprimaryondown` - Tells whether traffic will continue reaching backup vservers even after primary comes UP from DOWN state.
* `listenpolicy` - Listenpolicy configured for authentication vserver.
* `listenpriority` - Priority of listen policy for authentication vserver.
* `tcpprofilename` - The name of the TCP profile.
* `httpprofilename` - Name of the HTTP profile.
* `vstype` - Virtual Server Type, e.g. Load Balancing, Content Switch, Cache Redirection.
* `ngname` - Nodegroup devno to which this authentication vserver belongs to.
* `secondary` - Bind the authentication policy to the secondary chain.
* `groupextraction` - Bind the Authentication policy to a tertiary chain which will be used only for group extraction.


## Import

A authenticationvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationvserver.tf_authenticationvserver my_authenticationvserver
```
