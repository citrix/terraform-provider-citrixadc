---
subcategory: "VPN"
---

# Resource: vpnnexthopserver

The vpnnexthopserver resource is used to create Next Hop Server resource.


## Example usage

```hcl
resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name        = "tf_vpnnexthopserver"
  nexthopip   = "2.6.1.5"
  nexthopport = "200"
}
```


## Argument Reference

* `name` - (Required) Name for the Citrix Gateway appliance in the first DMZ.
* `nexthopport` - (Required) Port number of the Citrix Gateway proxy in the second DMZ.
* `nexthopfqdn` - (Optional) FQDN of the Citrix Gateway proxy in the second DMZ.
* `nexthopip` - (Optional) IP address of the Citrix Gateway proxy in the second DMZ.
* `resaddresstype` - (Optional) Address Type (IPV4/IPv6) of DNS name of nextHopServer FQDN.
* `secure` - (Optional) Use of a secure port, such as 443, for the double-hop configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnnexthopserver. It has the same value as the `name` attribute.


## Import

A vpnnexthopserver can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnnexthopserver.tf_vpnnexthopserver tf_vpnnexthopserver
```
