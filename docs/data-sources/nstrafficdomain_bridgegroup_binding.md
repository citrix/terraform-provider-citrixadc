---
subcategory: "NS"
---

# Data Source: nstrafficdomain_bridgegroup_binding

The nstrafficdomain_bridgegroup_binding data source allows you to retrieve information about a binding between nstrafficdomain and bridgegroup resources.


## Example usage

```terraform
data "citrixadc_nstrafficdomain_bridgegroup_binding" "tf_binding" {
  td          = 2
  bridgegroup = 2
}

output "td" {
  value = data.citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding.td
}

output "bridgegroup" {
  value = data.citrixadc_nstrafficdomain_bridgegroup_binding.tf_binding.bridgegroup
}
```


## Argument Reference

* `td` - (Required) Integer value that uniquely identifies a traffic domain. Minimum value =  1 Maximum value =  4094
* `bridgegroup` - (Required) ID of the configured bridge to bind to this traffic domain. More than one bridge group can be bound to a traffic domain, but the same bridge group cannot be a part of multiple traffic domains. Minimum value =  1 Maximum value =  1000


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstrafficdomain_bridgegroup_binding. It is the concatenation of `td` and `bridgegroup` attributes separated by comma.
