---
subcategory: "AAA"
---

# Resource: aaauser

The `citrixadc_aaauser` resource is used to create and manage AAA local user accounts on Citrix ADC.


## Example usage

### Using `password` (sensitive attribute - persisted in state)

```hcl
variable "aaauser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaauser" "tf_aaauser" {
  username = "john"
  password = var.aaauser_password
}
```

### Using `password_wo` (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the user password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "aaauser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaauser" "tf_aaauser" {
  username            = "john"
  password_wo         = var.aaauser_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_aaauser" "tf_aaauser" {
  username            = "john"
  password_wo         = var.aaauser_password
  password_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `username` - (Required) Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa user" or "my aaa user").
* `password` - (Optional, Sensitive) Password with which the user logs on. Required for any user account that does not exist on an external authentication server. If you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `loggedin` - (Optional) Show whether the user is logged in or not.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the aaauser resource. It has the same value as the `username` attribute.


## Import

A aaauser can be imported using its name, e.g.

```shell
terraform import citrixadc_aaauser.tf_aaauser john
```
