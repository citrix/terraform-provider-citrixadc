---
subcategory: "LSN"
---

# Resource: lsngroup_lsnsipalgprofile_binding

Associates an LSN (Large Scale NAT) SIP ALG profile with an LSN group so that SIP traffic processed by that group is handled according to the profile's Application Layer Gateway settings (for example, NAT rewriting of SIP and SDP payloads and pinhole management). Create one binding to attach a specific `citrixadc_lsnsipalgprofile` to a `citrixadc_lsngroup`.


## Example usage

```hcl
resource "citrixadc_lsngroup" "tf_lsngroup" {
  groupname     = "lsngroup1"
  clientname    = "lsnclient1"
  nattype       = "DYNAMIC"
  sessionlogging = "ENABLED"
}

resource "citrixadc_lsnsipalgprofile" "tf_lsnsipalgprofile" {
  sipalgprofilename = "sipalgprofile1"
  datasessionidletimeout = 120
}

resource "citrixadc_lsngroup_lsnsipalgprofile_binding" "tf_binding" {
  groupname         = citrixadc_lsngroup.tf_lsngroup.groupname
  sipalgprofilename = citrixadc_lsnsipalgprofile.tf_lsnsipalgprofile.sipalgprofilename
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1'). Changing this attribute forces a new resource to be created.
* `sipalgprofilename` - (Required) The name of the LSN SIP ALG Profile. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_lsnsipalgprofile_binding resource. It is the concatenation of the `groupname` and `sipalgprofilename` attributes (as `groupname:<value>,sipalgprofilename:<value>`, URL-encoded).


## Import

A lsngroup_lsnsipalgprofile_binding can be imported using its ID, which is the `groupname` and `sipalgprofilename` separated by a comma, e.g.

```shell
terraform import citrixadc_lsngroup_lsnsipalgprofile_binding.tf_binding groupname:lsngroup1,sipalgprofilename:sipalgprofile1
```
