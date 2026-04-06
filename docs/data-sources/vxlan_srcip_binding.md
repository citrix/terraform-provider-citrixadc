---
subcategory: "Network"
---

# Data Source: vxlan_srcip_binding

The vxlan_srcip_binding data source allows you to retrieve information about a source IP binding to a VXLAN.

## Example Usage

```terraform
data "citrixadc_vxlan_srcip_binding" "tf_binding" {
  vxlanid = 123
  srcip   = "11.22.33.44"
}

output "vxlanid" {
  value = data.citrixadc_vxlan_srcip_binding.tf_binding.vxlanid
}

output "srcip" {
  value = data.citrixadc_vxlan_srcip_binding.tf_binding.srcip
}
```

## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `srcip` - (Required) The source IP address to use in outgoing vxlan packets.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_srcip_binding. It is a system-generated identifier.
