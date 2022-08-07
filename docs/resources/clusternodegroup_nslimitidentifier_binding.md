---
subcategory: "Cluster"
---

# Resource: clusternodegroup_nslimitidentifier_binding

The clusternodegroup_nslimitidentifier_binding resource is used to create clusternodegroup_nslimitidentifier_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_nslimitidentifier_binding" "tf_clusternodegroup_nslimitidentifier_binding" {
  name           = "my_group"
  identifiername = "my_nslimitidentifier"
}

```


## Argument Reference

* `identifiername` - (Required) stream identifier  and rate limit identifier that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup to which you want to bind a cluster node or an entity. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_nslimitidentifier_binding. It is the concatenation of `name` and `identifiername` attributes separated by a comma.


## Import

A clusternodegroup_nslimitidentifier_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding my_group,my_nslimitidentifier
```
