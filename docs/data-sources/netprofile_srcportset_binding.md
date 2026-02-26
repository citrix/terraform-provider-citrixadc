---
subcategory: "Network"
---

# Data Source: netprofile_srcportset_binding

The netprofile_srcportset_binding data source allows you to retrieve information about the binding between a netprofile and a source port set.

## Example Usage

```terraform
data "citrixadc_netprofile_srcportset_binding" "tf_binding" {
  name         = "tf_netprofile"
  srcportrange = "2000"
}

output "name" {
  value = data.citrixadc_netprofile_srcportset_binding.tf_binding.name
}

output "srcportrange" {
  value = data.citrixadc_netprofile_srcportset_binding.tf_binding.srcportrange
}
```

## Argument Reference

* `name` - (Required) Name of the netprofile to which to bind port ranges.
* `srcportrange` - (Required) When the source port range is configured and associated with the netprofile bound to a service group, Citrix ADC will choose a port from the range configured for connection establishment at the backend servers.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netprofile_srcportset_binding. It is the concatenation of `name` and `srcportrange` attributes separated by comma.
