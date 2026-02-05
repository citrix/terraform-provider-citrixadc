---
subcategory: "SNMP"
---

# Data Source: citrixadc_snmpmib

The `citrixadc_snmpmib` data source is used to retrieve SNMP MIB configuration information from the Citrix ADC.

## Example usage

```hcl
data "citrixadc_snmpmib" "example" {
  ownernode = -1
}
```

## Argument Reference

* `ownernode` - (Required) ID of the cluster node for which we are retrieving the MIB. This is a mandatory argument to get SNMP MIB on CLIP. Use -1 for non-cluster setups.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the SNMP MIB resource.
* `contact` - Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address.
* `customid` - Custom identification number for the Citrix ADC.
* `location` - Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number.
* `name` - Name for this Citrix ADC.
