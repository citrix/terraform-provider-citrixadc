---
subcategory: "DNS"
---

# Resource: dnssrvrec

The dnssrvrec resource is used to create DNS srvrec.


## Example usage

```hcl
resource "citrixadc_dnssrvrec" "dnssrvrec" {
  domain   = "example.com"
  target   = "_sip._udp.example.com"
  priority = 1
  weight   = 1
  port     = 22
  ttl      = 3600
}
```


## Argument Reference

* `domain` - (Required) Domain name, which, by convention, is prefixed by the symbolic name of the desired service and the symbolic name of the desired protocol, each with an underscore (_) prepended. For example, if an SRV-aware client wants to discover a SIP service that is provided over UDP, in the domain example.com, the client performs a lookup for _sip._udp.example.com.
* `ecssubnet` - (Optional) Subnet for which the cached SRV record need to be removed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `port` - (Required) Port on which the target host listens for client requests.
* `priority` - (Required) Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.
* `target` - (Required) Target host for the specified service.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `type` - (Optional) Type of records to display. Available settings function as follows: * ADNS - Display all authoritative address records. * PROXY - Display all proxy address records. * ALL - Display all address records.
* `weight` - (Required) Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssrvrec. It has the same value as the `domain` attribute.


## Import

A dnssrvrec can be imported using its name, e.g.

```shell
terraform import citrixadc_dnssrvrec.dnssrvrec example.com
```
