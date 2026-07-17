---
subcategory: "NS"
---

# Resource: nsaptlicense_change

The nsaptlicense_change resource allocates Citrix ADC pooled APT/CADS license counts to the appliance against a registered license session. It maps to the NITRO nsaptlicense `change` action (invoked at `POST ?action=update`) and is used to draw down a number of available licenses from a pooled licensing server and bind them to this instance.

!> **DISRUPTIVE / NON-IDEMPOTENT.** Applying this resource invokes the NITRO change/update action, which consumes pooled licenses from the licensing server and changes the licensed capacity of the appliance. It is **not idempotent** — each apply re-runs the allocation action, which can over-allocate or exhaust the pool. Review the `countavailable` value carefully before applying, and treat this resource as a one-shot operational action rather than ordinary declarative configuration.

~> **Note.** This is an action-only resource. Create performs the allocation; Read, Update, and Delete are no-ops. There is no inverse NITRO API (no "un-allocate"), so deleting the resource only removes it from Terraform state — the allocated licenses remain on the appliance. Every attribute is marked `RequiresReplace`, so changing any argument forces the allocation action to run again as a replacement.


## Example usage

```hcl
resource "citrixadc_nsaptlicense_change" "tf_nsaptlicense_change" {
  id             = "1"
  sessionid      = "00000000-0000-0000-0000-000000000000"
  bindtype       = "CADS"
  countavailable = "5"
  serialno       = "ABC123XYZ789"
}
```


## Argument Reference

* `id` - (Required) License ID. This is the NITRO License ID used as the key for the change/update action and as the Terraform resource identifier. Changing this value forces the resource to be replaced.
* `sessionid` - (Required) Session ID. The license session against which the licenses are allocated. Changing this value forces the resource to be replaced.
* `bindtype` - (Required) Bind type. Changing this value forces the resource to be replaced.
* `countavailable` - (Required) The user can allocate one or more licenses. Ensure the value is less than (for partial allocation) or equal to the total number of available licenses. Changing this value forces the resource to be replaced.
* `licensedir` - (Optional) License Directory. Changing this value forces the resource to be replaced.
* `serialno` - (Optional) Hardware Serial Number/License Activation Code (LAC). This is a GET-only filter key; it is not part of the change/update action payload sent to the appliance. Changing this value forces the resource to be replaced.
* `useproxy` - (Optional) Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option. Changing this value forces the resource to be replaced.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsaptlicense_change resource. It has the same value as the configured `id` (NITRO License ID) attribute.
