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

* `id` - The id of the service_lbmonitor_binding. It is a system-generated identifier.
* `monstate` - The configured state (enable/disable) of the monitor on this server.
* `passive` - Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `weight` - Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.
