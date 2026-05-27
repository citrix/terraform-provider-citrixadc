---
subcategory: "API Definition"
---

# Resource: apispec

The apispec resource is used to create and manage Citrix ADC API specifications. An API spec describes the contract of an API (in formats such as OAS/Swagger, PROTO, or GraphQL) that the appliance uses for schema validation and routing.


## Example usage

```hcl
resource "citrixadc_apispec" "tf_apispec" {
  name           = "my_apispec"
  file           = "/nsconfig/apispec/petstore.yaml"
  type           = "OAS"
  skipvalidation = "NO"
}
```


## Argument Reference

* `name` - (Required) Name for the spec. Must begin with an ASCII alphanumeric or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. Cannot be changed after the spec is created; updating it forces resource replacement.
* `file` - (Required) Name of and, optionally, path to the api spec file. The spec file should be present on the appliance's hard-disk drive or solid-state drive. Storing a spec file in any location other than the default might cause inconsistency in a high availability setup. `/nsconfig/apispec/` is the default path.
* `encrypted` - (Optional) Specify the encrypted API spec. Must be in NetScaler format. This attribute is only honored on resource create; updating it forces resource replacement.
* `skipvalidation` - (Optional) Disabling openapi spec validation while adding it.
* `type` - (Optional) Input format of the spec file. The three formats supported by the appliance are: `PROTO`, `OAS/Swagger`, and `GRAPHQL`. Defaults to `"OAS"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the apispec. It has the same value as the `name` attribute.


## Import

An apispec can be imported using its name, e.g.

```shell
terraform import citrixadc_apispec.tf_apispec my_apispec
```
