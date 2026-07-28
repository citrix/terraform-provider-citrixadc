---
subcategory: "DNS"
---

# Resource: dnsproxyrecords_flush

This resource is used to flush cached DNS proxy records on the Citrix ADC.


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
