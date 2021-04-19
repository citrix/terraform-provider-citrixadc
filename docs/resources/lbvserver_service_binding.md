---
subcategory: "Load Balancing"
---

# Resource: lbvserver\_service\_binding

The lbvserver\_service\_binding resource is used to bind load balancing virtual servers to services.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  name = "tf_service"
  ip = "192.168.43.33"
  servicetype  = "HTTP"
  port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 1
}
```


## Argument Reference

* `servicename` - (Required) Service to bind to the virtual server.
* `weight` - (Optional) Weight to assign to the specified service.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_service\_binding. It is the concatenation of the `name` and `servicename` attributes separated by a comma.


## Import

A lbvserver\_service\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_service_binding.tf_binding tf_lbvserver,tf_service
```
