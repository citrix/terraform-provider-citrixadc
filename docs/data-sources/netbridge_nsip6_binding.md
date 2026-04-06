---
subcategory: "Network"
---

# Data Source: netbridge_nsip6_binding

The netbridge_nsip6_binding data source allows you to retrieve information about an IPv6 address binding to a network bridge.

## Example Usage

```terraform
data "citrixadc_netbridge_nsip6_binding" "tf_netbridge_nsip6_binding" {
  name      = "my_netbridge"
  ipaddress = "dea:97c5:d381:e72b::/64"
}

output "netmask" {
  value = data.citrixadc_netbridge_nsip6_binding.tf_netbridge_nsip6_binding.netmask
}

output "id" {
  value = data.citrixadc_netbridge_nsip6_binding.tf_netbridge_nsip6_binding.id
}
```

## Argument Reference

* `name` - (Required) The name of the network bridge.
* `ipaddress` - (Required) The subnet that is extended by this network bridge.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `netmask` - The network mask for the subnet.
* `id` - The id of the netbridge_nsip6_binding. It is a system-generated identifier.
