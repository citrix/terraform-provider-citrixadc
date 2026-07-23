---
subcategory: "NS"
---

# Resource: nschannelparam

This resource is used to manage global channel (link aggregation) parameters on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_nschannelparam" "tf_nschannelparam" {
  vfautorecover = "ENABLE"
}
```


## Argument Reference

* `vfautorecover` - (Required) VF autorecover mode for channels. Possible values: [ DISABLE, ENABLE ]. Defaults to `ENABLE`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nschannelparam. It is set to `nschannelparam-config`.


## Import

A singleton resource is imported using the constant id `nschannelparam-config`:

```shell
terraform import citrixadc_nschannelparam.tf_nschannelparam nschannelparam-config
```
