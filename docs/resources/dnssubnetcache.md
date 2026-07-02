---
subcategory: "DNS"
---

# Resource: dnssubnetcache

The dnssubnetcache resource performs the imperative `flush` action on the Citrix ADC, clearing entries from the DNS subnet (EDNS Client Subnet, or ECS) cache. When the ADC caches subnet-specific answers on behalf of ECS-aware resolvers, use this resource to evict those entries — for example after a back-end change so that subsequent client queries are resolved fresh instead of being served from the subnet cache. Set `all = true` to flush every ECS subnet, or set `ecssubnet` to flush the entries for one specific subnet.

This is an action-only resource. Applying it invokes `flush` once; there is no managed object to read back. The `all` and `ecssubnet` arguments are Optional-only (never Computed), because the flush action has no corresponding NITRO GET endpoint from which values could be populated. The Read and Delete operations are state-only no-ops (Delete simply removes the entry from Terraform state), and there is no Update operation — either argument forces resource replacement, which re-runs the `flush` action. Because nothing is persisted on the ADC as a queryable object, there is nothing to import.

The Citrix ADC CLI requires exactly one of `ecssubnet` or `all` to be supplied for a flush; specify one or the other, not both and not neither.


## Example usage

### Flush the entire DNS subnet (ECS) cache

```hcl
resource "citrixadc_dnssubnetcache" "tf_dnssubnetcache" {
  all = true
}
```

### Flush the entries for a specific ECS subnet

```hcl
resource "citrixadc_dnssubnetcache" "tf_dnssubnetcache" {
  ecssubnet = "192.0.2.0/24"
}
```


## Argument Reference

* `all` - (Optional) Flush all the ECS subnets from the DNS cache. Changing this value forces the `flush` action to re-run (resource replacement). Specify exactly one of `all` or `ecssubnet`.
* `ecssubnet` - (Optional) The specific ECS subnet whose entries should be flushed from the DNS cache (for example, `192.0.2.0/24`). Changing this value forces the `flush` action to re-run (resource replacement). Specify exactly one of `all` or `ecssubnet`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `dnssubnetcache-flush`. It is purely a Terraform state handle for this action; it is not a server-assigned key and cannot be used to look the resource up on the ADC.
