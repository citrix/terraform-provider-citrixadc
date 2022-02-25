---
subcategory: "Vpn"
---

# Resource: vpnvserver_vpnintranetapplication_binding

The vpnvserver_vpnintranetapplication_binding resource is used to bind vpnintranetapplication to vpnvserver Resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
resource "citrixadc_vpnvserver_vpnintranetapplication_binding" "tf_bind" {
  name                = citrixadc_vpnvserver.tf_vpnvserver.name
  intranetapplication = citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `intranetapplication` - (Required) The intranet VPN application.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnintranetapplication_binding. It is the concatenation of `name` and `intranetapplication` attributes seperated by comma.


## Import

A vpnvserver_vpnintranetapplication_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnintranetapplication_binding.tf_bind tf_examplevserver,tf_vpnintranetapplication
```
