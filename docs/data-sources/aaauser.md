---
subcategory: "AAA"
---

# Data Source `aaauser`

The aaauser data source allows you to retrieve information about AAA users.


## Example usage

```terraform
data "citrixadc_aaauser" "tf_aaauser" {
  username = "john"
}

output "username" {
  value = data.citrixadc_aaauser.tf_aaauser.username
}

output "loggedin" {
  value = data.citrixadc_aaauser.tf_aaauser.loggedin
}
```


## Argument Reference

* `username` - (Required) Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `loggedin` - Shows whether the user is logged in or not.
* `password` - Password with which the user logs on. Required for any user account that does not exist on an external authentication server.

## Attribute Reference

* `id` - The id of the aaauser. It has the same value as the `username` attribute.


## Import

A aaauser can be imported using its username, e.g.

```shell
terraform import citrixadc_aaauser.tf_aaauser john
```
