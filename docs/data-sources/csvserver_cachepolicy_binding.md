---
subcategory: "Content Switching"
---

# Data Source: csvserver_cachepolicy_binding

The csvserver_cachepolicy_binding data source allows you to retrieve information about a cache policy binding to a content switching virtual server.


## Example Usage

```terraform
data "citrixadc_csvserver_cachepolicy_binding" "tf_csvserver_cachepolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_cachepolicy"
  bindpoint  = "REQUEST"
}

output "bindpoint" {
  value = data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding.bindpoint
}

output "priority" {
  value = data.citrixadc_csvserver_cachepolicy_binding.tf_csvserver_cachepolicy_binding.priority
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
* `labelname` - Name of the label invoked.
* `priority` - Priority for the policy.
* `labeltype` - The invocation type.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
* `id` - The id of the csvserver_cachepolicy_binding. It is a system-generated identifier.
