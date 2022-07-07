---
subcategory: "Network"
---

# Resource: bridgetable

The bridgetable resource is used to create bridge table entry resource.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  vlan               = citrixadc_vlan.tf_vlan.vlanid
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_bridgetable" "tf_bridgetable" {
  mac       = "00:00:00:00:00:01"
  vxlan     = citrixadc_vxlan.tf_vxlan.vxlanid
  vtep      = "2.34.5.6"
  bridgeage = "300"
}
```


## Argument Reference

* `mac` - (Required) The MAC address of the target.
* `vtep` - (Required) The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.
* `vxlan` - (Required) The VXLAN to which this address is associated.
* `bridgeage` - (Optional) Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.
* `devicevlan` - (Optional) The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.
* `ifnum` - (Optional) INTERFACE  whose entries are to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `vlan` - (Optional) VLAN  whose entries are to be removed.
* `vni` - (Optional) The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint.  If omitted the value specified as vxlan will be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgetable It is the concatenation of `mac` , `vxlan` and `vtep` attributes separated by comma. 


## Import

A bridgetable can be imported using its is, e.g.

```shell
terraform import citrixadc_bridgetable.tf_bridgetable 00:00:00:00:00:01,123,2.34.5.6
```
