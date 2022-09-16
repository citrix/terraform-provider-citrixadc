---
subcategory: "Lsn"
---

# Resource: lsngroup_lsntransportprofile_binding

The lsngroup_lsntransportprofile_binding resource is used to create lsngroup_lsntransportprofile_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_lsntransportprofile_binding" "tf_lsngroup_lsntransportprofile_binding" {
  groupname            = "my_lsn_group"
  transportprofilename = "my_lsntransportfile"
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `transportprofilename` - (Required) Name of the LSN transport profile to bind to the specified LSN group. Bind a profile for each protocol for which you want to specify settings.  By default, one LSN transport profile with default settings for TCP, UDP, and ICMP protocols is bound to an LSN group during its creation. This profile is called a default transport.  An LSN transport profile that you bind to an LSN group overrides the default LSN transport profile for that protocol.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_lsntransportprofile_binding. It is the concatenation of `groupname` and `transportprofilename` attributes separated by a comma.


## Import

A lsngroup_lsntransportprofile_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_lsntransportprofile_binding.tf_lsngroup_lsntransportprofile_binding my_lsn_group,my_lsntransportfile
```
