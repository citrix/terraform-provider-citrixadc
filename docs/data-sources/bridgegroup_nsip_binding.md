---
subcategory: "Network"
---

# Data Source: bridgegroup_nsip_binding

The bridgegroup_nsip_binding data source allows you to retrieve information about the binding between a bridge group and an IPv4 address.

## Example Usage

```terraform
data "citrixadc_bridgegroup_nsip_binding" "tf_binding" {
  bridgegroup_id = 2
  ipaddress      = "2.2.2.3"
}

output "netmask" {
  value = data.citrixadc_bridgegroup_nsip_binding.tf_binding.netmask
}

output "ownergroup" {
  value = data.citrixadc_bridgegroup_nsip_binding.tf_binding.ownergroup
}
```

## Argument Reference

* `bridgegroup_id` - (Required) The integer that uniquely identifies the bridge group.
* `ipaddress` - (Required) The IP address assigned to the bridge group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup_nsip_binding. It is a system-generated identifier.
* `netmask` - The network mask for the subnet defined for the bridge group.
* `ownergroup` - The owner node group in a Cluster for this vlan.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
