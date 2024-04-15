---
subcategory: "Analytics"
---

# Resource: analyticsglobal_analyticsprofile_binding

The analyticsglobal_analyticsprofile_binding resource is used to bind analyticsprofile globally.


## Example usage

```hcl

resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
  name = "tf_analyticsprofile"
  type = "webinsight"
}

resource "citrixadc_analyticsglobal_analyticsprofile_binding" "tf_binding" {
  analyticsprofile = citrixadc_analyticsprofile.tf_analyticsprofile.name
}

```


## Argument Reference

* `analyticsprofile` - (Optional) Name of the analytics profile bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the analyticsglobal_analyticsprofile_binding. It has the same value as the `analyticsprofile` attribute.


## Import

A analyticsglobal_analyticsprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_analyticsglobal_analyticsprofile_binding.tf_binding tf_analyticsprofile
```
