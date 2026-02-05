---
subcategory: "VPN"
---

# Data Source: citrixadc_vpneula

The vpneula data source allows you to retrieve information about a VPN End User License Agreement (EULA) configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_vpneula" "tf_vpneula" {
  name = "my_eula"
}

output "eula_name" {
  value = data.citrixadc_vpneula.tf_vpneula.name
}
```

## Argument Reference

* `name` - (Required) Name for the eula

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpneula. It has the same value as the `name` attribute.
