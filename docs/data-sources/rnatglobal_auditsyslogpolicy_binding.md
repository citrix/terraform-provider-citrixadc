---
subcategory: "Network"
---

# Data Source: rnatglobal_auditsyslogpolicy_binding

The rnatglobal_auditsyslogpolicy_binding data source allows you to retrieve information about a rnatglobal_auditsyslogpolicy_binding.


## Example Usage

```terraform
data "citrixadc_rnatglobal_auditsyslogpolicy_binding" "tf_binding" {
  policy = "tf_auditsyslogpolicy"
}

output "policy" {
  value = data.citrixadc_rnatglobal_auditsyslogpolicy_binding.tf_binding.policy
}

output "priority" {
  value = data.citrixadc_rnatglobal_auditsyslogpolicy_binding.tf_binding.priority
}
```


## Argument Reference

* `policy` - (Required) The policy Name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatglobal_auditsyslogpolicy_binding. It is a system-generated identifier.
* `priority` - The priority of the policy.

