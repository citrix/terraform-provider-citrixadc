---
subcategory: "Authentication"
---

# Resource: authenticationdfaaction

The authenticationdfaaction resource is used to create Authentication dfa action Resource.


## Example usage

### Using passphrase (sensitive attribute - persisted in state)

```hcl
variable "authenticationdfaaction_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name       = "tf_dfaaction"
  serverurl  = "https://example.com/"
  clientid   = "cliId"
  passphrase = var.authenticationdfaaction_passphrase
}
```

### Using passphrase_wo (write-only/ephemeral - NOT persisted in state)

The `passphrase_wo` attribute provides an ephemeral path for the DFA server shared key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `passphrase_wo_version`.

```hcl
variable "authenticationdfaaction_passphrase" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name                  = "tf_dfaaction"
  serverurl             = "https://example.com/"
  clientid              = "cliId"
  passphrase_wo         = var.authenticationdfaaction_passphrase
  passphrase_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationdfaaction" "tf_dfaaction" {
  name                  = "tf_dfaaction"
  serverurl             = "https://example.com/"
  clientid              = "cliId"
  passphrase_wo         = var.authenticationdfaaction_passphrase
  passphrase_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name for the DFA action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.
* `passphrase` - (Optional, Sensitive) Key shared between the DFA server and the Citrix ADC. Required to allow the Citrix ADC to communicate with the DFA server. The value is persisted in Terraform state (encrypted). See also `passphrase_wo` for an ephemeral alternative.
* `passphrase_wo` - (Optional, Sensitive, WriteOnly) Same as `passphrase`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passphrase_wo_version`. If both `passphrase` and `passphrase_wo` are set, `passphrase_wo` takes precedence.
* `passphrase_wo_version` - (Optional) An integer version tracker for `passphrase_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
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
