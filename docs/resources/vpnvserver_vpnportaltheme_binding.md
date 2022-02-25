---
subcategory: "Vpn"
---

# Resource: vpnvserver_vpnportaltheme_binding

The vpnvserver_vpnportaltheme_binding resource is used to bind vpnportaltheme to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name           = "tf_exampleserver"
  servicetype    = "SSL"
  ipv46          = "3.3.3.3"
  port           = 443
}
resource "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name      = "tf_vpnportaltheme"
  basetheme = "X1"
}
resource "citrixadc_vpnvserver_vpnportaltheme_binding" "tf_bind" {
  name        = citrixadc_vpnvserver.tf_vpnvserver.name
  portaltheme = citrixadc_vpnportaltheme.tf_vpnportaltheme.name
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `portaltheme` - (Required) Name of the portal theme bound to VPN vserver


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnportaltheme_binding. It is the concatenation `name` and `portaltheme` attributes seperated by comma.


## Import

A vpnvserver_vpnportaltheme_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnportaltheme_binding.tf_bind tf_exampleserver,tf_vpnportaltheme
```
