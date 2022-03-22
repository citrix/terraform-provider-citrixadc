---
subcategory: "Authentication"
---

# Resource: authenticationtacacsaction

The authenticationtacacsaction resource is used to create authentication tacacsaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
  name            = "tf_tacacsaction"
  serverip        = "1.2.3.4"
  serverport      = 8080
  authtimeout     = 5
  authorization   = "ON"
  accounting      = "ON"
  auditfailedcmds = "ON"
  groupattrname   = "group"
}
```


## Argument Reference

* `name` - (Required) Name for the TACACS+ profile (action).  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after TACACS profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'y authentication action').
* `accounting` - (Optional) Whether the TACACS+ server is currently accepting accounting messages.
* `attribute1` - (Optional) Name of the custom attribute to be extracted from server and stored at index '1' (where '1' changes for each attribute)
* `attribute10` - (Optional) Name of the custom attribute to be extracted from server and stored at index '10' (where '10' changes for each attribute)
* `attribute11` - (Optional) Name of the custom attribute to be extracted from server and stored at index '11' (where '11' changes for each attribute)
* `attribute12` - (Optional) Name of the custom attribute to be extracted from server and stored at index '12' (where '12' changes for each attribute)
* `attribute13` - (Optional) Name of the custom attribute to be extracted from server and stored at index '13' (where '13' changes for each attribute)
* `attribute14` - (Optional) Name of the custom attribute to be extracted from server and stored at index '14' (where '14' changes for each attribute)
* `attribute15` - (Optional) Name of the custom attribute to be extracted from server and stored at index '15' (where '15' changes for each attribute)
* `attribute16` - (Optional) Name of the custom attribute to be extracted from server and stored at index '16' (where '16' changes for each attribute)
* `attribute2` - (Optional) Name of the custom attribute to be extracted from server and stored at index '2' (where '2' changes for each attribute)
* `attribute3` - (Optional) Name of the custom attribute to be extracted from server and stored at index '3' (where '3' changes for each attribute)
* `attribute4` - (Optional) Name of the custom attribute to be extracted from server and stored at index '4' (where '4' changes for each attribute)
* `attribute5` - (Optional) Name of the custom attribute to be extracted from server and stored at index '5' (where '5' changes for each attribute)
* `attribute6` - (Optional) Name of the custom attribute to be extracted from server and stored at index '6' (where '6' changes for each attribute)
* `attribute7` - (Optional) Name of the custom attribute to be extracted from server and stored at index '7' (where '7' changes for each attribute)
* `attribute8` - (Optional) Name of the custom attribute to be extracted from server and stored at index '8' (where '8' changes for each attribute)
* `attribute9` - (Optional) Name of the custom attribute to be extracted from server and stored at index '9' (where '9' changes for each attribute)
* `attributes` - (Optional) List of attribute names separated by ',' which needs to be fetched from tacacs server.  Note that preceeding and trailing spaces will be removed.  Attribute name can be 127 bytes and total length of this string should not cross 2047 bytes. These attributes have multi-value support separated by ',' and stored as key-value pair in AAA session
* `auditfailedcmds` - (Optional) The state of the TACACS+ server that will receive accounting messages.
* `authorization` - (Optional) Use streaming authorization on the TACACS+ server.
* `authtimeout` - (Optional) Number of seconds the Citrix ADC waits for a response from the TACACS+ server.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupattrname` - (Optional) TACACS+ group attribute name. Used for group extraction on the TACACS+ server.
* `serverip` - (Optional) IP address assigned to the TACACS+ server.
* `serverport` - (Optional) Port number on which the TACACS+ server listens for connections.
* `tacacssecret` - (Optional) Key shared between the TACACS+ server and the Citrix ADC.  Required for allowing the Citrix ADC to communicate with the TACACS+ server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationtacacsaction. It has the same value as the `name` attribute.


## Import

A authenticationtacacsaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationtacacsaction.tf_tacacsaction tf_tacacsaction
```
