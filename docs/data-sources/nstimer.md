---
subcategory: "Policy"
---

# Data Source `nstimer`

The nstimer data source allows you to retrieve information about a timer configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nstimer" "tf_nstimer" {
  name = "test_timer"
}

output "interval" {
  value = data.citrixadc_nstimer.tf_nstimer.interval
}

output "unit" {
  value = data.citrixadc_nstimer.tf_nstimer.unit
}
```


## Argument Reference

* `name` - (Required) Timer name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comments associated with this timer.
* `interval` - The frequency at which the policies bound to this timer are invoked.
* `newname` - The new name of the timer.
* `unit` - Timer interval unit. Can be SEC (seconds), MIN (minutes), or HOUR (hours).
