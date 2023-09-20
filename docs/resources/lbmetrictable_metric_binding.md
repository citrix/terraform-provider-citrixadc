---
subcategory: "Load Balancing"
---

# Resource: lbmetrictable_metric_binding

The lbmetrictable_metric_binding resource is used to bind metric to lb metrictable.


## Example usage

```hcl
resource "citrixadc_lbmetrictable" "Table" {
  metrictable = "Table-Custom"
}
resource "citrixadc_lbmetrictable_metric_binding" "tf_bind" {
  metric      = "2.3.6.4.5"
  metrictable = citrixadc_lbmetrictable.Table.metrictable
  snmpoid     = "1.2.3.6.5"
}

```


## Argument Reference

* `metric` - (Required) Name of the metric for which to change the SNMP OID. Minimum length =  1
* `Snmpoid` - (Required) New SNMP OID of the metric. Minimum length =  1
* `metrictable` - (Required) Name of the metric table. Minimum length =  1 Maximum length =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmetrictable_metric_binding. It is the concatenation of the `metrictable` and `metric` attributes separated by a comma.


## Import

A lbmetrictable_metric_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lbmetrictable_metric_binding.tf_bind Table-Custom,2.3.6.4.5
```
