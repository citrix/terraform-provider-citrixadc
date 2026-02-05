---
subcategory: "Network"
---

# Data Source: citrixadc_inatparam

This data source is used to retrieve information about INAT (Inbound NAT) parameters configuration.

## Example Usage

```hcl
data "citrixadc_inatparam" "example" {
  td = 0
}
```

## Argument Reference

* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the INAT parameter configuration.
* `nat46fragheader` - When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets.
* `nat46ignoretos` - Ignore TOS.
* `nat46v6mtu` - MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.
* `nat46v6prefix` - The prefix used for translating packets received from private IPv6 servers into IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
* `nat46zerochecksum` - Calculate checksum for UDP packets with zero checksum.
