---
subcategory: "Vpn"
---

# Resource: vpnvserver_vpnurl_binding

The vpnvserver_vpnurl_binding resource is used to bind vpnurl to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnurl" "url" {
  urlname          = "Firsturl"
  actualurl        = "www.citrix.com"
  appjson          = "xyz"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  linkname         = "Description"
  ssotype          = "unifiedgateway"
  vservername      = "server1"
}
resource "citrixadc_vpnvserver_vpnurl_binding" "tf_bind" {
  name    = citrixadc_vpnvserver.tf_vpnvserver.name
  urlname = citrixadc_vpnurl.url.urlname
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `urlname` - (Required) The intranet URL.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnurl_binding. It is the concatenation of `name` and `urlname` attribute seperated by comma.


## Import

A vpnvserver_vpnurl_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnurl_binding.tf_bind tf_examplevserver,Firsturl
```