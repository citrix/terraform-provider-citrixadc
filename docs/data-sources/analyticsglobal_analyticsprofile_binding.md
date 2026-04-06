---
subcategory: "Analytics"
---

# Data Source: analyticsglobal_analyticsprofile_binding

The analyticsglobal_analyticsprofile_binding data source allows you to retrieve information about the analytics profile that is bound globally.

## Example Usage

```terraform
data "citrixadc_analyticsglobal_analyticsprofile_binding" "tf_binding" {
  analyticsprofile = "my_analyticsprofile"
}

output "analyticsprofile" {
  value = data.citrixadc_analyticsglobal_analyticsprofile_binding.tf_binding.analyticsprofile
}

output "id" {
  value = data.citrixadc_analyticsglobal_analyticsprofile_binding.tf_binding.id
}
```

## Argument Reference

* `analyticsprofile` - (Required) Name of the analytics profile bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the analyticsglobal_analyticsprofile_binding. It has the same value as the `analyticsprofile` attribute.
