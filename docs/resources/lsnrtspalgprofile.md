---
subcategory: "LSN"
---

# Resource: lsnrtspalgprofile

The lsnrtspalgprofile resource is used to create lsnrtspalgprofile.


## Example usage

```hcl
resource "citrixadc_lsnrtspalgprofile" "tf_lsnrtspalgprofile" {
  rtspalgprofilename = "my_lsn_rtspalgprofile"
  rtspportrange      = 4200
  rtspidletimeout    = 150
}

```


## Argument Reference

* `rtspalgprofilename` - (Required) The name of the RTSPALG Profile.
* `rtspportrange` - (Required) port for the RTSP
* `rtspidletimeout` - (Optional) Idle timeout for the rtsp sessions in seconds.
* `rtsptransportprotocol` - (Optional) RTSP ALG Profile transport protocol type.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnrtspalgprofile. It has the same value as the `rtspalgprofilename` attribute.


## Import

A lsnrtspalgprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnrtspalgprofile.tf_lsnrtspalgprofile my_lsn_rtspalgprofile
```
