---
subcategory: "DNS"
---

# Data Source `dnsaction64`

The dnsaction64 data source allows you to retrieve information about a DNS64 action configuration.


## Example usage

```terraform
data "citrixadc_dnsaction64" "tf_dnsaction64" {
  actionname = "default_DNS64_action1"
}

output "prefix" {
  value = data.citrixadc_dnsaction64.tf_dnsaction64.prefix
}

output "mappedrule" {
  value = data.citrixadc_dnsaction64.tf_dnsaction64.mappedrule
}
```


## Argument Reference

The following arguments are supported:

* `actionname` - (Required) Name of the dns64 action.

## Attribute Reference

The following attributes are available:

* `id` - The id of the dnsaction64. It is a system-generated identifier.
* `actionname` - Name of the dns64 action.
* `prefix` - The dns64 prefix to be used if the after evaluating the rules.
* `mappedrule` - The expression to select the criteria for ipv4 addresses to be used for synthesis. Only if the mappedrule is evaluated to true the corresponding ipv4 address is used for synthesis using respective prefix, otherwise the A RR is discarded.
* `excluderule` - The expression to select the criteria for eliminating the corresponding ipv6 addresses from the response.
