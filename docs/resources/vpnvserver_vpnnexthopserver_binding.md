---
subcategory: "Vpn"
---

# Resource: vpnvserver_vpnnexthopserver_binding

The vpnvserver_vpnnexthopserver_binding resource is used to bind vpnnexthopserver to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_exampleserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name        = "tf_vpnnexthopserver"
  nexthopip   = "2.2.1.5"
  nexthopport = "200"
}
resource "citrixadc_vpnvserver_vpnnexthopserver_binding" "tf_bind" {
  name          = citrixadc_vpnvserver.tf_vpnvserver.name
  nexthopserver = citrixadc_vpnnexthopserver.tf_vpnnexthopserver.name
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `nexthopserver` - (Required) The name of the next hop server bound to the VPN virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnnexthopserver_binding. It is the concatenation of `name` and `nexthopserver` attributes seperated by comma.


## Import

A vpnvserver_vpnnexthopserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnnexthopserver_binding.tf_bind tf_exampleserver,tf_vpnnexthopserver
```
