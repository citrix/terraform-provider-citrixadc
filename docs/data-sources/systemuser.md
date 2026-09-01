---
subcategory: "System"
---

# Data Source: systemuser

The systemuser data source allows you to retrieve information about a system user configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_systemuser" "tf_user" {
  username = "tf_user"
}

output "timeout" {
  value = data.citrixadc_systemuser.tf_user.timeout
}

output "logging" {
  value = data.citrixadc_systemuser.tf_user.logging
}
```

## Argument Reference

* `username` - (Required) Name of the system user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `allowedmanagementinterface` - Allowed Management interfaces to the system user. By default user is allowed from both API and CLI interfaces. Possible values: [ CLI, API ]
* `externalauth` - Whether to use external authentication servers for the system user authentication or not. Possible values: [ ENABLED, DISABLED ]
* `id` - The id of the systemuser. It has the same value as the `username` attribute.
* `logging` - Users logging privilege. Possible values: [ ENABLED, DISABLED ]
* `maxsession` - Maximum number of client connection allowed per user.
* `password` - Password for the system user. Can include any ASCII character.
* `promptstring` - String to display at the command-line prompt.
* `timeout` - CLI session inactivity timeout, in seconds.

### Read-only systemuser metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_systemuser` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `encrypted` - Whether the password stored on the appliance is encrypted.
* `hashmethod` - Hash method used for the system user password (for example `SHA1`, `SHA512`, `PBKDF2`).
* `promptinheritedfrom` - From where the prompt has been inherited (`User`, `Group`, `Global`, `Climode`).
* `timeoutkind` - From where the timeout has been inherited (`User`, `Group`, `Global`, `Climode`).
* `allowedmanagementinterfacekind` - Value of allowed interface which can be inherited from `Global`, `Group` or `User` (`User`, `Group`, `Global`, `Climode`).
* `lastpwdchangetimestamp` - Timestamp for the last password change for the system user.
* `daystoexpirekind` - From where the daystoexpire value has been inherited (`User`, `Group`, `Global`, `Climode`).

## Import

A systemuser can be imported using its username, e.g.

```shell
terraform import citrixadc_systemuser.tf_user tf_user
```
