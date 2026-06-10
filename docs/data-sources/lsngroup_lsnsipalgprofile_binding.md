---
subcategory: "LSN"
---

# Data Source: lsngroup_lsnsipalgprofile_binding

The lsngroup_lsnsipalgprofile_binding data source allows you to retrieve information about an LSN SIP ALG profile binding on a given LSN group.


## Example usage

```terraform
data "citrixadc_lsngroup_lsnsipalgprofile_binding" "example" {
  groupname         = "lsngroup1"
  sipalgprofilename = "sipalgprofile1"
}

output "bound_sipalgprofile" {
  value = data.citrixadc_lsngroup_lsnsipalgprofile_binding.example.sipalgprofilename
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group whose binding you want to look up.
* `sipalgprofilename` - (Required) The name of the LSN SIP ALG Profile bound to the LSN group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_lsnsipalgprofile_binding resource. It is the concatenation of the `groupname` and `sipalgprofilename` attributes (as `groupname:<value>,sipalgprofilename:<value>`, URL-encoded).
* `groupname` - Name of the LSN group to which the SIP ALG profile is bound.
* `sipalgprofilename` - The name of the LSN SIP ALG Profile bound to the LSN group.
