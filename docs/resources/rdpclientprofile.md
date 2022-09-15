---
subcategory: "Rdp"
---

# Resource: rdpclientprofile

The rdpclientprofile resource is used to create rdpclientprofile.


## Example usage

```hcl
resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
  name              = "my_rdpclientprofile"
  rdpurloverride    = "ENABLE"
  redirectclipboard = "ENABLE"
  redirectdrives    = "ENABLE"
}
```


## Argument Reference

* `name` - (Required) The name of the rdp profile
* `addusernameinrdpfile` - (Optional) Add username in rdp file.
* `audiocapturemode` - (Optional) This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.
* `keyboardhook` - (Optional) This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.
* `multimonitorsupport` - (Optional) Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).
* `psk` - (Optional) Pre shared key value
* `randomizerdpfilename` - (Optional) Will generate unique filename everytime rdp file is downloaded by appending output of time() function in the format <rdpfileName>_<time>.rdp. This tries to avoid the pop-up for replacement of existing rdp file during each rdp connection launch, hence providing better end-user experience.
* `rdpcookievalidity` - (Optional) RDP cookie validity period. RDP cookie validity time is applicable for new connection and also for any re-connection that might happen, mostly due to network disruption or during fail-over.
* `rdpcustomparams` - (Optional) Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'
* `rdpfilename` - (Optional) RDP file name to be sent to End User
* `rdphost` - (Optional) Fully-qualified domain name (FQDN) of the RDP Listener.
* `rdplinkattribute` - (Optional) Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers(IP/FQDN) that a user can access, from an Authentication server attribute(Example: LDAP, SAML). Based on the list received, the RDP links will be generated and displayed to the user.             Note: The Attribute mentioned in the rdpLinkAttribute should be fetched through corresponding authentication method.
* `rdplistener` - (Optional) IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port
* `rdpurloverride` - (Optional) This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.
* `redirectclipboard` - (Optional) This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.
* `redirectcomports` - (Optional) This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.
* `redirectdrives` - (Optional) This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.
* `redirectpnpdevices` - (Optional) This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.
* `redirectprinters` - (Optional) This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.
* `videoplaybackmode` - (Optional) This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rdpclientprofile. It has the same value as the `name` attribute.


## Import

A rdpclientprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_rdpclientprofile.tf_rdpclientprofile my_rdpclientprofile
```
