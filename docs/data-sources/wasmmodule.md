---
subcategory: "WASM"
---

# Data Source: wasmmodule

The wasmmodule data source allows you to retrieve information about a WASM module.


## Example usage

```terraform
data "citrixadc_wasmmodule" "tf_wasmmodule" {
  name = "my_wasmmodule"
}

output "modulefile" {
  value = data.citrixadc_wasmmodule.tf_wasmmodule.modulefile
}

output "comment" {
  value = data.citrixadc_wasmmodule.tf_wasmmodule.comment
}
```


## Argument Reference

* `name` - (Required) The name of the WASM module file.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `modulefile` - File name of the WASM module.
* `signaturefile` - The SHA256 file contains the hash value used to validate the WASM module.
* `settingfile` - The WASM module filename contains module-specific configuration settings.
* `comment` - Any type of information about this WASM module.
* `id` - The id of the wasmmodule. It has the same value as the `name` attribute.
