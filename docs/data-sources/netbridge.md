---
subcategory: "Network"
---

# Data Source `netbridge`

The netbridge data source allows you to retrieve information about an existing network bridge.


## Example usage

```terraform
data "citrixadc_netbridge" "tf_netbridge" {
  name = "tf_netbridge_example"
}

output "vxlanvlanmap" {
  value = data.citrixadc_netbridge.tf_netbridge.vxlanvlanmap
}
```


## Argument Reference

* `name` - (Required) The name of the network bridge.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge. It has the same value as the `name` attribute.
* `vxlanvlanmap` - The vlan to vxlan mapping to be applied to this netbridge.


## Import

A netbridge can be imported using its name, e.g.

```shell
terraform import citrixadc_netbridge.tf_netbridge tf_netbridge_example
```
