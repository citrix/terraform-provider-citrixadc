---
subcategory: "System"
---

# Data Source: systemextramgmtcpu

The systemextramgmtcpu data source allows you to retrieve information about the extra management CPU configuration on the target ADC.


## Example usage

```terraform
data "citrixadc_systemextramgmtcpu" "tf_extramgmtcpu" {
}

output "nodeid" {
  value = data.citrixadc_systemextramgmtcpu.tf_extramgmtcpu.nodeid
}
```


## Argument Reference

This datasource does not require any arguments.


## Attribute Reference

The following attributes are available:

* `enabled` - (Required) Boolean value indicating the effective state of the extra management CPU.


## Attribute Reference

* `id` - The id of the systemextramgmtcpu. It is a system-generated identifier.
