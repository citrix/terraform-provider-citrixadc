---
subcategory: "Network"
---

# Data Source `ipv6`

The ipv6 data source allows you to retrieve information about IPv6 global configuration settings.


## Example usage

```terraform
data "citrixadc_ipv6" "tf_ipv6" {
  td = 0
}

output "ralearning" {
  value = data.citrixadc_ipv6.tf_ipv6.ralearning
}

output "ndbasereachtime" {
  value = data.citrixadc_ipv6.tf_ipv6.ndbasereachtime
}

output "routerredirection" {
  value = data.citrixadc_ipv6.tf_ipv6.routerredirection
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `dodad` - Enable the Citrix ADC to do Duplicate Address Detection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration.
* `natprefix` - Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
* `ndbasereachtime` - Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation.
* `ndretransmissiontime` - Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device.
* `ralearning` - Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers.
* `routerredirection` - Enable the Citrix ADC to do Router Redirection.
* `usipnatprefix` - IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on.
* `id` - The id of the ipv6 resource.
