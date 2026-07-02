---
subcategory: "DNS"
---

# Resource: dnsproxyrecords

The dnsproxyrecords resource performs the imperative `flush` action on the Citrix ADC, clearing DNS records that the appliance has cached while acting as a DNS proxy. Use it to force stale answers out of the proxy cache — for example after a back-end zone change, so that subsequent client queries are resolved fresh instead of being served from the cache. Omit both filters to flush the entire proxy cache, or narrow the flush with `type` (a specific record type such as `A`) and/or `negrectype` (only negative entries such as `NXDOMAIN`).

This is an action-only resource. Applying it invokes `flush` once; there is no managed object to read back. The `type` and `negrectype` arguments are Optional-only (never Computed), because the flush action has no corresponding NITRO GET endpoint from which values could be populated. The Read and Delete operations are state-only no-ops (Delete simply removes the entry from Terraform state), and there is no Update operation — every argument forces resource replacement, which re-runs the `flush` action. Because nothing is persisted on the ADC as a queryable object, there is nothing to import.


## Example usage

### Flush the entire DNS proxy records cache

```hcl
resource "citrixadc_dnsproxyrecords" "tf_dnsproxyrecords" {
}
```

### Flush only A records from the proxy cache

```hcl
resource "citrixadc_dnsproxyrecords" "tf_dnsproxyrecords" {
  type = "A"
}
```

### Flush only negative (NXDOMAIN) entries

```hcl
resource "citrixadc_dnsproxyrecords" "tf_dnsproxyrecords" {
  negrectype = "NXDOMAIN"
}
```


## Argument Reference

* `type` - (Optional) Filter the DNS records to be flushed. For example, `type = "A"` flushes only the A records from the proxy cache. Changing this value forces the `flush` action to re-run (resource replacement). Possible values: [ A, NS, CNAME, SOA, MX, AAAA, SRV, RRSIG, NSEC, DNSKEY, PTR, TXT, NAPTR, CAA ]
* `negrectype` - (Optional) Filter the negative DNS records (that is, `NXDOMAIN` and `NODATA` entries) to be flushed. For example, `negrectype = "NXDOMAIN"` flushes only the NXDOMAIN entries from the cache. Changing this value forces the `flush` action to re-run (resource replacement). Possible values: [ NXDOMAIN, NODATA ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `dnsproxyrecords-config`. It is purely a Terraform state handle for this action; it is not a server-assigned key and cannot be used to look the resource up on the ADC.
