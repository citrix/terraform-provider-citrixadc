---
subcategory: "Content Switching"
---

# Data Source: csvserver_rewritepolicy_binding

The csvserver_rewritepolicy_binding data source allows you to retrieve information about a rewrite policy binding to a content switching virtual server.


## Example Usage

```terraform
data "citrixadc_csvserver_rewritepolicy_binding" "tf_bind" {
  name       = "tf_csvserver"
  policyname = "tf_test_rewrite_policy"
  bindpoint  = "REQUEST"
}

output "bindpoint" {
  value = data.citrixadc_csvserver_rewritepolicy_binding.tf_bind.bindpoint
}

output "priority" {
  value = data.citrixadc_csvserver_rewritepolicy_binding.tf_bind.priority
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Required) The bindpoint to which the policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `priority` - Priority for the policy.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
* `id` - The id of the csvserver_rewritepolicy_binding. It is a system-generated identifier.
