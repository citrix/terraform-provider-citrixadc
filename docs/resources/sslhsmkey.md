---
subcategory: "SSL"
---

# Resource: sslhsmkey

The sslhsmkey resource is used to create SSL HSM key.


## Example usage

```hcl
resource "citrixadc_sslhsmkey" "demo_sslhsmkey" {
    hsmkeyname = "hsmk1"
    hsmtype = "SAFENET"
    serialnum = "116877xxxx465464"
    password = "xxxxxxx"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "sslhsmkey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslhsmkey" "demo_sslhsmkey" {
  hsmkeyname = "hsmk1"
  hsmtype    = "SAFENET"
  serialnum  = "116877xxxx465464"
  password   = var.sslhsmkey_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the HSM partition password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "sslhsmkey_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslhsmkey" "demo_sslhsmkey" {
  hsmkeyname          = "hsmk1"
  hsmtype             = "SAFENET"
  serialnum           = "116877xxxx465464"
  password_wo         = var.sslhsmkey_password
  password_wo_version = 1
}
```


## Argument Reference

* `hsmkeyname` - (Required) Name for the HSM key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the HSM key is created. 
* `hsmtype` - (Optional) Type of HSM. Possible values: THALES, SAFENET Default value: THALES
* `key` - (Optional) Name of the key. optionally, for THALES, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.
* `keystore` - (Optional) Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.
* `password` - (Optional, Sensitive) Password for a partition. Applies only to SAFENET HSM. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `serialnum` - (Optional) Serial number of the partition on which the key is present. Applies only to SAFENET HSM.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslhsmkey. It has the same value as the `hsmkeyname` attribute.
