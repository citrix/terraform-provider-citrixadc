---
subcategory: "System"
---

# Data Source: systemglobal_auditsyslogpolicy_binding

The systemglobal_auditsyslogpolicy_binding data source allows you to retrieve information about an audit syslog policy bound to the Citrix ADC system global scope. The binding is looked up by `policyname`.


## Example usage

```terraform
data "citrixadc_systemglobal_auditsyslogpolicy_binding" "syslog_global" {
  policyname = "syslogpol1"
}

output "syslog_binding_priority" {
  value = data.citrixadc_systemglobal_auditsyslogpolicy_binding.syslog_global.priority
}
```


## Argument Reference

* `policyname` - (Required) The name of the audit syslog policy bound to the system global scope. Used to filter the binding to read.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_auditsyslogpolicy_binding. It is a single token that has the same value as the `policyname` attribute.
* `priority` - The priority assigned to the policy at the global scope.
* `nextfactor` - On-success invoke label. Applicable for advanced authentication policy binding.
* `gotopriorityexpression` - Expression specifying the next policy to be evaluated if the current policy evaluates to TRUE (NEXT or END).
* `feature` - The feature checked while applying this configuration.
* `globalbindtype` - The global bind type for this binding (e.g., `SYSTEM_GLOBAL`).
