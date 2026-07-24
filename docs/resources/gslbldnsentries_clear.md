---
subcategory: "GSLB"
---

# Resource: gslbldnsentries_clear

This resource is used to clear the learned GSLB LDNS RTT entries from the Citrix ADC.


## Example usage

```hcl
# Flush the learned GSLB LDNS RTT entries on the appliance.
resource "citrixadc_gslbldnsentries_clear" "tf_gslbldnsentries_clear" {
}
```

To target a specific cluster node when clearing:

```hcl
resource "citrixadc_gslbldnsentries_clear" "tf_gslbldnsentries_clear" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing it forces resource replacement, which re-runs the `clear` action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbldnsentries_clear resource. It is set to `gslbldnsentries_clear`.
