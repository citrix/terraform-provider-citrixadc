---
subcategory: "NS"
---

# Resource: nsencryptionparams

The `nsencryptionparams` resource is used to configure the Citrix ADC encryption parameters, including the cipher method and encryption key value.


## Example usage

### Basic usage

```hcl
resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
  method = "AES256"
}
```

### Using keyvalue (sensitive attribute - persisted in state)

```hcl
variable "nsencryptionparams_keyvalue" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
  method   = "AES256"
  keyvalue = var.nsencryptionparams_keyvalue
}
```

### Using keyvalue_wo (write-only/ephemeral - NOT persisted in state)

The `keyvalue_wo` attribute provides an ephemeral path for the encryption key value. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the key value changes, increment `keyvalue_wo_version`.

```hcl
variable "nsencryptionparams_keyvalue" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
  method              = "AES256"
  keyvalue_wo         = var.nsencryptionparams_keyvalue
  keyvalue_wo_version = 1
}
```

To rotate the key, update the variable value and bump the version:

```hcl
resource "citrixadc_nsencryptionparams" "tf_nsencryptionparams" {
  method              = "AES256"
  keyvalue_wo         = var.nsencryptionparams_keyvalue
  keyvalue_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `method` - (Required) Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256. Possible values: [ NONE, RC4, DES3, AES128, AES192, AES256, DES, DES-CBC, DES-CFB, DES-OFB, DES-ECB, DES3-CBC, DES3-CFB, DES3-OFB, DES3-ECB, AES128-CBC, AES128-CFB, AES128-OFB, AES128-ECB, AES192-CBC, AES192-CFB, AES192-OFB, AES192-ECB, AES256-CBC, AES256-CFB, AES256-OFB, AES256-ECB ]
* `keyvalue` - (Optional, Sensitive) The base64-encoded key generation number, method, and key value. Note: Do not include this argument if you are changing the encryption method. To generate a new key value for the current encryption method, specify an empty string ("") as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup. The value is persisted in Terraform state (encrypted). See also `keyvalue_wo` for an ephemeral alternative.
* `keyvalue_wo` - (Optional, Sensitive, WriteOnly) Same as `keyvalue`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `keyvalue_wo_version`. If both `keyvalue` and `keyvalue_wo` are set, `keyvalue_wo` takes precedence.
* `keyvalue_wo_version` - (Optional) An integer version tracker for `keyvalue_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsencryptionparams. It is a unique string prefixed with "nsencryptionparams-config".


## Import

A nsencryptionparams can be imported using its id, e.g.

```shell
terraform import citrixadc_nsencryptionparams.tf_nsencryptionparams nsencryptionparams-config
```

