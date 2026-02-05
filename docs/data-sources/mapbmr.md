---
subcategory: "Network"
---

# Data Source `mapbmr`

The mapbmr data source allows you to retrieve information about MAP Basic Mapping Rule (BMR) configurations.


## Example usage

```terraform
data "citrixadc_mapbmr" "tf_mapbmr" {
  name = "bmr_rule_1"
}

output "name" {
  value = data.citrixadc_mapbmr.tf_mapbmr.name
}

output "ruleipv6prefix" {
  value = data.citrixadc_mapbmr.tf_mapbmr.ruleipv6prefix
}

output "psidoffset" {
  value = data.citrixadc_mapbmr.tf_mapbmr.psidoffset
}
```


## Argument Reference

* `name` - (Required) Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Basic Mapping Rule is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `eabitlength` - The Embedded Address (EA) bit field encodes the CE-specific IPv4 address and port information. The EA bit field, which is unique for a given Rule IPv6 prefix.
* `psidlength` - Length of Port Set Identifier Port Set Identifier(PSID) in Embedded Address (EA) bits.
* `psidoffset` - Start bit position of Port Set Identifier(PSID) value in Embedded Address (EA) bits.
* `ruleipv6prefix` - IPv6 prefix of Customer Edge(CE) device. MAP-T CE will send ipv6 packets with this ipv6 prefix as source ipv6 address prefix.

## Attribute Reference

* `id` - The id of the mapbmr. It has the same value as the `name` attribute.


## Import

A mapbmr can be imported using its name, e.g.

```shell
terraform import citrixadc_mapbmr.tf_mapbmr bmr_rule_1
```
