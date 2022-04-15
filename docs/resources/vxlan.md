---
subcategory: "Network"
---

# Resource: vxlan

The vxlan resource is used to create vxlan resource.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 40
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
```


## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `dynamicrouting` - (Optional) Enable dynamic routing on this VXLAN.
* `innervlantagging` - (Optional) Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.
* `ipv6dynamicrouting` - (Optional) Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
* `port` - (Optional) Specifies UDP destination port for VXLAN packets.
* `protocol` - (Optional) VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH
* `type` - (Optional) VXLAN encapsulation type. VXLAN, VXLANGPE
* `vlan` - (Optional) ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan. It has the same value as the `vxlanid` attribute.


## Import

A vxlan can be imported using its vxlanid, e.g.

```shell
terraform import citrixadc_vxlan.tf_vxlan 123
```
