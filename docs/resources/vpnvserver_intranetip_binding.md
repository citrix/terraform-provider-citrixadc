---
subcategory: "Vpn"
---

# Resource: vpnvserver_intranetip_binding

The vpnvserver_intranetip_binding resource is used to bind intranetip to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_intranetip_binding" "tf_bind" {
  name       = citrixadc_vpnvserver.tf_vpnvserver.name
  intranetip = "2.3.4.5"
  netmask    = "255.255.255.0"
}
```


## Argument Reference

* `intranetip` - (Required) The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.
* `name` - (Required) Name of the virtual server.
* `netmask` - (Optional) The netmask of the intranet IP address or range.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_intranetip_binding. It is the concatenation of `name` and `intranetip` attributes seperated by comma.


## Import

A vpnvserver_intranetip_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_intranetip_binding.tf_bind tf_vserverexample,2.3.4.5
```
