---
subcategory: "Basic"
---

# Data Source: service_dospolicy_binding

The service_dospolicy_binding data source allows you to retrieve information about the binding between a service and a DoS policy.


## Example usage

```terraform
data "citrixadc_service_dospolicy_binding" "tf_binding" {
  name       = "tf_service"
  policyname = "tf_dospolicy"
}

output "policyname" {
  value = data.citrixadc_service_dospolicy_binding.tf_binding.policyname
}
```


## Argument Reference

* `name` - (Required) Name of the service to which to bind a policy or monitor.
* `policyname` - (Required) The name of the policy bound to the service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the service_dospolicy_binding. It is the concatenation of `name` and `policyname` attributes separated by comma.
