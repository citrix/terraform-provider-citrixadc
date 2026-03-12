---
subcategory: "Load Balancing"
---

# Data Source: lbmonitor_metric_binding

The lbmonitor_metric_binding data source allows you to retrieve information about a specific metric binding to a load balancing monitor.


## Example usage

```terraform
data "citrixadc_lbmonitor_metric_binding" "tf_binding" {
  monitorname = "tf-monitor1"
  metric      = "metric1"
}

output "metricthreshold" {
  value = data.citrixadc_lbmonitor_metric_binding.tf_binding.metricthreshold
}

output "metricweight" {
  value = data.citrixadc_lbmonitor_metric_binding.tf_binding.metricweight
}
```


## Argument Reference

* `monitorname` - (Required) Name of the monitor.
* `metric` - (Required) Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor_metric_binding. It is a system-generated identifier.
* `metricthreshold` - Threshold to be used for that metric.
* `metricweight` - The weight for the specified service metric with respect to others.
