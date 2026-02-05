---
subcategory: "Load Balancing"
---

# Data Source `lbmetrictable`

The lbmetrictable data source allows you to retrieve information about a load balancing metric table.


## Example usage

```terraform
data "citrixadc_lbmetrictable" "tf_lbmetrictable" {
  metrictable = "my_lbmetrictable"
}

output "metric" {
  value = data.citrixadc_lbmetrictable.tf_lbmetrictable.metric
}

output "snmpoid" {
  value = data.citrixadc_lbmetrictable.tf_lbmetrictable.snmpoid
}
```


## Argument Reference

* `metrictable` - (Required) Name for the metric table. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `metric` - Name of the metric for which to change the SNMP OID.
* `snmpoid` - New SNMP OID of the metric.

## Attribute Reference

* `id` - The id of the lbmetrictable. It has the same value as the `metrictable` attribute.


## Import

A lbmetrictable can be imported using its name, e.g.

```shell
terraform import citrixadc_lbmetrictable.tf_lbmetrictable my_lbmetrictable
```
