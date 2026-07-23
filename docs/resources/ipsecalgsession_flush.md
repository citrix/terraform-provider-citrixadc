---
subcategory: "IPSECALG"
---

# Resource: ipsecalgsession_flush

This resource is used to flush active IPSec ALG sessions on the Citrix ADC.


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
* `sourceip_alg` - (Optional) Original source IP address. Use `sourceip` to scope the flush. Changing this attribute re-triggers the flush.
* `natip_alg` - (Optional) Natted source IP address. Use `natip` to scope the flush. Changing this attribute re-triggers the flush.
* `destip_alg` - (Optional) Destination IP address. Use `destip` to scope the flush. Changing this attribute re-triggers the flush.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ipsecalgsession_flush resource. It is set to `ipsecalgsession_flush`.
