---
subcategory: "High Availability"
---

# Data Source: hanode_routemonitor6_binding

The hanode_routemonitor6_binding data source allows you to retrieve information about a specific IPv6 route monitor binding configured on a High Availability (HA) node.

## Example Usage

```terraform
data "citrixadc_hanode_routemonitor6_binding" "tf_hanode_routemonitor6_binding" {
  hanode_id    = 0
  routemonitor = "fd7f:6bd8:cea9:f32d::/64"
}

output "hanode_id" {
  value = data.citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding.hanode_id
}

output "routemonitor" {
  value = data.citrixadc_hanode_routemonitor6_binding.tf_hanode_routemonitor6_binding.routemonitor
}
```

## Argument Reference

* `hanode_id` - (Required) Number that uniquely identifies the local node. The ID of the local node is always 0.
* `routemonitor` - (Required) The IPv6 address of the route to be monitored.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hanode_routemonitor6_binding. It is a system-generated identifier.
* `netmask` - The netmask for the IPv6 route monitor.
