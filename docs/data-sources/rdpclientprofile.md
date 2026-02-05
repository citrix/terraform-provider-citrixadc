---
subcategory: "RDP"
---

# Data Source: citrixadc_rdpclientprofile

The `citrixadc_rdpclientprofile` data source is used to retrieve information about a specific RDP client profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_rdpclientprofile" "example" {
  name = "my_rdpclientprofile"
}
```

## Argument Reference

* `name` - (Required) The name of the RDP client profile.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the RDP client profile.
* `addusernameinrdpfile` - Add username in RDP file.
* `audiocapturemode` - This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.
* `keyboardhook` - This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.
* `multimonitorsupport` - Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).
* `psk` - Pre shared key value.
* `randomizerdpfilename` - Will generate unique filename everytime rdp file is downloaded.
* `rdpcookievalidity` - RDP cookie validity period.
* `rdpcustomparams` - Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'.
* `rdpfilename` - RDP file name to be sent to End User.
* `rdphost` - Fully-qualified domain name (FQDN) of the RDP Listener.
* `rdplinkattribute` - Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers.
* `rdplistener` - IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port.
* `rdpurloverride` - This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.
* `rdpvalidateclientip` - This setting determines whether RDC launch is initiated by the valid client IP.
* `redirectclipboard` - This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.
* `redirectcomports` - This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.
* `redirectdrives` - This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.
* `redirectpnpdevices` - This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.
* `redirectprinters` - This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.
* `videoplaybackmode` - This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.
