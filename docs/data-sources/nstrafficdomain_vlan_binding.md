---
subcategory: "NS"
---

# Data Source: nstrafficdomain_vlan_binding

The nstrafficdomain_vlan_binding data source allows you to retrieve information about a binding between nstrafficdomain and vlan resources.


## Example usage

```terraform
data "citrixadc_nstrafficdomain_vlan_binding" "tf_binding" {
  td   = 2
  vlan = 20
}

output "td" {
  value = data.citrixadc_nstrafficdomain_vlan_binding.tf_binding.td
}

output "vlan" {
  value = data.citrixadc_nstrafficdomain_vlan_binding.tf_binding.vlan
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `vlan` - (Required) ID of the VLAN to bind to this traffic domain. More than one VLAN can be bound to a traffic domain, but the same VLAN cannot be a part of multiple traffic domains. Minimum value =  1 Maximum value =  4094


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_vlan_binding. It is the concatenation of `td` and `vlan` attributes separated by comma.
