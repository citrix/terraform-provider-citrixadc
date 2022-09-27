---
subcategory: "Autoscale"
---

# Resource: autoscalepolicy

The autoscalepolicy resource is used to create autoscalepolicy.


## Example usage

```hcl
resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
  name         = "my_autoscaleprofile"
  rule         = "true"
  action       = "my_autoscaleaction"
}
```


## Argument Reference

* `action` - (Required) The autoscale profile associated with the policy.
* `name` - (Reuqired) The name of the autoscale policy.
* `rule` - (Required) The rule associated with the policy.
* `comment` - (Optional) Comments associated with this autoscale policy.
* `logaction` - (Optional) The log action associated with the autoscale policy


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the autoscalepolicy. It has the same value as the `name` attribute.


## Import

A autoscalepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscalepolicy.tf_autoscalepolicy my_autoscaleprofile
```
