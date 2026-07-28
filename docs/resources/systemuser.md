---
subcategory: "System"
---

# Resource: systemuser

The systemuser resource is used to create users for the target ADC.


## Example usage

### Basic usage

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username                   = "tf_user"
  password                   = "tf_password"
  timeout                    = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "systemuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_systemuser" "tf_user" {
  username                   = "tf_user"
  password                   = var.systemuser_password
  timeout                    = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the user password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the password value changes, increment `password_wo_version`.

```hcl
variable "systemuser_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_systemuser" "tf_user" {
  username                   = "tf_user"
  password_wo                = var.systemuser_password
  password_wo_version        = 1
  timeout                    = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```

To rotate the password, update the variable value and bump the version:

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username                   = "tf_user"
  password_wo                = var.systemuser_password
  password_wo_version        = 2  # Bumped to trigger update
  timeout                    = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```

### External authentication user (no local password)

Users that authenticate against an external server (LDAP, RADIUS, TACACS) do not have a local password. Set `externalauth = "ENABLED"` and omit both `password` and `password_wo`.

```hcl
resource "citrixadc_systemuser" "ext_user" {
  username     = "ext_user"
  externalauth = "ENABLED"
}
```

### Command policy bindings

Command policies can be bound to the user inline using one or more `cmdpolicybinding` blocks.

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username = "tf_user"
  password = var.systemuser_password

  cmdpolicybinding {
    policyname = "superuser"
    priority   = 100
  }
}
```


## Password requirement

A local password (via `password` or `password_wo`) is **required only for local users** — users that are authenticated by the ADC itself. It is **not** required in these cases:

* `externalauth = "ENABLED"` — the user authenticates against an external server (LDAP/RADIUS/TACACS) and has no local password. Exempt at both plan and apply time.
* The user is `nsroot`. Exempt at both plan and apply time. A password for `nsroot` is not merely optional but **rejected** — see "Managing the admin (nsroot) user" below.

For any other local user, omitting both `password` and `password_wo` fails validation during `terraform plan`.

~> **Provider login account.** The account the provider uses to authenticate to the ADC is also treated as an admin account, and a password for it is **rejected** here (use [citrixadc_change_password](./change_password.md)). In the typical deployment this account is `nsroot`, which is fully supported. If the provider logs in as a *non-`nsroot`* account, it is only recognized as exempt at apply time — the plan-time check cannot see it, because the provider client is not configured during `terraform plan`/`validate`. Managing such a non-`nsroot` login account through this resource is therefore not recommended unless it authenticates externally (`externalauth = "ENABLED"`).

## Managing the admin (nsroot) user

Special handling applies to `nsroot` and to the account the provider uses to authenticate to the ADC:

* **Password changes are rejected (on both create and update).** Setting `password` or `password_wo` for `nsroot` (or the provider's login user) returns an error. Use the [citrixadc_change_password](./change_password.md) resource to change the admin password instead.
* **Create updates the existing user.** Because `nsroot` already exists on the ADC, applying a `citrixadc_systemuser` for it updates the existing account rather than creating a new one. This lets you manage attributes such as `timeout`, `promptstring`, or command policy bindings for `nsroot`.
* **Destroy does not remove the user.** Destroying a `citrixadc_systemuser` that manages `nsroot` only removes it from Terraform state; the `nsroot` account is left intact on the ADC.


## Argument Reference

* `username` - (Required) Name for a user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added — changing this forces a new resource to be created.
* `password` - (Optional, Sensitive) Password for the system user. Can include any ASCII character. The value is persisted in Terraform state in plaintext — it is marked sensitive (redacted from CLI/plan output), but Terraform state files are not encrypted. Use `password_wo` for an ephemeral alternative that is not stored in state. Required for local users unless `password_wo` is set; see the "Password requirement" section above.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Pair it with `password_wo_version` and increment that version to trigger a password rotation on a later apply (if omitted, `password_wo_version` defaults to `1`). If both `password` and `password_wo` are set, `password_wo` takes precedence. Required for local users unless `password` is set; see the "Password requirement" section above.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the password has changed and trigger an update. Defaults to `1`.
* `allowedmanagementinterface` - (Optional) Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface. Default value: [ NS_INTERFACE_ALL ]. Possible values = [ CLI, API ]
* `externalauth` - (Optional) Whether to use external authentication servers for the system user authentication or not. When set to `ENABLED`, the user has no local password, so `password`/`password_wo` are not required. Possible values: [ ENABLED, DISABLED ]
* `logging` - (Optional) Users logging privilege. Possible values: [ ENABLED, DISABLED ]
* `maxsession` - (Optional) Maximum number of client connection allowed per user.
* `promptstring` - (Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (\_), and the following variables: * %u - Will be replaced by the user name. * %h - Will be replaced by the hostname of the Citrix ADC. * %t - Will be replaced by the current time in 12-hour format. * %T - Will be replaced by the current time in 24-hour format. * %d - Will be replaced by the current date. * %s - Will be replaced by the state of the Citrix ADC. Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
* `timeout` - (Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.

### cmdpolicybinding

The `cmdpolicybinding` block binds a command policy to the user. It may be specified zero or more times and supports:

* `policyname` - (Optional) The name of the command policy to bind to the user.
* `priority` - (Optional) The priority (evaluation order) of the bound command policy.

~> If you use the separate [citrixadc_systemuser_systemcmdpolicy_binding](./systemuser_systemcmdpolicy_binding.md) resource to bind command policies to this user, do not also define `cmdpolicybinding` blocks here — otherwise the two would fight over the same bindings.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser. It has the same value as the `username` attribute.
* `hashedpassword` - (Computed) The hashed password as returned by the NITRO API. It is populated and managed by the provider to detect out-of-band password changes on the ADC, and should not be set in configuration.


## Import

A systemuser can be imported using its name, e.g.

```shell
terraform import citrixadc_systemuser.tf_user tf_user
```
