---
subcategory: "AAA"
---

# Data Source: aaagroup_vpnsessionpolicy_binding

The aaagroup_vpnsessionpolicy_binding data source allows you to retrieve information about a specific binding between an AAA group and a VPN session policy.

## Example Usage

```terraform
data "citrixadc_aaagroup_vpnsessionpolicy_binding" "example" {
  groupname = "my_group"
  policy    = "my_vpnsession_policy"
}

output "priority" {
  value = data.citrixadc_aaagroup_vpnsessionpolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_aaagroup_vpnsessionpolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `policy` - (Required) The policy name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the aaagroup_vpnsessionpolicy_binding. It is the concatenation of `groupname` and `policy` attributes separated by a comma.
* `type` - Bindpoint to which the policy is bound.
* `priority` - Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.

### Read-only aaagroup_vpnsessionpolicy_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_aaagroup_vpnsessionpolicy_binding` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `acttype` - Action type of the binding. Read-only value returned by the appliance.
