---
subcategory: "Integrated Caching"
---

# Data Source `cacheforwardproxy`

The cacheforwardproxy data source allows you to retrieve information about cache forward proxy configurations.


## Example usage

```terraform
data "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy" {
  ipaddress = "10.222.74.187"
}

output "port" {
  value = data.citrixadc_cacheforwardproxy.tf_cacheforwardproxy.port
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the Citrix ADC or a cache server for which the cache acts as a proxy. Requests coming to the Citrix ADC with the configured IP address are forwarded to the particular address, without involving the Integrated Cache in any way.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `port` - Port on the Citrix ADC or a server for which the cache acts as a proxy.
* `id` - The id of the cacheforwardproxy. It has the same value as the `ipaddress` attribute.
