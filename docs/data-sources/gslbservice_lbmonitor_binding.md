---
subcategory: "GSLB"
---

# Data Source: gslbservice_lbmonitor_binding

The gslbservice_lbmonitor_binding data source allows you to retrieve information about a GSLB service to monitor binding.


## Example usage

```terraform
data "citrixadc_gslbservice_lbmonitor_binding" "tf_gslbservice_lbmonitor_binding" {
  servicename  = "tf_gslb1vservice"
  monitor_name = "tf_monitor"
}

output "servicename" {
  value = data.citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding.servicename
}

output "monstate" {
  value = data.citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding.monstate
}

output "weight" {
  value = data.citrixadc_gslbservice_lbmonitor_binding.tf_gslbservice_lbmonitor_binding.weight
}
```


## Argument Reference

* `servicename` - (Required) Name of the GSLB service.
* `monitor_name` - (Required) Monitor name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbservice_lbmonitor_binding. It is a system-generated identifier.
* `monstate` - State of the monitor bound to gslb service.
* `weight` - Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.

### Read-only gslbservice_lbmonitor_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_gslbservice_lbmonitor_binding` resource). Each is Computed and is `null` when the appliance does not return it.

* `monstatparam2` - Second parameter for use with message code.
* `failedprobes` - Number of the current failed monitoring probes.
* `totalfailedprobes` - The total number of failed probs.
* `lastresponse` - Displays the gslb monitor status in string format.
* `monitor_state` - The running state of the monitor on this service (for example `UP`, `DOWN`, `UNKNOWN`, `DISABLED`).
* `monitortotalfailedprobes` - Total number of failed probes.
* `monstatcode` - The code indicating the monitor response.
* `monitorcurrentfailedprobes` - Total number of currently failed probes.
* `monstatparam1` - First parameter for use with message code.
* `monstatparam3` - Third parameter for use with message code.
* `responsetime` - Response time of this monitor.
* `monitortotalprobes` - Total number of probes sent to monitor this service.
