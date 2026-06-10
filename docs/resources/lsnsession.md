---
subcategory: "LSN"
---

# Resource: lsnsession

Flushes active Large Scale NAT (LSN) sessions on the Citrix ADC. This resource performs an imperative `flush` action: applying it tears down the LSN sessions that match the supplied optional filter selectors (NAT type, client name, subscriber network/netmask, IPv6 subscriber address, traffic domain, mapped NAT IP/port, or cluster node), so administrators can reclaim NAT mappings or clear stale state without rebooting.

This is an action-only resource. It does not manage a persistent server-side object:

* **Apply runs the flush action.** Creating the resource sends a `POST ?action=flush` request to the ADC. Any matching LSN sessions are removed immediately. This is a one-shot side-effect.
* **There is no read-back.** Read is a no-op. A flushed session is not a queryable managed object (NITRO exposes no get-by-name key for `lsnsession`), so the provider cannot re-resolve it or detect drift. State holds only the filters you supplied plus a synthetic ID.
* **There is no in-place update.** Every argument forces resource replacement. Changing any filter and re-applying re-runs the flush action against the new filter set.
* **Delete is state-only.** The `flush` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; no sessions are restored.

Because re-applying re-runs the flush, use this resource for deliberate, one-shot session-flush workflows rather than for declarative, drift-corrected configuration.


## Example usage

Flush LSN sessions, narrowing the selection with filters. The example flushes NAT44 sessions belonging to a specific LSN client, restricted to a subscriber network and a traffic domain:

```hcl
resource "citrixadc_lsnsession" "flush_nat44" {
  nattype    = "NAT44"
  clientname = "lsnclient1"
  network    = "192.168.1.0"
  netmask    = "255.255.255.0"
  td         = 0
}
```

Supplying fewer filters widens the set of sessions that get flushed. The following flushes all NAT44 sessions for a single subscriber address:

```hcl
resource "citrixadc_lsnsession" "flush_subscriber" {
  nattype = "NAT44"
  network = "192.168.1.25"
}
```


## Argument Reference

All arguments are optional filter selectors that narrow which LSN sessions are flushed. Changing any argument forces the resource to be recreated, which re-runs the flush action.

* `nattype` - (Optional) Type of sessions to flush. Defaults to `"NAT44"`. Possible values: [ NAT44, DS-Lite, NAT64 ]
* `clientname` - (Optional) Name of the LSN Client entity whose sessions should be flushed.
* `network` - (Optional) IP address or network address of subscriber(s) whose sessions should be flushed.
* `netmask` - (Optional) Subnet mask for the IP address specified by the `network` argument. Defaults to `"255.255.255.255"`.
* `network6` - (Optional) IPv6 address of the LSN subscriber or B4 device whose sessions should be flushed.
* `td` - (Optional) Traffic domain ID of the LSN client entity. Minimum value = 0 Maximum value = 4094.
* `natip` - (Optional) Mapped NAT IP address used in the LSN sessions to be flushed.
* `natport2` - (Optional) Mapped NAT port used in the LSN sessions to be flushed.
* `nodeid` - (Optional) Unique number that identifies the cluster node whose sessions should be flushed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `"lsnsession"`. The ADC does not assign an ID to the flush action; this value is purely a Terraform state handle and is not a NITRO lookup key.

This resource is action-only and represents no importable server-side object, so importing it is not meaningful and there is no Import section.
