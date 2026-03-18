---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_analyticsprofile_binding

The crvserver_analyticsprofile_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and an analytics profile.

## Example Usage

```terraform
data "citrixadc_crvserver_analyticsprofile_binding" "crvserver_analyticsprofile_binding" {
    name             = "my_vserver"
    analyticsprofile = "my_profile"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_analyticsprofile_binding.crvserver_analyticsprofile_binding.name
}

output "analytics_profile" {
  value = data.citrixadc_crvserver_analyticsprofile_binding.crvserver_analyticsprofile_binding.analyticsprofile
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the CR vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_analyticsprofile_binding. It is a system-generated identifier.
