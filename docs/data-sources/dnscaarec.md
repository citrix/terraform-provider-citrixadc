---
subcategory: "DNS"
---

# Data Source: dnscaarec

The dnscaarec data source allows you to retrieve information about a DNS Certification Authority Authorization (CAA) resource record.


## Example usage

```terraform
data "citrixadc_dnscaarec" "tf_dnscaarec" {
  domain   = "example.com"
  recordid = 12345
}

output "tag" {
  value = data.citrixadc_dnscaarec.tf_dnscaarec.tag
}

output "valuestring" {
  value = data.citrixadc_dnscaarec.tf_dnscaarec.valuestring
}

output "ttl" {
  value = data.citrixadc_dnscaarec.tf_dnscaarec.ttl
}
```


## Argument Reference

* `domain` - (Required) Domain name of the CAA record.
* `recordid` - (Required) Unique, internally generated record ID. View the details of the CAA record to obtain its record ID.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnscaarec. It is a composite of the `domain` and `recordid` attributes, formatted as `domain:<domain>,recordid:<recordid>`.
* `tag` - String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - `issue`, `issuewild` and `iodef`.
* `flag` - Flag associated with the CAA record.
* `valuestring` - Value associated with the chosen property tag in the CAA resource record.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `ecssubnet` - Subnet for which the cached CAA record needs to be removed.
