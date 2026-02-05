---
subcategory: "DNS"
---

# Data Source `dnscnamerec`

The dnscnamerec data source allows you to retrieve information about DNS CNAME (Canonical Name) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnscnamerec" "tf_dnscnamerec" {
  aliasname = "www.example.com"
}

output "aliasname" {
  value = data.citrixadc_dnscnamerec.tf_dnscnamerec.aliasname
}

output "canonicalname" {
  value = data.citrixadc_dnscnamerec.tf_dnscnamerec.canonicalname
}

output "ttl" {
  value = data.citrixadc_dnscnamerec.tf_dnscnamerec.ttl
}
```


## Argument Reference

* `aliasname` - (Required) Alias for the canonical domain name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnscnamerec. It is a unique identifier for the record.
* `canonicalname` - Canonical domain name.
* `ecssubnet` - Subnet for which the cached CNAME record need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
