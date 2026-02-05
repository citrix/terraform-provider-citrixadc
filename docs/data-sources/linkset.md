---
subcategory: "Network"
---

# Data Source `linkset`

The linkset data source allows you to retrieve information about Link sets.


## Example usage

```terraform
data "citrixadc_linkset" "tf_linkset" {
  linkset_id = "LS/1"
}

output "linkset_id" {
  value = data.citrixadc_linkset.tf_linkset.linkset_id
}

output "interfacebinding" {
  value = data.citrixadc_linkset.tf_linkset.interfacebinding
}
```


## Argument Reference

* `linkset_id` - (Required) Unique identifier for the linkset. Must be of the form LS/x, where x can be an integer from 1 to 32.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `interfacebinding` - Set of interface bindings for the linkset.
* `id` - The id of the linkset. It has the same value as the `linkset_id` attribute.


## Import

A linkset can be imported using its linkset_id, e.g.

```shell
terraform import citrixadc_linkset.tf_linkset LS/1
```
