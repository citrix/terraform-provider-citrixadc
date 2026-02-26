---
subcategory: "System"
---

# Data Source: systemglobal_auditnslogpolicy_binding

The systemglobal_auditnslogpolicy_binding data source allows you to retrieve information about a systemglobal_auditnslogpolicy_binding.


## Example Usage

```terraform
data "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
  policyname = "tf_auditnslogpolicy"
}

output "policyname" {
  value = data.citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding.priority
}
```


## Argument Reference

* `policyname` - (Required) The name of the  command policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_auditnslogpolicy_binding. It is a system-generated identifier.
* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation.
* `nextfactor` - On success invoke label. Applicable for advanced authentication policy binding.
* `priority` - The priority of the command policy.
