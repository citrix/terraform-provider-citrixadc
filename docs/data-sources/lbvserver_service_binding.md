---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_service_binding

The lbvserver_service_binding data source allows you to retrieve information about an existing binding between a load balancing virtual server and a service.


## Example Usage

```terraform
data "citrixadc_lbvserver_service_binding" "tf_binding" {
  name        = "tf_lbvserver"
  servicename = "tf_service"
}

output "weight" {
  value = data.citrixadc_lbvserver_service_binding.tf_binding.weight
}

output "order" {
  value = data.citrixadc_lbvserver_service_binding.tf_binding.order
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `servicename` - (Optional) Service to bind to the virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_service_binding. It is the concatenation of the `name` and `servicename` attributes separated by a comma.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `weight` - Weight to assign to the specified service.
