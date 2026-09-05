---
subcategory: "DNS"
---

# Data Source: dnssvcbrec

The dnssvcbrec data source allows you to retrieve information about DNS SVCB/HTTPS service binding records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnssvcbrec" "tf_dnssvcbrec" {
  domain     = "svcb.example.com"
  targetname = "target.example.com"
}

output "priority" {
  value = data.citrixadc_dnssvcbrec.tf_dnssvcbrec.priority
}

output "svcbtype" {
  value = data.citrixadc_dnssvcbrec.tf_dnssvcbrec.svcbtype
}

output "port" {
  value = data.citrixadc_dnssvcbrec.tf_dnssvcbrec.port
}

output "ttl" {
  value = data.citrixadc_dnssvcbrec.tf_dnssvcbrec.ttl
}
```


## Argument Reference

* `domain` - (Required) Domain name for the SVCB/HTTPS record.
* `targetname` - (Required) Target domain name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssvcbrec. It is a combination of `domain`, `targetname`, `priority` and `svcbtype`.
* `priority` - Service priority (0 for AliasMode, >0 for ServiceMode).
* `svcbtype` - Service type: SVCB or HTTPS.
* `alpn` - Comma-separated list of ALPN protocol identifiers.
* `encryptedclienthello` - Base64-encoded ECH configuration.
* `ipv4hint` - Comma-separated list of IPv4 hint addresses.
* `ipv6hint` - Comma-separated list of IPv6 hint addresses.
* `mandatory` - Comma-separated list of mandatory SvcParam keys.
* `nodefaultalpn` - Indicates no default ALPN protocols.
* `nodeid` - Unique number that identifies the cluster node.
* `port` - Port number for the service.
* `ttl` - Time to Live (TTL) in seconds.
