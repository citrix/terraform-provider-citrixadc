---
subcategory: "VPN"
---

# Resource: vpnglobal_vpnurl_binding

The vpnglobal_vpnurl_binding resource is used to bind vpnurl to global configuraton.


## Example usage

```hcl
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
resource "citrixadc_vpnglobal_vpnurl_binding" "tf_bind" {
  urlname = citrixadc_vpnurl.url.urlname
}
```


## Argument Reference

* `urlname` - (Required) The intranet url.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnurl_binding. It has the same value as the `urlname` attribute.


## Import

A vpnglobal_vpnurl_binding can be imported using its urlname, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnurl_binding.tf_bind Firsturl
```
