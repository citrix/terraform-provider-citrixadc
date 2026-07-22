---
subcategory: "Basic"
---

# Resource: locationdata_clear

The locationdata_clear resource performs the NITRO `locationdata` `clear` action, which clears the static location (GSLB geo) database from memory. It is an action resource: applying it performs the clear, there are no configurable arguments, and re-applying re-runs the action.


## Example usage

```hcl
resource "citrixadc_locationdata_clear" "tf_locationdata_clear" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the locationdata_clear resource. It is set to `locationdata_clear`.
