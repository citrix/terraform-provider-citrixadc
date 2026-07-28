---
subcategory: "Network"
---

# Data Source: rnatglobal_auditsyslogpolicy_binding

The rnatglobal_auditsyslogpolicy_binding data source allows you to retrieve information about an audit syslog policy bound to the global RNAT configuration.


## Example usage

```terraform
data "citrixadc_rnatglobal_auditsyslogpolicy_binding" "rnat_syslog" {
  policy = "syslogpol1"
}

output "rnat_syslog_binding_priority" {
  value = data.citrixadc_rnatglobal_auditsyslogpolicy_binding.rnat_syslog.priority
}
```


## Argument Reference

* `policy` - (Required) The name of the audit syslog policy bound to the RNAT global configuration. Used to filter the binding to read.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatglobal_auditsyslogpolicy_binding. It has the same value as the `policy` attribute.
* `priority` - The priority assigned to the policy at the RNAT global scope.
* `all` - Delete-only flag indicating whether all RNAT global bindings are removed on unbind.
