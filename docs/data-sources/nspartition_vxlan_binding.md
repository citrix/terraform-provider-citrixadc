---
subcategory: "NS"
---

# Data Source: nspartition_vxlan_binding

The nspartition_vxlan_binding data source allows you to retrieve information about a nspartition_vxlan_binding.


## Example Usage

```terraform
data "citrixadc_nspartition_vxlan_binding" "tf_binding" {
  partitionname = "tf_nspartition"
  vxlan         = 123
}

output "partitionname" {
  value = data.citrixadc_nspartition_vxlan_binding.tf_binding.partitionname
}

output "vxlan" {
  value = data.citrixadc_nspartition_vxlan_binding.tf_binding.vxlan
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `vxlan` - (Required) Identifier of the vxlan that is assigned to this partition.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition_vxlan_binding. It is the concatenation of the `partitionname` and `vxlan` attributes separated by a comma.
