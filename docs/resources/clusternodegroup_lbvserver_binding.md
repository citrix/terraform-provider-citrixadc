---
subcategory: "Cluster"
---

# Resource: clusternodegroup_lbvserver_binding

The clusternodegroup_lbvserver_binding resource is used to create clusternodegroup_lbvserver_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
  name = "my_test_group"
  vserver = "my_lbvserver"
}
```


## Argument Reference

* `vserver` - (Required) vserver that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_lbvserver_binding. It is the concatenation of `name` and `vserver` attributes separated by a comma.


## Import

A clusternodegroup_lbvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding my_test_group,my_lbvserver
```
