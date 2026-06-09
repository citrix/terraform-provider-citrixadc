---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_streamidentifier_binding

Retrieves information about an existing binding between a cluster node group and a stream identifier. Use this data source to confirm that a given stream identifier is associated with a node group and to reference the binding's computed `id` in other configuration.


## Example usage

```hcl
data "citrixadc_clusternodegroup_streamidentifier_binding" "ng_streamid" {
  name           = "ng1"
  identifiername = "streamid1"
}

output "binding_id" {
  value = data.citrixadc_clusternodegroup_streamidentifier_binding.ng_streamid.id
}
```


## Argument Reference

* `name` - (Required) Name of the node group to which the stream identifier is bound.
* `identifiername` - (Required) Name of the stream identifier bound to this node group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the binding. It is a comma-separated string of `key:value` pairs in the form `name:<name>,identifiername:<identifiername>`, where each value is URL-encoded.
