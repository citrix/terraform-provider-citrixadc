---
subcategory: "NS"
---

# Data Source: nsextension

The nsextension data source allows you to retrieve information about a NetScaler extension imported on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nsextension" "tf_nsextension" {
  name = "myextension"
}

output "comment" {
  value = data.citrixadc_nsextension.tf_nsextension.comment
}
```


## Argument Reference

* `name` - (Required) Name of the extension object to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about the extension object.
* `overwrite` - Indicates whether the existing file is overwritten on import.
* `src` - Local path to and name of, or URL for, the file from which the extension was imported. The appliance may not return this value.
* `trace` - Tracing level for extension execution. Possible values: [ off, calls, lines, all ]
* `tracefunctions` - Comma-separated list of extension functions being traced.
* `tracevariables` - Comma-separated list of variables being traced.
* `id` - The id of the nsextension. It has the same value as the `name` attribute.
