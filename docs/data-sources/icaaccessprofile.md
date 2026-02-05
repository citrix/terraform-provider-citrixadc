---
subcategory: "ICA"
---

# Data Source `icaaccessprofile`

The icaaccessprofile data source allows you to retrieve information about ICA access profiles for controlling ICA connections and features.


## Example usage

```terraform
data "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
  name = "my_ica_accessprofile"
}

output "connectclientlptports" {
  value = data.citrixadc_icaaccessprofile.tf_icaaccessprofile.connectclientlptports
}

output "wiaredirection" {
  value = data.citrixadc_icaaccessprofile.tf_icaaccessprofile.wiaredirection
}
```


## Argument Reference

* `name` - (Required) Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `clientaudioredirection` - Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input.
* `clientclipboardredirection` - Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server.
* `clientcomportredirection` - Allow Default access/Disable COM port redirection to and from the client.
* `clientdriveredirection` - Allow Default access/Disables drive redirection to and from the client.
* `clientprinterredirection` - Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session.
* `clienttwaindeviceredirection` - Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications.
* `clientusbdriveredirection` - Allow Default access/Disable the redirection of USB devices to and from the client.
* `connectclientlptports` - Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on.
* `draganddrop` - Allow default access or disable drag and drop between client and remote applications and desktops.
* `fido2redirection` - Allow default access or disable FIDO2 redirection.
* `localremotedatasharing` - Allow Default access/Disable file/data sharing via the Receiver for HTML5.
* `multistream` - Allow Default access/Disable the multistream feature for the specified users.
* `smartcardredirection` - Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD.
* `wiaredirection` - Allow default access or disable WIA scanner redirection.

## Attribute Reference

* `id` - The id of the icaaccessprofile. It has the same value as the `name` attribute.


## Import

An icaaccessprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_icaaccessprofile.tf_icaaccessprofile my_ica_accessprofile
```
