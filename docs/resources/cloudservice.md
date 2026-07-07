---
subcategory: "Cloud"
---

# Resource: cloudservice

The cloudservice resource performs the NITRO `cloudservice` `check` action, which checks the cloud service configuration on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the check, and there are no configurable arguments.

~> **NOTE** There is no NITRO GET endpoint for `cloudservice`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_cloudservice" "tf_cloudservice" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudservice resource. It is a synthetic value (`cloudservice-config`), since the NITRO `cloudservice` action exposes no readable object.
