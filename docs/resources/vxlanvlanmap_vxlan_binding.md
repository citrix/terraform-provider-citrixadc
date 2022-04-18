---
subcategory: "Network"
---

# Resource: vxlanvlanmap_vxlan_binding

The vxlanvlanmap_vxlan_binding resource is used to bind vxlan to vxlanvlanmap resource.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 40
  aliasname = "Management VLAN"
}
resource "citrixadc_vlan" "tf_vlan1" {
  vlanid    = 41
  aliasname = "Management VLAN"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_vxlanvlanmap_vxlan_binding" "tf_binding" {
  name = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
  vxlan = citrixadc_vxlan.tf_vxlan.vxlanid
  vlan = [citrixadc_vlan.tf_vlan.vlanid,citrixadc_vlan.tf_vlan1.vlanid]
}
```


## Argument Reference

* `name` - (Required) Name of the mapping table.
* `vxlan` - (Required) The VXLAN assigned to the vlan inside the cloud.
* `vlan` - (Optional) The vlan id or the range of vlan ids in the on-premise network.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlanvlanmap_vxlan_binding. It is the concatenation of `name`  and `vxlan` attributes seperated by comma.


## Import

A vxlanvlanmap_vxlan_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vxlanvlanmap_vxlan_binding.tf_binding tf_vxlanvlanmp,123
```
