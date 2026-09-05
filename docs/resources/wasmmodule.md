---
subcategory: "WASM"
---

# Resource: wasmmodule

The wasmmodule resource is used to create wasmmodule resource.

~> **NOTE:** The referenced `modulefile` (and its SHA256 `signaturefile`) must
already exist on the appliance filesystem before the wasmmodule can be created.


## Example usage

```hcl
resource "citrixadc_wasmmodule" "tf_wasmmodule" {
  name          = "my_wasmmodule"
  modulefile    = "my_module.wasm"
  signaturefile = "my_module.sha256"
  settingfile   = "my_module_settings.json"
  comment       = "example WASM module"
}
```


## Argument Reference

* `name` - (Required) The name of the WASM module file. This attribute is set only while adding a wasmmodule resource, but cannot be updated.
* `modulefile` - (Optional) File name of the WASM module. This attribute is set only while adding a wasmmodule resource, but cannot be updated.
* `signaturefile` - (Optional) The SHA256 file contains the hash value used to validate the WASM module. This attribute is set only while adding a wasmmodule resource, but cannot be updated.
* `settingfile` - (Optional) The WASM module filename contains module-specific configuration settings.
* `comment` - (Optional) Any type of information about this WASM module.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the wasmmodule. It has the same value as the `name` attribute.


## Import

A wasmmodule can be imported using its name, e.g.

```shell
terraform import citrixadc_wasmmodule.tf_wasmmodule my_wasmmodule
```
