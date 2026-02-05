---
subcategory: "Authentication"
---

# Data Source `authenticationtacacsaction`

The authenticationtacacsaction data source allows you to retrieve information about an existing TACACS+ authentication action.


## Example usage

```terraform
data "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
  name = "my_tacacsaction"
}

output "serverip" {
  value = data.citrixadc_authenticationtacacsaction.tf_tacacsaction.serverip
}

output "serverport" {
  value = data.citrixadc_authenticationtacacsaction.tf_tacacsaction.serverport
}

output "authorization" {
  value = data.citrixadc_authenticationtacacsaction.tf_tacacsaction.authorization
}
```


## Argument Reference

* `name` - (Required) Name for the TACACS+ profile (action). Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS profile is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationtacacsaction. It has the same value as the `name` attribute.
* `accounting` - Whether the TACACS+ server is currently accepting accounting messages.
* `attribute1` - Name of the custom attribute to be extracted from server and stored at index '1' (where '1' changes for each attribute)
* `attribute10` - Name of the custom attribute to be extracted from server and stored at index '10' (where '10' changes for each attribute)
* `attribute11` - Name of the custom attribute to be extracted from server and stored at index '11' (where '11' changes for each attribute)
* `attribute12` - Name of the custom attribute to be extracted from server and stored at index '12' (where '12' changes for each attribute)
* `attribute13` - Name of the custom attribute to be extracted from server and stored at index '13' (where '13' changes for each attribute)
* `attribute14` - Name of the custom attribute to be extracted from server and stored at index '14' (where '14' changes for each attribute)
* `attribute15` - Name of the custom attribute to be extracted from server and stored at index '15' (where '15' changes for each attribute)
* `attribute16` - Name of the custom attribute to be extracted from server and stored at index '16' (where '16' changes for each attribute)
* `attribute2` - Name of the custom attribute to be extracted from server and stored at index '2' (where '2' changes for each attribute)
* `attribute3` - Name of the custom attribute to be extracted from server and stored at index '3' (where '3' changes for each attribute)
* `attribute4` - Name of the custom attribute to be extracted from server and stored at index '4' (where '4' changes for each attribute)
* `attribute5` - Name of the custom attribute to be extracted from server and stored at index '5' (where '5' changes for each attribute)
* `attribute6` - Name of the custom attribute to be extracted from server and stored at index '6' (where '6' changes for each attribute)
* `attribute7` - Name of the custom attribute to be extracted from server and stored at index '7' (where '7' changes for each attribute)
* `attribute8` - Name of the custom attribute to be extracted from server and stored at index '8' (where '8' changes for each attribute)
* `attribute9` - Name of the custom attribute to be extracted from server and stored at index '9' (where '9' changes for each attribute)
* `attributes` - List of attribute names separated by ',' which needs to be fetched from tacacs server. Note that preceeding and trailing spaces will be removed. Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
* `auditfailedcmds` - The state of the TACACS+ server that will receive accounting messages.
* `authorization` - Use streaming authorization on the TACACS+ server.
* `authtimeout` - Number of seconds the Citrix ADC waits for a response from the TACACS+ server.
* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupattrname` - TACACS+ group attribute name. Used for group extraction on the TACACS+ server.
* `serverip` - IP address assigned to the TACACS+ server.
* `serverport` - Port number on which the TACACS+ server listens for connections.
* `tacacssecret` - Key shared between the TACACS+ server and the Citrix ADC. Required for allowing the Citrix ADC to communicate with the TACACS+ server.


## Import

A authenticationtacacsaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationtacacsaction.tf_tacacsaction my_tacacsaction
```
