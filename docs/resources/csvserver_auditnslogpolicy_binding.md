---
subcategory: "Content Switching"
---

# Resource: csvserver_auditnslogpolicy_binding

The csvserver_auditnslogpolicy_binding resource is used to bind an audit nslog policy to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
        name = "tf_csvserver"
        policyname = "tf_auditnslogpolicy"
        priority = 5
}
```


## Argument Reference

* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE (valid only for default-syntax policies such as application firewall, transform, integrated cache, rewrite, responder, and content switching).
* `labeltype` - (Optional) Type of label to be invoked. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label to be invoked.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_auditnslogpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding tf_csvserver,tf_auditnslogpolicy
```
