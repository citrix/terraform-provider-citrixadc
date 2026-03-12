---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnportaltheme_binding

The vpnvserver_vpnportaltheme_binding data source allows you to retrieve information about a vpnportaltheme binding to a vpnvserver.


## Example Usage

```terraform
data "citrixadc_vpnvserver_vpnportaltheme_binding" "tf_bind" {
  name        = "tf_exampleserver"
  portaltheme = "tf_vpnportaltheme"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpnportaltheme_binding.tf_bind.name
}

output "portaltheme" {
  value = data.citrixadc_vpnvserver_vpnportaltheme_binding.tf_bind.portaltheme
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `portaltheme` - (Required) Name of the portal theme bound to VPN vserver


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnportaltheme_binding. It is the concatenation of `name` and `portaltheme` attributes separated by comma.
