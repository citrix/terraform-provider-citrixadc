---
subcategory: "Basic"
---

# Data Source: servicegroup_servicegroupmember_binding

The servicegroup_servicegroupmember_binding data source allows you to retrieve information about a service group member binding.

## Example Usage

```terraform
data "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
  servicegroupname = "tf_servicegroup"
  ip               = "10.78.22.33"
  port             = 80
  servername       = "10.78.22.33"
}

output "servicegroupname" {
  value = data.citrixadc_servicegroup_servicegroupmember_binding.tf_binding.servicegroupname
}

output "order" {
  value = data.citrixadc_servicegroup_servicegroupmember_binding.tf_binding.order
}

output "weight" {
  value = data.citrixadc_servicegroup_servicegroupmember_binding.tf_binding.weight
}
```

## Argument Reference

* `servicegroupname` - (Required) Name of the service group.
* `ip` - (Optional) IP Address. Either `ip` or `servername` is required.
* `port` - (Optional) Server port number.
* `servername` - (Optional) Name of the server to which to bind the service group. Either `ip` or `servername` is required.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup_servicegroupmember_binding. It is the concatenation of the `servicegroupname`, the `ip` or `servername`, and the `port` attributes separated by a comma.
* `customserverid` - The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
* `dbsttl` - Specify the TTL for DNS record for domain based service. The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `hashid` - The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `nameserver` - Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `order` - Order number to be assigned to the servicegroup member.
* `serverid` - The identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - Initial state of the service group.
* `weight` - Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.

### Read-only servicegroup_servicegroupmember_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_servicegroup_servicegroupmember_binding` resource). They are GET-only/Computed and any attribute the appliance does not return is `null`.

* `svrstate` - The state of the service (for example `UP`, `DOWN`, `OUT OF SERVICE`).
* `tickssincelaststatechange` - Time in 10 millisecond ticks since the last state change.
* `statechangetimesec` - Time when the last state change occurred (seconds part).
* `trofsreason` - Reason the service group member is in TROFS, if applicable.
* `trofsdelay` - Delay before moving to TROFS.
* `orderstr` - Order number in string form assigned to the servicegroup member.
* `graceful` - Whether to wait for all existing connections to the service to terminate before shutting down the service.
* `svcitmpriority` - The priority of the FQDN service items for SRV server binding.
* `delay` - Time, in seconds, allocated for a shutdown of the services in the service group.
