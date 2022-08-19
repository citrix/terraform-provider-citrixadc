---
subcategory: "AAA"
---

# Resource: aaagroup_auditnslogpolicy_binding

The aaagroup_auditnslogpolicy_binding resource is used to create aaagroup_auditnslogpolicy_binding.


## Example usage

```hcl
# Since the auditnslogpolicy resource is not yet available on Terraform,
# the tf_auditnslogpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add audit nslogAction tf_auditnslogaction 1.1.1.1 -loglevel NONE
# add audit nslogPolicy tf_auditnslogpolicy ns_true tf_auditnslogaction

resource "citrixadc_aaagroup_auditnslogpolicy_binding" "tf_aaagroup_auditnslogpolicy_binding" {
  groupname = "my_group"
  policy    = "tf_auditnslogpolicy"
  priority  = 150
}
```


## Argument Reference

* `policy` - (Required) The policy name.
* `priority` - (Required) Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Minimum value =  0 Maximum value =  2147483647
* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: *  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_auditnslogpolicy_binding. It is the concatenation of  `groupname` and `policy` attributes separated by a comma.


## Import

A aaagroup_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_auditnslogpolicy_binding.tf_aaagroup_auditnslogpolicy_binding my_group,tf_auditnslogpolicy
```
