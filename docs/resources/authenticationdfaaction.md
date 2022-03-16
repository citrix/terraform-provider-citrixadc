---
subcategory: "Authentication"
---

# Resource: authenticationdfaaction

The authenticationdfaaction resource is used to create Authentication dfa action Resource.


## Example usage

```hcl
resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name       = "tf_dfaaction"
  serverurl  = "https://example.com/"
  clientid   = "cliId"
  passphrase = "secret"
}
```


## Argument Reference

* `name` - (Required) Name for the DFA action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.
* `passphrase` - (Required) Key shared between the DFA server and the Citrix ADC.  Required to allow the Citrix ADC to communicate with the DFA server.
* `serverurl` - (Required) DFA Server URL
* `clientid` - (Required) If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationdfaaction. It has the same value as the `name` attribute.


## Import

A authenticationdfaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationdfaaction.tf_dfaaction tf_dfaaction
```
