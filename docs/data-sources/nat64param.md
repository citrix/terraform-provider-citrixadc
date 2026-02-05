---
subcategory: "Network"
---

# Data Source `nat64param`

The nat64param data source allows you to retrieve information about NAT64 parameter configuration.


## Example usage

```terraform
data "citrixadc_nat64param" "tf_nat64param" {
  td = 0
}

output "nat64ignoretos" {
  value = data.citrixadc_nat64param.tf_nat64param.nat64ignoretos
}

output "nat64zerochecksum" {
  value = data.citrixadc_nat64param.tf_nat64param.nat64zerochecksum
}

output "nat64v6mtu" {
  value = data.citrixadc_nat64param.tf_nat64param.nat64v6mtu
}

output "nat64fragheader" {
  value = data.citrixadc_nat64param.tf_nat64param.nat64fragheader
}
```


## Argument Reference

The following arguments are required:

* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

The following attributes are available:

* `nat64fragheader` - When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets. Possible values: [ ENABLED, DISABLED ]
* `nat64ignoretos` - Ignore TOS. Possible values: [ YES, NO ]
* `nat64v6mtu` - MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.
* `nat64zerochecksum` - Calculate checksum for UDP packets with zero checksum. Possible values: [ ENABLED, DISABLED ]

## Attribute Reference

* `id` - The id of the nat64param. It is a system-generated identifier.
