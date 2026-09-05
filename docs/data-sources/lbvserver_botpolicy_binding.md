---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_botpolicy_binding

The lbvserver_botpolicy_binding data source allows you to retrieve information about a bot policy binding to a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_lbvserver_botpolicy_binding" "tf_lbvserver_botpolicy_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_botpolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_botpolicy_binding.tf_lbvserver_botpolicy_binding.gotopriorityexpression
}

output "labeltype" {
  value = data.citrixadc_lbvserver_botpolicy_binding.tf_lbvserver_botpolicy_binding.labeltype
}

output "invoke" {
  value = data.citrixadc_lbvserver_botpolicy_binding.tf_lbvserver_botpolicy_binding.invoke
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_botpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Priority.
* `bindpoint` - The bindpoint to which the policy is bound.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
