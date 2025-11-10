---
subcategory: "Network"
---

# Resource: route

The route resource is used to create routing rules.


## Example usage

```hcl
resource "citrixadc_route" "tf_route" {
    network    = "100.0.100.0"
    netmask    = "255.255.255.0"
    gateway    = "100.0.1.1"
    advertise  = "ENABLED"
}
```


## Argument Reference

* `network` - (Optional) IPv4 network address for which to add a route entry in the routing table of the Citrix ADC.
* `netmask` - (Optional) The subnet mask associated with the network address.
* `gateway` - (Optional) IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.
* `vlan` - (Optional) VLAN as the gateway for this route.
* `cost` - (Optional) Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `distance` - (Optional) Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.
* `cost1` - (Optional) The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.
* `weight` - (Optional) Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.
* `advertise` - (Optional) Advertise this route. Possible values: [ DISABLED, ENABLED ]
* `msr` - (Optional) Monitor this route using a monitor of type ARP or PING. Possible values: [ ENABLED, DISABLED ]
* `monitor` - (Optional) Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.
* `ownergroup` - (Optional) The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.
* `routetype` - (Optional) Protocol used by routes that you want to remove from the routing table of the Citrix ADC. Possible values: [ CONNECTED, STATIC, DYNAMIC, OSPF, ISIS, RIP, BGP ]
* `detail` - (Optional) Display a detailed view.
* `mgmt` - (Optional) Route in management plane.
* `protocol` - (Optional) Routing protocol used for advertising this route.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the route. It is the conatenation of the `network`, `netmask` and `gateway` attributes.


## Import

A route can be imported using its id value, e.g.

```shell
terraform import citrixadc_route.tf_route 100.0.100.0__255.255.255.0__100.0.1.1
```
