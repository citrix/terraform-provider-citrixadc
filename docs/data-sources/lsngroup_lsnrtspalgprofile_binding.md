---
subcategory: "LSN"
---

# Data Source: lsngroup_lsnrtspalgprofile_binding

The lsngroup_lsnrtspalgprofile_binding data source allows you to retrieve information about an existing binding between an LSN RTSP ALG profile and a Large Scale NAT (LSN) group.


## Example usage

```terraform
data "citrixadc_lsngroup_lsnrtspalgprofile_binding" "example" {
  groupname          = "lsngroup1"
  rtspalgprofilename = "rtspalgprofile1"
}

output "bound_rtspalgprofile" {
  value = data.citrixadc_lsngroup_lsnrtspalgprofile_binding.example.rtspalgprofilename
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group whose binding you want to look up.
* `rtspalgprofilename` - (Required) The name of the LSN RTSP ALG Profile bound to the group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_lsnrtspalgprofile_binding. It is a comma-separated string of `key:value` pairs in the form `groupname:<groupname>,rtspalgprofilename:<rtspalgprofilename>`, with each value URL-encoded.
* `groupname` - Name of the LSN group.
* `rtspalgprofilename` - The name of the LSN RTSP ALG Profile bound to the group.
