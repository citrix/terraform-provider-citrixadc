---
subcategory: "Stream"
---

# Data Source: streamidentifier_analyticsprofile_binding

The streamidentifier_analyticsprofile_binding data source allows you to retrieve information about a specific binding between a stream identifier and an analytics profile on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_streamidentifier_analyticsprofile_binding" "tf_binding" {
  name             = "tf_streamidentifier"
  analyticsprofile = "tf_analyticsprofile"
}

output "name" {
  value = data.citrixadc_streamidentifier_analyticsprofile_binding.tf_binding.name
}

output "analyticsprofile" {
  value = data.citrixadc_streamidentifier_analyticsprofile_binding.tf_binding.analyticsprofile
}
```

## Argument Reference

* `name` - (Required) The name of the stream identifier to which the analytics profile is bound.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the stream identifier.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the streamidentifier_analyticsprofile_binding.
