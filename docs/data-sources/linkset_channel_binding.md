---
subcategory: "Network"
---

# Data Source: linkset_channel_binding

The linkset_channel_binding data source allows you to retrieve information about a specific linkset channel binding.

## Example Usage

```terraform
data "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
  linkset_id = "LS/3"
  ifnum      = "LA/3"
}

output "id" {
  value = data.citrixadc_linkset_channel_binding.tf_linkset_channel_binding.id
}

output "linkset_id" {
  value = data.citrixadc_linkset_channel_binding.tf_linkset_channel_binding.linkset_id
}

output "ifnum" {
  value = data.citrixadc_linkset_channel_binding.tf_linkset_channel_binding.ifnum
}
```

## Argument Reference

* `linkset_id` - (Required) ID of the linkset to which to bind the interfaces.
* `ifnum` - (Required) The interfaces to be bound to the linkset.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the linkset_channel_binding. It is the concatenation of `linkset_id` and `ifnum` attributes separated by a comma.
