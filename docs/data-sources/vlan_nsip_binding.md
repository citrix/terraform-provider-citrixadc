---
subcategory: "Network"
---

# Data Source: vlan_nsip_binding

The vlan_nsip_binding data source allows you to retrieve information about an IP address binding to a VLAN.

## Example Usage

```terraform
data "citrixadc_vlan_nsip_binding" "tf_vlan_nsip_binding" {
  vlanid    = 40
  ipaddress = "10.222.74.145"
}

output "netmask" {
  value = data.citrixadc_vlan_nsip_binding.tf_vlan_nsip_binding.netmask
}

output "td" {
  value = data.citrixadc_vlan_nsip_binding.tf_vlan_nsip_binding.td
}
```

## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ipaddress` - (Required) The IP address assigned to the VLAN.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `netmask` - Subnet mask for the network address defined for this VLAN.
* `ownergroup` - The owner node group in a Cluster for this VLAN.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `id` - The id of the vlan_nsip_binding. It is a system-generated identifier.
