---
subcategory: "NS"
---

# Resource: nsmigration_stop

This resource stops NetScaler session migration.

Session migration is driven by three separate action resources — `citrixadc_nsmigration_start`, `citrixadc_nsmigration_stop`, and `citrixadc_nsmigration_complete`. This resource performs the *stop* step and takes no configurable arguments. Use the `citrixadc_nsmigration` data source to read the live migration state (including `dumpsession` and `migrationstatus`).

~> **NOTE:** Session migration is only available in an HA/migration deployment; it is not supported on a standalone appliance.


## Example usage

```hcl
resource "citrixadc_nsmigration_stop" "tf_stop" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsmigration_stop resource. It is set to `nsmigration_stop`.
