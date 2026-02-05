---
subcategory: "Authentication"
---

# Data Source `authenticationnoauthaction`

The authenticationnoauthaction data source allows you to retrieve information about authentication no-authentication actions.


## Example usage

```terraform
data "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
  name = "my_noauthaction"
}

output "defaultauthenticationgroup" {
  value = data.citrixadc_authenticationnoauthaction.tf_noauthaction.defaultauthenticationgroup
}
```


## Argument Reference

* `name` - (Required) Name for the new no-authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `defaultauthenticationgroup` - This is the group that is added to user sessions that match current policy.

## Attribute Reference

* `id` - The id of the authenticationnoauthaction. It has the same value as the `name` attribute.


## Import

A authenticationnoauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationnoauthaction.tf_noauthaction my_noauthaction
```
