---
subcategory: "NS"
---

# Data Source `nshostname`

The nshostname data source allows you to retrieve information about the hostname configured on the Citrix ADC appliance.


## Example usage

```terraform
data "citrixadc_nshostname" "example" {
}

output "adc_hostname" {
  value = data.citrixadc_nshostname.example.hostname
}
```


## Argument Reference

This datasource does not require any arguments. All attributes are optional and computed.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nshostname datasource.
* `hostname` - Host name for the Citrix ADC.
* `ownernode` - ID of the cluster node for which you are setting the hostname. Can be configured only through the cluster IP address.
