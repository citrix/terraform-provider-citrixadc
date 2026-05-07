---
subcategory: "NS"
---

# Resource: nshmackey

The nshmackey resource is used to create HMAC key resource.


## Example usage

### Using keyvalue (sensitive attribute - persisted in state)

```hcl
variable "nshmackey_keyvalue" {
  type      = string
  sensitive = true
}

resource "citrixadc_nshmackey" "tf_nshmackey" {
  name     = "tf_nshmackey"
  digest   = "MD4"
  keyvalue = var.nshmackey_keyvalue
  comment  = "Testing"
}
```

### Using keyvalue_wo (write-only/ephemeral - NOT persisted in state)

The `keyvalue_wo` attribute provides an ephemeral path for the HMAC key value. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `keyvalue_wo_version`.

```hcl
variable "nshmackey_keyvalue" {
  type      = string
  sensitive = true
}

resource "citrixadc_nshmackey" "tf_nshmackey" {
  name                = "tf_nshmackey"
  digest              = "MD4"
  keyvalue_wo         = var.nshmackey_keyvalue
  keyvalue_wo_version = 1
  comment             = "Testing"
}
```

To rotate the key, update the variable value and bump the version:

```hcl
resource "citrixadc_nshmackey" "tf_nshmackey" {
  name                = "tf_nshmackey"
  digest              = "MD4"
  keyvalue_wo         = var.nshmackey_keyvalue
  keyvalue_wo_version = 2  # Bumped to trigger update
  comment             = "Testing"
}
```


## Argument Reference

* `name` - (Required) Key name.  This follows the same syntax rules as other expression entity names: It must begin with an alpha character (A-Z or a-z) or an underscore (_). The rest of the characters must be alpha, numeric (0-9) or underscores. It cannot be re or xp (reserved for regular and XPath expressions). It cannot be an expression reserved word (e.g. SYS or HTTP). It cannot be used for an existing expression object (HTTP callout, patset, dataset, stringmap, or named expression). Minimum length =  1
* `digest` - (Optional) Digest (hash) function to be used in the HMAC computation. Possible values: [ MD2, MD4, MD5, SHA1, SHA224, SHA256, SHA384, SHA512 ]
* `keyvalue` - (Optional, Sensitive) The hex-encoded key to be used in the HMAC computation. The key can be any length (up to a Citrix ADC-imposed maximum of 255 bytes). If the length is less than the digest block size, it will be zero padded up to the block size. If it is greater than the block size, it will be hashed using the digest function to the block size. The block size for each digest is: MD2    - 16 bytes MD4    - 16 bytes MD5    - 16 bytes SHA1   - 20 bytes SHA224 - 28 bytes SHA256 - 32 bytes SHA384 - 48 bytes SHA512 - 64 bytes Note that the key will be encrypted when it it is saved There is a special key value AUTO which generates a new random key for the specified digest. This kind of key is intended for use cases where the NetScaler both generates and verifies an HMAC on  the same data. The value is persisted in Terraform state (encrypted). See also `keyvalue_wo` for an ephemeral alternative.
* `keyvalue_wo` - (Optional, Sensitive, WriteOnly) Same as `keyvalue`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `keyvalue_wo_version`. If both `keyvalue` and `keyvalue_wo` are set, `keyvalue_wo` takes precedence.
* `keyvalue_wo_version` - (Optional) An integer version tracker for `keyvalue_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `comment` - (Optional) Comments associated with this encryption key.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nshmackey. It has the same value as the `name` attribute.


## Import

A nshmackey can be imported using its name, e.g.

```shell
terraform import citrixadc_nshmackey.tf_nshmackey tf_nshmackey
```
