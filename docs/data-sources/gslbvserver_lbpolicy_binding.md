---
subcategory: "GSLB"
---

# Data Source: gslbvserver_lbpolicy_binding

The gslbvserver_lbpolicy_binding data source allows you to retrieve information about GSLB vserver load balancing policy bindings.


## Example Usage

```terraform
data "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
  name       = "tf_gslbvserver"
  policyname = "tf_pol"
}

output "priority" {
  value = data.citrixadc_gslbvserver_lbpolicy_binding.tf_bind.priority
}

output "type" {
  value = data.citrixadc_gslbvserver_lbpolicy_binding.tf_bind.type
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `policyname` - (Required) Name of the policy bound to the GSLB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_lbpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. o If gotoPriorityExpression is not present or if it is equal to END then the policy bank evaluation ends here o Else if the gotoPriorityExpression is equal to NEXT then the next policy in the priority order is evaluated. o Else gotoPriorityExpression is evaluated. The result of gotoPriorityExpression (which has to be a number) is processed as follows: - An UNDEF event is triggered if . gotoPriorityExpression cannot be evaluated . gotoPriorityExpression evaluates to number which is smaller than the maximum priority in the policy bank but is not same as any policy's priority . gotoPriorityExpression evaluates to a priority that is smaller than the current policy's priority - If the gotoPriorityExpression evaluates to the priority of the current policy then the next policy in the priority order is evaluated. - If the gotoPriorityExpression evaluates to the priority of a policy further ahead in the list then that policy will be evaluated next. This field is applicable only to rewrite and responder policies.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `priority` - Priority.
* `type` - The bindpoint to which the policy is bound

