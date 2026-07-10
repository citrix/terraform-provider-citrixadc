---
subcategory: "Audit"
---

# Data Source: auditsyslogglobal_auditsyslogpolicy_binding

The auditsyslogglobal_auditsyslogpolicy_binding data source allows you to retrieve information about the binding between auditsyslogglobal configuration and auditsyslogpolicy.

## Example Usage

```terraform
data "citrixadc_auditsyslogglobal_auditsyslogpolicy_binding" "tf_auditsyslogglobal_auditsyslogpolicy_binding" {
  policyname     = "tf_auditsyslogpolicy"
  globalbindtype = "SYSTEM_GLOBAL"
}

output "priority" {
  value = data.citrixadc_auditsyslogglobal_auditsyslogpolicy_binding.tf_auditsyslogglobal_auditsyslogpolicy_binding.priority
}

output "policyname" {
  value = data.citrixadc_auditsyslogglobal_auditsyslogpolicy_binding.tf_auditsyslogglobal_auditsyslogpolicy_binding.policyname
}
```

## Argument Reference

* `policyname` - (Required) Name of the audit syslog policy.
* `globalbindtype` - (Required) The global bind type identifier. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditsyslogglobal_auditsyslogpolicy_binding. It is the concatenation of  `globalbindtype` and `policyname` attributes separated by a comma.
* `priority` - Specifies the priority of the policy.
* `feature` - The feature to be checked while applying this config.
* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]
