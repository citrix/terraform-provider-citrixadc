---
subcategory: "Vpn"
---

# Resource: vpnvserver_appcontroller_binding

The vpnvserver_appcontroller_binding resource is used to bind appcontroller to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_newvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_appcontroller_binding" "tf_bind" {
  name          = citrixadc_vpnvserver.tf_vpnvserver.name
  appcontroller = "http://www.example.com"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `appcontroller` - (Required) Configured App Controller server in XenMobile deployment.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_appcontroller_binding. It is the concatenation of the `name` and `appcontroller` attributes separated by a comma.


## Import

A vpnvserver_appcontroller_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_appcontroller_binding.tf_bind tf_newvserver,http://www.example.com
```
