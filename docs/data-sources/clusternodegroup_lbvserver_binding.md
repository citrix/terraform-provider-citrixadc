---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_lbvserver_binding

The clusternodegroup_lbvserver_binding data source allows you to retrieve information about a load balancing virtual server bound to a cluster node group on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_clusternodegroup_lbvserver_binding" "example" {
  name    = "ng1"
  vserver = "lbvs1"
}

output "bound_vserver" {
  value = data.citrixadc_clusternodegroup_lbvserver_binding.example.vserver
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `vserver` - (Required) Name of the load balancing virtual server bound to this nodegroup.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_lbvserver_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,vserver:<vserver>`, with the values URL-encoded.
