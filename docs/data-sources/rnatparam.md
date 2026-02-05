---
subcategory: "Network"
---

# Data Source: citrixadc_rnatparam

The rnatparam data source allows you to retrieve information about RNAT parameters configuration.

## Example Usage

```terraform
data "citrixadc_rnatparam" "tf_rnatparam" {
}

output "srcippersistency" {
  value = data.citrixadc_rnatparam.tf_rnatparam.srcippersistency
}

output "tcpproxy" {
  value = data.citrixadc_rnatparam.tf_rnatparam.tcpproxy
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `srcippersistency` - Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip. Possible values: `ENABLED`, `DISABLED`.
* `tcpproxy` - Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features. Possible values: `ENABLED`, `DISABLED`.
* `id` - The id of the rnatparam resource.
