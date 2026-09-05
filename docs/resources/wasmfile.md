---
subcategory: "WASM"
---

# Resource: wasmfile

This resource is used to import a WASM module related file (module, signature, or setting) onto the Citrix ADC.

~> **Note:** A WASM file cannot be updated in place. Changing any attribute forces replacement (destroy + recreate).


## Example usage

```hcl
resource "citrixadc_wasmfile" "tf_wasmfile" {
  name     = "tf_wasmfile"
  src      = "http://wasm.example.com/modules/tf_wasmfile.wasm"
  filetype = "Module"
  comment  = "Imported WASM module for tf demo"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the WASM module/signature page object on the Citrix ADC. Minimum length = 1, Maximum length = 31. Forces replacement on change.
* `src` - (Required) Local path or URL (protocol, host, path, and file name) for the file from which to retrieve the imported file. The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Minimum length = 1, Maximum length = 2047. Forces replacement on change.
* `filetype` - (Optional) WASM file type to be imported. Default value: `Module`. Possible values: `Module`, `Signature`, `Setting`. Forces replacement on change.
* `comment` - (Optional) Any comments to preserve information about the WASM page object. Maximum length = 128. Forces replacement on change.
* `overwrite` - (Optional) Overwrites the existing file. This is a write-only option and is not read back from the appliance. Forces replacement on change.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `wasmfile`. It has the same value as the `name` attribute.


## Import

A wasmfile can be imported using its name, e.g.

```shell
terraform import citrixadc_wasmfile.tf_wasmfile <name>
```
