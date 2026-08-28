---
subcategory: "DNS"
---

# Resource: dnssvcbrec

The dnssvcbrec resource is used to create DNS SVCB/HTTPS service binding records.


## Example usage

```hcl
resource "citrixadc_dnssvcbrec" "dnssvcbrec" {
  domain     = "svcb.example.com"
  targetname = "target.example.com"
  priority   = 1
  svcbtype   = "HTTPS"
  alpn       = "h2,h3"
  port       = 443
  ttl        = 3600
}
```


## Argument Reference

* `domain` - (Required) Domain name for the SVCB/HTTPS record.
* `targetname` - (Required) Target domain name.
* `priority` - (Required) Service priority (0 for AliasMode, >0 for ServiceMode).
* `svcbtype` - (Optional) Service type: SVCB or HTTPS. Possible values = SVCB, HTTPS.
* `alpn` - (Optional) Comma-separated list of ALPN protocol identifiers.
* `encryptedclienthello` - (Optional) Base64-encoded ECH configuration.
* `ipv4hint` - (Optional) Comma-separated list of IPv4 hint addresses.
* `ipv6hint` - (Optional) Comma-separated list of IPv6 hint addresses.
* `mandatory` - (Optional) Comma-separated list of mandatory SvcParam keys.
* `nodefaultalpn` - (Optional) Indicates no default ALPN protocols.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `port` - (Optional) Port number for the service.
* `ttl` - (Optional) Time to Live (TTL) in seconds.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnssvcbrec. It is a combination of `domain`, `targetname`, `priority` and `svcbtype`.


## Import

A dnssvcbrec can be imported using its id (`domain,targetname,priority,svcbtype`), e.g.

```shell
terraform import citrixadc_dnssvcbrec.dnssvcbrec svcb.example.com,target.example.com,1,HTTPS
```
