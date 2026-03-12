---
subcategory: "Content Switching"
---

# Data Source: csvserver_analyticsprofile_binding

The csvserver_analyticsprofile_binding data source allows you to retrieve information about the binding between a content switching virtual server and an analytics profile.


## Example usage

```terraform
data "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
  name             = "tf_csvserver"
  analyticsprofile = "ns_analytics_global_profile"
}

output "id" {
  value = data.citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding.id
}

output "binding_name" {
  value = data.citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding.name
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the LB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_analyticsprofile_binding. It is the concatenation of the `name` and `analyticsprofile` attributes separated by a comma.
