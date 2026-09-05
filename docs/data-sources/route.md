---
subcategory: "Network"
---

# Data Source: route

The route data source allows you to retrieve information about a specific IPv4 route configuration.

## Example Usage

```terraform
data "citrixadc_route" "tf_route" {
  network = "100.0.100.0"
  netmask = "255.255.255.0"
  td      = 0
}

output "gateway" {
  value = data.citrixadc_route.tf_route.gateway
}

output "advertise" {
  value = data.citrixadc_route.tf_route.advertise
}
```

## Argument Reference

The following arguments are required:

* `network` - (Required) IPv4 network address for which to retrieve the route entry from the routing table of the Citrix ADC.
* `netmask` - (Required) The subnet mask associated with the network address.
* `td` - (Required) Integer value that uniquely identifies the traffic domain. If you do not specify an ID, the default traffic domain (ID 0) is used.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `advertise` - Advertise this route. Possible values: `ENABLED`, `DISABLED`.
* `cost` - Positive integer used by the routing algorithms to determine preference for using this route.
* `cost1` - The cost of a route is used to compare routes of the same type.
* `distance` - Administrative distance of this route.
* `gateway` - IP address of the gateway for this route.
* `monitor` - Name of the monitor configured to monitor this route.
* `msr` - Monitor this route using a monitor of type ARP or PING. Possible values: `ENABLED`, `DISABLED`.
* `ownergroup` - The owner node group in a Cluster for this route.
* `protocol` - Routing protocol used for advertising this route.
* `routetype` - Protocol used by routes.
* `vlan` - VLAN as the gateway for this route.
* `weight` - Positive integer used by the routing algorithms to determine preference for this route over others of equal cost.
* `id` - The id of the route resource.

### Read-only route metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_route` resource). They are GET-only / Computed. Any attribute the appliance does not return is `null`.

* `gatewayname` - The name of the gateway for this route. For a route other than a link load balancing (LLB) route, this value is null.
* `type` - State of the RNAT.
* `dynamic` - State of the route.
* `static` - Whether this is a static route.
* `permanent` - Whether this is a permanent route.
* `direct` - Whether this is a direct route.
* `nat` - Whether this is a NAT route.
* `lbroute` - Whether this is a link load balancing (LLB) route.
* `adv` - Whether this route is advertised.
* `tunnel` - Show whether it is a tunnel route or not.
* `data` - Internal data of this route.
* `data0` - Internal route type is stored, used for get.
* `flags` - If this route is dynamic, the name of the routing protocol from which it was learned.
* `routeowners` - In a cluster, the set of nodes from which this dynamic route has been learnt. A list of strings.
* `retain` - Internal retain value of this route.
* `ospf` - OSPF protocol.
* `isis` - ISIS protocol.
* `rip` - RIP protocol.
* `bgp` - BGP protocol.
* `dhcp` - DHCP protocol.
* `advospf` - Advertised through OSPF protocol.
* `advisis` - Advertised through ISIS protocol.
* `advrip` - Advertised through RIP protocol.
* `advbgp` - Advertised through BGP protocol.
* `state` - The state of the static route. Possible values: UP, DOWN.
* `totalprobes` - The total number of probes sent.
* `totalfailedprobes` - The total number of failed probes.
* `failedprobes` - Number of the current failed monitoring probes.
* `monstatcode` - The code indicating the monitor response.
* `monstatparam1` - First parameter used with the message code.
* `monstatparam2` - Second parameter used with the message code.
* `monstatparam3` - Third parameter used with the message code.
