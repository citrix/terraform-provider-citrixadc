---
subcategory: "Cluster"
---

# Resource: clusternodegroup

A cluster node group is a logical grouping of one or more cluster nodes that lets you control where spotted and partially-striped virtual servers are placed within a Citrix ADC cluster. By assigning a state (ACTIVE, SPARE, or PASSIVE) and a priority to the group, you decide which nodes own the traffic for the entities bound to the group and which nodes act as backups, giving you fine-grained control over traffic distribution and failover behavior in a cluster.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name     = "ng1"
  state    = "ACTIVE"
  priority = 10
  strict   = "YES"
  sticky   = "NO"
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Changing this attribute forces a new resource to be created.
* `state` - (Optional) State of the nodegroup. All the nodes binding to this nodegroup must have the same state. This value can be updated in place. Possible values: [ ACTIVE, SPARE, PASSIVE ]
* `priority` - (Optional) Priority of the nodegroup. This priority is used for all the nodes bound to the nodegroup for nodegroup selection. This value can be updated in place. Minimum value = 0 Maximum value = 31
* `strict` - (Optional) Specifies whether cluster nodes that are not part of the nodegroup will be used as backup for the nodegroup. When set to `NO`, a non-nodegroup cluster node is picked up to act as part of the nodegroup if one of the nodes goes down; when set to `YES`, no other cluster node is picked up to replace a failed node. This value can be updated in place. Defaults to `NO` on the server. Possible values: [ YES, NO ]
* `sticky` - (Optional) Specifies whether to preempt the traffic for the entities bound to the nodegroup when the owner node goes down and rejoins the cluster. When set to `YES`, only one node can be bound to the nodegroup and, after the owner node goes down, the backup node retains ownership even after the original node rejoins. This attribute is **create-only**: changing it forces the resource to be replaced (the NITRO set/PUT operation does not accept `sticky`). Defaults to `NO` on the server. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup. It has the same value as the `name` attribute.


## Import

A clusternodegroup can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup.tf_clusternodegroup ng1
```
