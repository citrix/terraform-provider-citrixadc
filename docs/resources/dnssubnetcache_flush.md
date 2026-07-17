---
subcategory: "DNS"
---

# Resource: dnssubnetcache_flush

The dnssubnetcache_flush resource performs the imperative `flush` action on the Citrix ADC, evicting entries from the DNS subnet (EDNS Client Subnet, or ECS) cache. When the ADC caches subnet-specific answers on behalf of ECS-aware resolvers, apply this resource to clear those entries — for example after a back-end change so that subsequent client queries are resolved fresh instead of being served from the subnet cache. Set `all = true` to flush every ECS subnet, or set `ecssubnet` to flush the entries for one specific subnet.

This is an action-only resource: applying it invokes the NITRO `flush` action once, and there is no persistent object to read back. The `all` and `ecssubnet` arguments are Optional-only (never Computed), because the flush action has no corresponding NITRO GET endpoint from which values could be populated — this is also why there is no data source. The Read and Update operations are state-only no-ops and Delete simply removes the entry from Terraform state; changing either argument forces resource replacement, which re-runs the `flush` action. Because nothing is persisted on the ADC as a queryable object, there is nothing to import.

The Citrix ADC requires exactly one of `ecssubnet` or `all` to be supplied for a flush: specify one or the other, not both and not neither. This is enforced at plan time.


## Example usage

### Flush the entire DNS subnet (ECS) cache

```hcl
resource "citrixadc_dnssubnetcache_flush" "flush_all" {
  all = true
}
```

### Flush the entries for a specific ECS subnet

```hcl
resource "citrixadc_dnssubnetcache_flush" "flush_subnet" {
  ecssubnet = "192.0.2.0/24"
}
```


## Argument Reference

Exactly one of `all` or `ecssubnet` must be specified.

* `all` - (Optional) Flush all the ECS subnets from the DNS cache. Changing this value forces the `flush` action to re-run (resource replacement).
* `ecssubnet` - (Optional) The specific ECS subnet whose entries should be flushed from the DNS cache (for example, `192.0.2.0/24`). Changing this value forces the `flush` action to re-run (resource replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `dnssubnetcache_flush`. It is purely a Terraform state handle for this action; it is not a server-assigned key and does not correspond to any object on the Citrix ADC.
