---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_clusternode_binding

The clusternodegroup_clusternode_binding data source allows you to retrieve information about a cluster node bound to a cluster nodegroup on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_clusternodegroup_clusternode_binding" "example" {
  name = "ng1"
  node = 0
}

output "bound_node" {
  value = data.citrixadc_clusternodegroup_clusternode_binding.example.node
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `node` - (Required) Node id of the cluster node bound to the nodegroup. This is an integer. Minimum value = `0` Maximum value = `31`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_clusternode_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,node:<node>`.
* `name` - Name of the nodegroup.
* `node` - Node id of the cluster node bound to the nodegroup.
