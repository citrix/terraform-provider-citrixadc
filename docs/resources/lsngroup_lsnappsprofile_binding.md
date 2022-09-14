---
subcategory: "Lsn"
---

# Resource: lsngroup_lsnappsprofile_binding

The lsngroup_lsnappsprofile_binding resource is used to create lsngroup_lsnappsprofile_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_lsnappsprofile_binding" "tf_lsngroup_lsnappsprofile_binding" {
  groupname       = "my_lsn_group"
  appsprofilename = "my_lsn_profile"
}
```


## Argument Reference

* `appsprofilename` - (Required) Name of the LSN application profile to bind to the specified LSN group. For each set of destination ports, bind a profile for each protocol for which you want to specify settings.  By default, one LSN application profile with default settings for TCP, UDP, and ICMP protocols for all destination ports is bound to an LSN group during its creation.  This profile is called a default application profile.  When you bind an LSN application profile, with a specified set of destination ports, to an LSN group, the bound profile overrides the default LSN application profile for that protocol at that set of destination ports.
* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_lsnappsprofile_binding. It is the concatenation of  `groupname` and `appsprofilename` attributes separated by a comma.


## Import

A lsngroup_lsnappsprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_lsnappsprofile_binding.tf_lsngroup_lsnappsprofile_binding my_lsn_group,my_lsn_profile
```
