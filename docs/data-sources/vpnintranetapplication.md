---
subcategory: "VPN"
---

# Data Source `vpnintranetapplication`

The vpnintranetapplication data source allows you to retrieve information about an intranet application configuration.


## Example usage

```terraform
data "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
}

output "protocol" {
  value = data.citrixadc_vpnintranetapplication.tf_vpnintranetapplication.protocol
}

output "destip" {
  value = data.citrixadc_vpnintranetapplication.tf_vpnintranetapplication.destip
}
```


## Argument Reference

* `intranetapplication` - (Required) Name of the intranet application.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `clientapplication` - Names of the client applications, such as PuTTY and Xshell.
* `destip` - Destination IP address, IP range, or host name of the intranet application. This address is the server IP address.
* `destport` - Destination TCP or UDP port number for the intranet application. Use a hyphen to specify a range of port numbers, for example 90-95.
* `hostname` - Name of the host for which to configure interception. The names are resolved during interception when users log on with the Citrix Gateway Plug-in.
* `interception` - Interception mode for the intranet application or resource. Correct value depends on the type of client software used to make connections. If the interception mode is set to TRANSPARENT, users connect with the Citrix Gateway Plug-in for Windows. With the PROXY setting, users connect with the Citrix Gateway Plug-in for Java.
* `iprange` - If you have multiple servers in your network, such as web, email, and file shares, configure an intranet application that includes the IP range for all the network applications. This allows users to access all the intranet applications contained in the IP address range.
* `netmask` - Destination subnet mask for the intranet application.
* `protocol` - Protocol used by the intranet application. If protocol is set to BOTH, TCP and UDP traffic is allowed.
* `spoofiip` - IP address that the intranet application will use to route the connection through the virtual adapter.
* `srcip` - Source IP address. Required if interception mode is set to PROXY. Default is the loopback address, 127.0.0.1.
* `srcport` - Source port for the application for which the Citrix Gateway virtual server proxies the traffic. If users are connecting from a device that uses the Citrix Gateway Plug-in for Java, applications must be configured manually by using the source IP address and TCP port values specified in the intranet application profile. If a port value is not set, the destination port value is used.
* `id` - The id of the vpnintranetapplication. It has the same value as the `intranetapplication` attribute.


## Import

A vpnintranetapplication can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnintranetapplication.tf_vpnintranetapplication tf_vpnintranetapplication
```
