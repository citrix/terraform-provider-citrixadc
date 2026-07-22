---
subcategory: "System"
---

# Resource: systemrestorepoint_create

Creates a restore point on a Citrix ADC. A restore point is a named snapshot that
captures the appliance configuration together with a tech-support bundle, giving
you a known-good point you can fall back to during upgrades, maintenance, or
troubleshooting.

Creating the resource snapshots the config, and destroying it deletes the
restore point from the appliance. The resource ID is the real object name
(the filename).

~> **NOTE:** The appliance enforces a MAXIMUM of 3 restore points. Attempting to
create a fourth restore point fails on the NITRO side until an existing restore
point is deleted.

~> **NOTE:** The `filename` attribute is immutable. Changing the filename
destroys the existing restore point and creates a new one.


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
