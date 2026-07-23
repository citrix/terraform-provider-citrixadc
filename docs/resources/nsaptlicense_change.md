---
subcategory: "NS"
---

# Resource: nsaptlicense_change

This resource is used to allocate pooled APT/CADS license counts to the Citrix ADC.

!> **DISRUPTIVE / NON-IDEMPOTENT:** Each apply consumes pooled licenses and re-runs the allocation, which can over-allocate or exhaust the pool. Review `countavailable` before applying.


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
* `serialno` - (Optional) Hardware Serial Number/License Activation Code (LAC). Changing this value forces the resource to be replaced.
* `useproxy` - (Optional) Specifies whether to use the licenseproxyserver to reach the internet. Make sure to configure licenseproxyserver to use this option. Changing this value forces the resource to be replaced.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsaptlicense_change resource. It has the same value as the configured `id` (NITRO License ID) attribute.
