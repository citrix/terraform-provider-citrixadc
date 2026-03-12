---
subcategory: "Content Switching"
---

# Data Source: csvserver_contentinspectionpolicy_binding

The csvserver_contentinspectionpolicy_binding data source allows you to retrieve information about bindings between a content switching vserver and content inspection policy.

## Example usage

```terraform
data "citrixadc_csvserver_contentinspectionpolicy_binding" "tf_csvserver_contentinspectionpolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_contentinspectionpolicy"
  bindpoint  = "REQUEST"
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_contentinspectionpolicy_binding.tf_csvserver_contentinspectionpolicy_binding.gotopriorityexpression
}

output "targetlbvserver" {
  value = data.citrixadc_csvserver_contentinspectionpolicy_binding.tf_csvserver_contentinspectionpolicy_binding.targetlbvserver
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

* `id` - The id of the csvserver_contentinspectionpolicy_binding. It is a system-generated identifier.
