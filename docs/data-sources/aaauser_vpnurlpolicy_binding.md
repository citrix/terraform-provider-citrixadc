---
subcategory: "AAA"
---

# Data Source: aaauser_vpnurlpolicy_binding

The aaauser_vpnurlpolicy_binding data source allows you to retrieve information about the binding between AAA users and VPN URL policies.

## Example Usage

```terraform
data "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
  username = "user1"
  policy   = "new_policy"
}

output "priority" {
  value = data.citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaauser_vpnurlpolicy_binding.tf_aaauser_vpnurlpolicy_binding.gotopriorityexpression
}
```

## Argument Reference

* `username` - (Required) User account to which the policy is bound.
* `policy` - (Required) The policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values:
  * NEXT - Evaluate the policy with the next higher priority number.
  * END - End policy evaluation.
  * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.
  * An expression that evaluates to a number.
* `type` - Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.
* `id` - The id of the aaauser_vpnurlpolicy_binding. It is a system-generated identifier.
