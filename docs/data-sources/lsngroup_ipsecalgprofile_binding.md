---
subcategory: "LSN"
---

# Data Source: lsngroup_ipsecalgprofile_binding

The lsngroup_ipsecalgprofile_binding data source allows you to retrieve information about an IPSec ALG profile that is bound to a Large Scale NAT (LSN) group.


## Example usage

```terraform
data "citrixadc_lsngroup_ipsecalgprofile_binding" "tf_binding" {
  groupname       = "lsngroup1"
  ipsecalgprofile = "ipsecalgprofile1"
}

output "bound_ipsecalgprofile" {
  value = data.citrixadc_lsngroup_ipsecalgprofile_binding.tf_binding.ipsecalgprofile
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group whose binding you want to look up.
* `ipsecalgprofile` - (Required) Name of the IPSec ALG profile bound to the specified LSN group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_ipsecalgprofile_binding resource. It is a comma-separated string of `key:value` pairs (URL-encoded), composed of the `groupname` and `ipsecalgprofile` attributes — for example, `groupname:lsngroup1,ipsecalgprofile:ipsecalgprofile1`.
* `groupname` - Name of the LSN group.
* `ipsecalgprofile` - Name of the IPSec ALG profile bound to the LSN group.
