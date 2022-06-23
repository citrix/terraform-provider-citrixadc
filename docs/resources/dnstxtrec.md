---
subcategory: "<fillme>"
---

# Resource: <fillme>

The <resource> resource is used to create <fillme>.


## Example usage

```hcl
<fillme>
```


## Argument Reference

* `String` - (Optional) Information to store in the TXT resource record. Enclose the string in single or double quotation marks. A TXT resource record can contain up to six strings, each of which can contain up to 255 characters. If you want to add a string of more than 255 characters, evaluate whether splitting it into two or more smaller strings, subject to the six-string limit, works for you.
* `domain` - (Optional) Name of the domain for the TXT record.
* `ecssubnet` - (Optional) Subnet for which the cached TXT record need to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `recordid` - (Optional) Unique, internally generated record ID. View the details of the TXT record to obtain its record ID. Mutually exclusive with the string parameter.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the <fillme>. It has the same value as the `name` attribute.


## Import

A <resource> can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
