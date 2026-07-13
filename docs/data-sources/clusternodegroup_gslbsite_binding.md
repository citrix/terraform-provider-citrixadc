---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_gslbsite_binding

This data source retrieves information about a specific cluster nodegroup to GSLB site binding.

## Example Usage

```hcl
data "citrixadc_clusternodegroup_gslbsite_binding" "example" {
  name     = "my_nodegroup"
  gslbsite = "my_gslb_site"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_gslbsite_binding.example.id
}
```

## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `gslbsite` - (Required) GSLB site that needs to be bound to this nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `name` and `gslbsite` attributes seperated by comma.
