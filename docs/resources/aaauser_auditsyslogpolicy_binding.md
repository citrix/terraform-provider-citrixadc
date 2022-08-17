---
subcategory: "AAA"
---

# Resource: aaauser_auditsyslogpolicy_binding

The aaauser_auditsyslogpolicy_binding resource is used to create aaauser_auditsyslogpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaauser_auditsyslogpolicy_binding" "tf_aaauser_auditsyslogpolicy_binding" {
  username = "user1"
  policy    = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
  priority  = 100
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "tf_syslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}
```


## Argument Reference

* `username` - (Required) User account to which to bind the policy. Minimum length =  1
* `policy` - (Required) The policy Name.
* `priority` - (Required) Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000. . Minimum value =  0 Maximum value =  2147483647
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: *  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaauser_auditsyslogpolicy_binding. It is the concatenation of  `username` and `policy` attributes separated by a comma.


## Import

A aaauser_auditsyslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaauser_auditsyslogpolicy_binding.tf_aaauser_auditsyslogpolicy_binding user1,tf_auditsyslogpolicy
```
