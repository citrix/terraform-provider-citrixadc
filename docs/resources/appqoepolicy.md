---
subcategory: "Appqoe"
---

# Resource: appqoepolicy

The appqoepolicy resource is used to create appqoepolicy.


## Example usage

```hcl
resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
  name   = "my_appqoepolicy"
  rule   = "true"
  action = "my_act"
}
```


## Argument Reference

* `action` - (Required) Configured AppQoE action to trigger
* `name` - (Required) Name of AppQoEpolicy
* `rule` - (Required) Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appqoepolicy. It has the same value as the `name` attribute.


## Import

A appqoepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_appqoepolicy.tf_appqoepolicy my_appqoepolicy
```
