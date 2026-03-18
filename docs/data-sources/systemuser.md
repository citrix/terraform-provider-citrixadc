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

## Import

A systemuser can be imported using its username, e.g.

```shell
terraform import citrixadc_systemuser.tf_user tf_user
```
