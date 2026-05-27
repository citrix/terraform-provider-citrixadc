---
subcategory: "API Definition"
---

# Data Source: apispec

The apispec data source allows you to retrieve information about an existing Citrix ADC API specification.


## Example usage

```terraform
data "citrixadc_apispec" "tf_apispec" {
  name = "my_apispec"
}

output "apispec_file" {
  value = data.citrixadc_apispec.tf_apispec.file
}

output "apispec_type" {
  value = data.citrixadc_apispec.tf_apispec.type
}
```


## Argument Reference

* `name` - (Required) Name of the API spec to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `encrypted` - Indicates whether the API spec is encrypted (NetScaler format).
* `file` - Name of and, optionally, path to the api spec file on the appliance's hard-disk drive or solid-state drive. `/nsconfig/apispec/` is the default path.
* `skipvalidation` - Indicates whether openapi spec validation was skipped while adding the spec.
* `type` - Input format of the spec file. One of `PROTO`, `OAS/Swagger`, or `GRAPHQL`.
* `id` - The id of the apispec. It has the same value as the `name` attribute.
