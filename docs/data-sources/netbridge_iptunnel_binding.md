---
subcategory: "Network"
---

# Data Source: netbridge_iptunnel_binding

The netbridge_iptunnel_binding data source allows you to retrieve information about the binding between a network bridge and an IP tunnel.

## Example Usage

```terraform
data "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
  name   = "tf_netbridge"
  tunnel = "tf_iptunnel"
}

output "name" {
  value = data.citrixadc_netbridge_iptunnel_binding.tf_binding.name
}

output "tunnel" {
  value = data.citrixadc_netbridge_iptunnel_binding.tf_binding.tunnel
}
```

## Argument Reference

* `name` - (Required) The name of the network bridge.
* `tunnel` - (Required) The name of the tunnel that is a part of this bridge.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_iptunnel_binding. It is the concatenation of `name` and `tunnel` attributes separated by comma.
