---
subcategory: "Audit"
---

# Data Source: auditnslogglobal_auditnslogpolicy_binding

The auditnslogglobal_auditnslogpolicy_binding data source allows you to retrieve information about the binding between the global audit nslog configuration and an audit nslog policy.

## Example Usage

```terraform
data "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
  policyname     = "my_auditnslogpolicy"
  globalbindtype = "SYSTEM_GLOBAL"
}

output "priority" {
  value = data.citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding.priority
}

output "policyname" {
  value = data.citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding.policyname
}
```

## Argument Reference

* `policyname` - (Required) Name of the audit nslog policy.
* `globalbindtype` - (Required) The global bind type. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogglobal_auditnslogpolicy_binding. It is the concatenation of the `globalbindtype` and `policyname` attributes separated by a comma.
* `priority` - Specifies the priority of the policy.
* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
