---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_nslimitidentifier_binding

Retrieves information about an existing binding between a cluster node group and an `nslimit` rate-limit identifier. Use this data source to confirm that a given rate-limit identifier is associated with a node group and to reference the binding's computed `id` in other configuration.


## Example usage

```hcl
data "citrixadc_clusternodegroup_nslimitidentifier_binding" "ng_ratelimit" {
  name           = "ng1"
  identifiername = "ratelimit1"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_nslimitidentifier_binding.ng_ratelimit.id
}
```


## Argument Reference

* `name` - (Required) Name of the node group to which the rate-limit identifier is bound.
* `identifiername` - (Required) Name of the `nslimit` rate-limit identifier bound to this node group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding. It is a comma-separated string of `key:value` pairs in the form `name:<name>,identifiername:<identifiername>`, where each value is URL-encoded.
