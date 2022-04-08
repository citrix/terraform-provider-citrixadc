---
subcategory: "Vpn"
---

# Resource: vpnvserver_intranetip6_binding

The vpnvserver_intranetip6_binding resource is used to bind intranetip6 to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_intranetip6_binding" "tf_bind" {
  name        = citrixadc_vpnvserver.tf_vpnvserver.name
  intranetip6 = "2.3.4.5"
  numaddr     = "45"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `intranetip6` - (Required) The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.
* `numaddr` - (Optional) The number of ipv6 addresses


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_intranetip6_binding. It is concatenation of `name` and `intranetip6` attributes seperated by comma.


## Import

A vpnvserver_intranetip6_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_intranetip6_binding.tf_bind tf_vserverexample,2.3.4.5
```
