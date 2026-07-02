---
subcategory: "SSL"
---

# Resource: sslpkcs12

The sslpkcs12 resource converts a certificate and private key to or from PKCS#12 format on the Citrix ADC. Use the import direction to extract a certificate-key pair from a PKCS#12 bundle into PEM format, or the export direction to package an existing PEM certificate and key into a PKCS#12 bundle.

This resource performs a one-shot convert action (`?action=convert`). The operation reads source files from, and writes the output file to, the appliance filesystem. It is **non-idempotent** and there is no NITRO GET endpoint to read the result back, so the converted output is not refreshed into Terraform state. For the same reason, importing this resource into Terraform is not meaningful, and destroying it only removes the Terraform state entry while the converted file remains on the appliance.


## Example usage

The secret attributes can be supplied either as sensitive attributes (persisted in state) or as write-only attributes (not persisted in state). Both `password`/`password_wo` and `pempassphrase`/`pempassphrase_wo` require that at least one variant be set.

### Using the sensitive attributes (persisted in state)

```hcl
variable "sslpkcs12_password" {
  type      = string
  sensitive = true
}

variable "sslpkcs12_pempassphrase" {
  type      = string
  sensitive = true
}

# Export an existing PEM certificate and key into a PKCS#12 bundle
resource "citrixadc_sslpkcs12" "tf_pkcs12" {
  outfile       = "exported.pfx"
  pkcs12file    = "exported.pfx"
  certfile      = "servercert1.pem"
  keyfile       = "servercert1.key"
  export        = true
  aes256        = true
  password      = var.sslpkcs12_password
  pempassphrase = var.sslpkcs12_pempassphrase
}
```

### Using the write-only attributes (not persisted in state)

The `password_wo` and `pempassphrase_wo` attributes provide an ephemeral path for the PKCS#12 and PEM pass phrases. The values are sent to the Citrix ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger a re-run when a value changes, increment the matching `_wo_version`.

```hcl
variable "sslpkcs12_password" {
  type      = string
  sensitive = true
}

variable "sslpkcs12_pempassphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslpkcs12" "tf_pkcs12" {
  outfile                  = "exported.pfx"
  pkcs12file               = "exported.pfx"
  certfile                 = "servercert1.pem"
  keyfile                  = "servercert1.key"
  export                   = true
  aes256                   = true
  password_wo              = var.sslpkcs12_password
  password_wo_version      = 1
  pempassphrase_wo         = var.sslpkcs12_pempassphrase
  pempassphrase_wo_version = 1
}
```


## Argument Reference

* `outfile` - (Required) Name for and, optionally, path to, the output file that contains the certificate and the private key after conversion. `/nsconfig/ssl/` is the default path. If importing, the certificate-key pair is stored in PEM format. If exporting, the certificate-key pair is stored in PKCS#12 format. Changing this attribute forces a new resource to be created.
* `password` - (Optional, Sensitive) Pass phrase for the PKCS#12 bundle. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Either `password` or `password_wo` must be specified. Changing this attribute forces a new resource to be created.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence. Either `password` or `password_wo` must be specified.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a re-run. Defaults to `1`.
* `pempassphrase` - (Optional, Sensitive) Pass phrase used to encrypt the PEM private key. The value is persisted in Terraform state (encrypted). See also `pempassphrase_wo` for an ephemeral alternative. Either `pempassphrase` or `pempassphrase_wo` must be specified. Changing this attribute forces a new resource to be created.
* `pempassphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `pempassphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `pempassphrase_wo_version`. If both `pempassphrase` and `pempassphrase_wo` are set, `pempassphrase_wo` takes precedence. Either `pempassphrase` or `pempassphrase_wo` must be specified.
* `pempassphrase_wo_version` - (Optional) An integer version tracker for `pempassphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a re-run. Defaults to `1`.
* `import` - (Optional) Convert the certificate and private key from PKCS#12 format to PEM format. Changing this attribute forces a new resource to be created.
* `export` - (Optional) Convert the certificate and private key from PEM format to PKCS#12 format. Changing this attribute forces a new resource to be created.
* `pkcs12file` - (Optional) Name for and, optionally, path to, the PKCS#12 file. If importing, specify the input file name that contains the certificate and the private key in PKCS#12 format. If exporting, specify the output file name. `/nsconfig/ssl/` is the default path. Changing this attribute forces a new resource to be created.
* `certfile` - (Optional) Certificate file to be converted from PEM to PKCS#12 format. Changing this attribute forces a new resource to be created.
* `keyfile` - (Optional) Name of the private key file to be converted from PEM to PKCS#12 format. If the key file is encrypted, the pass phrase used for encrypting the key is required. Changing this attribute forces a new resource to be created.
* `des` - (Optional) Encrypt the private key by using the DES algorithm in CBC mode during the import operation. Changing this attribute forces a new resource to be created.
* `des3` - (Optional) Encrypt the private key by using the Triple-DES algorithm in EDE CBC mode (168-bit key) during the import operation. Changing this attribute forces a new resource to be created.
* `aes256` - (Optional) Encrypt the private key by using the AES algorithm (256-bit key) during the import operation. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpkcs12. It is a synthetic identifier set to the value of the `outfile` attribute, because this action-only resource has no NITRO GET endpoint.
