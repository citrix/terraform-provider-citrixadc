---
subcategory: "Cluster"
---

# Resource: clusterpropstatus_clear

This resource is used to clear the configuration-propagation status counters of a Citrix ADC cluster.


## Example usage

Clear the propagation status counters:

```hcl
resource "citrixadc_clusterpropstatus_clear" "clear" {
}
```


## Argument Reference

This action-only resource takes no configurable arguments. The `clear` verb accepts no parameters. Each apply re-runs the clear action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusterpropstatus_clear resource. It is set to `clusterpropstatus_clear`.
