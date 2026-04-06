---
subcategory: "Audit"
---

# Data Source: auditnslogglobal_auditnslogpolicy_binding

The auditnslogglobal_auditnslogpolicy_binding data source allows you to retrieve information about the binding between auditnslogglobal configuration and auditnslogpolicy.

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
* `globalbindtype` - (Required) The global bind type identifier. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogglobal_auditnslogpolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
