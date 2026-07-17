---
subcategory: "LSN"
---

# Resource: lsnsession_flush

Flushes active Large Scale NAT (LSN) sessions on the Citrix ADC. Applying this resource invokes the NITRO `flush` action, which tears down the LSN sessions that match the supplied optional filter selectors (NAT type, LSN client name, subscriber network/netmask, IPv6 subscriber address, traffic domain, mapped NAT IP/port, or cluster node). Use it to reclaim NAT mappings or clear stale session state after changing LSN configuration, without rebooting the appliance.

This is an action-only resource. It does not create, read, or manage a persistent object on the appliance:

* **Apply runs the flush action.** Creating the resource sends a `POST ?action=flush` request to the ADC and removes any matching LSN sessions immediately. This is a one-shot side-effect.
* **There is no read-back.** NITRO exposes no GET endpoint for `lsnsession` flush state, so Read is a no-op and there is no corresponding data source. The provider cannot re-resolve the action or detect drift; state holds only the filters you supplied plus a synthetic ID.
* **There is no in-place update.** Every argument forces resource replacement, so changing any filter and re-applying re-runs the flush against the new filter set.
* **Delete is state-only.** The `flush` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; no sessions are restored.

Because re-applying re-runs the flush, use this resource for deliberate, one-shot session-flush workflows rather than for declarative, drift-corrected configuration.


## Example usage

### Flush LSN sessions with filters

The following flushes NAT44 sessions belonging to a specific LSN client, restricted to a subscriber network and a traffic domain:

```hcl
resource "citrixadc_lsnsession_flush" "flush_nat44" {
  nattype    = "NAT44"
  clientname = "lsnclient1"
  network    = "192.168.1.0"
  netmask    = "255.255.255.0"
  td         = 0
}
```

### Flush LSN sessions for a single subscriber

Supplying fewer filters widens the set of sessions that get flushed. The following flushes NAT44 sessions for a single subscriber address:

```hcl
resource "citrixadc_lsnsession_flush" "flush_subscriber" {
  nattype = "NAT44"
  network = "192.168.1.25"
}
```

### Flush LSN sessions by mapped NAT IP and port

```hcl
resource "citrixadc_lsnsession_flush" "flush_natip" {
  natip    = "192.0.2.100"
  natport2 = 5060
}
```


## Argument Reference

All arguments are optional filter selectors that narrow which LSN sessions are flushed. Changing any argument forces the resource to be recreated, which re-runs the flush action.

* `nattype` - (Optional) Type of sessions to flush. If omitted, NITRO applies its server-side default of `NAT44`. Possible values: [ NAT44, DS-Lite, NAT64 ]
* `clientname` - (Optional) Name of the LSN Client entity whose sessions should be flushed.
* `network` - (Optional) IP address or network address of subscriber(s) whose sessions should be flushed.
* `netmask` - (Optional) Subnet mask for the IP address specified by the `network` argument. Must be supplied together with `network`.
* `network6` - (Optional) IPv6 address of the LSN subscriber or B4 device whose sessions should be flushed.
* `td` - (Optional) Traffic domain ID of the LSN client entity whose sessions should be flushed.
* `natip` - (Optional) Mapped NAT IP address used in the LSN sessions to be flushed.
* `natport2` - (Optional) Mapped NAT port used in the LSN sessions to be flushed.
* `nodeid` - (Optional) Unique number that identifies the cluster node whose sessions should be flushed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `lsnsession_flush`. The ADC does not assign an ID to the flush action; this value is purely a Terraform state handle and is not a NITRO lookup key.
