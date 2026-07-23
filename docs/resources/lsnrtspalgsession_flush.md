---
subcategory: "LSN"
---

# Resource: lsnrtspalgsession_flush

This resource is used to flush an LSN RTSP ALG session on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_lsnrtspalgsession_flush" "tf_lsnrtspalgsession_flush" {
  sessionid = "10.102.43.13:6789"
}
```


## Argument Reference

* `sessionid` - (Required) Session ID for the RTSP call. Identifies the RTSP ALG session to flush. Changing this value forces the `flush` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing this value forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnrtspalgsession_flush resource. It is set to `lsnrtspalgsession_flush`.

This resource is action-only and represents no importable server-side object, so there is no Import section.
