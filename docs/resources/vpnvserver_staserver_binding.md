---
subcategory: "VPN"
---

# Resource: vpnvserver_staserver_binding

The vpnvserver_staserver_binding resource is used to bind staserver to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_staserver_binding" "tf_binding" {
  name           = citrixadc_vpnvserver.tf_vpnvserver.name
  staserver      = "http://www.example.com/"
  staaddresstype = "IPV4"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `staserver` - (Required) Configured Secure Ticketing Authority (STA) server.
* `staaddresstype` - (Optional) Type of the STA server address(ipv4/v6).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_staserver_binding. It is the concatenation of `name` and `staserver` attributes seperated by comma.


## Import

A vpnvserver_staserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_staserver_binding.tf_binding tf_vserver,http://www.example.com/
```
