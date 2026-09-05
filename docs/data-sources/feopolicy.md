---
subcategory: "Front-end-optimization"
---

# Data Source: feopolicy

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

### Read-only feopolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_feopolicy` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `builtin` - Flag to determine if the front end optimization policy is built-in or not. A list of strings.
* `feature` - The feature to be checked while applying this config.
* `hits` - Total number of hits.
* `undefhits` - Total number of undefined policy hits.
