---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_streamidentifier_binding

This data source retrieves information about a specific cluster nodegroup to stream identifier binding.

## Example Usage

```hcl
data "citrixadc_clusternodegroup_streamidentifier_binding" "example" {
  name           = "my_nodegroup"
  identifiername = "my_stream_identifier"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_streamidentifier_binding.example.id
}
```

## Argument Reference

* `name` - (Required) Name of the nodegroup to which you want to bind a cluster node or an entity.
* `identifiername` - (Required) Stream identifier and rate limit identifier that needs to be bound to this nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `name` and `identifiername` attributes seperated by comma.
