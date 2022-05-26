---
subcategory: "Network"
---

# Resource: bridgegroup_vlan_binding

The bridgegroup_vlan_binding resource is used to bind vlan to bridgegroup resource.


## Example usage

```hcl
resource "citrixadc_bridgegroup" "tf_bridgegroup" {
  bridgegroup_id     = 2
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_bridgegroup_vlan_binding" "tf_binding" {
  bridgegroup_id = citrixadc_bridgegroup.tf_bridgegroup.bridgegroup_id
  vlan           = citrixadc_vlan.tf_vlan.vlanid
}
```


## Argument Reference

* `bridgegroup_id` - (Required) The integer that uniquely identifies the bridge group. Minimum value =  1 Maximum value =  1000
* `vlan` - (Required) Names of all member VLANs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the bridgegroup_vlan_binding.It is the concatenation of `bridgegroup_id` and `vlan` attributes separated by comma.


## Import

A bridgegroup_vlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_bridgegroup_vlan_binding.tf_csaction 2,20
```
