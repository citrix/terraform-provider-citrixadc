---
subcategory: "Network"
---

# Data Source: vxlan

The vxlan data source allows you to retrieve information about a VXLAN configuration.

## Example usage

```terraform
data "citrixadc_vxlan" "example" {
  vxlanid = 123
}

output "vlan" {
  value = data.citrixadc_vxlan.example.vlan
}

output "port" {
  value = data.citrixadc_vxlan.example.port
}
```

## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `dynamicrouting` - Enable dynamic routing on this VXLAN.
* `innervlantagging` - Specifies whether Citrix ADC should generate VXLAN packets with inner VLAN tag.
* `ipv6dynamicrouting` - Enable all IPv6 dynamic routing protocols on this VXLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
* `port` - Specifies UDP destination port for VXLAN packets.
* `protocol` - VXLAN-GPE next protocol. RESERVED, IPv4, IPv6, ETHERNET, NSH.
* `type` - VXLAN encapsulation type. VXLAN, VXLANGPE.
* `vlan` - ID of VLANs whose traffic is allowed over this VXLAN. If you do not specify any VLAN IDs, the Citrix ADC allows traffic of all VLANs that are not part of any other VXLANs.
