---
subcategory: "VPN"
---

# Data Source: vpnpcoipprofile

The vpnpcoipprofile data source allows you to retrieve information about a VPN PCoIP profile.

## Example usage

```terraform
data "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
  name = "tf_vpnpcoipprofile"
}

output "conserverurl" {
  value = data.citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile.conserverurl
}

output "sessionidletimeout" {
  value = data.citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile.sessionidletimeout
}
```

## Argument Reference

* `name` - (Required) Name of PCoIP profile

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnpcoipprofile. It is the same as the `name` attribute.
* `conserverurl` - Connection server URL
* `icvverification` - ICV verification for PCOIP transport packets.
* `sessionidletimeout` - PCOIP Idle Session timeout
