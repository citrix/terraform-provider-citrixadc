---
subcategory: "VPN"
---

# Data Source: vpnvserver_authenticationpolicy_binding

The vpnvserver_authenticationpolicy_binding data source allows you to retrieve information about an authentication policy bound to a VPN (NetScaler Gateway) virtual server.


## Example usage

```terraform
data "citrixadc_vpnvserver_authenticationpolicy_binding" "tf_bind" {
  name      = "gatewayvserver1"
  policy    = "ldapauthpolicy1"
  bindpoint = "REQUEST"
}

output "priority" {
  value = data.citrixadc_vpnvserver_authenticationpolicy_binding.tf_bind.priority
}
```


## Argument Reference

* `name` - (Required) Name of the VPN virtual server.
* `policy` - (Required) The name of the authentication policy bound to the VPN virtual server.
* `bindpoint` - (Required) Bind point to which the policy is bound. Used together with `name` and `policy` as a lookup key. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST, AAA_REQUEST, AAA_RESPONSE ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advanced VPN session policies. An expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.
* `groupextraction` - Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `priority` - Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.
* `id` - The id of the vpnvserver_authenticationpolicy_binding. It is a comma-separated list of `key:value` pairs in the form `bindpoint:<bindpoint>,name:<name>,policy:<policy>`, where each value is URL-encoded.
