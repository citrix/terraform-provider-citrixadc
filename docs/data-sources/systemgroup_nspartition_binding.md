---
subcategory: "System"
---

# Data Source: systemgroup_nspartition_binding

The systemgroup_nspartition_binding data source allows you to retrieve information about a system group partition binding.

## Example Usage

```terraform
data "citrixadc_systemgroup_nspartition_binding" "tf_systemgroup_nspartition_binding" {
  groupname     = "tf_systemgroup"
  partitionname = "tf_nspartition"
}

output "groupname" {
  value = data.citrixadc_systemgroup_nspartition_binding.tf_systemgroup_nspartition_binding.groupname
}

output "partitionname" {
  value = data.citrixadc_systemgroup_nspartition_binding.tf_systemgroup_nspartition_binding.partitionname
}
```

## Argument Reference

* `groupname` - (Required) Name of the system group.
* `partitionname` - (Required) Name of the Partition to bind to the system group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup_nspartition_binding. It has the format `<groupname>,<partitionname>`.
