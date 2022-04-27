---
subcategory: "Application Firewall"
---

# Resource: appfwglobal_auditsyslogpolicy_binding

The appfwglobal_auditsyslogpolicy_binding resource is used to bind auditsyslogpolicy to appfwglobal configuration.


## Example usage

```hcl
resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "tf_syslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
resource "citrixadc_auditsyslogpolicy" "tf_policy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}
resource "citrixadc_appfwglobal_auditsyslogpolicy_binding" "tf_binding" {
  policyname = citrixadc_auditsyslogpolicy.tf_policy.name
  priority   = 90
  state      = "DISABLED"
  type       = "NONE"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.  An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is smaller than the current policy's priority number. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - (Optional) Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows: * reqvserver. Invoke the unnamed policy label associated with the specified request virtual server. * policylabel. Invoke the specified user-defined policy label.
* `priority` - (Optional) The priority of the policy.
* `state` - (Optional) Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.
* `type` - (Optional) Bind point to which to policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwglobal_auditsyslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A appfwglobal_auditsyslogpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_appfwglobal_auditsyslogpolicy_binding.tf_binding tf_auditsyslogpolicy
```
