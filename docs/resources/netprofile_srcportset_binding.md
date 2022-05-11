---
subcategory: "Network"
---

# Resource: netprofile_srcportset_binding

The netprofile_srcportset_binding resource is used to bind srcportset to netprofile.


## Example usage

```hcl
resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_netprofile_srcportset_binding" "tf_binding" {
  name         = citrixadc_netprofile.tf_netprofile.name
  srcportrange = "2000"
}
```


## Argument Reference

* `name` - (Required) Name of the netprofile to which to bind port ranges.
* `srcportrange` - (Required) When the source port range is configured and associated with the netprofile bound to a service group, Citrix ADC will choose a port from the range configured for connection establishment at the backend servers.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netprofile_srcportset_binding. It is the concatenation of `name` and `srcportrange` attributes separated by comma.


## Import

A netprofile_srcportset_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_netprofile_srcportset_binding.tf_binding tf_netprofile,2000
```
