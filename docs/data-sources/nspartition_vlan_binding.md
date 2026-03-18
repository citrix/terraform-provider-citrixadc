---
subcategory: "NS"
---

# Data Source: nspartition_vlan_binding

The nspartition_vlan_binding data source allows you to retrieve information about a binding between nspartition and vlan resources.


## Example usage

```terraform
data "citrixadc_nspartition_vlan_binding" "tf_binding" {
  partitionname = "tf_nspartition"
  vlan          = 20
}

output "partitionname" {
  value = data.citrixadc_nspartition_vlan_binding.tf_binding.partitionname
}

output "vlan" {
  value = data.citrixadc_nspartition_vlan_binding.tf_binding.vlan
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `vlan` - (Required) Identifier of the vlan that is assigned to this partition.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition_vlan_binding. It is the concatenation of `partitionname` and `vlan` attributes separated by comma.
