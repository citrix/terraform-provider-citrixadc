---
subcategory: "Cluster"
---

# Resource: clusternodegroup_gslbvserver_binding

The clusternodegroup_gslbvserver_binding resource is used to create clusternodegroup_gslbvserver_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_gslbvserver_binding" "tf_clusternodegroup_gslbvserver_binding" {
  name = "my_gslb_group"
  vserver = "my_gslbvserver"
}
```


## Argument Reference

* `vserver` - (Required) vserver that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_gslbvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.


## Import

A clusternodegroup_gslbvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_gslbvserver_binding.tf_clusternodegroup_gslbvserver_binding my_gslb_group,my_gslbvserver
```
