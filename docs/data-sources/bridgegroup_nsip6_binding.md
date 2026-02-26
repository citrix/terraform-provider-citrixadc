---
subcategory: "Network"
---

# Data Source: bridgegroup_nsip6_binding

The bridgegroup_nsip6_binding data source allows you to retrieve information about the binding between a bridge group and an IPv6 address.

## Example Usage

```terraform
data "citrixadc_bridgegroup_nsip6_binding" "tf_binding" {
  bridgegroup_id = 2
  ipaddress      = "2001:db8:100::fb/64"
}

output "netmask" {
  value = data.citrixadc_bridgegroup_nsip6_binding.tf_binding.netmask
}

output "ownergroup" {
  value = data.citrixadc_bridgegroup_nsip6_binding.tf_binding.ownergroup
}
```

## Argument Reference

* `bridgegroup_id` - (Required) The integer that uniquely identifies the bridge group.
* `ipaddress` - (Optional) The IP address assigned to the bridge group. Used as an additional filter.
* `netmask` - (Optional) A subnet mask associated with the network address. Used as an additional filter.
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan. Used as an additional filter.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Used as an additional filter.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup_nsip6_binding. It is a system-generated identifier.
* `ipaddress` - The IP address assigned to the bridge group.
* `netmask` - A subnet mask associated with the network address.
* `ownergroup` - The owner node group in a Cluster for this vlan.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
