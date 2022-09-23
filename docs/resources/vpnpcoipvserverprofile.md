---
subcategory: "VPN"
---

# Resource: vpnpcoipvserverprofile

The vpnpcoipvserverprofile resource is used to create PCoIP vserver profile resource.


## Example usage

```hcl
resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
  name        = "tf_vpnpcoipvserverprofile"
  logindomain = "domainname"
  udpport     = "802"
}
```


## Argument Reference

* `name` - (Required) name of PCoIP vserver profile
* `logindomain` - (Required) Login domain for PCoIP users
* `udpport` - (Optional) UDP port for PCoIP data traffic


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnpcoipvserverprofile. It has the same value as the `name` attribute.


## Import

A vpnpcoipvserverprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnpcoipvserverprofile.tf_vpnpcoipvserverprofile tf_vpnpcoipvserverprofile
```
