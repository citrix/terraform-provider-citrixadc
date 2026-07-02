---
subcategory: "Utility"
---

# Resource: filesystemencryption

Enables full file system encryption on a Citrix ADC appliance, protecting the data at rest on the `/flash` and `/var` file systems with a user-supplied passphrase. Applying this resource runs the `enable` action; destroying it runs the `disable` action. Use this when compliance or physical-security requirements mandate encryption of on-disk configuration and data.

~> **WARNING: This resource performs a real, potentially destructive, platform-gated operation.** Enabling file system encryption zeroes out the `/flash` and `/var` directories (`ntimes0flash` / `ntimes0var` times) and re-encrypts the file system. It is only supported on hardware/platforms that report `supportedstate = "ENABLED"`. Attempting to enable it on an unsupported platform will fail, and running it in production without understanding the consequences can render the appliance temporarily unavailable or lead to data loss. Verify `supportedstate` (for example via the `citrixadc_filesystemencryption` data source) before applying, and ensure you have retained the passphrase — it is required to disable encryption again.

-> **Note:** This is an action-only resource. There is no update endpoint on the NITRO API. Every attribute is marked `RequiresReplace`, so any change to a configured value forces Terraform to destroy the resource (running `disable`) and recreate it (running `enable`).


## Example usage

### Using passphrase (sensitive attribute - persisted in state)

```hcl
variable "filesystemencryption_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption" "example" {
  ntimes0flash = 3
  ntimes0var   = 3
  passphrase   = var.filesystemencryption_passphrase
}
```

### Using passphrase_wo (write-only/ephemeral - NOT persisted in state)

The `passphrase_wo` attribute provides an ephemeral path for the encryption passphrase. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the passphrase value changes, increment `passphrase_wo_version`.

```hcl
variable "filesystemencryption_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_filesystemencryption" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_passphrase
  passphrase_wo_version = 1
}
```

To rotate the passphrase, update the variable value and bump the version (this forces a destroy/recreate, which re-runs the encryption action):

```hcl
resource "citrixadc_filesystemencryption" "example" {
  ntimes0flash          = 3
  ntimes0var            = 3
  passphrase_wo         = var.filesystemencryption_passphrase
  passphrase_wo_version = 2 # Bumped to trigger recreation
}
```


## Argument Reference

* `ntimes0flash` - (Required) Number of times the `/flash` directory has to be written with 0s (zeroed out) before encryption. Minimum value = 0 Maximum value = 16. Changing this value forces recreation.
* `ntimes0var` - (Required) Number of times the `/var` directory has to be written with 0s (zeroed out) before encryption. Minimum value = 0 Maximum value = 16. Changing this value forces recreation.
* `passphrase` - (Optional, Sensitive) Encryption passphrase used to enable and disable file system encryption. The value is persisted in Terraform state (encrypted). Retain this value — it is required to later `disable` encryption. Either `passphrase` or `passphrase_wo` must be specified. See also `passphrase_wo` for an ephemeral alternative. Changing this value forces recreation.
* `passphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `passphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passphrase_wo_version`. If both `passphrase` and `passphrase_wo` are set, `passphrase_wo` takes precedence. Note: because the write-only passphrase is not stored in state, it is not available at destroy time — persist `passphrase` instead if you rely on Terraform to `disable` encryption. Changing this value forces recreation.
* `passphrase_wo_version` - (Optional) An integer version tracker for `passphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the passphrase has changed and trigger a recreation. Changing this value forces recreation.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Used to target the read (GET) operation to a specific node in a cluster. Changing this value forces recreation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the filesystemencryption resource. It is a static string, `"filesystemencryption-config"`, because the resource is a nameless singleton.
* `supportedstate` - Whether the platform supports file system encryption. Possible values: [ DISABLED, ENABLED, UNKNOWN ]. Encryption can only be enabled when this is `ENABLED`.
* `effectivestate` - The current encrypted state of the file system. Possible values: [ ENABLED, DISABLED ].


## Import

This is a nameless singleton resource. It can be imported using its static ID:

```shell
terraform import citrixadc_filesystemencryption.example filesystemencryption-config
```

Note: the `passphrase` / `passphrase_wo` value cannot be recovered from the appliance and will not be populated by import. Set it in configuration after importing.
