---
subcategory: "Policy"
---

# Data Source: citrixadc_policyexpression

The policyexpression data source is used to retrieve information about an existing policy expression.

## Example usage

```terraform
data "citrixadc_policyexpression" "tf_policyexpression" {
  name = "tf_policyexpression"
}

output "policyexpression_id" {
  value = data.citrixadc_policyexpression.tf_policyexpression.id
}

output "policyexpression_value" {
  value = data.citrixadc_policyexpression.tf_policyexpression.value
}

output "policyexpression_comment" {
  value = data.citrixadc_policyexpression.tf_policyexpression.comment
}
```

## Argument Reference

* `name` - (Required) Unique name for the expression. Each expression name must be unique within its type.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the policy expression (combination of name and type).
* `clientsecuritymessage` - Message to display if the expression fails client security check.
* `comment` - Comments associated with the expression.
* `value` - Expression string that defines the policy expression logic.
