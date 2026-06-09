---
subcategory: "Cluster"
---

# Resource: clusternodegroup_clusternode_binding

Assigns a specific cluster node to a cluster nodegroup on the Citrix ADC. Nodegroups let you pin spotted and partially-striped virtual servers to a defined subset of cluster nodes; binding a node to the nodegroup adds that node to the set of nodes on which those vservers are placed and activated.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name = "ng1"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
  name = citrixadc_clusternodegroup.tf_clusternodegroup.name
  node = 0
}
```

~> **Note** The `node` value must refer to an existing cluster node (added via `add cluster node`). Binding fails on the ADC if the node id does not correspond to a configured cluster node.


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Changing this forces a new resource to be created.
* `node` - (Required) Node id of the cluster node to bind to the nodegroup. This is an integer. Minimum value = `0` Maximum value = `31`. Changing this forces a new resource to be created.

~> **Note** This binding has no NITRO update endpoint and every attribute forces replacement. Any change to `name` or `node` recreates the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_clusternode_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,node:<node>`. The `name` value is URL-encoded inside the id; `node` is a plain integer.


## Import

A clusternodegroup_clusternode_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding "name:ng1,node:0"
```
