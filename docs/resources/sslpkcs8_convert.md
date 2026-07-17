---
subcategory: "SSL"
---

# Resource: sslpkcs8_convert

The sslpkcs8_convert resource converts a private key file to PKCS#8 format on the Citrix ADC, reading the input key in PEM or DER format and writing the converted PKCS#8 key to an output file on the appliance filesystem.

This resource performs a one-shot convert action (`?action=convert`). The operation reads the source key from, and writes the output file to, the appliance filesystem. It is **non-idempotent** and there is no NITRO GET endpoint to read the result back, so the converted output is not refreshed into Terraform state, and destroying the resource only removes the Terraform state entry while the converted file remains on the appliance.


## Example usage

### Using the sensitive attribute (persisted in state)

```hcl
variable "sslpkcs8_convert_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslpkcs8_convert" "tf_pkcs8" {
  keyfile   = "servercert1.key"
  pkcs8file = "servercert1_pkcs8.key"
  keyform   = "PEM"
  password  = var.sslpkcs8_convert_password
}
```

### Using the write-only attribute (not persisted in state)

The `password_wo` attribute provides an ephemeral path for the key pass phrase. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger a re-run when the value changes, increment `password_wo_version`.

```hcl
variable "sslpkcs8_convert_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslpkcs8_convert" "tf_pkcs8" {
  keyfile             = "servercert1.key"
  pkcs8file           = "servercert1_pkcs8.key"
  keyform             = "PEM"
  password_wo         = var.sslpkcs8_convert_password
  password_wo_version = 1
}
```


## Argument Reference

* `keyfile` - (Required) Name of and, optionally, path to the input key file to be converted from PEM or DER format to PKCS#8 format. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `pkcs8file` - (Required) Name for and, optionally, path to, the output file where the PKCS#8 format key file is stored. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `keyform` - (Optional) Format in which the key file is stored on the appliance. Possible values: [ DER, PEM ]. Defaults to `"PEM"`. Changing this attribute forces a new resource to be created.
* `password` - (Optional, Sensitive) Password to assign to the file if the key is encrypted. Applies only for PEM format files. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Changing this attribute forces a new resource to be created.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a re-run. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpkcs8_convert. It is a static string `"sslpkcs8_convert"`, because this action-only resource has no NITRO GET endpoint.
