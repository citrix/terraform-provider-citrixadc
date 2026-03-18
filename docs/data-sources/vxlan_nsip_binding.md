---
subcategory: "Network"
---

# Data Source: vxlan_nsip_binding

The vxlan_nsip_binding data source allows you to retrieve information about the binding between a VXLAN and an NSIP (NetScaler IP).

## Example Usage

```terraform
data "citrixadc_vxlan_nsip_binding" "tf_binding" {
  vxlanid   = 123
  ipaddress = "10.222.74.146"
}

output "vxlanid" {
  value = data.citrixadc_vxlan_nsip_binding.tf_binding.vxlanid
}

output "ipaddress" {
  value = data.citrixadc_vxlan_nsip_binding.tf_binding.ipaddress
}

output "netmask" {
  value = data.citrixadc_vxlan_nsip_binding.tf_binding.netmask
}
```

## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `ipaddress` - (Required) The IP address assigned to the VXLAN.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_nsip_binding. It is a system-generated identifier.
* `netmask` - Subnet mask for the network address defined for this VXLAN.
