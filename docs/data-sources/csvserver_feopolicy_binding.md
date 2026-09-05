---
subcategory: "Content Switching"
---

# Data Source: csvserver_feopolicy_binding

The csvserver_feopolicy_binding data source allows you to retrieve information about a binding between a content switching virtual server and a Front End Optimization (FEO) policy.


## Example usage

```terraform
data "citrixadc_csvserver_feopolicy_binding" "tf_bind" {
  name       = "tf_csvserver"
  policyname = "tf_feopolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_feopolicy_binding.tf_bind.gotopriorityexpression
}

output "targetlbvserver" {
  value = data.citrixadc_csvserver_feopolicy_binding.tf_bind.targetlbvserver
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_feopolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `bindpoint` - The bindpoint to which the policy is bound.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - Name of the label to be invoked.
* `labeltype` - Type of label to be invoked.
* `priority` - Priority for the policy.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
