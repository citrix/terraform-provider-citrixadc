---
subcategory: "Appflow"
---

# Data Source: appflowaction_analyticsprofile_binding

The appflowaction_analyticsprofile_binding data source allows you to retrieve information about a specific binding between an appflow action and an analytics profile.

## Example Usage

```terraform
data "citrixadc_appflowaction_analyticsprofile_binding" "tf_appflowaction_analyticsprofile_binding" {
  name             = "test_action"
  analyticsprofile = "my_analyticsprofile"
}

output "name" {
  value = data.citrixadc_appflowaction_analyticsprofile_binding.tf_appflowaction_analyticsprofile_binding.name
}

output "analyticsprofile" {
  value = data.citrixadc_appflowaction_analyticsprofile_binding.tf_appflowaction_analyticsprofile_binding.analyticsprofile
}
```

## Argument Reference

* `name` - (Required) Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `analyticsprofile` - (Required) Analytics profile to be bound to the appflow action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowaction_analyticsprofile_binding. It is a system-generated identifier.
