---
subcategory: "Database"
---

# Resource: dbuser

The dbuser resource is used to create and manage database users on Citrix ADC.


## Example usage

### Basic usage

```hcl
resource "citrixadc_dbuser" "tf_dbuser" {
  username = "dbuser1"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "dbuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser" {
  username = "dbuser1"
  password = var.dbuser_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the database password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the password value changes, increment `password_wo_version`.

```hcl
variable "dbuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_dbuser" "tf_dbuser" {
  username           = "dbuser1"
  password_wo        = var.dbuser_password
  password_wo_version = 1
}
```

To rotate the password, update the variable value and bump the version:

```hcl
resource "citrixadc_dbuser" "tf_dbuser" {
  username           = "dbuser1"
  password_wo        = var.dbuser_password
  password_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `username` - (Required) Name of the database user. Must be the same as the user name specified in the database.
* `password` - (Optional, Sensitive) Password for logging on to the database. Must be the same as the password specified in the database. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `loggedin` - (Optional) Display the names of all database users currently logged on to the Citrix ADC.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbuser. It has the same value as the `username` attribute.


## Import

A dbuser can be imported using its name, e.g.

```shell
terraform import citrixadc_dbuser.tf_dbuser dbuser1
```
