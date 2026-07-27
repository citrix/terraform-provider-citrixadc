---
subcategory: "Utility"
---

# Resource: change_password

This resource is used to change the password of a Citrix ADC user. The
`first_time_password_reset` flag selects the operation:

* `true` - performs the default password reset at first login (via the `login`
  endpoint), authenticating with the current password and setting the new one.
* `false` - changes the password of an existing user at any later time (via a
  `systemuser` update, equivalent to `set system user <username> <new_password>`).

It is a one-shot action: every attribute forces a new resource, and there is no
read-back or delete operation on the appliance.

Both secrets support a write-only (ephemeral) path so that values sourced from a
variable, Vault, or other ephemeral source are **not persisted in Terraform state**.


## Example usage

### Using `password` / `new_password` (sensitive attributes - persisted in state)

```hcl
resource "citrixadc_change_password" "tf_change_password" {
  username                  = "nsroot"
  password                  = "secret"
  new_password              = "verysecret"
  first_time_password_reset = false
}
```

### Using the write-only variants (ephemeral - NOT persisted in state)

The `_wo` attributes provide an ephemeral path for the credentials. The values are
sent to the ADC but are **not stored in Terraform state**, reducing the risk of secret
exposure. To re-run the operation when a value changes (secret rotation), increment
the corresponding `_wo_version`.

```hcl
variable "current_password" {
  type      = string
  sensitive = true
}

variable "new_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_change_password" "tf_change_password" {
  username                  = "nsroot"
  password_wo               = var.current_password
  password_wo_version       = 1
  new_password_wo           = var.new_password
  new_password_wo_version   = 1
  first_time_password_reset = false
}
```

~>
│ Error: [ERROR] nitro-go: Failed to create resource of type login, name=, err=failed: 599 Netscaler specific error ({ "errorcode": 278, "message": "Invalid argument", "severity": "ERROR" })
If you observe the above error, it may be because `first_time_password_reset` is set to `true` but you want to change the password at a later point. In that case, set `first_time_password_reset` to `false`.


## Argument Reference

* `username` - (Required) User name for the operation.
* `password` - (Optional, Sensitive) The default (current) password. The value is persisted in Terraform state. See also `password_wo` for an ephemeral alternative. Either `password` or `password_wo` must be specified.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Should be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and re-run the operation. Defaults to `1`.
* `new_password` - (Optional, Sensitive) The new password. The value is persisted in Terraform state. See also `new_password_wo` for an ephemeral alternative. Either `new_password` or `new_password_wo` must be specified.
* `new_password_wo` - (Optional, Sensitive, WriteOnly) Same as `new_password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Should be used together with `new_password_wo_version`. If both `new_password` and `new_password_wo` are set, `new_password_wo` takes precedence.
* `new_password_wo_version` - (Optional) An integer version tracker for `new_password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and re-run the operation. Defaults to `1`.
* `first_time_password_reset` - (Required) Boolean value. Set to `true` to perform the default password reset operation at first login, or `false` to change the password of an existing user at any later time.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the change_password resource. It has the value `change_password-<username>`.
