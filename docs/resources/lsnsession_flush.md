---
subcategory: "LSN"
---

# Resource: lsnsession_flush

This resource is used to flush active Large Scale NAT (LSN) sessions on the Citrix ADC.


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

* `id` - The id of the lsnsession_flush resource. It is set to `lsnsession_flush`.
