---
subcategory: "DNS"
---

# Data Source `dnssrvrec`

The dnssrvrec data source allows you to retrieve information about DNS SRV (Service) records configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnssrvrec" "tf_dnssrvrec" {
  domain = "example.com"
  target = "_sip._udp.example.com"
}

output "priority" {
  value = data.citrixadc_dnssrvrec.tf_dnssrvrec.priority
}

output "weight" {
  value = data.citrixadc_dnssrvrec.tf_dnssrvrec.weight
}

output "port" {
  value = data.citrixadc_dnssrvrec.tf_dnssrvrec.port
}

output "ttl" {
  value = data.citrixadc_dnssrvrec.tf_dnssrvrec.ttl
}
```


## Argument Reference

* `domain` - (Required) Domain name, which, by convention, is prefixed by the symbolic name of the desired service and the symbolic name of the desired protocol, each with an underscore (_) prepended. For example, if an SRV-aware client wants to discover a SIP service that is provided over UDP, in the domain example.com, the client performs a lookup for _sip._udp.example.com.
* `target` - (Required) Target host for the specified service.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssrvrec. It is a combination of domain, target, and type.
* `ecssubnet` - Subnet for which the cached SRV record need to be removed.
* `nodeid` - Unique number that identifies the cluster node.
* `port` - Port on which the target host listens for client requests.
* `priority` - Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.
* `ttl` - Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
* `weight` - Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.
