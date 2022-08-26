---
subcategory: "autoscale"
---

# Resource: autoscalepolicy

The `autoscalepolicy` resource is used to create autoscalepolicy.


## Example usage

```hcl
resource "citrixadc_autoscalepolicy" "policy1" {
    name = "policy1"
    rule = "true"
    action = "action1"
}
```


## Argument Reference

* `action` - (Optional) The autoscale profile associated with the policy.
* `comment` - (Optional) Comments associated with this autoscale policy.
* `logaction` - (Optional) The log action associated with the autoscale policy
* `name` - (Optional) The name of the autoscale policy.
* `newname` - (Optional) The new name of the autoscale policy.
* `rule` - (Optional) The rule associated with the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `autoscalepolicy`. It has the same value as the `name` attribute.


## Import

A `autoscalepolicy` can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
