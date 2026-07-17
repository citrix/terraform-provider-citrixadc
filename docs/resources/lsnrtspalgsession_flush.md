---
subcategory: "LSN"
---

# Resource: lsnrtspalgsession_flush

The lsnrtspalgsession_flush resource performs the imperative `flush` action on the Citrix ADC, clearing a Large Scale NAT (LSN) RTSP ALG (Application Layer Gateway) session. Use it to forcibly tear down the data-channel state the ADC tracks for a streaming RTSP call — for example to recover stuck media flows or to reclaim ALG resources during troubleshooting. The session to clear is selected by its `sessionid`.

This is an action-only resource. It does not manage a persistent server-side object:

* **Apply runs the flush action.** Creating the resource sends a `POST ?action=flush` request to the ADC. The matching RTSP ALG session is cleared immediately.
* **There is no read-back.** Read is a no-op. A flushed RTSP ALG session is a transient runtime object, not a stably keyed managed resource, so the provider cannot re-resolve it. State holds only the values you supplied plus a synthetic ID.
* **There is no in-place update.** Every argument forces resource replacement. Changing `sessionid` and re-applying re-runs the flush against the new session.
* **Delete is state-only.** The `flush` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; nothing is restored on the ADC.

Because re-applying re-runs the flush, use this resource for deliberate, one-shot session-flush workflows rather than for declarative, drift-corrected configuration. To inspect RTSP ALG sessions without clearing them, use the `citrixadc_lsnrtspalgsession` data source instead.


## Example usage

```hcl
resource "citrixadc_lsnrtspalgsession_flush" "tf_lsnrtspalgsession_flush" {
  sessionid = "10.102.43.13:6789"
}
```


## Argument Reference

* `sessionid` - (Required) Session ID for the RTSP call. Identifies the RTSP ALG session to flush. Changing this value forces the `flush` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node. Note: `nodeid` is a GET-only cluster filter and is **not** included in the `flush` action payload; it is primarily useful on the corresponding `lsnrtspalgsession` data source. Changing this value forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `lsnrtspalgsession_flush`. The ADC does not assign an ID to the flush action, so this value is purely a Terraform state handle and cannot be used to look the resource up on the ADC.

This resource is action-only and represents no importable server-side object, so there is no Import section.
