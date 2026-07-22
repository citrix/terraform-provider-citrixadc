---
subcategory: "Basic"
---

# Resource: dbsmonitors_restart

The dbsmonitors_restart resource performs the NITRO `dbsmonitors` `restart` action, which restarts the database (DBS) monitors on the Citrix ADC. It is an action-only, zero-attribute resource: applying it triggers the restart, and there are no configurable arguments.


## Example usage

```hcl
resource "citrixadc_dbsmonitors_restart" "tf_dbsmonitors_restart" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbsmonitors_restart resource. It is set to `dbsmonitors_restart`.
