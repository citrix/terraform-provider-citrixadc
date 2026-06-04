---
subcategory: "AAA"
---

# Data Source: aaagroup_intranetip6_binding

The `aaagroup_intranetip6_binding` data source allows you to retrieve information about an intranet IPv6 range bound to an AAA group on the Citrix ADC, such as how many addresses are allocated to the group starting from a given IPv6 address.


## Example usage

```terraform
data "citrixadc_aaagroup_intranetip6_binding" "tf_aaagroup_intranetip6_binding" {
  groupname   = "my_group"
  intranetip6 = "2001:db8::1"
}

output "numaddr" {
  value = data.citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding.numaddr
}
```


## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `intranetip6` - (Required) The intranet IPv6 address bound to the group, used to look up the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `numaddr` - Number of IPv6 addresses bound, starting with `intranetip6`.
* `gotopriorityexpression` - Not applicable to the intranet IPv6 branch of this binding; the ADC does not populate a meaningful value for it.
* `id` - The id of the aaagroup_intranetip6_binding. It is a comma-separated set of `key:value` pairs in the form `groupname:<groupname>,intranetip6:<intranetip6>,numaddr:<numaddr>`, with the `intranetip6` value URL-encoded (colons become `%3A`).
