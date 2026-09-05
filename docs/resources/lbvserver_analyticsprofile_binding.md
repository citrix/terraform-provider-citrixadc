---
subcategory: "Load Balancing"
---

# Resource: lbvserver_analyticsprofile_binding

The lbvserver_analyticsprofile_binding resource is used to bound Analytics Profile to lbvserver.


## Example usage

```hcl
resource "citrixadc_lbvserver_analyticsprofile_binding" "demo_bind" {
  name             = "test-server"
  analyticsprofile = "ns_analytics_global_profile"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `analyticsprofile` - (Required) Name of the analytics profile bound to the LB vserver.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_analyticsprofile_binding. It is the concatenation of the `name` and `analyticsprofile` attributes separated by a comma.


## Import

A lbvserver_analyticsprofile_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_analyticsprofile_binding.tf_lbvserver_analyticsprofile_binding test-server,ns_analytics_global_profile
```
