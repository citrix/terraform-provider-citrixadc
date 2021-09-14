---
subcategory: "Content Switching"
---

# Resource: csvserver_analyticsprofile_binding

The csvserver_analyticsprofile_binding resource is used to bound Analytics Profile to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	analyticsprofile = "ns_analytics_global_profile"
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "1.1.1.2"
	port = 80
	servicetype = "HTTP"
}
```


## Argument Reference

* `analyticsprofile` - (Required) Name of the analytics profile bound to the LB vserver.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_analyticsprofile_binding. It is the concatenation of the `name` and `analyticsprofile` attributes separated by a comma.


## Import

A csvserver_analyticsprofile_binding can be imported using its name, e.g.

```shellterraform import terraform import citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding tf_csvserver,ns_analytics_global_profile
```
