---
subcategory: "LSN"
---

# Data Source `lsnrtspalgprofile`

The lsnrtspalgprofile data source allows you to retrieve information about LSN RTSP Application Layer Gateway profiles.


## Example usage

```terraform
data "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
  rtspalgprofilename = "my_lsn_rtspalgprofile_ds"
}

output "rtspportrange" {
  value = data.citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile.rtspportrange
}

output "rtspidletimeout" {
  value = data.citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile.rtspidletimeout
}
```


## Argument Reference

* `rtspalgprofilename` - (Required) The name of the RTSPALG Profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `rtspidletimeout` - Idle timeout for the rtsp sessions in seconds.
* `rtspportrange` - port for the RTSP
* `rtsptransportprotocol` - RTSP ALG Profile transport protocol type.

## Attribute Reference

* `id` - The id of the lsnrtspalgprofile. It has the same value as the `rtspalgprofilename` attribute.


## Import

A lsnrtspalgprofile can be imported using its rtspalgprofilename, e.g.

```shell
terraform import citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile my_lsn_rtspalgprofile_ds
```
