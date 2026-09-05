---
subcategory: "Load Balancing"
---

# Resource: lbmonitor\_metric\_binding

The lbmonitor\_metric\_binding resource is used to bind a metric to a load balancing monitor.


## Example usage

```hcl
resource "citrixadc_lbmonitor_metric_binding" "tf_binding" {
  monitorname     = "tf-monitor1"
  metric          = "metric1"
  metricthreshold = 100
  metricweight    = 50
}
```


## Argument Reference

* `monitorname` - (Required) Name of the monitor.
* `metric` - (Required) Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.
* `metricthreshold` - (Optional) Threshold to be used for that metric.
* `metricweight` - (Optional) The weight for the specified service metric with respect to others.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor\_metric\_binding. It is the concatenation of the `monitorname` and `metric` attributes separated by a comma.


## Import

A lbmonitor_metric_binding can be imported using the concatenation of the `monitorname` and `metric` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_lbmonitor_metric_binding.tf_binding tf-monitor1,metric1
```
