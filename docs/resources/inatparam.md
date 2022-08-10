---
subcategory: "Network"
---

# Resource: inatparam

The inatparam resource is used to create inatparam.


## Example usage

```hcl
resource "citrixadc_inatparam" "tf_inatparam" {
  nat46ignoretos    = "NO"
  nat46zerochecksum = "ENABLED"
  nat46v6mtu        = "1400"
}
```


## Argument Reference

* `nat46v6prefix` - (Optional) The prefix used for translating packets received from private IPv6 servers into IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `nat46ignoretos` - (Optional) Ignore TOS. Possible values: [ YES, NO ]
* `nat46zerochecksum` - (Optional) Calculate checksum for UDP packets with zero checksum. Possible values: [ ENABLED, DISABLED ]
* `nat46v6mtu` - (Optional) MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error. Minimum value =  1280 Maximum value =  9216
* `nat46fragheader` - (Optional) When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the inatparam. It is a unique string prefixed with "tf-inatparam-".
