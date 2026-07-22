---
subcategory: "SSL"
---

# Data Source: sslglobal_sslpolicy_binding

The sslglobal_sslpolicy_binding data source allows you to retrieve information about the global binding of an SSL policy on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslglobal_sslpolicy_binding" "tf_binding" {
  policyname = "tf_sslpolicy"
  type       = "CONTROL_OVERRIDE"
  priority   = 100
}

output "gotopriorityexpression" {
  value = data.citrixadc_sslglobal_sslpolicy_binding.tf_binding.gotopriorityexpression
}
```


## Argument Reference

* `policyname` - (Required) The name for the SSL policy.
* `type` - (Required) Global bind point to which the policy is bound.
* `priority` - (Required) The priority of the policy binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslglobal_sslpolicy_binding. It is the concatenation of the `policyname`, `type`, and `priority` attributes, formatted as `policyname:<policyname>,type:<type>,priority:<priority>`.
* `globalbindtype` - The global bind point type returned by the appliance.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `invoke` - Whether policies bound to a virtual server, service, or policy label are invoked.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke.
