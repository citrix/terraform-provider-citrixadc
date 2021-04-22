---
subcategory: "DNS"
---

# Resource: dnssoarec

The dnssoarec resource is used to create dns SOA records.


## Example usage

```hcl
resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "hello.com"
	originserver  = "10.2.3.5"
	contact =  "other"
	expire = 1800
	refresh = 4000
}
```


## Argument Reference

* `domain` - (Required) Domain name for which to add the SOA record.
* `originserver` - (Optional) Domain name of the name server that responds authoritatively for the domain.
* `contact` - (Optional) Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.
* `serial` - (Optional) The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.
* `refresh` - (Optional) Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.
* `retry` - (Optional) Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.
* `expire` - (Optional) Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.
* `minimum` - (Optional) Default time to live (TTL) for all records in the zone. Can be overridden for individual records.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `ecssubnet` - (Optional) Subnet for which the cached SOA record need to be removed.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records. Possible values: [ ALL, ADNS, PROXY ]
* `nodeid` - (Optional) Unique number that identifies the cluster node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssoarec. It has the same value as the `domain` attribute.


## Import

A dnssoarec can be imported using its domain, e.g.

```shell
terraform import citrixadc_dnssoarec.tf_dnssoarec hello.com
```
