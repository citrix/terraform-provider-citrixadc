---
subcategory: "Content Switching"
---

# Data Source: csvserver_authorizationpolicy_binding

The csvserver_authorizationpolicy_binding data source allows you to retrieve information about a csvserver_authorizationpolicy_binding.


## Example Usage

```terraform
data "citrixadc_csvserver_authorizationpolicy_binding" "tf_csvserver_authorizationpolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_authorizationpolicy"
}

output "name" {
  value = data.citrixadc_csvserver_authorizationpolicy_binding.tf_csvserver_authorizationpolicy_binding.name
}

output "policyname" {
  value = data.citrixadc_csvserver_authorizationpolicy_binding.tf_csvserver_authorizationpolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_csvserver_authorizationpolicy_binding.tf_csvserver_authorizationpolicy_binding.priority
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the csvserver_authorizationpolicy_binding. It is a system-generated identifier.
* `priority` - Priority for the policy.
* `invoke` - Invoke flag.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.

