---
subcategory: "DNS"
---

# Data Source `dnssoarec`

The dnssoarec data source allows you to retrieve information about DNS SOA (Start of Authority) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnssoarec" "tf_dnssoarec" {
  domain = "test.com"
}

output "originserver" {
  value = data.citrixadc_dnssoarec.tf_dnssoarec.originserver
}

output "contact" {
  value = data.citrixadc_dnssoarec.tf_dnssoarec.contact
}

output "refresh" {
  value = data.citrixadc_dnssoarec.tf_dnssoarec.refresh
}
```


## Argument Reference

* `domain` - (Required) Domain name for which to retrieve the SOA record.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssoarec. It is based on the domain name.
* `contact` - Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.
* `ecssubnet` - Subnet for which the cached SOA record need to be removed.
* `expire` - Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.
* `minimum` - Default time to live (TTL) for all records in the zone. Can be overridden for individual records.
* `nodeid` - Unique number that identifies the cluster node.
* `originserver` - Domain name of the name server that responds authoritatively for the domain.
* `refresh` - Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.
* `retry` - Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.
* `serial` - The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
