---
subcategory: "Load Balancing"
---

# Resource: lbgroup

The lbgroup resource is used to create an lb group entity.


## Example usage

```hcl
resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 15.0
	persistmask = "255.255.254.0"
	cookiename = "tf_cookie_1"
	v6persistmasklen = 96
	timeout = 15.0
}
```


## Argument Reference

* `name` - (Required) Name of the load balancing virtual server group.
* `persistencetype` - (Optional) Type of persistence for the group. Available settings function as follows: * SOURCEIP - Create persistence sessions based on the client IP. * COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client. * RULE - Create persistence sessions based on a user defined rule. * NONE - Disable persistence for the group. Possible values: [ SOURCEIP, COOKIEINSERT, RULE, NONE ]
* `persistencebackup` - (Optional) Type of backup persistence for the group. Possible values: [ SOURCEIP, NONE ]
* `backuppersistencetimeout` - (Optional) Time period, in minutes, for which backup persistence is in effect.
* `persistmask` - (Optional) Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.
* `cookiename` - (Optional) Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
* `v6persistmasklen` - (Optional) Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.
* `cookiedomain` - (Optional) Domain attribute for the HTTP cookie.
* `timeout` - (Optional) Time period for which a persistence session is in effect.
* `rule` - (Optional) Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `usevserverpersistency` - (Optional) Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers. Possible values: [ ENABLED, DISABLED ]
* `mastervserver` - (Optional) When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions.
* `newname` - (Optional) New name for the load balancing virtual server group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbgroup. It has the same value as the `name` attribute.


## Import

A lbgroup can be imported using its name, e.g.

```shell
terraform import citrixadc_lbgroup.tf_lbgroup tf_lbgroup
```
