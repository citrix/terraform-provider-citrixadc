---
subcategory: "VPN"
---

# Resource: vpnvserver_vpneula_binding

The vpnvserver_vpneula_binding resource is used to bind vpneula to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpneula" "tf_vpneula" {
  name = "tf_vpneula"
}
resource "citrixadc_vpnvserver_vpneula_binding" "tf_bind" {
  name = citrixadc_vpnvserver.tf_vpnvserver.name
  eula = citrixadc_vpneula.tf_vpneula.name
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `eula` - (Required) Name of the EULA bound to VPN vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpneula_binding. It is the concatenation  of `name` and `eula` attributes serprated by comma.


## Import

A vpnvserver_vpneula_binding can be imported using its id , e.g.

```shell
terraform import citrixadc_vpnvserver_vpneula_binding.tf_bind tf_examplevserver,tf_vpneula
```
