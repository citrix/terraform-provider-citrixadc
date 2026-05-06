---
subcategory: "AAA"
---

# Resource: aaatacacsparams

The aaatacacsparams resource is used to update aaatacacsparams.


## Example usage

### Using tacacssecret (sensitive attribute - persisted in state)

```hcl
variable "aaatacacsparams_tacacssecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
  serverip      = "10.222.74.159"
  serverport    = 49
  authtimeout   = 3
  authorization = "on"
  tacacssecret  = var.aaatacacsparams_tacacssecret
}
```

### Using tacacssecret_wo (write-only/ephemeral - NOT persisted in state)

The `tacacssecret_wo` attribute provides an ephemeral path for the TACACS+ shared secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `tacacssecret_wo_version`.

```hcl
variable "aaatacacsparams_tacacssecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
  serverip                = "10.222.74.159"
  serverport              = 49
  authtimeout             = 3
  authorization           = "on"
  tacacssecret_wo         = var.aaatacacsparams_tacacssecret
  tacacssecret_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
  serverip                = "10.222.74.159"
  serverport              = 49
  authtimeout             = 3
  authorization           = "on"
  tacacssecret_wo         = var.aaatacacsparams_tacacssecret
  tacacssecret_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `accounting` - (Optional) Send accounting messages to the TACACS+ server. Possible values: [ on, off ]
* `auditfailedcmds` - (Optional) The option for sending accounting messages to the TACACS+ server. Possible values: [ on, off ]
* `authorization` - (Optional) Use streaming authorization on the TACACS+ server. Possible values: [ on, off ]
* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server. Defaults to `3`.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupattrname` - (Optional) TACACS+ group attribute name. Used for group extraction on the TACACS+ server.
* `serverip` - (Optional) IP address of your TACACS+ server.
* `serverport` - (Optional) Port number on which the TACACS+ server listens for connections. Defaults to `49`.
* `tacacssecret` - (Optional, Sensitive) Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server. The value is persisted in Terraform state (encrypted). See also `tacacssecret_wo` for an ephemeral alternative.
* `tacacssecret_wo` - (Optional, Sensitive, WriteOnly) Same as `tacacssecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `tacacssecret_wo_version`. If both `tacacssecret` and `tacacssecret_wo` are set, `tacacssecret_wo` takes precedence.
* `tacacssecret_wo_version` - (Optional) An integer version tracker for `tacacssecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaatacacsparams. It is a unique string prefixed with `aaatacacsparams-config`.