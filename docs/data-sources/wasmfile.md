---
subcategory: "WASM"
---

# Data Source: wasmfile

The wasmfile data source allows you to retrieve information about a WASM module related file configured on the Citrix ADC.


## Example usage

```hcl
# Look up an existing WASM file by name
data "citrixadc_wasmfile" "example_wasmfile" {
  name = citrixadc_wasmfile.tf_wasmfile.name
}

# Use the data source outputs
output "wasmfile_id" {
  value = data.citrixadc_wasmfile.example_wasmfile.id
}

output "wasmfile_filetype" {
  value = data.citrixadc_wasmfile.example_wasmfile.filetype
}
```


## Argument Reference

The following argument is required:

* `name` - (Required) Name of the WASM file to retrieve.


## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The id of the `wasmfile`. It has the same value as the `name` attribute.
* `src` - Local path or URL from which the WASM object was imported.
* `filetype` - WASM file type. Possible values: `Module`, `Signature`, `Setting`.
* `comment` - Any comments to preserve information about the WASM page object.
* `overwrite` - Write-only option; not read back from the appliance.
