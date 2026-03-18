---
subcategory: "Application Firewall"
---

# Data Source: appfwglobal_auditnslogpolicy_binding

The appfwglobal_auditnslogpolicy_binding data source allows you to retrieve information about the binding between appfwglobal configuration and auditnslogpolicy.

## Example Usage

```terraform
data "citrixadc_appfwglobal_auditnslogpolicy_binding" "tf_binding" {
  policyname = "my_auditnslogpolicy"
  type       = "NONE"
}

output "state" {
  value = data.citrixadc_appfwglobal_auditnslogpolicy_binding.tf_binding.state
}

output "policyname" {
  value = data.citrixadc_appfwglobal_auditnslogpolicy_binding.tf_binding.policyname
}
```

## Argument Reference

* `policyname` - (Required) Name of the policy.
* `type` - (Required) Bind point to which to policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is smaller than the current policy's priority number. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `id` - The id of the appfwglobal_auditnslogpolicy_binding. It is a system-generated identifier.
* `priority` - The priority of the policy.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows: * reqvserver. Invoke the unnamed policy label associated with the specified request virtual server. * policylabel. Invoke the specified user-defined policy label.
* `state` - Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.
