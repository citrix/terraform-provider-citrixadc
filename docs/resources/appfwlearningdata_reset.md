---
subcategory: "Application Firewall"
---

# Resource: appfwlearningdata_reset

The appfwlearningdata_reset resource clears the Citrix ADC Application-Firewall learned-data on demand. It is an action-only resource: applying it invokes the NITRO `reset` action on `appfwlearningdata`, which purges all learned-data databases and zeroes the transaction count. Use it to discard accumulated learning (for example, after tuning App-Firewall profiles or before starting a fresh learning cycle) so that subsequently learned rules are not skewed by stale samples.

Each apply performs the reset.

~> **WARNING** Applying this resource clears (resets) all App-Firewall learned data. This is a disruptive, non-reversible side effect. Use it deliberately.


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
