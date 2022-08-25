---
subcategory: "Traffic Management"
---

# Resource: tmglobal_auditnslogpolicy_binding

The tmglobal_auditnslogpolicy_binding resource is used to create tmglobal_auditnslogpolicy_binding.


## Example usage

```hcl
resource "citrixadc_tmglobal_auditnslogpolicy_binding" "tf_tmglobal_auditnslogpolicy_binding" {
  policyname = "tf_auditnslogpolicy"
  priority   = 100
}
```


## Argument Reference

* `policyname` - (Required) The name of the policy.
* `priority` - (Required) The priority of the policy.
* `gotopriorityexpression` - (Optional) Applicable only to advance tmsession policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmglobal_auditnslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A tmglobal_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_tmglobal_auditnslogpolicy_binding.tf_tmglobal_auditnslogpolicy_binding tf_auditnslogpolicy
```
