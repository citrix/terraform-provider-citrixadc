---
subcategory: "Autoscale"
---

# Data Source: autoscalepolicy

The autoscalepolicy data source allows you to retrieve information about autoscale policies.


## Example usage

```terraform
data "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
  name = "my_autoscalepolicy"
}

output "rule" {
  value = data.citrixadc_autoscalepolicy.tf_autoscalepolicy.rule
}

output "action" {
  value = data.citrixadc_autoscalepolicy.tf_autoscalepolicy.action
}
```


## Argument Reference

* `name` - (Required) The name of the autoscale policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - The autoscale profile associated with the policy.
* `comment` - Comments associated with this autoscale policy.
* `logaction` - The log action associated with the autoscale policy.
* `newname` - The new name of the autoscale policy.
* `rule` - The rule associated with the policy.
* `id` - The id of the autoscalepolicy. It has the same value as the `name` attribute.

### Read-only autoscalepolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_autoscalepolicy` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `hits` - Number of hits.
* `undefhits` - Number of Undef hits.
* `priority` - Specifies the priority of the policy.
