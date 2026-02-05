---
subcategory: "AAA"
---

# Data Source `aaatacacsparams`

The aaatacacsparams data source allows you to retrieve information about global TACACS+ parameters configuration.


## Example usage

```terraform
data "citrixadc_aaatacacsparams" "tf_aaatacacsparams" {
}

output "serverip" {
  value = data.citrixadc_aaatacacsparams.tf_aaatacacsparams.serverip
}

output "serverport" {
  value = data.citrixadc_aaatacacsparams.tf_aaatacacsparams.serverport
}

output "authorization" {
  value = data.citrixadc_aaatacacsparams.tf_aaatacacsparams.authorization
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `serverip` - IP address of your TACACS+ server.
* `serverport` - Port number on which the TACACS+ server listens for connections.
* `authtimeout` - Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server.
* `authorization` - Use streaming authorization on the TACACS+ server.
* `tacacssecret` - Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.
* `accounting` - Send accounting messages to the TACACS+ server.
* `auditfailedcmds` - The option for sending accounting messages to the TACACS+ server.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupattrname` - TACACS+ group attribute name. Used for group extraction on the TACACS+ server.

## Attribute Reference

* `id` - The id of the aaatacacsparams. It is a system-generated identifier.
