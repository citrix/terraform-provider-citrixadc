---
subcategory: "Network"
---

# Data Source: vxlan_nsip6_binding

The vxlan_nsip6_binding data source allows you to retrieve information about an IPv6 address binding to a VXLAN.


## Example usage

```terraform
data "citrixadc_vxlan_nsip6_binding" "tf_binding" {
  vxlanid   = 123
  ipaddress = "2001:db8:100::fb/64"
}

output "vxlanid" {
  value = data.citrixadc_vxlan_nsip6_binding.tf_binding.vxlanid
}

output "ipaddress" {
  value = data.citrixadc_vxlan_nsip6_binding.tf_binding.ipaddress
}
```


## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `ipaddress` - (Required) The IPv6 address assigned to the VXLAN.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_nsip6_binding. It is a system-generated identifier.
* `netmask` - Subnet mask for the network address defined for this VXLAN.
