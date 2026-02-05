---
subcategory: "Network"
---

# Data Source `nstrafficdomain`

The nstrafficdomain data source allows you to retrieve information about a traffic domain configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nstrafficdomain" "tf_trafficdomain" {
  td = 2
}

output "aliasname" {
  value = data.citrixadc_nstrafficdomain.tf_trafficdomain.aliasname
}

output "vmac" {
  value = data.citrixadc_nstrafficdomain.tf_trafficdomain.vmac
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `aliasname` - Name of traffic domain being added.
* `vmac` - Associate the traffic domain with a VMAC address instead of with VLANs.
