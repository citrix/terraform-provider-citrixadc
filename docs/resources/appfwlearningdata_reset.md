---
subcategory: "Application Firewall"
---

# Resource: appfwlearningdata_reset

This resource is used to reset (clear) the Application Firewall learned data on the Citrix ADC.

!> **WARNING:** Applying this resource clears all App-Firewall learned data. This action is irreversible.


## Example usage

The `reset` action takes no arguments; it always clears the entire learned-data table.

```hcl
resource "citrixadc_appfwlearningdata_reset" "reset_all" {}
```


## Argument Reference

This resource has no configurable arguments. The NITRO `reset` action carries an empty payload and the equivalent CLI command `reset appfw learningdata` takes no parameters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwlearningdata_reset resource. It is set to `appfwlearningdata_reset`.
