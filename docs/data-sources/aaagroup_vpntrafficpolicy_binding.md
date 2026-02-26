---
subcategory: "AAA"
---

# Data Source: aaagroup_vpntrafficpolicy_binding

The aaagroup_vpntrafficpolicy_binding data source allows you to retrieve information about a specific binding between an AAA group and a VPN traffic policy.

## Example Usage

```terraform
data "citrixadc_aaagroup_vpntrafficpolicy_binding" "example" {
  groupname = "my_group"
  policy    = "tf_vpntrafficpolicy"
}

output "priority" {
  value = data.citrixadc_aaagroup_vpntrafficpolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_vpntrafficpolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `policy` - (Required) The policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `id` - The id of the aaagroup_vpntrafficpolicy_binding. It is a system-generated identifier.
* `type` - Bindpoint to which the policy is bound.
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
