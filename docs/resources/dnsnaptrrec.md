---
subcategory: "DNS"
---

# Resource: dnsnaptrrec

The dnsnaptrrec resource is used to create DNS naptrrec.


## Example usage

```hcl
resource "citrixadc_dnsnaptrrec" "dnsnaptrrec" {
  domain      = "example.com"
  order       = 10
  preference  = 2
  ttl         = 3600
  replacement = "example1.com"
}
```


## Argument Reference

* `domain` - (Required) Name of the domain for the NAPTR record.
* `order` - (Required) An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest
* `preference` - (Required) An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.
* `ecssubnet` - (Optional) Subnet for which the cached NAPTR record need to be removed.
* `flags` - (Optional) flags for this NAPTR.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `recordid` - (Optional) Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying domain name and all other naptr record attributes as was supplied during the add command.
* `regexp` - (Optional) The regular expression, that specifies the substitution expression for this NAPTR
* `replacement` - (Optional) The replacement domain name for this NAPTR.
* `services` - (Optional) Service Parameters applicable to this delegation path.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsnaptrrec. It has the same value as the `domain` attribute.


## Import

A <resource> can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsnaptrrec.dnsnaptrrec example.com
```
