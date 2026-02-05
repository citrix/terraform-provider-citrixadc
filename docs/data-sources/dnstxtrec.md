---
subcategory: "DNS"
---

# Data Source `dnstxtrec`

The dnstxtrec data source allows you to retrieve information about DNS TXT (Text) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnstxtrec" "tf_dnstxtrec" {
  domain = "example.com"
}

output "domain" {
  value = data.citrixadc_dnstxtrec.tf_dnstxtrec.domain
}

output "string" {
  value = data.citrixadc_dnstxtrec.tf_dnstxtrec.string
}

output "ttl" {
  value = data.citrixadc_dnstxtrec.tf_dnstxtrec.ttl
}
```


## Argument Reference

* `domain` - (Required) Name of the domain for the TXT record.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnstxtrec. It is a unique identifier for the record.
* `string` - Information stored in the TXT resource record. A TXT resource record can contain up to six strings, each of which can contain up to 255 characters.
* `ecssubnet` - Subnet for which the cached TXT record need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `recordid` - Unique, internally generated record ID.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add a TXT record, with a TTL of 36000, to the domain name example.com, the TTLs of all the TXT records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
