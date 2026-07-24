---
subcategory: "LSN"
---

# Resource: lsnsipalgcall_flush

This resource is used to flush an LSN SIP ALG call on the Citrix ADC.


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
