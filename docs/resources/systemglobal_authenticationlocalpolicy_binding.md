---
subcategory: "System"
---

# Resource: systemglobal_authenticationlocalpolicy_binding

The systemglobal_authenticationlocalpolicy_binding resource is used to bind an authenticationlocalpolicy to systemglobal.


## Example usage

```hcl
resource "citrixadc_authenticationlocalpolicy" "tf_authenticationlocalpolicy" {
  name = "tf_authenticationlocalpolicy"
  rule = "ns_true"
}

resource "citrixadc_systemglobal_authenticationlocalpolicy_binding" "tf_bind" {
  policyname = citrixadc_authenticationlocalpolicy.tf_authenticationlocalpolicy.name
  priority   = 50
}
```


## Argument Reference

* `policyname` - (Required) The name of the authentication local policy.
* `priority` - (Required) The priority of the policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
* `feature` - (Optional) The feature to be checked while applying this config.
* `globalbindtype` - (Optional) The global bind type. Defaults to `"SYSTEM_GLOBAL"`.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation.
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_authenticationlocalpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_authenticationlocalpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_systemglobal_authenticationlocalpolicy_binding.tf_bind tf_authenticationlocalpolicy
```
