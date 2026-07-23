---
subcategory: "DNS"
---

# Resource: dnssubnetcache_flush

This resource is used to flush the DNS subnet (ECS) cache on the Citrix ADC.


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
