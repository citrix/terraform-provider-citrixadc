---
subcategory: "Policy"
---

# Data Source: policyexpression

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

### Read-only policyexpression metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_policyexpression` resource). They are Computed / GET-only, and any attribute the appliance does not return is `null`.

* `hits` - The total number of hits.
* `pihits` - The total number of hits.
* `type1` - The type of expression. This is for output only. Possible values: `CLASSIC`, `ADVANCED`.
* `isdefault` - A value of true is returned if it is a default policy expression.
* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type. A list of strings. Possible values: `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`.
* `feature` - The feature to be checked while applying this config.
