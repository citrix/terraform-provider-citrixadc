---
subcategory: "LLDP"
---

# Resource: lldpneighbors_clear

This resource is used to clear the learned LLDP neighbor table on the Citrix ADC.


## Example usage

### Clear the learned neighbor table

The clear action operates on the entire learned neighbor table and takes no arguments.

```hcl
resource "citrixadc_lldpneighbors_clear" "tf_lldpneighbors_clear" {
}
```

To clear the table again on a subsequent run, taint the resource first:

```shell
terraform taint citrixadc_lldpneighbors_clear.tf_lldpneighbors_clear
terraform apply
```


## Argument Reference

The clear action takes no arguments, so this resource has no required configuration. The following optional attributes are carried in the schema for parity with the `citrixadc_lldpneighbors` data source filters. They do **not** scope the clear action (which always operates on all learned neighbors). Changing either attribute forces a new clear action to be performed.

* `ifnum` - (Optional) Interface name. Retained for parity with the `citrixadc_lldpneighbors` data source filter; it does not scope the clear action. Changing this attribute re-triggers the clear.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Retained for parity with the data source filter; it does not scope the clear action. Changing this attribute re-triggers the clear.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lldpneighbors_clear resource. It is set to `lldpneighbors_clear`.
