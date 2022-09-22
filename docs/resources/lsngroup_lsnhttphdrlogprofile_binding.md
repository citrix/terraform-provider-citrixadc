---
subcategory: "LSN"
---

# Resource: lsngroup_lsnhttphdrlogprofile_binding

The lsngroup_lsnhttphdrlogprofile_binding resource is used to create lsngroup_lsnhttphdrlogprofile_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_lsnhttphdrlogprofile_binding" "tf_lsngroup_lsnhttphdrlogprofile_binding" {
  groupname             = "my_lsn_group"
  httphdrlogprofilename = "my_httplogprofile"
}

```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `httphdrlogprofilename` - (Required) The name of the LSN HTTP header logging Profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_lsnhttphdrlogprofile_binding. It is the concatenation of `groupname` and `httphdrlogprofilename` attributes separated by a comma.


## Import

A lsngroup_lsnhttphdrlogprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_lsnhttphdrlogprofile_binding.tf_lsngroup_lsnhttphdrlogprofile_binding my_lsn_group,my_httplogprofile
```
