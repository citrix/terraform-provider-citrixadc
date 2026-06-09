---
subcategory: "Cluster"
---

# Resource: clusternodegroup_streamidentifier_binding

Associates a stream identifier with a cluster node group. A stream identifier collects and aggregates real-time traffic statistics (request counts, bandwidth, response times) for the flows it tracks. Binding a stream identifier to a node group scopes that data collection to the spotted nodes that own the group, so the cluster maintains the identifier's statistics consistently for the traffic handled by that group.

This binding is immutable: both `name` and `identifiername` form the resource identity, and changing either forces Terraform to destroy and recreate the binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "ng" {
  name = "ng1"
}

resource "citrixadc_streamidentifier" "si" {
  name         = "streamid1"
  selectorname = "streamselector1"
  samplecount  = 10
  sort         = "CONNECTIONS"
}

resource "citrixadc_clusternodegroup_streamidentifier_binding" "ng_streamid" {
  name           = citrixadc_clusternodegroup.ng.name
  identifiername = citrixadc_streamidentifier.si.name
}
```


## Argument Reference

* `name` - (Required) Name of the node group to which you want to bind the stream identifier. Changing this forces a new resource to be created.
* `identifiername` - (Required) Name of the stream identifier (and rate-limit identifier) that needs to be bound to this node group. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `clusternodegroup_streamidentifier_binding` resource. It is a comma-separated string of `key:value` pairs in the form `name:<name>,identifiername:<identifiername>`, where each value is URL-encoded.


## Import

A `clusternodegroup_streamidentifier_binding` can be imported using its ID, which is the composite of the `name` and `identifiername` attributes in the form `name:<name>,identifiername:<identifiername>`, e.g.

```shell
terraform import citrixadc_clusternodegroup_streamidentifier_binding.ng_streamid name:ng1,identifiername:streamid1
```
