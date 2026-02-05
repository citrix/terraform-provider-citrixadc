---
subcategory: "Network"
---

# Data Source: citrixadc_route

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
