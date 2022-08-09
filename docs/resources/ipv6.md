---
subcategory: "Network"
---

# Resource: ipv6

The ipv6 resource is used to create ipv6.


## Example usage

```hcl
resource "citrixadc_ipv6" "tf_ipv6" {
  ralearning        = "DISABLED"
  ndbasereachtime   = 4000
  routerredirection = "ENABLED"
}
```


## Argument Reference

* `ralearning` - (Optional) Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers. Possible values: [ ENABLED, DISABLED ]
* `routerredirection` - (Optional) Enable the Citrix ADC to do Router Redirection. Possible values: [ ENABLED, DISABLED ]
* `ndbasereachtime` - (Optional) Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation. Minimum value =  1
* `ndretransmissiontime` - (Optional) Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device. Minimum value =  1
* `natprefix` - (Optional) Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `dodad` - (Optional) Enable the Citrix ADC to do Duplicate Address Detection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration. Possible values: [ ENABLED, DISABLED ]
* `usipnatprefix` - (Optional) IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipv6. It is a unique string prefixed with "tf-ipv6-"