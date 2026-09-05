---
subcategory: "Network"
---

# Data Source: route6

The route6 data source allows you to retrieve information about a specific IPv6 route configuration.

## Example Usage

```terraform
data "citrixadc_route6" "tf_route6" {
  network = "2001:db8:85a3::/64"
  td      = 0
}

output "vlan" {
  value = data.citrixadc_route6.tf_route6.vlan
}

output "weight" {
  value = data.citrixadc_route6.tf_route6.weight
}
```

## Argument Reference

The following arguments are required:

* `network` - (Required) IPv6 network address for which to retrieve the route entry from the routing table of the Citrix ADC.
* `td` - (Required) Integer value that uniquely identifies the traffic domain. If you do not specify an ID, the default traffic domain (ID 0) is used.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `advertise` - Advertise this route. Possible values: `ENABLED`, `DISABLED`.
* `cost` - Positive integer used by the routing algorithms to determine preference for using this route.
* `distance` - Administrative distance of this route from the appliance.
* `gateway` - The gateway for this route. The value is either an IPv6 address or null.
* `monitor` - Name of the monitor (of type ND6 or PING) configured to monitor this route.
* `msr` - Monitor this route with a monitor of type ND6 or PING. Possible values: `ENABLED`, `DISABLED`.
* `ownergroup` - The owner node group in a Cluster for this route6.
* `routetype` - Type of IPv6 routes.
* `vlan` - Integer value that uniquely identifies a VLAN through which the Citrix ADC forwards the packets for this route.
* `vxlan` - Integer value that uniquely identifies a VXLAN through which the Citrix ADC forwards the packets for this route.
* `weight` - Positive integer used by the routing algorithms to determine preference for this route over others of equal cost.
* `id` - The id of the route6 resource.

### Read-only route6 metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_route6` resource). They are GET-only / Computed. Any attribute the appliance does not return is `null`.

* `gatewayname` - The name of the gateway for this route.
* `type` - State of the RNAT.
* `dynamic` - Whether this route is dynamically learned or not.
* `data` - Internal data of this route.
* `flags` - For a dynamic route, the routing protocol from which the route was learned.
* `state` - Whether this route is UP or DOWN.
* `totalprobes` - The total number of probes sent.
* `totalfailedprobes` - The total number of failed probes.
* `failedprobes` - Current number of failed monitoring probes.
* `monstatcode` - The code indicating the monitor response.
* `monstatparam1` - First parameter for use with message code.
* `monstatparam2` - Second parameter for use with message code.
* `monstatparam3` - Third parameter for use with message code.
* `data1` - Internal data of this route. Possible values = ENABLED, DISABLED.
* `routeowners` - In a cluster, the set of nodes from which this dynamic route has been learnt. A list of strings.
* `retain` - Internal retain value of this route.
* `static` - Static route.
* `permanent` - Permanent Route.
* `connected` - Connected Route.
* `ospfv3` - For a dynamic route, the routing protocol from which the route was learned.
* `isis` - If this route is dynamic then which routing protocol was it learnt from.
* `active` - For a dynamic route, the routing protocol from which the route was learned.
* `bgp` - For a dynamic route, the routing protocol from which the route was learned.
* `rip` - For a dynamic route, the routing protocol from which the route was learned.
* `raroute` - For a dynamic route, the routing protocol from which the route was learned.
