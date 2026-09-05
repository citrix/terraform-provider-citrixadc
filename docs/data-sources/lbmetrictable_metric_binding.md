---
subcategory: "Load Balancing"
---

# Data Source: lbmetrictable_metric_binding

The lbmetrictable_metric_binding data source allows you to retrieve information about metric table bindings.


## Example usage

```terraform
data "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
  metric      = "2.3.6.4.5"
  metrictable = "Table-Custom"
}

output "snmpoid" {
  value = data.citrixadc_lbmetrictable_metric_binding.tf_bind.snmpoid
}

output "metric" {
  value = data.citrixadc_lbmetrictable_metric_binding.tf_bind.metric
}
```


## Argument Reference

* `metric` - (Required) Name of the metric for which to change the SNMP OID.
* `metrictable` - (Required) Name of the metric table.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmetrictable_metric_binding. It is the concatenation of the `metrictable` and `metric` attributes separated by a comma.
* `snmpoid` - New SNMP OID of the metric.

### Read-only lbmetrictable_metric_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_lbmetrictable_metric_binding` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `metrictype` - Indication if it is a configured or internal metric. Possible values = INTERNAL, CONFIGURED.
