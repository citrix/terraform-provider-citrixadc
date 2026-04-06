---
subcategory: "Network"
---

# Data Source: nd6

The nd6 data source allows you to retrieve information about IPv6 Neighbor Discovery (ND6) entries.

## Example usage

```terraform
data "citrixadc_nd6" "tf_nd6_ds_data" {
  neighbor = "2001::5"
}

output "mac" {
  value = data.citrixadc_nd6.tf_nd6_ds_data.mac
}

output "ifnum" {
  value = data.citrixadc_nd6.tf_nd6_ds_data.ifnum
}
```

## Argument Reference

* `neighbor` - (Required) Link-local IPv6 address of the adjacent network device to add to the ND6 table.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nd6. It has the same value as the `neighbor` attribute.
* `ifnum` - Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.
* `mac` - MAC address of the adjacent network device.
* `nodeid` - Unique number that identifies the cluster node.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `vlan` - Integer value that uniquely identifies the VLAN on which the adjacent network device exists.
* `vtep` - IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable.
* `vxlan` - ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable.

## Import

A nd6 can be imported using its neighbor, e.g.

```shell
terraform import citrixadc_nd6.tf_nd6 2001::5
```
