---
subcategory: "Front-end-optimization"
---

# Resource: feopolicy

The feopolicy resource is used to createfeopolicy.


## Example usage

```hcl
resource "citrixadc_feopolicy" "tf_feopolicy" {
  name   = "my_feopolicy"
  action = "my_feoaction"
  rule   = "true"
}

```


## Argument Reference

* `name` - (Required) The name of the front end optimization policy. Minimum length =  1
* `rule` - (Required) The rule associated with the front end optimization policy.
* `action` - (Required) The front end optimization action that has to be performed when the rule matches. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the feopolicy. It has the same value as the `name` attribute.


## Import

A feopolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
