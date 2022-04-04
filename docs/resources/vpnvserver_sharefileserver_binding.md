---
subcategory: "Vpn"
---

# Resource: vpnvserver_sharefileserver_binding

The vpnvserver_sharefileserver_binding resource is used to bind sharefileserver to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_sharefileserver_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  sharefile = "3.3.4.3:90"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `sharefile` - (Required) Configured ShareFile server in XenMobile deployment. Format IP:PORT / FQDN:PORT


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_sharefileserver_binding. It is the concatenation of `name` and `sharefile` attributes seperated by comma.


## Import

A vpnvserver_sharefileserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_sharefileserver_binding.tf_bind tf_vpnvserver,3.3.4.3:90
```
