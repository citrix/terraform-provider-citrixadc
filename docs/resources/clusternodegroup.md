---
subcategory: "Cluster"
---

# Resource: clusternodegroup

The clusternodegroup resource is used to create clusternodegroup.


## Example usage

```hcl
<fillme>
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1
* `strict` - (Optional) Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup. * Enabled - When one of the nodes goes down, no other cluster node is picked up to replace it. When the node comes up, it will continue being part of the nodegroup. * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced. Possible values: [ YES, NO ]
* `sticky` - (Optional) Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster. * Enabled - When owner node goes down, backup node will become the owner node and takes the traffic for the entities bound to the nodegroup. When bound node rejoins the cluster, traffic for the entities bound to nodegroup will not be steered back to this bound node. Current owner will have the ownership till it goes down. * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced. Possible values: [ YES, NO ]
* `state` - (Optional) State of the nodegroup. All the nodes binding to this nodegroup must have the same state. ACTIVE/SPARE/PASSIVE. Possible values: [ ACTIVE, SPARE, PASSIVE ]
* `priority` - (Optional) Priority of Nodegroup. This priority is used for all the nodes bound to the nodegroup for Nodegroup selection. Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup. It has the same value as the `name` attribute.

