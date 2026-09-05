---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_spilloverpolicy_binding

The lbvserver_spilloverpolicy_binding data source allows you to retrieve information about the binding between a load balancing virtual server and a spillover policy.


## Example Usage

```terraform
data "citrixadc_lbvserver_spilloverpolicy_binding" "tf_binding" {
  name       = "demo_lb"
  policyname = "demo_spilloverpolicy"
}

output "bindpoint" {
  value = data.citrixadc_lbvserver_spilloverpolicy_binding.tf_binding.bindpoint
}

output "priority" {
  value = data.citrixadc_lbvserver_spilloverpolicy_binding.tf_binding.priority
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_spilloverpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `bindpoint` - The bindpoint to which the policy is bound.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies.
* `priority` - Priority.
