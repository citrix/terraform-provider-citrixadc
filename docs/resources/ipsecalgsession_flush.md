---
subcategory: "IPSECALG"
---

# Resource: ipsecalgsession_flush

The ipsecalgsession_flush resource flushes active IPSec ALG (Application Layer Gateway) sessions on the Citrix ADC. It is an action-only resource: applying it invokes the NITRO `?action=flush` action, which clears entries from the IPSec ALG session table so that new IPSec traffic re-establishes through the current configuration. This is useful for clearing stale mappings, for troubleshooting, or for forcing sessions off a particular source, NAT, or destination IP address.

This resource does not create, read, or manage a persistent object on the appliance. The IPSec ALG session table is populated by the traffic engine, not by the configuration API, and there is no inverse "un-flush" action. Each apply performs the flush; changing any scope argument forces the resource to be recreated, which re-runs the flush. Removing the resource simply drops it from Terraform state and does not restore previously flushed sessions.

-> **Note:** The IPSec ALG feature must be in use for sessions to exist. Flushing when no sessions match the supplied scope is a valid no-op.


## Example usage

### Flush all IPSec ALG sessions

With no scope arguments, all IPSec ALG sessions are flushed.

```hcl
resource "citrixadc_ipsecalgsession_flush" "flush_all" {}
```

### Flush IPSec ALG sessions by source IP

```hcl
resource "citrixadc_ipsecalgsession_flush" "flush_by_source" {
  sourceip = "192.168.10.5"
}
```

### Flush IPSec ALG sessions by NAT IP and destination IP

```hcl
resource "citrixadc_ipsecalgsession_flush" "flush_by_natip_destip" {
  natip  = "10.100.0.20"
  destip = "203.0.113.40"
}
```


## Argument Reference

All arguments are optional and scope the flush action. Supplying none flushes all sessions. Changing any of these arguments forces the resource to be recreated (re-runs the flush).

* `sourceip` - (Optional) Original source IP address. Restricts the flush to sessions matching this source. Changing this attribute re-triggers the flush.
* `natip` - (Optional) Natted source IP address. Restricts the flush to sessions matching this NAT IP. Changing this attribute re-triggers the flush.
* `destip` - (Optional) Destination IP address. Restricts the flush to sessions matching this destination. Changing this attribute re-triggers the flush.
* `sourceip_alg` - (Optional) Original source IP address. This is the NITRO GET filter name; it is accepted for compatibility but is **not** sent in the flush payload. Use `sourceip` to scope the flush. Changing this attribute re-triggers the flush.
* `natip_alg` - (Optional) Natted source IP address. This is the NITRO GET filter name; it is accepted for compatibility but is **not** sent in the flush payload. Use `natip` to scope the flush. Changing this attribute re-triggers the flush.
* `destip_alg` - (Optional) Destination IP address. This is the NITRO GET filter name; it is accepted for compatibility but is **not** sent in the flush payload. Use `destip` to scope the flush. Changing this attribute re-triggers the flush.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `ipsecalgsession_flush`. It does not correspond to any object on the Citrix ADC.
