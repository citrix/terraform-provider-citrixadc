---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_service_binding

This data source retrieves information about a specific cluster nodegroup to service binding.

## Example Usage

```hcl
data "citrixadc_clusternodegroup_service_binding" "example" {
  name    = "my_nodegroup"
  service = "my_service"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_service_binding.example.id
}
```

## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `service` - (Required) Name of the service that needs to be bound to this nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `name` and `service` attributes seperated by comma.
