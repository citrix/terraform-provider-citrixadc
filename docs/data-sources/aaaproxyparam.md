---
subcategory: "AAA"
---

# Data Source: aaaproxyparam

The aaaproxyparam data source allows you to retrieve information about the AAA
proxy parameters configuration.


## Example usage

```terraform
data "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
}

output "proxy" {
  value = data.citrixadc_aaaproxyparam.tf_aaaproxyparam.proxy
}

output "proxyauthorization" {
  value = data.citrixadc_aaaproxyparam.tf_aaaproxyparam.proxyauthorization
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `proxy` - IP address and Port of the proxy server used for HTTP access. Configured in `ipaddress:port` format or as a URL.
* `proxyauthorization` - Indicates whether the Proxy-Authorization header will be sent. Possible values: [ disabled, basic ]
* `proxyusername` - Username that will be sent as part of the Basic Proxy-Authorization header.
* `proxypassword` - Password that will be sent as part of the Basic Proxy-Authorization header. (Sensitive; returned by the appliance only in encrypted form.)
* `id` - The id of the aaaproxyparam. It is a system-generated identifier.
