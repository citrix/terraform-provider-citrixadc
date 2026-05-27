---
subcategory: "API Definition"
---

# Data Source: apispecfile

The apispecfile data source allows you to retrieve information about an API specification file that has been imported onto the Citrix ADC.


## Example usage

```terraform
data "citrixadc_apispecfile" "tf_apispecfile" {
  name = "my_apispecfile"
}

output "apispecfile_src" {
  value = data.citrixadc_apispecfile.tf_apispecfile.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported API spec file to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `src` - URL from which the spec file was originally imported (protocol, host, and path including the file name).
* `overwrite` - Whether an existing spec file of the same name was overwritten during import. This is a create-time input to the NITRO `Import` action and is not echoed back by the appliance; the value reflects the prior plan/state rather than a live attribute on the ADC.
* `id` - The id of the apispecfile. It has the same value as the `name` attribute.
