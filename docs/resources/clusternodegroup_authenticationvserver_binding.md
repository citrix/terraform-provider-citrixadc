---
subcategory: "Cluster"
---

# Resource: clusternodegroup_authenticationvserver_binding

The clusternodegroup_authenticationvserver_binding resource is used to create clusternodegroup_authenticationvserver_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
  name     = "my_authentication_group"
  vserver = "my_authentication_server"
}
```


## Argument Reference

* `vserver` - (Required) vserver that need to be bound to this nodegroup.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_authenticationvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.


## Import

A clusternodegroup_authenticationvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding my_authentication_group,my_authentication_server
```
