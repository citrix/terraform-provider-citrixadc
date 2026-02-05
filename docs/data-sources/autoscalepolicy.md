---
subcategory: "Autoscale"
---

# Data Source `autoscalepolicy`

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

## Attribute Reference

* `id` - The id of the autoscalepolicy. It has the same value as the `name` attribute.


## Import

An autoscalepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscalepolicy.tf_autoscalepolicy my_autoscalepolicy
```
