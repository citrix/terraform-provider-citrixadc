---
subcategory: "Network"
---

# Resource: mapbmr

The mapbmr resource is used to create MAP-T Basic Mapping rule resource.


## Example usage

```hcl
resource "citrixadc_mapbmr" "tf_mapbmr" {
  name           = "tf_mapbmr"
  ruleipv6prefix = "2001:db8:abcd:12::/64"
  psidoffset     = 6
  eabitlength    = 16
  psidlength     = 8
}
```


## Argument Reference

* `name` - (Required) Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8" ). The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device. Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE. Minimum length =  1 Maximum length =  127
* `ruleipv6prefix` - (Required) IPv6 prefix of Customer Edge(CE) device.MAP-T CE will send ipv6 packets with this ipv6 prefix as source ipv6 address prefix.
* `psidoffset` - (Optional) Start bit position  of Port Set Identifier(PSID) value in Embedded Address (EA) bits. Minimum value =  1 Maximum value =  15
* `eabitlength` - (Optional) The Embedded Address (EA) bit field encodes the CE-specific IPv4 address and port information.  The EA bit field, which is unique for a given Rule IPv6 prefix. Minimum value =  2 Maximum value =  47
* `psidlength` - (Optional) Length of Port Set IdentifierPort Set Identifier(PSID) in Embedded Address (EA) bits. Minimum value =  1 Maximum value =  16


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapbmr. It has the same value as the `name` attribute.


## Import

A mapbmr can be imported using its name, e.g.

```shell
terraform import citrixadc_mapbmr.tf_mapbmr tf_mapbmr
```
