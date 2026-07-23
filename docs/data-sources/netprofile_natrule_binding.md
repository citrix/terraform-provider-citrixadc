---
subcategory: "Network"
---

# Data Source: netprofile_natrule_binding

The netprofile_natrule_binding data source allows you to retrieve information about a binding between a netprofile and a NAT rule.


## Example usage

```terraform
data "citrixadc_netprofile_natrule_binding" "tf_binding" {
  name    = "tf_netprofile"
  natrule = "10.10.10.10"
}

output "name" {
  value = data.citrixadc_netprofile_natrule_binding.tf_binding.name
}

output "rewriteip" {
  value = data.citrixadc_netprofile_natrule_binding.tf_binding.rewriteip
}
```

## Argument Reference

* `name` - (Required) Name of the netprofile to which to bind port ranges.
* `natrule` - (Required) IPv4 network address on whose traffic you want the Citrix ADC to do rewrite ip prefix.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netprofile_natrule_binding. It is a system-generated identifier.
* `rewriteip` - IP address used to rewrite the network address prefix.
* `netmask` - Subnet mask associated with the network address.
