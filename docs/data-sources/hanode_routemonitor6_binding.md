---
subcategory: "High Availability"
---

# Data Source: hanode_routemonitor6_binding

The hanode_routemonitor6_binding data source allows you to retrieve information about a specific IPv6 route monitor binding configured on a High Availability (HA) node.

## Example usage

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
* `routemonitor` - (Required) The IP address (IPv4 or IPv6).
* `netmask` - (Optional) The netmask.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hanode_routemonitor6_binding. It is the concatenation of `hanode_id` and `routemonitor` attributes separated by a comma.

### Read-only hanode_routemonitor6_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_hanode_routemonitor6_binding` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `routemonitorstate` - State for the route monitor (`UP`, `DOWN`).
* `flags` - The flags for this entry.
