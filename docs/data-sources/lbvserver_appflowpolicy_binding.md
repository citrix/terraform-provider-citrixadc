---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_appflowpolicy_binding

The lbvserver_appflowpolicy_binding data source allows you to retrieve information about an AppFlow policy binding to a load balancing virtual server.

## Example usage

```terraform
data "citrixadc_lbvserver_appflowpolicy_binding" "tf_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_appflowpolicy"
}

output "bindpoint" {
  value = data.citrixadc_lbvserver_appflowpolicy_binding.tf_binding.bindpoint
}

output "priority" {
  value = data.citrixadc_lbvserver_appflowpolicy_binding.tf_binding.priority
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_appflowpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `bindpoint` - The bindpoint to which the policy is bound.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service.
* `priority` - Priority.
