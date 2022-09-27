---
subcategory: "Integrated Caching"
---

# Resource: cacheforwardproxy

The cacheforwardproxy resource is used to createcacheforwardproxy.


## Example usage

```hcl
resource "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy" {
  ipaddress  = "10.222.74.185"
  port        = 5000
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the Citrix ADC or a cache server for which the cache acts as a proxy. Requests coming to the Citrix ADC with the configured IP address are forwarded to the particular address, without involving the Integrated Cache in any way.
* `port` - (Required) Port on the Citrix ADC or a server for which the cache acts as a proxy


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheforwardproxy. It has the same value as the `ipaddress` attribute.


## Import

A cacheforwardproxy can be imported using its name, e.g.

```shell
terraform import citrixadc_cacheforwardproxy.tf_cacheforwardproxy 10.222.74.185
```
