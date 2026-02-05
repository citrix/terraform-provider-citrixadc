---
subcategory: "Network"
---

# Data Source: citrixadc_route6

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
