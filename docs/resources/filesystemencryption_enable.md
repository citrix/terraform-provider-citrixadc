---
subcategory: "Utility"
---

# Resource: filesystemencryption_enable

Enables full file system encryption on a Citrix ADC appliance, protecting the data at rest on the `/flash` and `/var` file systems with a user-supplied passphrase. Applying this resource invokes the NITRO `enable` action. Use this when compliance or physical-security requirements mandate encryption of the on-disk configuration and data. Disabling encryption is a separate action, modelled by the `citrixadc_filesystemencryption_disable` resource.

~> **WARNING: This resource performs a real, potentially destructive, platform-gated operation.** Enabling file system encryption zeroes out the `/flash` and `/var` directories (`ntimes0flash` / `ntimes0var` times) and re-encrypts the file system. It is only supported on hardware/platforms that report `supportedstate = "ENABLED"`. Attempting to enable it on an unsupported platform will fail, and running it in production without understanding the consequences can render the appliance temporarily unavailable or lead to data loss. Verify `supportedstate` (for example via the `citrixadc_filesystemencryption` data source) before applying, and retain the passphrase — it is required to later `disable` encryption.

-> **Note:** This is an action-only resource. There is no NITRO GET, update, or inverse endpoint for the `enable` action, so Read, Update, and Delete are no-ops. Every attribute is marked `RequiresReplace`, so any change to a configured value re-triggers the `enable` action.


## Example usage

### Using passphrase (sensitive attribute - persisted in state)

```hcl
variable "filesystemencryption_enable_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption_enable" "example" {
  ntimes0flash = 3
  ntimes0var   = 3
  passphrase   = var.filesystemencryption_enable_passphrase
}
```

### Using passphrase_wo (write-only/ephemeral - NOT persisted in state)

The `passphrase_wo` attribute provides an ephemeral path for the encryption passphrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger a re-run of the action when the passphrase value changes, increment `passphrase_wo_version`.

```hcl
variable "filesystemencryption_enable_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption_enable" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_enable_passphrase
  passphrase_wo_version = 1
}
```

To rotate the passphrase, update the variable value and bump the version (this re-triggers the `enable` action):

```hcl
resource "citrixadc_filesystemencryption_enable" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_enable_passphrase
  passphrase_wo_version = 2 # Bumped to re-run the action
}
```


## Argument Reference

* `ntimes0flash` - (Required) Number of times the `/flash` directory has to be written with 0s (zeroed out) before encryption. Changing this attribute re-triggers the `enable` action.
* `ntimes0var` - (Required) Number of times the `/var` directory has to be written with 0s (zeroed out) before encryption. Changing this attribute re-triggers the `enable` action.
* `passphrase` - (Optional, Sensitive) Encryption passphrase used to enable file system encryption. The value is persisted in Terraform state (encrypted). Retain this value — it is required to later `disable` encryption. Either `passphrase` or `passphrase_wo` must be specified. See also `passphrase_wo` for an ephemeral alternative. Changing this attribute re-triggers the `enable` action.
* `passphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `passphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passphrase_wo_version`. If both `passphrase` and `passphrase_wo` are set, `passphrase_wo` takes precedence. Either `passphrase` or `passphrase_wo` must be specified. Changing this attribute re-triggers the `enable` action.
* `passphrase_wo_version` - (Optional) An integer version tracker for `passphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the passphrase has changed and re-trigger the action. Changing this attribute re-triggers the `enable` action.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Used to target a specific node in a cluster. Changing this attribute re-triggers the `enable` action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `filesystemencryption_enable`. It does not correspond to any object on the Citrix ADC.
