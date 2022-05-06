---
subcategory: "Basic"
---

# Resource: service_dospolicy_binding

The service_dospolicy_binding resource is used to bind dospolicy to service.


## Example usage

```hcl

resource "citrixadc_service" "tf_service" {
  servicetype         = "HTTP"
  name                = "tf_service"
  ipaddress           = "10.77.33.22"
  ip                  = "10.77.33.22"
  port                = "80"
  state               = "ENABLED"
  wait_until_disabled = true
}
resource "citrixadc_service_dospolicy_binding" "tf_binding" {
  name       = citrixadc_service.tf_service.name
  policyname = "tf_dospolicy"
}
```


## Argument Reference

* `name` - (Required) Name of the service to which to bind a policy or monitor.
* `policyname` - (Required) The name of the policyname for which this service is bound


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the service_dospolicy_binding. It is the concatenation of `name` and `policyname` attributes separated by comma.


## Import

A service_dospolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_service_dospolicy_binding.tf_binding tf_service,tf_dospolicy
```
