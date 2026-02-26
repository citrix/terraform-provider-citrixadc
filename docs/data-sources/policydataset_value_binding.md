---
subcategory: "Policy"
---

# Data Source: policydataset_value_binding

The policydataset_value_binding data source allows you to retrieve information about a specific value binding in a policy dataset.

## Example Usage

```terraform
data "citrixadc_policydataset_value_binding" "example" {
  name     = "my_dataset"
  value    = "100"
  endrange = "150"
}

output "index" {
  value = data.citrixadc_policydataset_value_binding.example.index
}

output "comment" {
  value = data.citrixadc_policydataset_value_binding.example.comment
}
```

## Argument Reference

* `name` - (Required) Name of the dataset to which to bind the value.
* `value` - (Required) Value of the specified type that is associated with the dataset. For ipv4 and ipv6, value can be a subnet using the slash notation address/n, where address is the beginning of the subnet and n is the number of left-most bits set in the subnet mask, defining the end of the subnet. The start address will be masked by the subnet mask if necessary, for example for 192.128.128.0/10, the start address will be 192.128.0.0.
* `endrange` - (Required) The dataset entry is a range from <value> through <end_range>, inclusive. endRange cannot be used if value is an ipv4 or ipv6 subnet and endRange cannot itself be a subnet.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this dataset or a data bound to this dataset.
* `id` - The id of the policydataset_value_binding. It is a system-generated identifier.
* `index` - The index of the value (ipv4, ipv6, number) associated with the set.
