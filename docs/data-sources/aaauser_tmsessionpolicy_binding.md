---
subcategory: "AAA"
---

# Data Source: aaauser_tmsessionpolicy_binding

The aaauser_tmsessionpolicy_binding data source allows you to retrieve information about an aaauser_tmsessionpolicy_binding.


## Example Usage

```terraform
data "citrixadc_aaauser_tmsessionpolicy_binding" "tf_aaauser_tmsessionpolicy_binding" {
  username = "user1"
  policy   = "my_tmsession_policy"
}

output "username" {
  value = data.citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding.username
}

output "priority" {
  value = data.citrixadc_aaauser_tmsessionpolicy_binding.tf_aaauser_tmsessionpolicy_binding.priority
}
```


## Argument Reference

* `username` - (Required) User account to which to bind the policy.
* `policy` - (Required) The policy Name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaauser_tmsessionpolicy_binding. It is a system-generated identifier.
* `type` - Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000.
