---
subcategory: "Load Balancing"
---

# Data Source: lbmetrictable_metric_binding

The lbmetrictable_metric_binding data source allows you to retrieve information about metric table bindings.


## Example Usage

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

* `id` - The id of the lbmetrictable_metric_binding. It is a system-generated identifier.
* `snmpoid` - New SNMP OID of the metric.
