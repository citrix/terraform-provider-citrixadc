---
subcategory: "API Definition"
---

# Resource: apiprofile_apispec_binding

Associates an API specification (an imported OpenAPI/Swagger definition) with an API profile on the Citrix ADC. Binding an apispec to an apiprofile tells the profile which API schema to use when validating and protecting matching API traffic, so an apiprofile is only enforced against the endpoints and payloads described by the specs bound to it.


## Example usage

```hcl
resource "citrixadc_apispec" "tf_apispec" {
  name = "petstore_spec"
}

resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "my_apiprofile"
  apivisibility = "ENABLED"
}

resource "citrixadc_apiprofile_apispec_binding" "tf_binding" {
  name    = citrixadc_apiprofile.tf_apiprofile.name
  apispec = citrixadc_apispec.tf_apispec.name
}
```


## Argument Reference

* `name` - (Required) Name of the API profile in which to bind the API apispec(s). This value cannot be changed after the resource is created; updating it forces resource replacement.
* `apispec` - (Required) Name of the API spec to bind to the profile. This value cannot be changed after the resource is created; updating it forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the apiprofile_apispec_binding. It is a comma-separated string of `key:value` pairs (with URL-encoded values) in the form `apispec:<apispec>,name:<name>`.


## Import

An apiprofile_apispec_binding can be imported using its id, in the format `apispec:<apispec>,name:<name>`, e.g.

```shell
terraform import citrixadc_apiprofile_apispec_binding.tf_binding apispec:petstore_spec,name:my_apiprofile
```
