---
subcategory: "VPN"
---

# Resource: vpnglobal_vpnnexthopserver_binding

The vpnglobal_vpnnexthopserver_binding resource is used to bind vpnnexthopserver to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnnexthopserver" "tf_vpnnexthopserver" {
  name        = "tf_vpnnexthopserver"
  nexthopip   = "2.6.1.5"
  nexthopport = "200"
}
resource "citrixadc_vpnglobal_vpnnexthopserver_binding" "tf_bind" {
  nexthopserver = citrixadc_vpnnexthopserver.tf_vpnnexthopserver.name
}
```


## Argument Reference

* `nexthopserver` - (Required) The name of the next hop server bound to vpn global.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnnexthopserver_binding. It has the same value as the `nexthopserver` attribute.


## Import

A vpnglobal_vpnnexthopserver_binding can be imported using its nexthopserver, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnnexthopserver_binding.tf_bind tf_vpnnexthopserver
```
