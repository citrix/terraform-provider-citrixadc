---
subcategory: "Vpn"
---

# Resource: vpnglobal_vpnsessionpolicy_binding

The vpnglobal_vpnsessionpolicy_binding resource is used to bind vpn sessionpolicy to the global configuration.


## Example usage

```hcl
resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
  name                       = "newsession"
  sesstimeout                = "10"
  defaultauthorizationaction = "ALLOW"
}

resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
  name   = "tf_vpnsessionpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
}
resource "citrixadc_vpnglobal_vpnsessionpolicy_binding" "tf_bind" {
  policyname = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
  priority = 20
  secondary = true
  feature = "SYSTEM"
}
```


## Argument Reference

* `policyname` - (Required) The name of the policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
* `feature` - (Optional) The feature to be checked while applying this config
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnsessionpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A vpnglobal_vpnsessionpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnsessionpolicy_binding.tf_bind tf_vpnsessionpolicy
```
