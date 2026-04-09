---
subcategory: "System"
---

# Resource: systemuser

The systemuser resource is used to create users for the target ADC.


## Example usage

### Basic usage

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username                  = "tf_user"
  password                  = "tf_password"
  timeout                   = 200
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
  username                  = "tf_user"
  password                  = var.systemuser_password
  timeout                   = 200
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
  username                  = "tf_user"
  password_wo               = var.systemuser_password
  password_wo_version       = 1
  timeout                   = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```

To rotate the password, update the variable value and bump the version:

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username                  = "tf_user"
  password_wo               = var.systemuser_password
  password_wo_version       = 2  # Bumped to trigger update
  timeout                   = 200
  allowedmanagementinterface = ["CLI", "API"]
}
```


## Argument Reference

* `username` - (Required) Name for a user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.
* `password` - (Optional, Sensitive) Password for the system user. Can include any ASCII character. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the password has changed and trigger an update. Defaults to `1`.
* `allowedmanagementinterface` - (Optional) Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. If management interface for a user is set to API, then user is not allowed to access NS through CLI. GUI interface will come under API interface. Default value: [ NS_INTERFACE_ALL ]. Possible values = [ CLI, API ]
* `externalauth` - (Optional) Whether to use external authentication servers for the system user authentication or not. Possible values: [ ENABLED, DISABLED ]
* `logging` - (Optional) Users logging privilege. Possible values: [ ENABLED, DISABLED ]
* `maxsession` - (Optional) Maximum number of client connection allowed per user.
* `promptstring` - (Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (\_), and the following variables: * %u - Will be replaced by the user name. * %h - Will be replaced by the hostname of the Citrix ADC. * %t - Will be replaced by the current time in 12-hour format. * %T - Will be replaced by the current time in 24-hour format. * %d - Will be replaced by the current date. * %s - Will be replaced by the state of the Citrix ADC. Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
* `timeout` - (Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser. It has the same value as the `username` attribute.


## Import

A systemuser can be imported using its name, e.g.

```shell
terraform import citrixadc_systemuser.tf_user tf_user
```
