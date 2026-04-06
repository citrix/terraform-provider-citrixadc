---
subcategory: "GSLB"
---

# Data Source: gslbvserver_spilloverpolicy_binding

The gslbvserver_spilloverpolicy_binding data source allows you to retrieve information about a spillover policy binding to a GSLB virtual server.

## Example Usage

```terraform
data "citrixadc_gslbvserver_spilloverpolicy_binding" "tf_gslbvserver_spilloverpolicy_binding" {
  name       = "gslb_vserver"
  policyname = "tf_spilloverpolicy"
}

output "priority" {
  value = data.citrixadc_gslbvserver_spilloverpolicy_binding.tf_gslbvserver_spilloverpolicy_binding.priority
}

output "type" {
  value = data.citrixadc_gslbvserver_spilloverpolicy_binding.tf_gslbvserver_spilloverpolicy_binding.type
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `policyname` - (Required) Name of the policy bound to the GSLB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_spilloverpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `priority` - Priority.
* `type` - The bindpoint to which the policy is bound.
