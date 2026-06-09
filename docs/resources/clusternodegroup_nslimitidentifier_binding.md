---
subcategory: "Cluster"
---

# Resource: clusternodegroup_nslimitidentifier_binding

Associates an `nslimit` rate-limit identifier with a cluster node group. Binding a rate-limit identifier to a node group lets the cluster apply that identifier's rate-limiting (throttling) policy on the spotted nodes that own the node group, so traffic handled by the group is counted and limited consistently across the cluster.

This binding is immutable: both `name` and `identifiername` form the resource identity, and changing either forces Terraform to destroy and recreate the binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup" "ng" {
  name = "ng1"
}

resource "citrixadc_nslimitidentifier" "rate_limit" {
  limitidentifier = "ratelimit1"
  threshold       = 100
  mode            = "REQUEST_RATE"
}

resource "citrixadc_clusternodegroup_nslimitidentifier_binding" "ng_ratelimit" {
  name           = citrixadc_clusternodegroup.ng.name
  identifiername = citrixadc_nslimitidentifier.rate_limit.limitidentifier
}
```


## Argument Reference

* `name` - (Required) Name of the node group to which you want to bind the rate-limit identifier. Changing this forces a new resource to be created.
* `identifiername` - (Required) Name of the `nslimit` rate-limit identifier that needs to be bound to this node group. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the `clusternodegroup_nslimitidentifier_binding` resource. It is a comma-separated string of `key:value` pairs in the form `name:<name>,identifiername:<identifiername>`, where each value is URL-encoded.


## Import

A `clusternodegroup_nslimitidentifier_binding` can be imported using its ID, which is the composite of the `name` and `identifiername` attributes in the form `name:<name>,identifiername:<identifiername>`, e.g.

```shell
terraform import citrixadc_clusternodegroup_nslimitidentifier_binding.ng_ratelimit name:ng1,identifiername:ratelimit1
```
