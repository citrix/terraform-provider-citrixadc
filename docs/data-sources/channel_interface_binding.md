---
subcategory: "Network"
---

# Data Source: channel\_interface\_binding

The channel\_interface\_binding data source allows you to retrieve information about the interfaces bound to a link aggregation (LA) channel on the Citrix ADC.


## Example usage

```hcl
data "citrixadc_channel_interface_binding" "tf_bind" {
  channelid = "LA/1"
  ifnum     = ["1/3"]
}

output "bound_interfaces" {
  value = data.citrixadc_channel_interface_binding.tf_bind.ifnum
}
```


## Argument Reference

* `channelid` - (Required) ID of the LA channel or the cluster LA channel whose interface bindings you want to look up. Specify an LA channel in `LA/x` notation, a cluster LA channel in `CLA/x` notation, or a link-redundant channel in `LR/x` notation.
* `ifnum` - (Required) A list of interfaces bound to the channel. For an LA channel of a Citrix ADC, interfaces are in `C/U` notation (for example, `1/3`); for a cluster configuration, in `N/C/U` notation (for example, `2/1/3`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the channel\_interface\_binding resource. It is a composite key of the form `id:<channelid>,ifnum:<ifnum>`, with each value URL-encoded.
