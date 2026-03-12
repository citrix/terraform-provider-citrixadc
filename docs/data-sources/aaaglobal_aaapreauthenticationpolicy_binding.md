---
subcategory: "AAA"
---

# Data Source: aaaglobal_aaapreauthenticationpolicy_binding

The `aaaglobal_aaapreauthenticationpolicy_binding` data source allows you to retrieve information about a specific binding between the global AAA configuration and a preauthentication policy. This binding determines which preauthentication policies are applied globally and their priority order.

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

* `id` - The id of the aaaglobal_aaapreauthenticationpolicy_binding. It is a system-generated identifier.
* `priority` - Priority of the bound policy.
