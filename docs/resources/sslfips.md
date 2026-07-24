---
subcategory: "SSL"
---

# Resource: sslfips

This resource is used to configure the FIPS Hardware Security Module (HSM) on a Citrix ADC FIPS appliance.

!> **WARNING:** Requires a dedicated FIPS appliance with an on-board HSM. Setting `inithsm` initializes the HSM and ERASES all FIPS key and certificate data; the operation is irreversible.


## Example usage

### Using the sensitive password attributes (persisted in state)

```hcl
variable "sslfips_sopassword" {
  type      = string
  sensitive = true
}

variable "sslfips_oldsopassword" {
  type      = string
  sensitive = true
}

variable "sslfips_userpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfips" "tf_sslfips" {
  inithsm       = "Level-2"
  hsmlabel      = "myhsm"
  sopassword    = var.sslfips_sopassword
  oldsopassword = var.sslfips_oldsopassword
  userpassword  = var.sslfips_userpassword
}
```

### Using the write-only password attributes (NOT persisted in state)

The `*_wo` attributes provide an ephemeral path for the FIPS passwords. The values are sent to the Citrix ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when a value changes, increment the corresponding `*_wo_version`.

```hcl
variable "sslfips_sopassword" {
  type      = string
  sensitive = true
}

variable "sslfips_oldsopassword" {
  type      = string
  sensitive = true
}

variable "sslfips_userpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslfips" "tf_sslfips" {
  inithsm                  = "Level-2"
  hsmlabel                 = "myhsm"
  sopassword_wo            = var.sslfips_sopassword
  sopassword_wo_version    = 1
  oldsopassword_wo         = var.sslfips_oldsopassword
  oldsopassword_wo_version = 1
  userpassword_wo          = var.sslfips_userpassword
  userpassword_wo_version  = 1
}
```

To rotate a secret, update the variable value and bump the corresponding version:

```hcl
resource "citrixadc_sslfips" "tf_sslfips" {
  inithsm                  = "Level-2"
  hsmlabel                 = "myhsm"
  sopassword_wo            = var.sslfips_sopassword
  sopassword_wo_version    = 2  # Bumped to trigger update
  oldsopassword_wo         = var.sslfips_oldsopassword
  oldsopassword_wo_version = 1
  userpassword_wo          = var.sslfips_userpassword
  userpassword_wo_version  = 1
}
```

## Argument Reference

* `inithsm` - (Required) FIPS initialization level. The appliance currently supports Level-2 (FIPS 140-2). Possible values: [ Level-2 ]. Note: applying this value performs a destructive HSM initialization (see the warning above).
* `hsmlabel` - (Optional) Label to identify the Hardware Security Module (HSM).
* `fipsfw` - (Optional) Path to the FIPS firmware file. Used to drive the FIPS firmware update action. Note: the NITRO GET response does not echo this value back; the provider preserves the user-configured value in state.
* `sopassword` - (Optional, Sensitive) Security officer password that will be in effect after you have configured the HSM. The value is persisted in Terraform state (encrypted). See also `sopassword_wo` for an ephemeral alternative. Either `sopassword` or `sopassword_wo` must be specified.
* `sopassword_wo` - (Optional, Sensitive, WriteOnly) Same as `sopassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `sopassword_wo_version`. If both `sopassword` and `sopassword_wo` are set, `sopassword_wo` takes precedence.
* `sopassword_wo_version` - (Optional) An integer version tracker for `sopassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `oldsopassword` - (Optional, Sensitive) Old password for the security officer. The value is persisted in Terraform state (encrypted). See also `oldsopassword_wo` for an ephemeral alternative. Either `oldsopassword` or `oldsopassword_wo` must be specified.
* `oldsopassword_wo` - (Optional, Sensitive, WriteOnly) Same as `oldsopassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `oldsopassword_wo_version`. If both `oldsopassword` and `oldsopassword_wo` are set, `oldsopassword_wo` takes precedence.
* `oldsopassword_wo_version` - (Optional) An integer version tracker for `oldsopassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `userpassword` - (Optional, Sensitive) The Hardware Security Module's (HSM) User password. The value is persisted in Terraform state (encrypted). See also `userpassword_wo` for an ephemeral alternative. Either `userpassword` or `userpassword_wo` must be specified.
* `userpassword_wo` - (Optional, Sensitive, WriteOnly) Same as `userpassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `userpassword_wo_version`. If both `userpassword` and `userpassword_wo` are set, `userpassword_wo` takes precedence.
* `userpassword_wo_version` - (Optional) An integer version tracker for `userpassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslfips. It is set to `sslfips-config`.

## Import

A sslfips can be imported using its id (the synthetic singleton constant), e.g.

```shell
terraform import citrixadc_sslfips.tf_sslfips sslfips-config
```
