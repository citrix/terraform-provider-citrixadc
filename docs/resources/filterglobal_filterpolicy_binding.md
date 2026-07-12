---
subcategory: "Filter"
---

# Resource: filterglobal_filterpolicy_binding

The filterglobal_filterpolicy_binding resource is used to bind a filter policy to global, in order to apply it to the entire traffic handled by the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_filterglobal_filterpolicy_binding" "tf_filterglobal" {
  policyname = citrixadc_filterpolicy.tf_filterpolicy.name
  priority   = 200
  state      = "ENABLED"
}

resource "citrixadc_filterpolicy" "tf_filterpolicy" {
  name      = "tf_filterpolicy"
  reqaction = "DROP"
  rule      = "REQ.HTTP.URL CONTAINS http://abcd.com"
}
```


## Argument Reference

* `policyname` - (Required) Name of the filter policy.
* `priority` - (Optional) Specifies the priority of the policy.
* `state` - (Optional) State of the binding. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the filterglobal_filterpolicy_binding. It has the format `policyname:<policyname>`.


## Import

A filterglobal_filterpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_filterglobal_filterpolicy_binding.tf_filterglobal policyname:tf_filterpolicy
```
