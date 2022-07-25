---
subcategory: "CMP"
---

# Resource: cmpglobal_cmppolicy_binding

The cmpglobal_cmppolicy_binding resource is used to bind a compression policy globally.


## Example usage

```hcl
resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_cmpglobal_binding" "tf_cmpglobal_binding" {
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 100
}
```


## Argument Reference

* `policyname` - (Required) Name of the HTTP compression policy to bind globally.
* `priority` - (Required) The priority for the policy binding.
* `state` - (Optional) The current state of the policy binding. This attribute is relevant only for CLASSIC policies.
Possible values = `ENABLED`, `DISABLED`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the compression policy. It has the same value as the `policyname` attribute.


## Import

A cmpglobal_cmppolicy_binding can be imported using its policyname, e.g.

```shell
terraform import cmpglobal_cmppolicy_binding.tf_cmpglobal_binding tf_cmppolicy
```
