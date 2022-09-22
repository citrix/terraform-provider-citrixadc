---
subcategory: "VPN"
---

# Resource: vpnpcoipprofile

The vpnpcoipprofile resource is used to create vpn PCoIP profile resource.


## Example usage

```hcl
resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
  name               = "tf_vpnpcoipprofile"
  conserverurl       = "http://www.example.com"
  sessionidletimeout = 80
}

```


## Argument Reference

* `name` - (Required) name of PCoIP profile
* `conserverurl` - (Required) Connection server URL
* `icvverification` - (Optional) ICV verification for PCOIP transport packets.
* `sessionidletimeout` - (Optional) PCOIP Idle Session timeout


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnpcoipprofile. It has the same value as the `name` attribute.


## Import

A vpnpcoipprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnpcoipprofile.tf_vpnpcoipprofile tf_vpnpcoipprofile
```
