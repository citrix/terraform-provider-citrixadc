---
subcategory: "AAA"
---

# Data Source: aaagroup

The aaagroup data source allows you to retrieve information about AAA groups.


## Example usage

```terraform
data "citrixadc_aaagroup" "tf_aaagroup" {
  groupname = "test_group"
}

output "weight" {
  value = data.citrixadc_aaagroup.tf_aaagroup.weight
}
```


## Argument Reference

* `groupname` - (Required) Name for the group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `weight` - Weight of the group.
* `id` - The id of the aaagroup. It has the same value as the `groupname` attribute.
