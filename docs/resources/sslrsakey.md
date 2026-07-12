---
subcategory: "SSL"
---

# Resource: sslrsakey

The sslrsakey resource is used to create ssl rsakey file.


## Example usage

```hcl
resource "citrixadc_sslrsakey" "tf_sslrsakey" {
  keyfile          = "/nsconfig/ssl/key1.pem"
  bits             = 2048
  aes256           = true
  password         = "SecretPassword"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "sslrsakey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslrsakey" "tf_sslrsakey" {
  keyfile  = "/nsconfig/ssl/key1.pem"
  bits     = 2048
  aes256   = true
  password = var.sslrsakey_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the key encryption pass phrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "sslrsakey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslrsakey" "tf_sslrsakey" {
  keyfile             = "/nsconfig/ssl/key1.pem"
  bits                = 2048
  aes256              = true
  password_wo         = var.sslrsakey_password
  password_wo_version = 1
}
```


## Argument Reference

* `keyfile` - (Required) Name for and, optionally, path to the RSA key. /nsconfig/ssl/ is the default path. Maximum length =  63
* `bits` - (Required) Size, in bits, of the RSA key. Minimum value: 512 Maximum value: 4096
* `exponent` - (Optional) Public exponent for the RSA key. The exponent is part of the cipher algorithm and is required for creating the RSA key.  Possible values: 3, F4 Default value: F4
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ] Default value: PEM
* `aes256` - (Optional) Encrypt the generated RSA key by using the AES algorithm.
* `des` - (Optional) Encrypt the generated RSA key by using the DES algorithm.
* `des3` - (Optional) Encrypt the generated RSA key by using the Triple-DES algorithm.
* `password` - (Optional, Sensitive) Pass phrase to use for encryption if AES256, DES or DES3 option is selected. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) A user-managed integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed; this causes the write-only secret to be re-sent, which forces the resource to be replaced. This attribute is user-controlled and has no default value.
* `pkcs8` - (Optional) Create the private key in PKCS#8 format.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslrsakey. It is a unique string prefixed with "tf-sslrsakey-"

