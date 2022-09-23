---
subcategory: "VPN"
---

# Resource: vpnportaltheme

The vpnportaltheme resource is used to create vpn portal theme resource.


## Example usage

```hcl
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
```


## Argument Reference

* `name` - (Required) Name of the uitheme
* `basetheme` - (Required) 0


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnportaltheme. It has the same value as the `name` attribute.


## Import

A vpnportaltheme can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnportaltheme.tf_vpnportaltheme tf_vpnportaltheme
```
