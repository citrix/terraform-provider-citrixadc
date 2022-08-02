---
subcategory: "Cluster"
---

# Resource: clusternodegroup_csvserver_binding

The clusternodegroup_csvserver_binding resource is used to create clusternodegroup_csvserver_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
  name = "my_cs_group"
  vserver = "my_csvserver"
}
```


## Argument Reference

* `vserver` - (Required) vserver that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_csvserver_binding. It has the same value as the `name` attribute.


## Import

A clusternodegroup_csvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding my_cs_group,my_csvserver
```
