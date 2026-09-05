---
subcategory: "Basic"
---

# Data Source: service_lbmonitor_binding

The service_lbmonitor_binding data source allows you to retrieve information about the binding between a service and a load balancing monitor.


## Example usage

```terraform
data "citrixadc_service_lbmonitor_binding" "tf_binding" {
  name         = "tf_service"
  monitor_name = "tf_monitor"
}

output "monstate" {
  value = data.citrixadc_service_lbmonitor_binding.tf_binding.monstate
}

output "weight" {
  value = data.citrixadc_service_lbmonitor_binding.tf_binding.weight
}
```


## Argument Reference

* `name` - (Required) Name of the service to which to bind a monitor.
* `monitor_name` - (Required) The monitor Names.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the service_lbmonitor_binding. It is the concatenation of `name` and `monitor_name` attributes separated by comma.
* `monstate` - The configured state (enable/disable) of the monitor on this server.
* `passive` - Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `weight` - Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.

### Read-only service_lbmonitor_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_service_lbmonitor_binding` resource). They are GET-only/Computed monitor runtime metadata. Any attribute the appliance does not return is `null`.

* `monitortotalfailedprobes` - Total number of failed probes.
* `lastresponse` - The string form of monstatcode.
* `failedprobes` - Number of the current failed monitoring probes.
* `monstatparam2` - Second parameter for use with message code.
* `totalprobes` - The total number of probes sent.
* `dup_weight` - The weight of the monitor.
* `monitortotalprobes` - Total number of probes sent to monitor this service.
* `monstatparam1` - First parameter for use with message code.
* `monitorcurrentfailedprobes` - Total number of currently failed probes.
* `monstatcode` - The code indicating the monitor response.
* `monitor_state` - The running state of the monitor on this service (for example UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, DISABLED).
* `totalfailedprobes` - The total number of failed probes.
* `dup_state` - State value from table (ENABLED or DISABLED).
* `responsetime` - Response time of this monitor.
* `monstatparam3` - Third parameter for use with message code.
