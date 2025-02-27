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

## Argument Reference

* `keyfile` - (Required) Name for and, optionally, path to the RSA key. /nsconfig/ssl/ is the default path. Maximum length =  63
* `curve` - (Required) Curve id to generate ECDSA key. Only P_256 and P_384 are supported. Possible values: [ P_256, P_384 ]
* `keyform` - (Optional) Format in which the key is stored on the appliance. Possible values: [ DER, PEM ] Default value: PEM
* `aes256` - (Optional) Encrypt the generated RSA key by using the AES algorithm.
* `des` - (Optional) Encrypt the generated RSA key by using the DES algorithm.
* `des3` - (Optional) Encrypt the generated RSA key by using the Triple-DES algorithm.
* `password` - (Optional) Pass phrase to use for encryption if AES256, DES or DES3 option is selected. Maximum value: 31
* `pkcs8` - (Optional) Create the private key in PKCS#8 format.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslecdsakey. It is a unique string prefixed with "tf-sslecdsakey-"

