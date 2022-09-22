---
subcategory: "VPN"
---

# Resource: vpnglobal_vpnclientlessaccesspolicy_binding

The vpnglobal_vpnclientlessaccesspolicy_binding resource is used to bind vpnclientlessaccesspolicy to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
  name        = "tf_vpnclientlessaccesspolicy"
  profilename = "ns_cvpn_default_profile"
  rule        = "true"
}
resource "citrixadc_vpnglobal_vpnclientlessaccesspolicy_binding" "tf_bind" {
  policyname     = citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy.name
  priority       = 90
  globalbindtype = "RNAT_GLOBAL"
  secondary      = "false"
  type           = "RES_OVERRIDE"
}
```


## Argument Reference

* `policyname` - (Required) The name of the policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
* `feature` - (Optional) The feature to be checked while applying this config
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.
* `type` - (Optional) Bindpoint to which the policy is bound


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnclientlessaccesspolicy_binding. It has the same value as the `policyname` attribute.


## Import

A vpnglobal_vpnclientlessaccesspolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnclientlessaccesspolicy_binding.tf_bind tf_vpnclientlessaccesspolicy
```
