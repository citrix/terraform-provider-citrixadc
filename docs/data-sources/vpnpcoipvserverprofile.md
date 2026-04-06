---
subcategory: "VPN"
---

# Data Source: vpnpcoipvserverprofile

The vpnpcoipvserverprofile data source allows you to retrieve information about a VPN PCoIP vserver profile.

## Example usage

```terraform
data "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
  name = "tf_vpnpcoipvserverprofile"
}

output "logindomain" {
  value = data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile.logindomain
}

output "udpport" {
  value = data.citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile.udpport
}
```

## Argument Reference

* `name` - (Required) Name of PCoIP vserver profile

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnpcoipvserverprofile. It is the same as the `name` attribute.
* `logindomain` - Login domain for PCoIP users
* `udpport` - UDP port for PCoIP data traffic
