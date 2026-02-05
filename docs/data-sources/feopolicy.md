---
subcategory: "Front End Optimization"
---

# Data Source `feopolicy`

The feopolicy data source allows you to retrieve information about an existing front end optimization policy.


## Example usage

```terraform
data "citrixadc_feopolicy" "tf_feopolicy" {
  name = "my_feopolicy"
}

output "id" {
  value = data.citrixadc_feopolicy.tf_feopolicy.id
}

output "name" {
  value = data.citrixadc_feopolicy.tf_feopolicy.name
}

output "action" {
  value = data.citrixadc_feopolicy.tf_feopolicy.action
}

output "rule" {
  value = data.citrixadc_feopolicy.tf_feopolicy.rule
}
```


## Argument Reference

* `name` - (Required) The name of the front end optimization policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the feopolicy. It has the same value as the `name` attribute.
* `action` - The front end optimization action that has to be performed when the rule matches.
* `rule` - The rule associated with the front end optimization policy.


## Import

A feopolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_feopolicy.tf_feopolicy my_feopolicy
```
