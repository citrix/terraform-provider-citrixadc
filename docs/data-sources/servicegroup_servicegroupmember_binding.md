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
* `ip` - (Optional) IP Address. Either IP Address or servername is required.
* `port` - (Required) Server port number.
* `servername` - (Optional) Name of the server to which to bind the service group. Either IP Address or servername is required.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup_servicegroupmember_binding. It is a system-generated identifier.
* `customserverid` - The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
* `dbsttl` - Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `hashid` - The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `nameserver` - Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `order` - Order number to be assigned to the servicegroup member.
* `serverid` - The identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - Initial state of the service group.
* `weight` - Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
