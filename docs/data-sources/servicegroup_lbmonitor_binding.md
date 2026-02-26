---
subcategory: "Basic"
---

# Data Source: servicegroup_lbmonitor_binding

The servicegroup_lbmonitor_binding data source allows you to retrieve information about a specific binding between a servicegroup and a load balancing monitor.

## Example Usage

```terraform
data "citrixadc_servicegroup_lbmonitor_binding" "example" {
  servicegroupname = "my_servicegroup"
  monitor_name     = "my_monitor"
}

output "weight" {
  value = data.citrixadc_servicegroup_lbmonitor_binding.example.weight
}

output "monstate" {
  value = data.citrixadc_servicegroup_lbmonitor_binding.example.monstate
}

output "state" {
  value = data.citrixadc_servicegroup_lbmonitor_binding.example.state
}
```

## Argument Reference

* `servicegroupname` - (Required) Name of the service group.
* `monitor_name` - (Required) Monitor name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `customserverid` - Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.
* `dbsttl` - Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `hashid` - Unique numerical identifier used by hash based load balancing methods to identify a service.
* `port` - Port number of the service. Each service must have a unique port number.
* `id` - The id of the servicegroup_lbmonitor_binding. It is a system-generated identifier.
* `monstate` - Monitor state.
* `nameserver` - Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `order` - Order number to be assigned to the servicegroup member.
* `passive` - Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `serverid` - The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - Initial state of the service after binding.
* `weight` - Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
