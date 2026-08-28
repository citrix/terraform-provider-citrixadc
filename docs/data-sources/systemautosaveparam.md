---
subcategory: "System"
---

# Data Source: systemautosaveparam

The systemautosaveparam data source allows you to retrieve information about the
system autosave parameters configuration.


## Example usage

```terraform
data "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
}

output "status" {
  value = data.citrixadc_systemautosaveparam.tf_systemautosaveparam.status
}

output "periodicsave" {
  value = data.citrixadc_systemautosaveparam.tf_systemautosaveparam.periodicsave
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `status` - Configure autosave feature. Possible values: [ DEFAULT, DISABLED, ENABLED ]
* `periodicsave` - Enable or disable periodic save of autosave configuration. Possible values: [ ENABLED, DISABLED ]
* `periodicsavefrequency` - Frequency in multiple of 60 minutes for periodic save of autosave configuration. Default value is 720 minutes.
* `id` - The id of the systemautosaveparam. It is a system-generated identifier.
