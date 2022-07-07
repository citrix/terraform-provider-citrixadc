---
subcategory: "DNS"
---

# Resource: dnsaction64

The dnsaction64 resource is used to create DNS action64.


## Example usage

```hcl
resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.TYPE.EQ(A)"
    excluderule = "DNS.RR.TYPE.EQ(AAAA)"
}

```


## Argument Reference

* `actionname` - (Required) Name of the dns64 action.
* `prefix` - (Required) The dns64 prefix to be used if the after evaluating the rules
* `excluderule` - (Optional) The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.
* `mappedrule` - (Optional) The expression to select the criteria for ipv4 addresses to be used for synthesis.                       Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix,                        otherwise the A RR is discarded


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsaction64. It has the same value as the `actionname` attribute.


## Import

A dnsaction64 can be imported using its name, e.g.

```shell
terraform import citrixadc_dnsaction64.dnsaction64 default_DNS64_action1
```
