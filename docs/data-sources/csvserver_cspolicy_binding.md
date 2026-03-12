---
subcategory: "Content Switching"
---

# Data Source: csvserver_cspolicy_binding

The csvserver_cspolicy_binding data source allows you to retrieve information about a specific binding between a content switching virtual server and a content switching policy.

## Example Usage

```terraform
data "citrixadc_csvserver_cspolicy_binding" "example" {
  name       = "tf_csvserver"
  policyname = "tf_cspolicy"
}

output "priority" {
  value = data.citrixadc_csvserver_cspolicy_binding.example.priority
}

output "targetlbvserver" {
  value = data.citrixadc_csvserver_cspolicy_binding.example.targetlbvserver
}

output "bindpoint" {
  value = data.citrixadc_csvserver_cspolicy_binding.example.bindpoint
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver. When specified, the data source will filter the results to match this policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_cspolicy_binding. It is a system-generated identifier.
* `priority` - Priority for the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetlbvserver` - Target vserver name.
