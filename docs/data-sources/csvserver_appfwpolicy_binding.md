---
subcategory: "Content Switching"
---

# Data Source: csvserver_appfwpolicy_binding

The csvserver_appfwpolicy_binding data source allows you to retrieve information about the binding between a content switching virtual server and an application firewall policy.

## Example Usage

```terraform
data "citrixadc_csvserver_appfwpolicy_binding" "demo_binding" {
  name       = "demo_csvserver"
  policyname = "demo_appfwpolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_csvserver_appfwpolicy_binding.demo_binding.gotopriorityexpression
}

output "bindpoint" {
  value = data.citrixadc_csvserver_appfwpolicy_binding.demo_binding.bindpoint
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

* `id` - The id of the csvserver_appfwpolicy_binding. It is a system-generated identifier.
