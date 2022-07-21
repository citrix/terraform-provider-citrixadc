---
subcategory: "DNS"
---

# Resource: dnsmxrec

The dnsmxrec resource is used to create DNS mxRec.


## Example usage

```hcl
resource "citrixadc_dnsmxrec" "dnsmxrec" {
  domain = "example.com"
  mx     = "mail.example.com"
  pref   = 5
  ttl    = 3600
}
```


## Argument Reference

* `domain` - (Required) Domain name for which to add the MX record.
* `mx` - (Required) Host name of the mail exchange server.
* `pref` - (Required) Priority number to assign to the mail exchange server. A domain name can have multiple mail servers, with a priority number assigned to each server. The lower the priority number, the higher the mail server's priority. When other mail servers have to deliver mail to the specified domain, they begin with the mail server with the lowest priority number, and use other configured mail servers, in priority order, as backups.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `ecssubnet` - (Optional) Subnet for which the cached MX record need to be removed.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsmxrec. It has the same value as the `domain` attribute.


## Import

A dnsmxrec can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsmxrec.dnsmxrec example.com
```
