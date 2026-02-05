---
subcategory: "Authentication"
---

# Data Source `authenticationdfaaction`

The authenticationdfaaction data source allows you to retrieve information about authentication DFA (Decentralized Factor Authentication) actions.


## Example usage

```terraform
data "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name = "my_dfaaction"
}

output "serverurl" {
  value = data.citrixadc_authenticationdfaaction.tf_dfaaction.serverurl
}

output "clientid" {
  value = data.citrixadc_authenticationdfaaction.tf_dfaaction.clientid
}
```


## Argument Reference

* `name` - (Required) Name for the DFA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `clientid` - If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `passphrase` - Key shared between the DFA server and the Citrix ADC. Required to allow the Citrix ADC to communicate with the DFA server.
* `serverurl` - DFA Server URL.

## Attribute Reference

* `id` - The id of the authenticationdfaaction. It has the same value as the `name` attribute.


## Import

A authenticationdfaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationdfaaction.tf_dfaaction my_dfaaction
```
