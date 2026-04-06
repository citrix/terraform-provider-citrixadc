---
subcategory: "Network"
---

# Data Source: vxlanvlanmap_vxlan_binding

The vxlanvlanmap_vxlan_binding data source allows you to retrieve information about the binding between a vxlanvlanmap and a vxlan.

## Example usage

```terraform
data "citrixadc_vxlanvlanmap_vxlan_binding" "tf_binding" {
  name  = "tf_vxlanvlanmp"
  vxlan = 123
}

output "vlan" {
  value = data.citrixadc_vxlanvlanmap_vxlan_binding.tf_binding.vlan
}
```

## Argument Reference

* `name` - (Required) Name of the mapping table.
* `vxlan` - (Required) The VXLAN assigned to the vlan inside the cloud.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlanvlanmap_vxlan_binding. It is a system-generated identifier.
* `vlan` - The vlan id or the range of vlan ids in the on-premise network.
