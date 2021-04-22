---
subcategory: "DNS"
---

# Resource: dnsnsrec

The dnsnsrec resource is used to create dns name server records.


## Example usage

```hcl
resource "citrixadc_dnsnsrec" "tf_dnsnsrec" {
    domain = "www.test.com"
    nameserver = "192.168.1.100"
}
```


## Argument Reference

* `domain` - (Optional) Domain name.
* `nameserver` - (Optional) Host name of the name server to add to the domain.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsnsrec. It is the concatenation of the `domain` and `nameserver` attributes separated by a comma.


## Import

A dnsnsrec can be imported using its id, e.g.

```shell
terraform import citrixadc_dnsnsrec.tf_dnsnsrec www.test.com,192.168.1.100
```
