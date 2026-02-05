---
page_title: "citrixadc_dnsnsrec Data Source - terraform-provider-citrixadc"
subcategory: "DNS"
description: |-
  Data source to retrieve DNS NS record information from Citrix ADC.
---

# citrixadc_dnsnsrec (Data Source)

This data source is used to retrieve DNS NS (Name Server) record information from Citrix ADC.

## Example Usage

```terraform
# Query DNS NS record by domain and type
data "citrixadc_dnsnsrec" "example" {
  domain = "example.com"
}

# Reference the data source
output "nameserver" {
  value = data.citrixadc_dnsnsrec.example.nameserver
}

output "ttl" {
  value = data.citrixadc_dnsnsrec.example.ttl
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) Domain name for which to retrieve NS records.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the DNS NS record resource.
* `domain` - Domain name.
* `nameserver` - Host name of the name server for the domain.
* `ttl` - Time to Live (TTL), in seconds, for the record. The TTL is the time for which the record must be cached by DNS proxies.
* `ecssubnet` - Subnet for which the cached name server record applies.
* `nodeid` - Unique number that identifies the cluster node.
