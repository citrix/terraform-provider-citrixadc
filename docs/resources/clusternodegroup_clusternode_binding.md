---
subcategory: "Cluster"
---

# Resource: clusternodegroup_clusternode_binding

The clusternodegroup_clusternode_binding resource is used to create clusternodegroup_clusternode_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
  name = "my_group"
  node = 2
}

```


## Argument Reference

* `node` - (Required) Nodes in the nodegroup. Minimum value =  0 Maximum value =  31
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_clusternode_binding. It is the concatenation of `name` and `node` attributes seperated by a comma.


## Import

A clusternodegroup_clusternode_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding my_group,2
```
