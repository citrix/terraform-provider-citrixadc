---
subcategory: "VPN"
---

# Data Source: vpnglobal_authenticationpolicy_binding

The vpnglobal_authenticationpolicy_binding data source allows you to retrieve information about an advanced authentication policy bound to the global VPN bind point.

## Example Usage

```terraform
data "citrixadc_vpnglobal_authenticationpolicy_binding" "tf_bind" {
  policyname = "advauth_pol1"
}

output "priority" {
  value = data.citrixadc_vpnglobal_authenticationpolicy_binding.tf_bind.priority
}

output "secondary" {
  value = data.citrixadc_vpnglobal_authenticationpolicy_binding.tf_bind.secondary
}
```

## Argument Reference

* `policyname` - (Required) The name of the policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Bind the Authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.
* `id` - The id of the vpnglobal_authenticationpolicy_binding. It has the same value as the `policyname` attribute.
* `priority` - Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.
