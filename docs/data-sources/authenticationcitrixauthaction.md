---
subcategory: "Authentication"
---

# Data Source `authenticationcitrixauthaction`

The authenticationcitrixauthaction data source allows you to retrieve information about authentication Citrix auth actions.


## Example usage

```terraform
data "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
  name = "my_citrixauthaction"
}

output "authenticationtype" {
  value = data.citrixadc_authenticationcitrixauthaction.tf_citrixauthaction.authenticationtype
}

output "authentication" {
  value = data.citrixadc_authenticationcitrixauthaction.tf_citrixauthaction.authentication
}
```


## Argument Reference

* `name` - (Required) Name for the new Citrix Authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `authentication` - Authentication needs to be disabled for searching user object without performing authentication.
* `authenticationtype` - Type of the Citrix Authentication implementation. Default implementation uses Citrix Cloud Connector.

## Attribute Reference

* `id` - The id of the authenticationcitrixauthaction. It has the same value as the `name` attribute.


## Import

A authenticationcitrixauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcitrixauthaction.tf_citrixauthaction my_citrixauthaction
```
