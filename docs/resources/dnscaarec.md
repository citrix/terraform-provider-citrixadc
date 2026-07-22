---
subcategory: "DNS"
---

# Resource: dnscaarec

Publishes a DNS Certification Authority Authorization (CAA) resource record on the Citrix ADC. CAA records let a domain owner declare which certificate authorities are permitted to issue certificates for a domain, helping prevent unauthorized certificate issuance.

CAA resource records cannot be modified on the ADC. Changing any attribute forces the record to be destroyed and recreated.

## Example usage

```hcl
resource "citrixadc_dnscaarec" "dnscaarec" {
  domain      = "example.com"
  tag         = "issue"
  valuestring = "letsencrypt.org"
  flag        = "NONE"
  ttl         = 3600
}
```


## Argument Reference

* `domain` - (Required) Domain name of the CAA record. Changing this forces a new resource to be created.
* `valuestring` - (Required) Value associated with the chosen property tag in the CAA resource record. Enclose the string in single or double quotation marks. Changing this forces a new resource to be created.
* `tag` - (Optional) String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - `issue`, `issuewild` and `iodef`. Defaults to `"issue"`. Changing this forces a new resource to be created.
* `flag` - (Optional) Flag associated with the CAA record. Possible values: [ NONE, CRITICAL ]. Changing this forces a new resource to be created.
* `ttl` - (Optional) Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600. Defaults to `3600`. Changing this forces a new resource to be created.
* `ecssubnet` - (Optional) Subnet for which the cached CAA record needs to be removed. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnscaarec. It is a composite of the `domain` and `recordid` attributes, formatted as `domain:<domain>,recordid:<recordid>`.
* `recordid` - Unique, internally generated record ID assigned by the Citrix ADC. This value is the server-assigned identifier of the CAA record and is required to delete the record (Delete targets the record by `domain` plus `recordid`). If you import an existing record, this value is populated from the ADC. Multiple CAA records can share the same domain, so the `recordid` is what distinguishes an individual record.


## Import

A dnscaarec can be imported using its id (`domain:<domain>,recordid:<recordid>`), e.g.

```shell
terraform import citrixadc_dnscaarec.dnscaarec domain:example.com,recordid:12345
```
