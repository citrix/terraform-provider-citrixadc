---
subcategory: "Traffic Management"
---

# Data Source: tmsessionpolicy

The tmsessionpolicy data source allows you to retrieve information about a TM session policy.

## Example usage

```terraform
data "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
  name = "my_tmsession_policy"
}

output "rule" {
  value = data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy.rule
}

output "action" {
  value = data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy.action
}
```

## Argument Reference

* `name` - (Required) Name for the session policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to be applied to connections that match this policy.
* `rule` - Expression, against which traffic is evaluated. Both classic and advance expressions are supported in default partition but only advance expressions in non-default partition.
* `id` - The id of the tmsessionpolicy. It has the same value as the `name` attribute.

### Read-only tmsessionpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_tmsessionpolicy` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `builtin` - Indicates that a variable is a built-in (SYSTEM INTERNAL) type (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this configuration.
* `expressiontype` - Type of policy (`Classic Policy` / `Advanced Policy`).
* `hits` - Number of hits.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
