---
subcategory: "Authentication"
---

# Resource: authenticationstorefrontauthaction

The authenticationstorefrontauthaction resource is used to create authentication storefrontauthaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
  name                       = "tf_storefront"
  serverurl                  = "http://www.example.com/"
  domain                     = "domainname"
  defaultauthenticationgroup = "group_name"
}
```


## Argument Reference

* `name` - (Required) Name for the Storefront Authentication action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `serverurl` - (Required) URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/.  Authentication endpoints are learned dynamically by Gateway.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `domain` - (Optional) Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationstorefrontauthaction. It has the same value as the `name` attribute.


## Import

A authenticationstorefrontauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationstorefrontauthaction.tf_storefront tf_storefront
```
