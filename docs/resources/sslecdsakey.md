---
subcategory: "SSL"
---

# Resource: sslecdsakey

The sslecdsakey resource is used to create ssl ecdsakey file.


## Example usage

```hcl
resource "citrixadc_sslecdsakey" "tf_sslecdsakey" {
  keyfile  = "/nsconfig/ssl/demoecdsa.pem"
  curve    = "P_256"
  aes256   = true
  password = "SecretPassword"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "sslecdsakey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslecdsakey" "tf_sslecdsakey" {
  keyfile  = "/nsconfig/ssl/demoecdsa.pem"
  curve    = "P_256"
  aes256   = true
  password = var.sslecdsakey_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the key encryption pass phrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "sslecdsakey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslecdsakey" "tf_sslecdsakey" {
  keyfile             = "/nsconfig/ssl/demoecdsa.pem"
  curve               = "P_256"
  aes256              = true
  password_wo         = var.sslecdsakey_password
  password_wo_version = 1
}
```

## Argument Reference

* `keyfile` - (Required) Name for and, optionally, path to the RSA key. /nsconfig/ssl/ is the default path. Maximum length =  63
* `curve` - (Required) Curve id to generate ECDSA key. Only P_256 and P_384 are supported. Possible values: [ P_256, P_384 ]
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ] Default value: PEM
* `aes256` - (Optional) Encrypt the generated RSA key by using the AES algorithm.
* `des` - (Optional) Encrypt the generated RSA key by using the DES algorithm.
* `des3` - (Optional) Encrypt the generated RSA key by using the Triple-DES algorithm.
* `password` - (Optional, Sensitive) Pass phrase to use for encryption if DES or DES3 option is selected. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) A user-managed integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger the write-only secret to be re-sent. For this resource, changing the value forces the resource to be replaced.
* `pkcs8` - (Optional) Create the private key in PKCS#8 format.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslecdsakey. It is a unique string prefixed with "tf-sslecdsakey-"

