---
subcategory: "NS"
---

# Resource: nschannelparam

Configures the global channel (link aggregation / VF) parameters on the Citrix ADC. The `vfautorecover` setting controls whether channels automatically recover their virtual function (VF) interfaces - that is, whether member interfaces are brought back into an aggregated channel automatically once they become available again, rather than requiring manual intervention.

This is a singleton resource: a single channel parameter configuration always exists on the appliance, so this resource has no create or delete operation on the ADC - applying it updates the existing configuration, and destroying it only removes the resource from Terraform state.


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
