---
subcategory: "AAA"
---

# Resource: aaatacacsparams

The aaatacacsparams resource is used to update aaatacacsparams.


## Example usage

```hcl
resource "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
  serverip      = "10.222.74.159"
  serverport    = 50
  authtimeout   = 6
  authorization = "ON"
}
```


## Argument Reference

* `serverip` - (Optional) IP address of your TACACS+ server. Minimum length =  1
* `serverport` - (Optional) Port number on which the TACACS+ server listens for connections. Minimum value =  1
* `authtimeout` - (Optional) Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server. Minimum value =  1
* `tacacssecret` - (Optional) Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server. Minimum length =  1
* `authorization` - (Optional) Use streaming authorization on the TACACS+ server. Possible values: [ on, off ]
* `accounting` - (Optional) Send accounting messages to the TACACS+ server. Possible values: [ on, off ]
* `auditfailedcmds` - (Optional) The option for sending accounting messages to the TACACS+ server. Possible values: [ on, off ]
* `groupattrname` - (Optional) TACACS+ group attribute name.Used for group extraction on the TACACS+ server.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaatacacsparams. It is a unique string prefixed with  `tf-aaatacacsparams-`.