---
subcategory: "Cloud"
---

# Resource: cloudparaminternal

This resource is used to manage internal cloud parameters on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_cloudparaminternal" "tf_cloudparaminternal" {
  nonftumode = "YES"
}
```


## Argument Reference

* `nonftumode` - (Optional) Indicates whether the management GUI is in first-time-user (FTU) mode or not. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudparaminternal. It is set to `cloudparaminternal-config`.


## Import

A cloudparaminternal can be imported using its id, e.g.

```shell
terraform import citrixadc_cloudparaminternal.tf_cloudparaminternal cloudparaminternal-config
```
