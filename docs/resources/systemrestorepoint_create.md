---
subcategory: "System"
---

# Resource: systemrestorepoint_create

This resource is used to create a restore point on a Citrix ADC.


## Example usage

```hcl
resource "citrixadc_systemrestorepoint_create" "tf_systemrestorepoint_create" {
  filename = "pre-upgrade-restorepoint"
}
```


## Argument Reference

* `filename` - (Required) Name of the restore point. Changing this value forces a
  new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The identifier of the restore point. It has the same value as the
  `filename` attribute.


## Import

A systemrestorepoint_create can be imported using its filename, e.g.

```shell
terraform import citrixadc_systemrestorepoint_create.tf_systemrestorepoint_create pre-upgrade-restorepoint
```
