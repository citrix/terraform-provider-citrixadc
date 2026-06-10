---
subcategory: "LSN"
---

# Resource: lsngroup_lsnrtspalgprofile_binding

Associates an LSN RTSP ALG profile with a Large Scale NAT (LSN) group so that RTSP control and media traffic handled by that group is processed according to the profile's Application Layer Gateway settings. Bind a profile when subscribers behind the LSN group use RTSP-based applications (such as streaming media) that require the ADC to rewrite embedded transport addresses.


## Example usage

```hcl
resource "citrixadc_lsnrtspalgprofile" "rtspalgprofile1" {
  rtspalgprofilename     = "rtspalgprofile1"
  rtspidletimeout        = 120
  rtspportrange          = "554"
  rtsptransportprotocol  = "TCP"
}

resource "citrixadc_lsngroup" "lsngroup1" {
  groupname    = "lsngroup1"
  clientname   = "client1"
  nattype      = "NAT44"
}

resource "citrixadc_lsngroup_lsnrtspalgprofile_binding" "tf_binding" {
  groupname          = citrixadc_lsngroup.lsngroup1.groupname
  rtspalgprofilename = citrixadc_lsnrtspalgprofile.rtspalgprofile1.rtspalgprofilename
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1'). Changing this attribute forces a new resource to be created.
* `rtspalgprofilename` - (Required) The name of the LSN RTSP ALG Profile. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_lsnrtspalgprofile_binding. It is a comma-separated string of `key:value` pairs in the form `groupname:<groupname>,rtspalgprofilename:<rtspalgprofilename>`, with each value URL-encoded.


## Import

A lsngroup_lsnrtspalgprofile_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lsngroup_lsnrtspalgprofile_binding.tf_binding groupname:lsngroup1,rtspalgprofilename:rtspalgprofile1
```
