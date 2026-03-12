---
subcategory: "NS"
---

# Data Source: nspartition_bridgegroup_binding

The nspartition_bridgegroup_binding data source allows you to retrieve information about a binding between nspartition and bridgegroup resources.


## Example usage

```terraform
data "citrixadc_nspartition_bridgegroup_binding" "tf_binding" {
  partitionname = "tf_nspartition"
  bridgegroup   = 2
}

output "partitionname" {
  value = data.citrixadc_nspartition_bridgegroup_binding.tf_binding.partitionname
}

output "bridgegroup" {
  value = data.citrixadc_nspartition_bridgegroup_binding.tf_binding.bridgegroup
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `bridgegroup` - (Required) Identifier of the bridge group that is assigned to this partition.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nspartition_bridgegroup_binding. It is the concatenation of `partitionname` and `bridgegroup` attributes separated by comma.
