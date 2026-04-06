---
subcategory: "Network"
---

# Data Source: mapbmr_bmrv4network_binding

The mapbmr_bmrv4network_binding data source allows you to retrieve information about a specific mapbmr_bmrv4network_binding resource.


## Example Usage

```terraform
data "citrixadc_mapbmr_bmrv4network_binding" "tf_binding" {
  name    = "tf_mapbmr"
  network = "1.2.3.0"
}

output "name" {
  value = data.citrixadc_mapbmr_bmrv4network_binding.tf_binding.name
}

output "network" {
  value = data.citrixadc_mapbmr_bmrv4network_binding.tf_binding.network
}
```


## Argument Reference

* `name` - (Required) Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the MAP Basic Mapping Rule is created.
* `network` - (Required) IPv4 NAT address range of Customer Edge (CE). parameter.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the mapbmr_bmrv4network_binding. It is a system-generated identifier.
* `netmask` - Subnet mask for the IPv4 address specified in the Network
