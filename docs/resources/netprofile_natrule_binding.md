---
subcategory: "Network"
---

# Resource: netprofile_natrule_binding

The netprofile_natrule_binding resource is used to bind natrule to netprofile.


## Example usage

```hcl
resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_netprofile_natrule_binding" "tf_binding" {
  name      = citrixadc_netprofile.tf_netprofile.name
  natrule   = "10.10.10.10"
  netmask   = "255.255.255.255"
  rewriteip = "3.3.3.3"
}
```


## Argument Reference

* `name` - (Required) Name of the netprofile to which to bind port ranges.
* `natrule` - (Required) IPv4 network address on whose traffic you want the Citrix ADC to do rewrite ip prefix.
* `netmask` - (Optional) 0
* `rewriteip` - (Optional) 0


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netprofile_natrule_binding. Itis the concatenation of `name` and `natrule` attributes separated by comma.


## Import

A netprofile_natrule_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_netprofile_natrule_binding.tf_binding tf_netprofile,10.10.10.10
```
