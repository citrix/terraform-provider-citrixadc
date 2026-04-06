---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_appflowpolicy_binding

The lbvserver_appflowpolicy_binding data source allows you to retrieve information about an AppFlow policy binding to a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_lbvserver_appflowpolicy_binding" "tf_lbvserver_appflowpolicy_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_appflowpolicy"
}

output "bindpoint" {
  value = data.citrixadc_lbvserver_appflowpolicy_binding.tf_lbvserver_appflowpolicy_binding.bindpoint
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_appflowpolicy_binding.tf_lbvserver_appflowpolicy_binding.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_appflowpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `priority` - Priority.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service.
