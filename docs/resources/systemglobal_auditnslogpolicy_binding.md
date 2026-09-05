---
subcategory: "System"
---

# Resource: systemglobal_auditnslogpolicy_binding

The systemglobal_auditnslogpolicy_binding resource is used to bind an auditnslogpolicy to systemglobal.


## Example usage

```hcl
resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
  policyname = "tf_auditnslogpolicy"
  priority   = 50
}
```


## Argument Reference

* `policyname` - (Required) The name of the auditnslog policy.
* `priority` - (Required) The priority of the policy.
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. This is a list of strings.
* `feature` - (Optional) The feature to be checked while applying this config.
* `globalbindtype` - (Optional) The global bind type. Defaults to `SYSTEM_GLOBAL`.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation.
* `nextfactor` - (Optional) On success invoke label. Applicable for advanced authentication policy binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_auditnslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A systemglobal_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding tf_auditnslogpolicy
```
