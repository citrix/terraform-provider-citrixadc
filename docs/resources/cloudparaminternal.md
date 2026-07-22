---
subcategory: "Cloud"
---

# Resource: cloudparaminternal

Configures internal cloud parameters on the Citrix ADC. This resource controls the appliance's first-time-user (FTU) mode indicator, which the management GUI consults to decide whether to present the initial cloud-onboarding setup workflow to the administrator.

This is a singleton resource: a single `cloudparaminternal` configuration always exists on the appliance. Creating this resource updates the existing global configuration rather than adding a new object, so there is no delete operation and no name key.

Note: The corresponding NITRO GET/show operation (`show cloud paramInternal`) is platform-gated. On platforms that do not support it, NITRO returns "Operation not supported on this platform". The provider treats this as a non-fatal read: it preserves the value you configured in Terraform state rather than failing the apply. On those platforms the provider cannot read the value back from the ADC.


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
