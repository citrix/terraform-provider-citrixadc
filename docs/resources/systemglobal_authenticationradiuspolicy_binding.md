---
subcategory: "System"
---

# Resource: systemglobal_authenticationradiuspolicy_binding

The systemglobal_authenticationradiuspolicy_binding resource is used to bind an authenticationradiuspolicy to systemglobal.


## Example usage

```hcl
resource "citrixadc_systemglobal_authenticationradiuspolicy_binding" "tf_systemglobal_authenticationradiuspolicy_binding" {
  policyname = citrixadc_authenticationradiuspolicy.tf_radiuspolicy.name
  priority   = 50
}

resource "citrixadc_authenticationradiusaction" "tf_radiusaction" {
  name         = "tf_radiusaction"
  radkey       = "secret"
  serverip     = "1.2.3.4"
  serverport   = 8080
  authtimeout  = 2
  radnasip     = "DISABLED"
  passencoding = "chap"
}
resource "citrixadc_authenticationradiuspolicy" "tf_radiuspolicy" {
  name      = "tf_radiuspolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationradiusaction.tf_radiusaction.name
}
```


## Argument Reference

* `policyname` - (Required) The name of the command policy.
* `priority` - (Required) The priority of the command policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
* `feature` - (Optional) The feature to be checked while applying this config.
* `globalbindtype` - (Optional) The global bind type for the binding.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation.
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_authenticationradiuspolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_authenticationradiuspolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_systemglobal_authenticationradiuspolicy_binding.tf_systemglobal_authenticationradiuspolicy_binding tf_radiuspolicy
```
