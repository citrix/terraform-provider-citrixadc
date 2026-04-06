---
subcategory: "Network"
---

# Data Source: netbridge_nsip_binding

The netbridge_nsip_binding data source allows you to retrieve information about an IPv4 address binding to a network bridge.

## Example Usage

```terraform
data "citrixadc_netbridge_nsip_binding" "tf_netbridge_nsip_binding" {
  name      = "my_netbridge"
  ipaddress = "10.222.74.128"
}

output "ipaddress" {
  value = data.citrixadc_netbridge_nsip_binding.tf_netbridge_nsip_binding.ipaddress
}

output "netmask" {
  value = data.citrixadc_netbridge_nsip_binding.tf_netbridge_nsip_binding.netmask
}
```

## Argument Reference

* `name` - (Required) The name of the network bridge.
* `ipaddress` - (Required) The subnet that is extended by this network bridge.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_nsip_binding. It is a system-generated identifier.
* `netmask` - The network mask for the subnet.
