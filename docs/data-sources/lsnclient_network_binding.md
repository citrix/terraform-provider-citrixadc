---
subcategory: "LSN"
---

# Data Source: lsnclient_network_binding

This data source is used to retrieve information about a specific `lsnclient_network_binding` configuration.

## Example usage

```hcl
data "citrixadc_lsnclient_network_binding" "example" {
  clientname = "my_lsnclient"
  network    = "10.222.74.160"
}

output "binding_id" {
  value = data.citrixadc_lsnclient_network_binding.example.id
}

output "network" {
  value = data.citrixadc_lsnclient_network_binding.example.network
}

output "netmask" {
  value = data.citrixadc_lsnclient_network_binding.example.netmask
}
```

## Argument Reference

* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created.
* `network` - (Required) IPv4 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding. It is a system-generated identifier.
* `netmask` - Subnet mask for the IPv4 address specified in the Network parameter.
* `td` - ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.
