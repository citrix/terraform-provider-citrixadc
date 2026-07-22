---
subcategory: "LSN"
---

# Resource: lsnsipalgcall_flush

The lsnsipalgcall_flush resource performs the imperative `flush` action on the Citrix ADC, clearing a Large Scale NAT (LSN) SIP ALG (Application Layer Gateway) call. Use it to forcibly tear down the signaling and media pinhole state the ADC tracks for an active SIP call — for example to recover a stuck call, free ALG resources, or reset SIP ALG state during troubleshooting. The call to clear is selected by its `callid`.

This is an action resource: applying it performs the `flush` action against the SIP ALG call selected by `callid`; it does not manage a persistent object, so re-applying re-runs the flush. Every argument forces resource replacement.

To inspect SIP ALG calls without clearing them, use the `citrixadc_lsnsipalgcall` data source instead.


## Example usage

```hcl
resource "citrixadc_lsnsipalgcall_flush" "tf_lsnsipalgcall_flush" {
  callid = "12345-abcde"
}
```


## Argument Reference

* `callid` - (Required) Call ID for the SIP call. Identifies the SIP ALG call to flush. Changing this value forces the `flush` action to re-run (resource replacement).
* `nodeid` - (Optional) Unique number that identifies the cluster node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnsipalgcall_flush resource. It is set to `lsnsipalgcall_flush`.

This resource is action-only and represents no importable server-side object, so there is no Import section.
