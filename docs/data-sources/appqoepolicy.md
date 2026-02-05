---
subcategory: "Appqoe"
---

# Data Source: citrixadc_appqoepolicy

The appqoepolicy data source is used to retrieve information about an existing AppQoE (Application Quality of Experience) policy configured on the Citrix ADC.

## Example usage

```hcl

# Retrieve the appqoepolicy data
data "citrixadc_appqoepolicy" "tf_appqoepolicy" {
  name = citrixadc_appqoepolicy.tf_appqoepolicy.name
}

# Reference the datasource attributes in other resources
output "policy_rule" {
  value = data.citrixadc_appqoepolicy.tf_appqoepolicy.rule
}

output "policy_action" {
  value = data.citrixadc_appqoepolicy.tf_appqoepolicy.action
}

# Use in conditional logic
locals {
  is_always_true_policy = data.citrixadc_appqoepolicy.tf_appqoepolicy.rule == "true"
}
```

## Argument Reference

* `name` - (Required) Name of the AppQoE policy to retrieve.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The id of the appqoepolicy. It has the same value as the `name` attribute.
* `action` - The configured AppQoE action to trigger when the policy rule evaluates to true.
* `rule` - Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.
