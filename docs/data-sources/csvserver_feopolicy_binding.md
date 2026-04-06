---
subcategory: "Content Switching"
---

# Data Source: csvserver_feopolicy_binding

The `csvserver_feopolicy_binding` data source allows you to retrieve information about a specific binding between a Content Switching virtual server and a Front-End Optimization (FEO) policy.


## Example usage

```terraform
data "citrixadc_csvserver_feopolicy_binding" "tf_csvserver_feopolicy_binding" {
  name       = "my_csvserver"
  policyname = "my_feopolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding.gotopriorityexpression
}

output "targetlbvserver" {
  value = data.citrixadc_csvserver_feopolicy_binding.tf_csvserver_feopolicy_binding.targetlbvserver
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Name of the policy bound to the virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the csvserver_feopolicy_binding. It is a system-generated identifier.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - Name of the label to be invoked.
* `priority` - Priority for the policy.
* `labeltype` - Type of label to be invoked.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
