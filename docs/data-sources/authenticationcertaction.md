---
subcategory: "Authentication"
---

# Data Source `authenticationcertaction`

The authenticationcertaction data source allows you to retrieve information about authentication certificate actions.


## Example usage

```terraform
data "citrixadc_authenticationcertaction" "tf_certaction" {
  name = "my_certaction"
}

output "twofactor" {
  value = data.citrixadc_authenticationcertaction.tf_certaction.twofactor
}

output "usernamefield" {
  value = data.citrixadc_authenticationcertaction.tf_certaction.usernamefield
}
```


## Argument Reference

* `name` - (Required) Name for the client cert authentication server profile (action). Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after certificate action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupnamefield` - Client-cert field from which the group is extracted. Must be set to either "Subject" and "Issuer" (include both sets of double quotation marks). Format: <field>:<subfield>
* `twofactor` - Enables or disables two-factor authentication. Two factor authentication is client cert authentication followed by password authentication.
* `usernamefield` - Client-cert field from which the username is extracted. Must be set to either "Subject" and "Issuer" (include both sets of double quotation marks). Format: <field>:<subfield>.

## Attribute Reference

* `id` - The id of the authenticationcertaction. It has the same value as the `name` attribute.
