---
subcategory: "Content Switching"
---

# Data Source: csvserver_appqoepolicy_binding

The csvserver_appqoepolicy_binding data source allows you to retrieve information about the binding between a content switching virtual server and an AppQoE policy.

## Example Usage

```terraform
data "citrixadc_csvserver_appqoepolicy_binding" "tf_binding" {
  name       = "tf_csvserver"
  policyname = "tf_appqoepolicy"
}

output "labelname" {
  value = data.citrixadc_csvserver_appqoepolicy_binding.tf_binding.labelname
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_appqoepolicy_binding.tf_binding.gotopriorityexpression
}

output "labeltype" {
  value = data.citrixadc_csvserver_appqoepolicy_binding.tf_binding.labeltype
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `priority` - Priority for the policy.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1. Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.

## Attribute Reference

* `id` - The id of the csvserver_appqoepolicy_binding. It is a system-generated identifier.
