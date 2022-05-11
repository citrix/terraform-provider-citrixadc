---
subcategory: "Basic"
---

# Resource: service_lbmonitor_binding

The service_lbmonitor_binding resource is used to bind lbmonitor to service.


## Example usage

```hcl
resource "citrixadc_service" "tf_service" {
  servicetype         = "HTTP"
  name                = "tf_service"
  ipaddress           = "10.77.33.22"
  ip                  = "10.77.33.22"
  port                = "80"
  state               = "ENABLED"
  wait_until_disabled = true
}
resource "citrixadc_lbmonitor" "tf_monitor" {
  monitorname = "tf_monitor"
  type        = "HTTP"
}
resource "citrixadc_service_lbmonitor_binding" "tf_binding" {
  name         = citrixadc_service.tf_service.name
  monitor_name = citrixadc_lbmonitor.tf_monitor.monitorname
  monstate     = "ENABLED"
  weight       = 2
}
```


## Argument Reference

* `name` - (Required) Name of the service to which to bind a policy or monitor.
* `monitor_name` - (Required) The monitor Names.
* `monstate` - (Optional) The configured state (enable/disable) of the monitor on this server.
* `passive` - (Optional) Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
* `weight` - (Optional) Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the service_lbmonitor_binding. It is the concatenation of `name` and `monitor_name` attributes separated by comma.


## Import

A service_lbmonitor_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_service_lbmonitor_binding.tf_binding tf_service,tf_monitor
```
