---
subcategory: "Network"
---

# Resource: route6

The route6 resource is used to create route6.


## Example usage

```hcl
resource "citrixadc_route6" "tf_route6" {
  network  = "2001:db8:85a3::/64"
  vlan     = 2
  weight   = 5
  distance = 3
}
```


## Argument Reference

* `network` - (Required) IPv6 network address for which to add a route entry to the routing table of the Citrix ADC.
* `gateway` - (Optional) The gateway for this route. The value for this parameter is either an IPv6 address or null.
* `vlan` - (Optional) Integer value that uniquely identifies a VLAN through which the Citrix ADC forwards the packets for this route. Minimum value =  0 Maximum value =  4094
* `vxlan` - (Optional) Integer value that uniquely identifies a VXLAN through which the Citrix ADC forwards the packets for this route. Minimum value =  1 Maximum value =  16777215
* `weight` - (Optional) Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference. Minimum value =  1 Maximum value =  65535
* `distance` - (Optional) Administrative distance of this route from the appliance. Minimum value =  1 Maximum value =  254
* `cost` - (Optional) Positive integer used by the routing algorithms to determine preference for this route. The lower the cost, the higher the preference. Minimum value =  0 Maximum value =  65535
* `advertise` - (Optional) Advertise this route. Possible values: [ DISABLED, ENABLED ]
* `msr` - (Optional) Monitor this route with a monitor of type ND6 or PING. Possible values: [ ENABLED, DISABLED ]
* `monitor` - (Optional) Name of the monitor, of type ND6 or PING, configured on the Citrix ADC to monitor this route. Minimum length =  1
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `ownergroup` - (Optional) The owner node group in a Cluster for this route6. If owner node group is not specified then the route is treated as Striped route. Minimum length =  1
* `routetype` - (Optional) Type of IPv6 routes to remove from the routing table of the Citrix ADC. Possible values: [ CONNECTED, STATIC, DYNAMIC, OSPF, ISIS, BGP, RIP, ND-RA-ROUTE, FIB6 ]
* `detail` - (Optional) To get a detailed view.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the route6. It has the same value as the `network` attribute.


## Import

A route6 can be imported using its name, e.g.

```shell
terraform import citrixadc_route6.tf_route6 2001:db8:85a3::/64
```
