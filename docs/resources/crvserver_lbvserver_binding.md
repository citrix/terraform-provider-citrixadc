---
subcategory: "cache Redirection"
---

# Resource: crvserver_lbvserver_binding

The crvserver_lbvserver_binding resource is used to create CRvserver LBvserver Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_lbvserver"
  servicetype = "HTTP"
  ipv46       = "192.0.0.0"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_service" "tf_service" {
  lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
  name        = "tf_service"
  port        = 8081
  ip          = "10.33.4.5"
  servicetype = "HTTP"
  cachetype   = "TRANSPARENT"
}
resource "citrixadc_crvserver_lbvserver_binding" "crvserver_lbvserver_binding" {
  name      = citrixadc_crvserver.crvserver.name
  lbvserver = citrixadc_lbvserver.foo_lbvserver.name
  depends_on = [
    citrixadc_service.tf_service
  ]
}
```


## Argument Reference

* `lbvserver` - (Optional) The Default target server name.
* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_lbvserver_binding. It has the same value as the `name` attribute.


## Import

A crvserver_lbvserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_lbvserver_binding.crvserver_lbvserver_binding my_vserver,test_lbvserver
```
