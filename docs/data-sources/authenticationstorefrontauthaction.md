---
subcategory: "Authentication"
---

# Data Source `authenticationstorefrontauthaction`

The authenticationstorefrontauthaction data source allows you to retrieve information about an existing Storefront authentication action.


## Example usage

```terraform
data "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
  name = "my_storefront_action"
}

output "serverurl" {
  value = data.citrixadc_authenticationstorefrontauthaction.tf_storefront.serverurl
}

output "domain" {
  value = data.citrixadc_authenticationstorefrontauthaction.tf_storefront.domain
}

output "defaultauthenticationgroup" {
  value = data.citrixadc_authenticationstorefrontauthaction.tf_storefront.defaultauthenticationgroup
}
```


## Argument Reference

* `name` - (Required) Name for the Storefront Authentication action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationstorefrontauthaction. It has the same value as the `name` attribute.
* `serverurl` - URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/. Authentication endpoints are learned dynamically by Gateway.
* `domain` - Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.


## Import

A authenticationstorefrontauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationstorefrontauthaction.tf_storefront my_storefront_action
```
