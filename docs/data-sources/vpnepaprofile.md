---
subcategory: "SSL VPN"
---

# Data Source: vpnepaprofile

The vpnepaprofile data source allows you to retrieve information about an existing EPA device profile configured on the Citrix ADC, looked up by its name.


## Example usage

```terraform
data "citrixadc_vpnepaprofile" "example" {
  name = "tf_vpnepaprofile"
}

output "vpnepaprofile_filename" {
  value = data.citrixadc_vpnepaprofile.example.filename
}
```


## Argument Reference

* `name` - (Required) Name of the device profile to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnepaprofile. It has the same value as the `name` attribute.
* `filename` - Filename of the device-profile data XML.
* `data` - Device-profile data XML.
