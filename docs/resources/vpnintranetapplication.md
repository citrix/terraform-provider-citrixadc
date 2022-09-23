---
subcategory: "VPN"
---

# Resource: vpnintranetapplication

The vpnintranetapplication resource is used to create vpn intranet application resource.


## Example usage

```hcl
resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "TCP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
```


## Argument Reference

* `intranetapplication` - (Required) Name of the intranet application.
* `clientapplication` - (Optional) Names of the client applications, such as PuTTY and Xshell.
* `destip` - (Optional) Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.
* `destport` - (Optional) Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.
* `hostname` - (Optional) Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.
* `interception` - (Optional) Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.
* `iprange` - (Optional) If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.
* `netmask` - (Optional) Destination subnet mask for the intranet application.
* `protocol` - (Optional) Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.
* `spoofiip` - (Optional) IP address that the intranet application will use to route the connection through the virtual adapter.
* `srcip` - (Optional) Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.
* `srcport` - (Optional) Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnintranetapplication. It has the same value as the `intranetapplication` attribute.


## Import

A vpnintranetapplication can be imported using its intranetapplication, e.g.

```shell
terraform import citrixadc_vpnintranetapplication.tf_vpnintranetapplication tf_vpnintranetapplication
```
