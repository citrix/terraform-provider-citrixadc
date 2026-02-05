---
subcategory: "AAA"
---

# Data Source `aaassoprofile`

The aaassoprofile data source allows you to retrieve information about SSO Profiles for single sign-on configuration.


## Example usage

```terraform
data "citrixadc_aaassoprofile" "tf_aaassoprofile" {
  name = "myssoprofile"
}

output "username" {
  value = data.citrixadc_aaassoprofile.tf_aaassoprofile.username
}

output "password" {
  value = data.citrixadc_aaassoprofile.tf_aaassoprofile.password
  sensitive = true
}
```


## Argument Reference

* `name` - (Required) Name for the SSO Profile. Must begin with an ASCII alphabetic or underscore (_) character.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `username` - Name for the user. Must begin with a letter, number, or the underscore (_) character.
* `password` - Password with which the user logs on. Required for Single sign on to external server.

## Attribute Reference

* `id` - The id of the aaassoprofile. It has the same value as the `name` attribute.


## Import

A aaassoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_aaassoprofile.tf_aaassoprofile myssoprofile
```
