---
subcategory: "LSN"
---

# Resource: lsnsipalgcall_flush

The lsnsipalgcall_flush resource performs the imperative `flush` action on the Citrix ADC, clearing a Large Scale NAT (LSN) SIP ALG (Application Layer Gateway) call. Use it to forcibly tear down the signaling and media pinhole state the ADC tracks for an active SIP call — for example to recover a stuck call, free ALG resources, or reset SIP ALG state during troubleshooting. The call to clear is selected by its `callid`.

This is an action-only resource. It does not manage a persistent server-side object:

* **Apply runs the flush action.** Creating the resource sends a `POST ?action=flush` request to the ADC. The matching SIP ALG call is flushed immediately.
* **There is no read-back.** Read is a no-op. A flushed SIP ALG call is a transient runtime object, not a stably keyed managed resource, so the provider cannot re-resolve it. State holds only the values you supplied plus a synthetic ID.
* **There is no in-place update.** Every argument forces resource replacement. Changing `callid` and re-applying re-runs the flush against the new call.
* **Delete is state-only.** The `flush` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; nothing is restored on the ADC.

Because re-applying re-runs the flush, use this resource for deliberate, one-shot call-flush workflows rather than for declarative, drift-corrected configuration. To inspect SIP ALG calls without clearing them, use the `citrixadc_lsnsipalgcall` data source instead.


## Example usage

```hcl
resource "citrixadc_lsnsipalgcall_flush" "tf_lsnsipalgcall_flush" {
  callid = "12345-abcde"
}
```


## Argument Reference

* `callid` - (Required) Call ID for the SIP call. Identifies the SIP ALG call to flush. Changing this value forces the `flush` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node. Note: `nodeid` is a GET-only cluster filter and is **not** included in the `flush` action payload; it is primarily useful on the corresponding `lsnsipalgcall` data source.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `lsnsipalgcall_flush`. The ADC does not assign an ID to the flush action, so this value is not a server-assigned key and cannot be used to look the resource up on the ADC.

This resource is action-only and represents no importable server-side object, so there is no Import section.
</content>
</invoke>
