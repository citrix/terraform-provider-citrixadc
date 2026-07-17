---
subcategory: "Basic"
---

# Resource: dbsmonitors_restart

The dbsmonitors_restart resource performs the NITRO `dbsmonitors` `restart` action, which restarts the database (DBS) monitors on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the restart, and there are no configurable arguments.

~> **NOTE** There is no NITRO GET endpoint for `dbsmonitors`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_dbsmonitors_restart" "tf_dbsmonitors_restart" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbsmonitors_restart resource. It is a synthetic value (`dbsmonitors_restart`), since the NITRO `dbsmonitors` action exposes no readable object.
