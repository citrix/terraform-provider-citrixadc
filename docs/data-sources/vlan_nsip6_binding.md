---
subcategory: "Network"
---

# Data Source: vlan_nsip6_binding

The vlan_nsip6_binding data source allows you to retrieve information about an IPv6 address binding to a VLAN.

## Example Usage

```terraform
data "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
  vlanid    = 2
  ipaddress = "2001::a/96"
}

output "netmask" {
  value = data.citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding.netmask
}

output "ownergroup" {
  value = data.citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding.ownergroup
}
```

## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ipaddress` - (Required) The IPv6 address assigned to the VLAN.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `netmask` - Subnet mask for the network address defined for this VLAN.
* `ownergroup` - The owner node group in a Cluster for this VLAN.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `id` - The id of the vlan_nsip6_binding. It is a system-generated identifier.
