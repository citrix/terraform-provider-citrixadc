---
subcategory: "Basic"
---

# Resource: servicegroup\_lbmonitor\_binding

The servicegroup\_lbmonitor\_binding resource is used to bind servicegroups to load balancing monitors.

~> If you are using this resource to bind lbmonitors to a servicegroup,
do not define the `lbmonitor` attribute in the servicegroup resource.


## Example usage

```hcl
resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
}

resource "citrixadc_lbmonitor" "tf_monitor" {
  monitorname = "tf_monitor"
  type        = "HTTP"
}

resource "citrixadc_servicegroup_lbmonitor_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  monitorname      = citrixadc_lbmonitor.tf_monitor.monitorname
  weight           = 80
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the service group.
* `monitorname` - (Required) Monitor name.
* `customserverid` - (Optional) Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.
* `dbsttl` - (Optional) Specify the TTL for DNS record for domain based service. The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `hashid` - (Optional) Unique numerical identifier used by hash based load balancing methods to identify a service.
* `monstate` - (Optional) Monitor state.
* `nameserver` - (Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `order` - (Optional) Order number to be assigned to the servicegroup member.
* `passive` - (Optional) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `port` - (Optional) Port number of the service. Each service must have a unique port number.
* `serverid` - (Optional) The identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - (Optional) Initial state of the service after binding.
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup\_lbmonitor\_binding. It is the concatenation of the `servicegroupname` and `monitorname` attributes separated by a comma.


## Import

A servicegroup\_lbmonitor\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_servicegroup_lbmonitor_binding.tf_binding tf_servicegroup,tf_monitor
```
