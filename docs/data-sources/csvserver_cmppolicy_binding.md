---
subcategory: "Content Switching"
---

# Data Source: csvserver_cmppolicy_binding

The csvserver_cmppolicy_binding data source allows you to retrieve information about bindings between a content switching vserver and compression policy.

## Example usage

```terraform
data "citrixadc_csvserver_cmppolicy_binding" "tf_bind" {
  name       = "tf_csvserver"
  policyname = "tf_cmppolicy"
  bindpoint  = "REQUEST"
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_cmppolicy_binding.tf_bind.gotopriorityexpression
}

output "targetlbvserver" {
  value = data.citrixadc_csvserver_cmppolicy_binding.tf_bind.targetlbvserver
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
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.

## Attribute Reference

* `id` - The id of the csvserver_cmppolicy_binding. It is a system-generated identifier.
