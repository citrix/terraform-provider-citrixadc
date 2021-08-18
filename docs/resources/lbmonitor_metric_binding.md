---
subcategory: "Load Balancing"
---

# Resource: lbmonitor\_metric\_binding

The lbmonitor\_metric\_binding resource is used bind load balancing monitor to metric.


## Example usage

```hcl
resource citrixadc_lbmonitor_metric_binding demo_binding1 {
 monitorname = "mload2"
 metric = "demometric"
 metricthreshold = 100
}
```


## Argument Reference

* `metric` - (Required) Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation.
* `metricweight` - (Optional) The weight for the specified service metric with respect to others.
* `metricthreshold` - (Optional) Threshold to be used for that metric.
* `monitorname` - (Required) Name of the monitor.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor\_metric\_binding. It is the concatenation of the `monitorname` and `metric` attributes separated by a comma.


```shell
terraform import citrixadc_lbmonitor_metric_binding.tf_lbmonitor_metric_binding tf_lbmonitor_metric_binding
```
