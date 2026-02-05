---
subcategory: "DNS"
---

# Data Source `dnsnaptrrec`

The dnsnaptrrec data source allows you to retrieve information about DNS NAPTR (Naming Authority Pointer) records.


## Example usage

```terraform
data "citrixadc_dnsnaptrrec" "tf_dnsnaptrrec" {
  domain = "example.com"
}

output "order" {
  value = data.citrixadc_dnsnaptrrec.tf_dnsnaptrrec.order
}

output "preference" {
  value = data.citrixadc_dnsnaptrrec.tf_dnsnaptrrec.preference
}

output "replacement" {
  value = data.citrixadc_dnsnaptrrec.tf_dnsnaptrrec.replacement
}
```


## Argument Reference

* `domain` - (Required) Name of the domain for the NAPTR record.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsnaptrrec. It is a composite of `domain` and `type`.
* `ecssubnet` - Subnet for which the cached NAPTR record need to be removed.
* `flags` - flags for this NAPTR.
* `nodeid` - Unique number that identifies the cluster node.
* `order` - An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest.
* `preference` - An integer specifying the preference of this NAPTR among NAPTR records having same order. Lower the number, higher the preference.
* `recordid` - Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying domain name and all other naptr record attributes as was supplied during the add command.
* `regexp` - The regular expression, that specifies the substitution expression for this NAPTR.
* `replacement` - The replacement domain name for this NAPTR.
* `services` - Service Parameters applicable to this delegation path.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
