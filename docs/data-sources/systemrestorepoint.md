---
subcategory: "System"
---

# Data Source: systemrestorepoint

The systemrestorepoint data source allows you to retrieve information about an
existing restore point (a named snapshot of the appliance configuration plus a
tech-support bundle) on a Citrix ADC.


## Example usage

```terraform
data "citrixadc_systemrestorepoint" "example" {
  filename = "pre-upgrade-restorepoint"
}

output "restorepoint_filename" {
  value = data.citrixadc_systemrestorepoint.example.filename
}
```


## Argument Reference

* `filename` - (Required) Name of the restore point to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The identifier of the restore point. It has the same value as the
  `filename` attribute.
* `filename` - Name of the restore point.
