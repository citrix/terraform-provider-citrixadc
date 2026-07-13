---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_lbvserver_binding

This data source retrieves information about a specific cluster nodegroup to load balancing virtual server binding.

## Example Usage

```hcl
data "citrixadc_clusternodegroup_lbvserver_binding" "example" {
  name    = "my_nodegroup"
  vserver = "my_lbvserver"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_lbvserver_binding.example.id
}
```

## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `vserver` - (Required) Load balancing virtual server that needs to be bound to this nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `name` and `vserver` attributes seperated by comma.
