---
subcategory: "AAA"
---

# Data Source: systemuser_nspartition_binding

The systemuser_nspartition_binding data source allows you to retrieve information about the binding between a system user and a partition.

## Example Usage

```terraform
data "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
  username      = "george"
  partitionname = "tf_nspartition"
}

output "username" {
  value = data.citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding.username
}

output "partitionname" {
  value = data.citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding.partitionname
}
```

## Argument Reference

* `username` - (Required) Name of the system-user entry to which to bind the command policy.
* `partitionname` - (Required) Name of the Partition to bind to the system user.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser_nspartition_binding. It is the concatenation of `username` and `partitionname` attributes separated by a comma.
