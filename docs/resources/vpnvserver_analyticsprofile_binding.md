---
subcategory: "Vpn"
---

# Resource: vpnvserver_analyticsprofile_binding

The vpnvserver_analyticsprofile_binding resource is used to bind analyticsprofile to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name           = "tf_vserver"
  servicetype    = "SSL"
  ipv46          = "3.3.3.3"
  port           = 443
}
resource "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
  name             = citrixadc_vpnvserver.tf_vpnvserver.name
  analyticsprofile = "new_profile"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the VPN Vserver


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_analyticsprofile_binding. It is the concatenation of the `name` and `analyticsprofile` attributes separated by a comma.


## Import

A vpnvserver_analyticsprofile_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_analyticsprofile_binding.tf_bind tf_vserver,new_profile
```
