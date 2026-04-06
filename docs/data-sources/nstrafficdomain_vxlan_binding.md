---
subcategory: "NS"
---

# Data Source: nstrafficdomain_vxlan_binding

The nstrafficdomain_vxlan_binding data source allows you to retrieve information about the binding between a traffic domain and a VXLAN.

## Example Usage

```terraform
data "citrixadc_nstrafficdomain_vxlan_binding" "tf_binding" {
  td    = 2
  vxlan = 123
}

output "td" {
  value = data.citrixadc_nstrafficdomain_vxlan_binding.tf_binding.td
}

output "vxlan" {
  value = data.citrixadc_nstrafficdomain_vxlan_binding.tf_binding.vxlan
}
```

## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain.
* `vxlan` - (Required) ID of the VXLAN to bind to this traffic domain. More than one VXLAN can be bound to a traffic domain, but the same VXLAN cannot be a part of multiple traffic domains.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_vxlan_binding. It is the concatenation of `td` and `vxlan` attributes separated by comma.
