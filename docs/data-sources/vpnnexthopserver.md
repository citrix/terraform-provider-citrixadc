---
subcategory: "VPN"
---

# Data Source `vpnnexthopserver`

The vpnnexthopserver data source allows you to retrieve information about a Citrix Gateway next hop server configuration used in double-hop DMZ deployments.


## Example usage

```terraform
data "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name = "tf_vpnnexthopserver"
}

output "nexthopip" {
  value = data.citrixadc_vpnnexthopserver.tf_vpnnexthopserver.nexthopip
}

output "nexthopport" {
  value = data.citrixadc_vpnnexthopserver.tf_vpnnexthopserver.nexthopport
}
```


## Argument Reference

* `name` - (Required) Name for the Citrix Gateway appliance in the first DMZ.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `nexthopfqdn` - FQDN of the Citrix Gateway proxy in the second DMZ.
* `nexthopip` - IP address of the Citrix Gateway proxy in the second DMZ.
* `nexthopport` - Port number of the Citrix Gateway proxy in the second DMZ.
* `resaddresstype` - Address Type (IPV4/IPv6) of DNS name of nextHopServer FQDN.
* `secure` - Use of a secure port, such as 443, for the double-hop configuration.
* `id` - The id of the vpnnexthopserver. It has the same value as the `name` attribute.


## Import

A vpnnexthopserver can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnnexthopserver.tf_vpnnexthopserver tf_vpnnexthopserver
```
