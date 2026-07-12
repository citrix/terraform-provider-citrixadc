---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_filterpolicy_binding

The lbvserver_filterpolicy_binding data source allows you to retrieve information about the binding between a load balancing virtual server and a filter policy.


## Example Usage

```terraform
data "citrixadc_lbvserver_filterpolicy_binding" "tf_bind" {
  name       = "tf_lbvserver"
  policyname = "tf_filterpolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_filterpolicy_binding.tf_bind.gotopriorityexpression
}

output "priority" {
  value = data.citrixadc_lbvserver_filterpolicy_binding.tf_bind.priority
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_filterpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `priority` - Priority.
