---
subcategory: "Network"
---

# Resource: channel\_interface\_binding

This resource is used to bind physical interfaces to an LA, cluster LA, or LR channel.


## Example usage

```hcl
resource "citrixadc_channel" "tf_channel" {
  channelid = "LA/1"
  ifnum     = ["1/1", "1/2"]
  speed     = "10000"
}

resource "citrixadc_channel_interface_binding" "tf_bind" {
  channelid = citrixadc_channel.tf_channel.channelid
  ifnum     = ["1/3", "1/4"]
}
```


## Argument Reference

* `channelid` - (Required) ID of the LA channel or the cluster LA channel to which you want to bind interfaces. Specify an LA channel in `LA/x` notation, where `x` can range from 1 to 8, a cluster LA channel in `CLA/x` notation, or a link-redundant channel in `LR/x` notation, where `x` can range from 1 to 4. Changing this value forces a new resource to be created.
* `ifnum` - (Required) A list of interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration. For an LA channel of a Citrix ADC, specify an interface in `C/U` notation (for example, `1/3`). For an LA channel of a cluster configuration, specify an interface in `N/C/U` notation (for example, `2/1/3`). `C` can be `0` (management interface), `1` (1 Gbps port), or `10` (10 Gbps port). `U` is a unique integer representing an interface within a particular port group. `N` is the ID of the node to which an interface belongs in a cluster configuration. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the channel\_interface\_binding resource. It is a composite key of the form `id:<channelid>,ifnum:<ifnum>`, where each value is URL-encoded (because channel and interface identifiers contain `/`). For example, the channel `LA/1` bound to interface `1/3` yields `id:LA%2F1,ifnum:1%2F3`.


## Import

A channel\_interface\_binding can be imported using its `id`, e.g.

```shell
terraform import citrixadc_channel_interface_binding.tf_bind "id:LA%2F1,ifnum:1%2F3"
```
