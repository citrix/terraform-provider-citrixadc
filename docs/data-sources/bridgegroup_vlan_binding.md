---
subcategory: "Network"
---

# Data Source: bridgegroup_vlan_binding

The bridgegroup_vlan_binding data source allows you to retrieve information about the binding between a bridge group and a VLAN.

## Example Usage

```terraform
data "citrixadc_bridgegroup_vlan_binding" "tf_binding" {
  bridgegroup_id = 2
  vlan           = 20
}

output "bridgegroup_id" {
  value = data.citrixadc_bridgegroup_vlan_binding.tf_binding.bridgegroup_id
}

output "vlan" {
  value = data.citrixadc_bridgegroup_vlan_binding.tf_binding.vlan
}
```

## Argument Reference

* `bridgegroup_id` - (Required) The integer that uniquely identifies the bridge group.
* `vlan` - (Required) The VLAN ID to bind to the bridge group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup_vlan_binding. It is a system-generated identifier.
