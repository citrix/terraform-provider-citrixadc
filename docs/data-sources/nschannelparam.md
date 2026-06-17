---
subcategory: "NS"
---

# Data Source: nschannelparam

The nschannelparam data source allows you to retrieve the global channel (link aggregation / VF) parameters configured on the Citrix ADC appliance, such as the VF autorecover mode.


## Example usage

```terraform
data "citrixadc_nschannelparam" "example" {
}

output "vf_autorecover" {
  value = data.citrixadc_nschannelparam.example.vfautorecover
}
```


## Argument Reference

This datasource is a singleton and does not require any arguments. All attributes are computed.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nschannelparam datasource. Set to the constant string `nschannelparam-config`.
* `vfautorecover` - VF autorecover mode for channels. Possible values: `DISABLE`, `ENABLE`.
