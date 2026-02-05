---
subcategory: "Load Balancing"
---

# Data Source `lbgroup`

The lbgroup data source allows you to retrieve information about a load balancing virtual server group.


## Example usage

```terraform
data "citrixadc_lbgroup" "tf_lbgroup" {
  name = "my_lbgroup"
}

output "persistencetype" {
  value = data.citrixadc_lbgroup.tf_lbgroup.persistencetype
}

output "timeout" {
  value = data.citrixadc_lbgroup.tf_lbgroup.timeout
}
```


## Argument Reference

* `name` - (Required) Name of the load balancing virtual server group.

## Attribute Reference

The following attributes are available:

* `name` - Name of the load balancing virtual server group.
* `persistencetype` - Type of persistence for the group. Available settings function as follows:
  * SOURCEIP - Create persistence sessions based on the client IP.
  * COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client.
  * RULE - Create persistence sessions based on a user defined rule.
  * NONE - Disable persistence for the group.
* `persistencebackup` - Type of backup persistence for the group.
* `backuppersistencetimeout` - Time period, in minutes, for which backup persistence is in effect.
* `persistmask` - Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.
* `cookiename` - Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
* `cookiedomain` - Domain attribute for the HTTP cookie.
* `timeout` - Time period for which a persistence session is in effect.
* `rule` - Expression, or name of a named expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI:
  * If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
  * If the expression itself includes double quotation marks, escape the quotations by using the \ character.
  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `usevserverpersistency` - Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers.
* `mastervserver` - When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions.
* `v6persistmasklen` - Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.
* `newname` - New name for the load balancing virtual server group.
* `id` - The id of the lbgroup. It is a system-generated identifier.
