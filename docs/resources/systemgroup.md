---
subcategory: "System"
---

# Resource: systemgroup

The systemgroup resource is used to create user groups.


## Example usage

```hcl
resource "citrixadc_systemgroup" "tf_systemgroup" {
  groupname    = "tf_systemgroup"
  timeout      = 999
  promptstring = "bye>"
}
```


## Argument Reference

* `groupname` - (Optional) Name for the group. 
* `promptstring` - (Optional) String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (\_), and the following variables: * %u - Will be replaced by the user name. * %h - Will be replaced by the hostname of the Citrix ADC. * %t - Will be replaced by the current time in 12-hour format. * %T - Will be replaced by the current time in 24-hour format. * %d - Will be replaced by the current date. * %s - Will be replaced by the state of the Citrix ADC. Note: The 63-character limit for the length of the string does not apply to the characters that replace the variables.
* `timeout` - (Optional) CLI session inactivity timeout, in seconds. If Restrictedtimeout argument of system parameter is enabled, Timeout can have values in the range [300-86400] seconds.If Restrictedtimeout argument of system parameter is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.
* `allowedmanagementinterface` - (Optional) Allowed Management interfaces of the system users in the group. By default allowed from both API and CLI interfaces. If management interface for a group is set to API, then all users under this group will not allowed to access NS through CLI. GUI interface will come under API interface.
Default value: NS_INTERFACE_ALL
Possible values = CLI, API
* `systemusers` - (Optional) A set of user names to bind to this group. (deprecates soon)

!>
[**DEPRECATED**] Please use [`systemgroup_systemuser_binding`](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs/resources/systemgroup_systemuser_binding) to bind `systemuser` to `systemgroup` insted of this resource. The support for binding `systemuser` to `systemgroup` in this resource will get deprecated soon.


* `cmdpolicybinding` - (Optional) A set of command policies to bing to this group. Attributes are detailed below (deprecates soon)

!>
[**DEPRECATED**] Please use [`systemgroup_systemcmdpolicy_binding`](https://registry.terraform.io/providers/citrix/citrixadc/latest/docs/resources/systemgroup_systemcmdpolicy_binding) to bind `systemcmdpolicy` to `systemgroup` insted of this resource. The support for binding `systemcmdpolicy` to `systemgroup` in this resource will get deprecated soon.

In a command policy block the following attributes are allowed:

* `policyname` - (Optional) Name of the policy to bind.
* `priority` - (Optional) Priority for the biding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup. It has the same value as the `groupname` attribute.


## Import

A systemgroup can be imported using its groupname, e.g.

```shell
terraform import citrixadc_systemgroup.tf_systemgroup tf_systemgroup
```
