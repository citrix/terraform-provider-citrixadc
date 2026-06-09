---
subcategory: "Cluster"
---

# Resource: clusternodegroup_service_binding

Binds a load balancing service to a cluster node group on the Citrix ADC. A node group scopes a set of cluster nodes and the entities owned by those nodes; binding a service to the node group pins that service to the node group's spotted/striped placement, so traffic for the service is handled by the nodes in the group rather than being distributed across the entire cluster.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
  name = "ng1"
}

resource "citrixadc_service" "tf_service" {
  name        = "svc1"
  ip          = "10.10.10.10"
  servicetype = "HTTP"
  port        = 80
}

resource "citrixadc_clusternodegroup_service_binding" "tf_clusternodegroup_service_binding" {
  name    = citrixadc_clusternodegroup.tf_clusternodegroup.name
  service = citrixadc_service.tf_service.name
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Changing this forces a new resource to be created.
* `service` - (Required) Name of the service bound to this nodegroup. Changing this forces a new resource to be created.

~> **Note** This binding has no NITRO update endpoint, and both attributes force replacement. Any change to `name` or `service` recreates the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_service_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,service:<service>`, where each value is URL-encoded so that any special characters do not collide with the `key:value` and comma delimiters.


## Import

A clusternodegroup_service_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_clusternodegroup_service_binding.tf_clusternodegroup_service_binding "name:ng1,service:svc1"
```
