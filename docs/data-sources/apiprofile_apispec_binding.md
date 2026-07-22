---
subcategory: "API Definition"
---

# Data Source: apiprofile_apispec_binding

The apiprofile_apispec_binding data source allows you to retrieve information about the binding between an API specification and an API profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_apiprofile_apispec_binding" "tf_binding" {
  name    = "my_apiprofile"
  apispec = "petstore_spec"
}

output "binding_id" {
  value = data.citrixadc_apiprofile_apispec_binding.tf_binding.id
}
```


## Argument Reference

* `name` - (Required) Name of the API profile in which the API apispec is bound.
* `apispec` - (Required) Name of the API spec bound to the profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the apiprofile_apispec_binding. It is a comma-separated string of `key:value` pairs (with URL-encoded values) in the form `apispec:<apispec>,name:<name>`.
