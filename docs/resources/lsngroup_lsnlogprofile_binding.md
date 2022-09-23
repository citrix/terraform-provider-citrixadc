---
subcategory: "LSN"
---

# Resource: lsngroup_lsnlogprofile_binding

The lsngroup_lsnlogprofile_binding resource is used to create lsngroup_lsnlogprofile_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_lsnlogprofile_binding" "tf_lsngroup_lsnlogprofile_binding" {
  groupname      = "my_lsn_group"
  logprofilename = "my_lsn_logprofile"
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `logprofilename` - (Required) The name of the LSN logging Profile.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_lsnlogprofile_binding. It is the concatenation of `groupname` and `logprofilename` attributes separated by a comma.


## Import

A lsngroup_lsnlogprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_lsnlogprofile_binding.tf_lsngroup_lsnlogprofile_binding my_lsn_group,my_lsn_logprofile
```
