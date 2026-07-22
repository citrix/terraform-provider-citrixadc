---
subcategory: "DNS"
---

# Resource: dnssubnetcache_flush

The dnssubnetcache_flush resource performs the imperative `flush` action on the Citrix ADC, evicting entries from the DNS subnet (EDNS Client Subnet, or ECS) cache. When the ADC caches subnet-specific answers on behalf of ECS-aware resolvers, apply this resource to clear those entries — for example after a back-end change so that subsequent client queries are resolved fresh instead of being served from the subnet cache. Set `all = true` to flush every ECS subnet, or set `ecssubnet` to flush the entries for one specific subnet.

This is an action resource: applying it performs the `flush` action; it does not manage a persistent object, so re-applying re-runs the action.

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

* `id` - The id of the dnssubnetcache_flush resource. It is set to `dnssubnetcache_flush`.
