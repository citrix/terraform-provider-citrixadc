---
subcategory: "VPN"
---

# Data Source: vpnvserver_cachepolicy_binding

The vpnvserver_cachepolicy_binding data source allows you to retrieve information about a specific cachepolicy binding to a vpnvserver.


## Example usage

```terraform
data "citrixadc_vpnvserver_cachepolicy_binding" "tf_bind_data" {
  name      = "tf_examplecom"
  policy    = "tf_cachepolicy"
  bindpoint = "REQUEST"
}

output "priority" {
  value = data.citrixadc_vpnvserver_cachepolicy_binding.tf_bind_data.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnvserver_cachepolicy_binding.tf_bind_data.gotopriorityexpression
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `policy` - (Required) The name of the policy, if any, bound to the VPN virtual server.
* `bindpoint` - (Required) Bindpoint to which the policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Next priority expression.
* `groupextraction` - Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `id` - The id of the vpnvserver_cachepolicy_binding. It is the concatenation of `name`, `policy` and `bindpoint` attributes.
* `priority` - Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.
