---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_analyticsprofile_binding

The lbvserver_analyticsprofile_binding data source allows you to retrieve information about Analytics Profile bindings to lbvserver.


## Example usage

```terraform
data "citrixadc_lbvserver_analyticsprofile_binding" "tf_binding" {
  name = "test_server"
  analyticsprofile = "ns_analytics_global_profile"
}

output "name" {
  value = data.citrixadc_lbvserver_analyticsprofile_binding.tf_binding.name
}

output "analyticsprofile" {
  value = data.citrixadc_lbvserver_analyticsprofile_binding.tf_binding.analyticsprofile
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the LB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_analyticsprofile_binding. It is a system-generated identifier.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
