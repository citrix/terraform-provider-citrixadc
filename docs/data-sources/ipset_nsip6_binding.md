---
subcategory: "Network"
---

# Data Source: ipset_nsip6_binding

The ipset_nsip6_binding data source allows you to retrieve information about ipset_nsip6_binding.


## Example Usage

```terraform
data "citrixadc_ipset_nsip6_binding" "tf_ipset_nsip6_binding" {
  name      = "tf_ipset"
  ipaddress = "2003:db8:100::fb/64"
}

output "name" {
  value = data.citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding.name
}

output "ipaddress" {
  value = data.citrixadc_ipset_nsip6_binding.tf_ipset_nsip6_binding.ipaddress
}
```


## Argument Reference

* `name` - (Required) Name of the IP set to which to bind IP addresses.
* `ipaddress` - (Required) One or more IP addresses bound to the IP set.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipset_nsip6_binding. It is a system-generated identifier.
