---
subcategory: "Basic"
---

# Resource: servicegroup_lbmonitor_binding

The servicegroup_lbmonitor_binding resource is used to bind servicegroups to load balancing monitors.

~> If you are using this resource to bind lbmonitors to a servicegroup,
do not define the `lbmonitor` attribute in the servicegroup resource.

## Example usage

```hcl
resource "citrixadc_servicegroup_lbmonitor_binding" "tf_binding" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
}
```


## Argument Reference

* `monitorname` - (Optional) Monitor name.
* `monstate` - (Optional) Monitor state. Possible values: [ ENABLED, DISABLED ]
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
* `passive` - (Optional) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `servicegroupname` - (Optional) Name of the service group.
* `port` - (Optional) Port number of the service. Each service must have a unique port number. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `customserverid` - (Optional) Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.
* `serverid` - (Optional) The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `state` - (Optional) Initial state of the service after binding. Possible values: [ ENABLED, DISABLED ]
* `hashid` - (Optional) Unique numerical identifier used by hash based load balancing methods to identify a service.
* `nameserver` - (Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `dbsttl` - (Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup_lbmonitor_binding. It is the concatenation of the `servicegroupname` and `monitorname` attributes


## Import

A servicegroup_lbmonitor_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_servicegroup_lbmonitor_binding.tf_binding tf_csaction tf_servicegroup,tf_monitorname
```
