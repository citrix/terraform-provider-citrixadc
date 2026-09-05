---
subcategory: "Content Switching"
---

# Data Source: csvserver_tmtrafficpolicy_binding

This data source is used to retrieve information about a specific `csvserver_tmtrafficpolicy_binding` configuration.

## Example usage

```hcl
data "citrixadc_csvserver_tmtrafficpolicy_binding" "example" {
  name       = "tf_csvserver"
  policyname = "tf_tmttrafficpolicy"
}

output "binding_id" {
  value = data.citrixadc_csvserver_tmtrafficpolicy_binding.example.id
}

output "binding_priority" {
  value = data.citrixadc_csvserver_tmtrafficpolicy_binding.example.priority
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `priority` - Priority for the policy.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `bindpoint` - Bind point at which the policy needs to be bound.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labeltype` - Type of label to be invoked.
* `labelname` - Name of the label to be invoked.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if the policy rule is evaluated to be TRUE.
