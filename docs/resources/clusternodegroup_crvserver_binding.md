---
subcategory: "Cluster"
---

# Resource: clusternodegroup_crvserver_binding

The clusternodegroup_crvserver_binding resource is used to create clusternodegroup_crvserver_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
  name = "my_cr_group"
  vserver = "my_crvserver"
}
```


## Argument Reference

* `vserver` - (Required) vserver that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_crvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.


## Import

A clusternodegroup_crvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding my_cr_group,my_crvserver
```
