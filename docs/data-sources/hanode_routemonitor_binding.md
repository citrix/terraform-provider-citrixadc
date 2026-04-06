---
subcategory: "High Availability"
---

# Data Source: hanode_routemonitor_binding

This data source is used to retrieve information about a specific `hanode_routemonitor_binding` configuration.

## Example usage

```hcl
data "citrixadc_hanode_routemonitor_binding" "example" {
  hanode_id    = 0
  routemonitor = "10.222.74.128"
  netmask      = "255.255.255.192"
}

output "binding_id" {
  value = data.citrixadc_hanode_routemonitor_binding.example.id
}

output "binding_netmask" {
  value = data.citrixadc_hanode_routemonitor_binding.example.netmask
}
```

## Argument Reference

* `hanode_id` - (Required) Number that uniquely identifies the local node. The ID of the local node is always 0.
* `routemonitor` - (Required) The IP address (IPv4 or IPv6).
* `netmask` - (Required) The netmask.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding.
