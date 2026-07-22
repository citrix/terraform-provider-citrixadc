---
subcategory: "LSN"
---

# Resource: lsnrtspalgsession_flush

The lsnrtspalgsession_flush resource performs the imperative `flush` action on the Citrix ADC, clearing a Large Scale NAT (LSN) RTSP ALG (Application Layer Gateway) session. Use it to forcibly tear down the data-channel state the ADC tracks for a streaming RTSP call — for example to recover stuck media flows or to reclaim ALG resources during troubleshooting. The session to clear is selected by its `sessionid`.

This is an action resource: applying it performs the flush; it does not manage a persistent object, so re-applying re-runs the action. Every argument forces resource replacement; changing `sessionid` and re-applying re-runs the flush against the new session.

Because re-applying re-runs the flush, use this resource for deliberate, one-shot session-flush workflows rather than for declarative, drift-corrected configuration. To inspect RTSP ALG sessions without clearing them, use the `citrixadc_lsnrtspalgsession` data source instead.


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
