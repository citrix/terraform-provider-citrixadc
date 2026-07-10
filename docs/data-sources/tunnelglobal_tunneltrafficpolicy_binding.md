---
subcategory: "Tunnel"
---

# Data Source: tunnelglobal_tunneltrafficpolicy_binding

The `citrixadc_tunnelglobal_tunneltrafficpolicy_binding` data source allows you to retrieve information about a specific binding between the global tunnel configuration and a tunnel traffic policy.


## Example usage

```terraform
data "citrixadc_tunnelglobal_tunneltrafficpolicy_binding" "tf_tunnelglobal_tunneltrafficpolicy_binding" {
  policyname = "my_tunneltrafficpolicy"
}

output "policy_priority" {
  value = data.citrixadc_tunnelglobal_tunneltrafficpolicy_binding.tf_tunnelglobal_tunneltrafficpolicy_binding.priority
}

output "policy_state" {
  value = data.citrixadc_tunnelglobal_tunneltrafficpolicy_binding.tf_tunnelglobal_tunneltrafficpolicy_binding.state
}
```


## Argument Reference

* `policyname` - (Required) Policy name.
* `type` - (Optional) Bind point to which the policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, NONE ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tunnelglobal_tunneltrafficpolicy_binding. It is the concatenation of `policyname` and `type` attributes separated by a comma.
* `feature` - The feature to be checked while applying this config.
* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Priority.
* `state` - Current state of the binding. If the binding is enabled, the policy is active.
