---
subcategory: "Traffic Management"
---

# Data Source: tmtrafficpolicy

The tmtrafficpolicy data source allows you to retrieve information about a TM traffic policy.

## Example usage

```terraform
data "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
  name = "my_tmtraffic_policy"
}

output "rule" {
  value = data.citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.rule
}

output "action" {
  value = data.citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.action
}
```

## Argument Reference

* `name` - (Required) Name for the traffic policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the action to apply to requests or connections that match this policy.
* `rule` - Name of the Citrix ADC named expression, or an expression, that the policy uses to determine whether to apply certain action on the current traffic.

## Attribute Reference

* `id` - The id of the tmtrafficpolicy. It has the same value as the `name` attribute.

## Import

A tmtrafficpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy my_tmtraffic_policy
```
