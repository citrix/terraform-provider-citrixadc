---
subcategory: "Network"
---

# Data Source: nstrafficdomain

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

### Read-only nstrafficdomain metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_nstrafficdomain` resource) and are computed. Any attribute the appliance does not return is `null`.

* `state` - The state of the traffic domain (for example `ENABLED`, `DISABLED`).
