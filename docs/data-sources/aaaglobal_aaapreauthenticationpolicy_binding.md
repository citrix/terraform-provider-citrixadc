---
subcategory: "AAA"
---

# Data Source: aaaglobal_aaapreauthenticationpolicy_binding

The aaaglobal_aaapreauthenticationpolicy_binding data source allows you to retrieve information about the binding between the global AAA configuration and a preauthentication policy.


## Example usage

```terraform
data "citrixadc_aaaglobal_aaapreauthenticationpolicy_binding" "tf_aaaglobal_aaapreauthenticationpolicy_binding" {
  policy = "my_policy"
}

output "policy_name" {
  value = data.citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_aaaglobal_aaapreauthenticationpolicy_binding.policy
}

output "policy_priority" {
  value = data.citrixadc_aaaglobal_aaapreauthenticationpolicy_binding.tf_aaaglobal_aaapreauthenticationpolicy_binding.priority
}
```

## Argument Reference

* `policy` - (Required) Name of the policy to be unbound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaglobal_aaapreauthenticationpolicy_binding. It has the same value as the `policy` attribute.
* `priority` - Priority of the bound policy.

### Read-only aaaglobal_aaapreauthenticationpolicy_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_aaaglobal_aaapreauthenticationpolicy_binding` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `bindpolicytype` - Bound policy type.
