---
subcategory: "Content Switching"
---

# Data Source: csvserver_spilloverpolicy_binding

The csvserver_spilloverpolicy_binding data source allows you to retrieve information about a csvserver spilloverpolicy binding.


## Example usage

```terraform
data "citrixadc_csvserver_spilloverpolicy_binding" "tf_csvserver_spilloverpolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_spilloverpolicy"
}

output "bindpoint" {
  value = data.citrixadc_csvserver_spilloverpolicy_binding.tf_csvserver_spilloverpolicy_binding.bindpoint
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_spilloverpolicy_binding.tf_csvserver_spilloverpolicy_binding.gotopriorityexpression
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_spilloverpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - Name of the label to be invoked.
* `priority` - Priority for the policy.
* `labeltype` - Type of label to be invoked.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.
