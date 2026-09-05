---
subcategory: "NS"
---

# Data Source: nsmigration

The nsmigration data source allows you to retrieve information about the
Migration operation resource configuration.


## Example usage

```terraform
data "citrixadc_nsmigration" "tf_nsmigration" {
}

output "dumpsession" {
  value = data.citrixadc_nsmigration.tf_nsmigration.dumpsession
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `dumpsession` - Displays the current active migrated session details, if DUMPSESSION option is YES. Possible values: [ YES, NO ]
* `id` - The id of the nsmigration. It is a system-generated identifier.
