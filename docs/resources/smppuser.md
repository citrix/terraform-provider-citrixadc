---
subcategory: "SMPP"
---

# Resource: smppuser

The smppuser resource is used to create smppuser.


## Example usage

### Basic usage

```hcl
resource "citrixadc_smppuser" "tf_smppuser" {
  username = "user1"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "smppuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_smppuser" "tf_smppuser" {
  username = "user1"
  password = var.smppuser_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the SMPP user password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the password value changes, increment `password_wo_version`.

```hcl
variable "smppuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_smppuser" "tf_smppuser" {
  username          = "user1"
  password_wo       = var.smppuser_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_smppuser" "tf_smppuser" {
  username          = "user1"
  password_wo       = var.smppuser_password
  password_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `password` - (Optional, Sensitive) Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `username` - (Required) Name of the SMPP user. Must be the same as the user name specified in the SMPP server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the smppuser. It has the same value as the `name` attribute.


## Import

A smppuser can be imported using its name, e.g.

```shell
terraform import citrixadc_smppuser.tf_smppuser user1
```
