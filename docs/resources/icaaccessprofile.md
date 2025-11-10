---
subcategory: "Ica"
---

# Resource: icaaccessprofile

The icaaccessprofile resource is used to create icaaccessprofile.


## Example usage

```hcl
resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
  name                   = "my_ica_accessprofile"
  connectclientlptports  = "DEFAULT"
  localremotedatasharing = "DEFAULT"
}

```


## Argument Reference

* `name` - (Required) Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA accessprofile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica accessprofile" or 'my ica accessprofile'). Each of the features can be configured as DEFAULT/DISABLED. Here, DISABLED means that the policy settings on the backend XenApp/XenDesktop server are overridden and the Citrix ADC makes the decision to deny access. Whereas DEFAULT means that the Citrix ADC allows the request to reach the XenApp/XenDesktop that takes the decision to allow/deny access based on the policy configured on it. For example, if ClientAudioRedirection is enabled on the backend XenApp/XenDesktop server, and the configured profile has ClientAudioRedirection as DISABLED, the Citrix ADC makes the decision to deny the request irrespective of the configuration on the backend. If the configured profile has ClientAudioRedirection as DEFAULT, then the Citrix ADC forwards the requests to the backend XenApp/XenDesktop server.It then makes the decision to allow/deny access based on the policy configured on it. Minimum length =  1
* `connectclientlptports` - (Optional) Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on. Possible values: [ DEFAULT, DISABLED ]
* `clientaudioredirection` - (Optional) Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input. Possible values: [ DEFAULT, DISABLED ]
* `localremotedatasharing` - (Optional) Allow Default access/Disable file/data sharing via the Receiver for HTML5. Possible values: [ DEFAULT, DISABLED ]
* `clientclipboardredirection` - (Optional) Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server. Possible values: [ DEFAULT, DISABLED ]
* `clientcomportredirection` - (Optional) Allow Default access/Disable COM port redirection to and from the client. Possible values: [ DEFAULT, DISABLED ]
* `clientdriveredirection` - (Optional) Allow Default access/Disables drive redirection to and from the client. Possible values: [ DEFAULT, DISABLED ]
* `clientprinterredirection` - (Optional) Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session. Possible values: [ DEFAULT, DISABLED ]
* `multistream` - (Optional) Allow Default access/Disable the multistream feature for the specified users. Possible values: [ DEFAULT, DISABLED ]
* `clientusbdriveredirection` - (Optional) Allow Default access/Disable the redirection of USB devices to and from the client. Possible values: [ DEFAULT, DISABLED ]
* `clienttwaindeviceredirection` - (Optional) Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications
* `draganddrop` - (Optional) Allow default access or disable drag and drop between client and remote applications and desktops
* `fido2redirection` - (Optional) Allow default access or disable FIDO2 redirection
* `smartcardredirection` - (Optional) Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD
* `wiaredirection` - (Optional) Allow default access or disable WIA scanner redirection


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icaaccessprofile. It has the same value as the `name` attribute.


## Import

A icaaccessprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_icaaccessprofile.tf_icaaccessprofile my_ica_accessprofile
```
