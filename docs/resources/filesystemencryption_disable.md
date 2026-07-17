---
subcategory: "Utility"
---

# Resource: filesystemencryption_disable

Disables full file system encryption on a Citrix ADC appliance, decrypting the data at rest on the `/flash` and `/var` file systems. Applying this resource invokes the NITRO `disable` action and requires the same passphrase that was supplied when encryption was enabled. Use this to reverse a previous `citrixadc_filesystemencryption_enable`, for example when decommissioning an appliance or lifting an encryption requirement.

~> **WARNING: This resource performs a real, potentially destructive, platform-gated operation.** Disabling file system encryption zeroes out the `/flash` and `/var` directories (`ntimes0flash` / `ntimes0var` times) and decrypts the file system. Running it in production without understanding the consequences can render the appliance temporarily unavailable or lead to data loss. You must supply the passphrase that was used to enable encryption; without it the operation cannot complete.

-> **Note:** This is an action-only resource. There is no NITRO GET, update, or inverse endpoint for the `disable` action, so Read, Update, and Delete are no-ops. Every attribute is marked `RequiresReplace`, so any change to a configured value re-triggers the `disable` action.


## Example usage

### Using passphrase (sensitive attribute - persisted in state)

```hcl
variable "filesystemencryption_disable_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption_disable" "example" {
  ntimes0flash = 3
  ntimes0var   = 3
  passphrase   = var.filesystemencryption_disable_passphrase
}
```

### Using passphrase_wo (write-only/ephemeral - NOT persisted in state)

The `passphrase_wo` attribute provides an ephemeral path for the encryption passphrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger a re-run of the action when the passphrase value changes, increment `passphrase_wo_version`.

```hcl
variable "filesystemencryption_disable_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption_disable" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_disable_passphrase
  passphrase_wo_version = 1
}
```

To supply a different passphrase, update the variable value and bump the version (this re-triggers the `disable` action):

```hcl
resource "citrixadc_filesystemencryption_disable" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_disable_passphrase
  passphrase_wo_version = 2 # Bumped to re-run the action
}
```


## Argument Reference

* `ntimes0flash` - (Required) Number of times the `/flash` directory has to be written with 0s (zeroed out) during the operation. Changing this attribute re-triggers the `disable` action.
* `ntimes0var` - (Required) Number of times the `/var` directory has to be written with 0s (zeroed out) during the operation. Changing this attribute re-triggers the `disable` action.
* `passphrase` - (Optional, Sensitive) Encryption passphrase used to disable file system encryption. It must match the passphrase that was supplied when encryption was enabled. The value is persisted in Terraform state (encrypted). Either `passphrase` or `passphrase_wo` must be specified. See also `passphrase_wo` for an ephemeral alternative. Changing this attribute re-triggers the `disable` action.
* `passphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `passphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passphrase_wo_version`. If both `passphrase` and `passphrase_wo` are set, `passphrase_wo` takes precedence. Either `passphrase` or `passphrase_wo` must be specified. Changing this attribute re-triggers the `disable` action.
* `passphrase_wo_version` - (Optional) An integer version tracker for `passphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the passphrase has changed and re-trigger the action. Changing this attribute re-triggers the `disable` action.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Used to target a specific node in a cluster. Changing this attribute re-triggers the `disable` action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `filesystemencryption_disable`. It does not correspond to any object on the Citrix ADC.
