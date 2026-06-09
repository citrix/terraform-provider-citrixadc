---
subcategory: "Cluster"
---

# Data Source: clusternodegroup_service_binding

The clusternodegroup_service_binding data source allows you to retrieve information about a service that is bound to a cluster node group on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_clusternodegroup_service_binding" "example" {
  name    = "ng1"
  service = "svc1"
}

output "bound_service" {
  value = data.citrixadc_clusternodegroup_service_binding.example.service
}
```


## Argument Reference

* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `service` - (Required) Name of the service bound to this nodegroup.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_service_binding. It is a comma-separated set of `key:value` pairs in the form `name:<name>,service:<service>`, with each value URL-encoded.
* `name` - Name of the nodegroup.
* `service` - Name of the service bound to this nodegroup.
