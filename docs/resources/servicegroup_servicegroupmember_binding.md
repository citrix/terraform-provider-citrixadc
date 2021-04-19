---
subcategory: "Basic"
---

# Resource: servicegroup\_servicegroupmember\_binding

The servicegroup\_servicegroupmember\_binding resource is used to bind service members to servicegroups.


## Example usage

```hcl
resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale        = "DNS"
}

resource "citrixadc_servicegroup_servicegroupmember_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  ip               = "10.78.22.33"
  port             = 80
}
```


## Argument Reference

* `ip` - (Optional) IP Address.
* `port` - (Optional) Server port number. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
* `servername` - (Optional) Name of the server to which to bind the service group.
* `customserverid` - (Optional) The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
* `serverid` - (Optional) The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - (Optional) Initial state of the service group. Possible values: [ ENABLED, DISABLED ]
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `nameserver` - (Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `dbsttl` - (Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `servicegroupname` - (Required) Name of the service group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup\_servicegroupmember\_binding. It is the concatenation of three components separated by comma. First component is the `servicegroupname`. Second component is the `ip` or the `servername` attribute. Last optional component is the `port` attribute.
