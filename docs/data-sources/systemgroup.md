---
subcategory: "System"
---

# Data Source: systemgroup

The systemgroup data source allows you to retrieve information about system user groups on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_systemgroup" "example" {
  groupname = "my_systemgroup"
}

output "timeout" {
  value = data.citrixadc_systemgroup.example.timeout
}

output "promptstring" {
  value = data.citrixadc_systemgroup.example.promptstring
}

output "daystoexpire" {
  value = data.citrixadc_systemgroup.example.daystoexpire
}
```

## Argument Reference

* `groupname` - (Required) Name for the group. Must begin with a letter, number, hash(#) or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `allowedmanagementinterface` - Allowed Management interfaces of the system users in the group. By default allowed from both API and CLI interfaces. If management interface for a group is set to API, then all users under this group will not allowed to access NS through CLI. GUI interface will come under API interface.
* `daystoexpire` - Password days to expire for system groups. The daystoexpire value ranges from 30 to 255.
* `promptstring` - String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:
  * %u - Will be replaced by the user name.
  * %h - Will be replaced by the hostname of the Citrix ADC.
  * %t - Will be replaced by the current time in 12-hour format.
  * %T - Will be replaced by the current time in 24-hour format.
  * %d - Will be replaced by the current date.
  * %s - Will be replaced by the state of the Citrix ADC.

  Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
* `timeout` - CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds. If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.
* `warnpriorndays` - Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.
* `id` - The id of the systemgroup. It has the same value as the `groupname` attribute.

## Import

A systemgroup can be imported using its groupname, e.g.

```shell
terraform import citrixadc_systemgroup.example my_systemgroup
```
