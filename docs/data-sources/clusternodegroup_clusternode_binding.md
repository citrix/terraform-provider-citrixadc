---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_clusternode_binding

This data source retrieves information about a specific cluster nodegroup to clusternode binding.

## Example Usage

```hcl
data "citrixadc_clusternodegroup_clusternode_binding" "example" {
  name = "my_nodegroup"
  node = 0
}

output "binding_state" {
  value = data.citrixadc_clusternodegroup_clusternode_binding.example.state
}
```

## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `node` - (Required) Nodes in the nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `name` and `node` attributes seperated by comma.
* `state` - State of the node in the nodegroup.
