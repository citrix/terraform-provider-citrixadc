---
subcategory: "Cluster"
---

# Data Source: clusternodegroup

The clusternodegroup data source allows you to retrieve information about an existing cluster node group, such as its state, priority, and backup behavior, by looking it up by name.


## Example usage

```terraform
data "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name = "ng1"
}

output "state" {
  value = data.citrixadc_clusternodegroup.tf_clusternodegroup.state
}

output "priority" {
  value = data.citrixadc_clusternodegroup.tf_clusternodegroup.priority
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup. It has the same value as the `name` attribute.
* `state` - State of the nodegroup. All the nodes binding to this nodegroup must have the same state. Possible values: [ ACTIVE, SPARE, PASSIVE ]
* `priority` - Priority of the nodegroup. This priority is used for all the nodes bound to the nodegroup for nodegroup selection.
* `strict` - Specifies whether cluster nodes that are not part of the nodegroup will be used as backup for the nodegroup. Possible values: [ YES, NO ]
* `sticky` - Specifies whether to preempt the traffic for the entities bound to the nodegroup when the owner node goes down and rejoins the cluster. When enabled, only one node can be bound to the nodegroup. Possible values: [ YES, NO ]
