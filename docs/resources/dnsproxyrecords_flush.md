---
subcategory: "DNS"
---

# Resource: dnsproxyrecords_flush

The dnsproxyrecords_flush resource performs the imperative `flush` action on the Citrix ADC, clearing DNS records that the appliance has cached while acting as a DNS proxy. Use it to force stale answers out of the proxy cache — for example after a back-end zone change, so that subsequent client queries are resolved fresh instead of being served from the cache. Omit both filters to flush the entire proxy cache, or narrow the flush with `type` (a specific record type such as `A`) and/or `negrectype` (only negative entries such as `NXDOMAIN`).

This is an action resource: applying it performs the `flush` action; it does not manage a persistent object, so re-applying re-runs the action.


## Example usage

### Flush the entire DNS proxy records cache

```hcl
resource "citrixadc_dnsproxyrecords_flush" "tf_dnsproxyrecords_flush" {
}
```

### Flush only A records from the proxy cache

```hcl
resource "citrixadc_dnsproxyrecords_flush" "tf_dnsproxyrecords_flush" {
  type = "A"
}
```

### Flush only negative (NXDOMAIN) entries

```hcl
resource "citrixadc_dnsproxyrecords_flush" "tf_dnsproxyrecords_flush" {
  negrectype = "NXDOMAIN"
}
```


## Argument Reference

* `type` - (Optional) Filter the DNS records to be flushed. For example, `type = "A"` flushes only the A records from the proxy cache. Changing this value forces the `flush` action to re-run (resource replacement). Possible values: [ A, NS, CNAME, SOA, MX, AAAA, SRV, RRSIG, NSEC, DNSKEY, PTR, TXT, NAPTR, CAA ]
* `negrectype` - (Optional) Filter the negative DNS records (that is, `NXDOMAIN` and `NODATA` entries) to be flushed. For example, `negrectype = "NXDOMAIN"` flushes only the NXDOMAIN entries from the cache. Changing this value forces the `flush` action to re-run (resource replacement). Possible values: [ NXDOMAIN, NODATA ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsproxyrecords_flush resource. It is set to `dnsproxyrecords_flush`.
