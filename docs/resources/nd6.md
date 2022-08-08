---
subcategory: "Network"
---

# Resource: nd6

The nd6 resource is used to create nd6.


## Example usage

```hcl
resource "citrixadc_nd6" "tf_nd6" {
  neighbor = "2001::3"
  mac      = "e6:ec:41:50:b1:d1"
  ifnum    = "LO/1"
}
```


## Argument Reference

* `neighbor` - (Required) Link-local IPv6 address of the adjacent network device to add to the ND6 table.
* `mac` - (Required) MAC address of the adjacent network device.
* `ifnum` - (Optional) Interface through which the adjacent network device is available, specified in slot/port notation (for example, 1/3). Use spaces to separate multiple entries.
* `vlan` - (Optional) Integer value that uniquely identifies the VLAN on which the adjacent network device exists. Minimum value =  1 Maximum value =  4094
* `vxlan` - (Optional) ID of the VXLAN on which the IPv6 address of this ND6 entry is reachable. Minimum value =  1 Maximum value =  16777215
* `vtep` - (Optional) IP address of the VXLAN tunnel endpoint (VTEP) through which the IPv6 address of this ND6 entry is reachable. Minimum length =  1
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `nodeid` - (Optional) Unique number that identifies the cluster node. Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nd6. It has the same value as the `neighbor` attribute.


## Import

A nd6 can be imported using its name, e.g.

```shell
terraform import citrixadc_nd6.tf_nd6 2001::3
```
