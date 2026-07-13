---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_authenticationvserver_binding

The clusternodegroup_authenticationvserver_binding data source allows you to retrieve information about a binding between a cluster nodegroup and an authentication vserver.

## Example Usage

```terraform
data "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
  name    = "my_tf_group"
  vserver = "my_authentication_server"
}

output "id" {
  value = data.citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding.id
}

output "name" {
  value = data.citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding.name
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `vserver` - (Required) Name of the authentication vserver that is bound to this nodegroup.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_authenticationvserver_binding. It is the concatenation of the `name` and `vserver` attributes separated by a comma.
