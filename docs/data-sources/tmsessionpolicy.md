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

## Attribute Reference

* `id` - The id of the tmsessionpolicy. It has the same value as the `name` attribute.

## Import

A tmsessionpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tmsessionpolicy.tf_tmsessionpolicy my_tmsession_policy
```
