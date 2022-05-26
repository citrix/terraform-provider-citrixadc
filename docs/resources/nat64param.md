---
subcategory: "Network"
---

# Resource: nat64param

The nat64param resource is used to create NAT64 parameter resource.


## Example usage

```hcl
resource "citrixadc_nat64param" "tf_nat64param" {
  nat64ignoretos    = "NO"
  nat64zerochecksum = "ENABLED"
  nat64v6mtu        = 1280
  nat64fragheader   = "ENABLED"
}
```


## Argument Reference

* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `nat64ignoretos` - (Optional) Ignore TOS. Possible values: [ YES, NO ]
* `nat64zerochecksum` - (Optional) Calculate checksum for UDP packets with zero checksum. Possible values: [ ENABLED, DISABLED ]
* `nat64v6mtu` - (Optional) MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error. Minimum value =  1280 Maximum value =  9216
* `nat64fragheader` - (Optional) When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nat64param. It is a unique string prefixed with "tf-nat64param-"

