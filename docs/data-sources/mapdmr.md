---
subcategory: "Network"
---

# Data Source `mapdmr`

The mapdmr data source allows you to retrieve information about MAP Default Mapping Rule (DMR) configurations.


## Example usage

```terraform
data "citrixadc_mapdmr" "tf_mapdmr" {
  name = "dmr_rule_1"
}

output "name" {
  value = data.citrixadc_mapdmr.tf_mapdmr.name
}

output "bripv6prefix" {
  value = data.citrixadc_mapdmr.tf_mapdmr.bripv6prefix
}
```


## Argument Reference

* `name` - (Required) Name for the Default Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Default Mapping Rule is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `bripv6prefix` - IPv6 prefix of Border Relay (Citrix ADC) device. MAP-T CE will send ipv6 packets to this ipv6 prefix. The DMR IPv6 prefix length SHOULD be 64 bits long by default and in any case MUST NOT exceed 96 bits.

## Attribute Reference

* `id` - The id of the mapdmr. It has the same value as the `name` attribute.


## Import

A mapdmr can be imported using its `name`, e.g.

```shell
terraform import citrixadc_mapdmr.tf_mapdmr dmr_rule_1
```
