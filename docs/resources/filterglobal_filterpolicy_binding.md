---
subcategory: "Filter"
---

# Resource: filterglobal_filterpolicy_binding

The filterglobal_filterpolicy_binding resource is used to bind a filter policy globally.


## Example usage

```hcl
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}

resource "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
    policyname = citrixadc_filterpolicy.tf_filterpolicy.name
    priority = 200
    state = "ENABLED"
}
```


## Argument Reference

* `policyname` - (Required) The name of the filter policy.
* `priority` - (Optional) The priority of the policy.
* `state` - (Optional) State of the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the filterglobal_filterpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A filterglobal_filterpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal tf_filterpolicy
```
