---
subcategory: "Network"
---

# Data Source: netbridge_vlan_binding

The netbridge_vlan_binding data source allows you to retrieve information about a VLAN binding to a network bridge.

## Example Usage

```terraform
data "citrixadc_netbridge_vlan_binding" "tf_binding" {
  name = "tf_netbridge"
  vlan = 20
}

output "name" {
  value = data.citrixadc_netbridge_vlan_binding.tf_binding.name
}

output "vlan" {
  value = data.citrixadc_netbridge_vlan_binding.tf_binding.vlan
}
```

## Argument Reference

* `name` - (Required) The name of the network bridge.
* `vlan` - (Required) The VLAN that is extended by this network bridge.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_vlan_binding. It is a system-generated identifier.
