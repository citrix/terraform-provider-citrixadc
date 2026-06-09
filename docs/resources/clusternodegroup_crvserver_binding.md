---
subcategory: "Cluster"
---

# Resource: clusternodegroup_crvserver_binding

Binds a cache redirection (CR) virtual server to a cluster node group on the Citrix ADC. A node group lets you pin entities to a specific subset of cluster nodes; binding a CR vserver to a node group controls where the vserver is placed (spotted or partially striped) so that its traffic is processed only by the nodes in that group rather than being striped across the entire cluster.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name = "ng1"
}

resource "citrixadc_crvserver" "tf_crvserver" {
  name        = "crvs1"
  servicetype = "HTTP"
  ipv46       = "10.10.10.10"
  port        = 80
}

resource "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
  name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
  vserver = citrixadc_crvserver.tf_crvserver.name
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Changing this forces a new resource to be created.
* `vserver` - (Required) Name of the cache redirection virtual server that is bound to this nodegroup. Changing this forces a new resource to be created.

~> **Note** This binding has no NITRO update endpoint and both attributes force replacement. Any change to `name` or `vserver` recreates the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_crvserver_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,vserver:<vserver>`. The values are URL-encoded inside the id so that any reserved characters do not collide with the `key:value` and comma delimiters.


## Import

A clusternodegroup_crvserver_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding "name:ng1,vserver:crvs1"
```
