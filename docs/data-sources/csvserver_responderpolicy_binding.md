---
subcategory: "Content Switching"
---

# Data Source: csvserver_responderpolicy_binding

The csvserver_responderpolicy_binding data source allows you to retrieve information about a csvserver_responderpolicy_binding.

## Example Usage

```terraform
data "citrixadc_csvserver_responderpolicy_binding" "tf_bind" {
  name       = "tf_csvserver"
  policyname = "tf_responder_policy"
}

output "name" {
  value = data.citrixadc_csvserver_responderpolicy_binding.tf_bind.name
}

output "priority" {
  value = data.citrixadc_csvserver_responderpolicy_binding.tf_bind.priority
}

output "bindpoint" {
  value = data.citrixadc_csvserver_responderpolicy_binding.tf_bind.bindpoint
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_responderpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `priority` - Priority for the policy.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
