---
subcategory: "Basic"
---

# Resource: locationdata

The locationdata resource performs the NITRO `locationdata` `clear` action, which clears the static location (GSLB geo) database from memory. It is an action-only, zero-attribute resource: applying it triggers the clear, and there are no configurable arguments.

~> **NOTE** There is no NITRO GET endpoint for `locationdata`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_locationdata" "tf_locationdata" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationdata resource. It is a synthetic value (`locationdata-config`), since the NITRO `locationdata` action exposes no readable object.
