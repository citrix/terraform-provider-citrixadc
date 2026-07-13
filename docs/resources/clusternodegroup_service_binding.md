---
subcategory: "Cluster"
---

# Resource: clusternodegroup_service_binding

The clusternodegroup_service_binding resource is used to create clusternodegroup_service_binding.


## Example usage

```hcl
resource "citrixadc_clusternodegroup_service_binding" "tf_clusternodegroup_service_binding" {
  name    = "my_gslb_group"
  service = citrixadc_service.tf_service.name
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    servicetype = "ADNS"
    ip = "10.77.33.22"
    port = "53"
}

```


## Argument Reference

* `service` - (Required) name of the service bound to this nodegroup. The servicetype must be ADNS.
* `name` - (Required) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternodegroup_service_binding. It is the concatenation of `name` and `service` attributes saparated by a comma.


## Import

A clusternodegroup_service_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternodegroup_service_binding.tf_clusternodegroup_service_binding my_gslb_group,tf_service
```
