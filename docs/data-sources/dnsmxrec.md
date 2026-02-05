---
subcategory: "DNS"
---

# Data Source `dnsmxrec`

The dnsmxrec data source allows you to retrieve information about DNS MX (Mail Exchange) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsmxrec" "tf_dnsmxrec" {
  domain = "example.com"
  type   = "ALL"
}

output "mx" {
  value = data.citrixadc_dnsmxrec.tf_dnsmxrec.mx
}

output "pref" {
  value = data.citrixadc_dnsmxrec.tf_dnsmxrec.pref
}

output "ttl" {
  value = data.citrixadc_dnsmxrec.tf_dnsmxrec.ttl
}
```


## Argument Reference

* `domain` - (Required) Domain name for which to retrieve the MX record.
* `type` - (Required) Type of records to display. Available settings function as follows:
  * `ADNS` - Display all authoritative address records.
  * `PROXY` - Display all proxy address records.
  * `ALL` - Display all address records.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsmxrec. It is a combination of domain and type.
* `ecssubnet` - Subnet for which the cached MX record need to be removed.
* `mx` - Host name of the mail exchange server.
* `nodeid` - Unique number that identifies the cluster node.
* `pref` - Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
