---
subcategory: "IPSECALG"
---

# Resource: ipsecalgsession

Flushes active IPSec ALG sessions on the Citrix ADC. Applying this resource fires a `?action=flush` NITRO action so operators can clear stale or unwanted IPSec ALG session-table entries, optionally scoped to a specific source, NAT, or destination IP address. A bare apply with no attributes flushes all IPSec ALG sessions.

~> **This resource performs an action, not persistent configuration.** The IPSec ALG session table is populated by the traffic engine, not by the config API. This resource has no persistent NITRO object: Create fires the flush, and Read, Update, and Delete are no-ops. Any change to a scope attribute forces the resource to be recreated, which re-runs the flush. Removing the resource simply drops it from Terraform state; it does not restore previously flushed sessions.

-> **Note:** The IPSec ALG feature must be in use for sessions to exist. Flushing when no sessions match the supplied scope is a valid no-op.

## Example usage

Flush all IPSec ALG sessions:

```hcl
resource "citrixadc_ipsecalgsession" "flush_all" {
}
```

Flush only the sessions for a specific original source IP address:

```hcl
resource "citrixadc_ipsecalgsession" "flush_by_source" {
  sourceip = "192.168.10.5"
}
```

Flush by NAT IP or destination IP address:

```hcl
resource "citrixadc_ipsecalgsession" "flush_by_dest" {
  natip  = "10.100.0.20"
  destip = "203.0.113.40"
}
```

## Argument Reference

All arguments are optional and scope the flush action. Supplying none flushes all sessions. Changing any of these attributes forces the resource to be recreated (re-runs the flush).

* `sourceip` - (Optional) Original source IP address. Restricts the flush to sessions matching this source.
* `natip` - (Optional) Natted source IP address. Restricts the flush to sessions matching this NAT IP.
* `destip` - (Optional) Destination IP address. Restricts the flush to sessions matching this destination.
* `sourceip_alg` - (Optional) Original source IP address. CLI-name twin of `sourceip`; scopes the flush by original source.
* `natip_alg` - (Optional) Natted source IP address. CLI-name twin of `natip`; scopes the flush by NAT IP.
* `destip_alg` - (Optional) Destination IP address. CLI-name twin of `destip`; scopes the flush by destination.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the flush action. It is derived from the supplied scope (for example `flush:sourceip:192.168.10.5`), or `flush-all` when no scope is provided. It does not correspond to a persistent object on the Citrix ADC.
