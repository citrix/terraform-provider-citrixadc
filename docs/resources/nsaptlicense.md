---
subcategory: "NS"
---

# Resource: nsaptlicense

The nsaptlicense resource allocates Citrix ADC pooled APT/CADS license counts to the appliance against a registered license session. Use it to draw down a number of available licenses from a pooled licensing server and bind them to this instance.

!> **DISRUPTIVE / NON-IDEMPOTENT.** This resource maps to the NITRO `update` action (`POST ?action=update`). Applying it consumes pooled licenses from the licensing server and changes the licensed capacity of the appliance. It is **not idempotent** — every create or replace re-runs the allocation action, which can over-allocate or exhaust the pool. Review the `countavailable` value carefully before applying, and treat this resource as a one-shot operational action rather than ordinary declarative configuration. Deleting the resource only removes it from Terraform state; the allocated licenses remain on the appliance.

~> **Note.** Because allocating licenses is disruptive, in-place updates are a no-op: the `id` and `serialno` attributes are marked `RequiresReplace`, so changing them forces the allocation action to run again as a replacement.


## Example usage

```hcl
resource "citrixadc_nsaptlicense" "tf_nsaptlicense" {
  id             = "1"
  sessionid      = "00000000-0000-0000-0000-000000000000"
  bindtype       = "CADS"
  countavailable = "5"
  serialno       = "ABC123XYZ789"
}
```


## Argument Reference

* `id` - (Required) License ID. This is the NITRO License ID used as the key for the allocation action and as the Terraform resource identifier. Changing this value forces the resource to be replaced.
* `sessionid` - (Required) Session ID. The license session against which the licenses are allocated.
* `bindtype` - (Required) Bind type.
* `countavailable` - (Required) The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses.
* `serialno` - (Optional, Computed) Hardware Serial Number/License Activation Code (LAC). This is a GET-only filter key used to read the allocated license record back from the appliance; it is not part of the allocation action payload. Changing this value forces the resource to be replaced.
* `licensedir` - (Optional, Computed) License Directory.
* `useproxy` - (Optional, Computed) Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsaptlicense resource. It has the same value as the configured `id` (NITRO License ID) attribute.


## Import

An nsaptlicense resource can be imported using its License ID, e.g.

```shell
terraform import citrixadc_nsaptlicense.tf_nsaptlicense 1
```
