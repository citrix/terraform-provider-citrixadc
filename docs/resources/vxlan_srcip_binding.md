---
subcategory: "Network"
---

# Resource: vxlan_srcip_binding

The vxlan_srcip_binding resource is used to bind srcip to vxlan resource.


## Example usage

```hcl
resource "citrixadc_nsip" "tf_srcip" {
  ipaddress = "11.22.33.44"
  type      = "SNIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_vxlan_srcip_binding" "tf_binding" {
  vxlanid = citrixadc_vxlan.tf_vxlan.vxlanid
  srcip   = citrixadc_nsip.tf_srcip.ipaddress
}
```


## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `srcip` - (Required) The source IP address to use in outgoing vxlan packets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_srcip_binding. It is the concatenation of `vxlanid` and `srcip` attributes seperated by comma.


## Import

A vxlan_srcip_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vxlan_srcip_binding.tf_binding 123,11.22.33.44
```
