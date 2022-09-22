---
subcategory: "VPN"
---

# Resource: vpnglobal_vpnintranetapplication_binding

The vpnglobal_vpnintranetapplication_binding resource is used to bind vpnintranetapplication to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnintranetapplication" "tf_vpnintranetapplication" {
  intranetapplication = "tf_vpnintranetapplication"
  protocol            = "UDP"
  destip              = "2.3.6.5"
  interception        = "TRANSPARENT"
}
resource "citrixadc_vpnglobal_vpnintranetapplication_binding" "tf_bind" {
  intranetapplication =  citrixadc_vpnintranetapplication.tf_vpnintranetapplication.intranetapplication
}

```


## Argument Reference

* `intranetapplication` - (Required) The intranet vpn application.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpnintranetapplication_binding. It has the same value as the `intranetapplication` attribute.


## Import

A vpnglobal_vpnintranetapplication_binding can be imported using its intranetapplication, e.g.

```shell
terraform import citrixadc_vpnglobal_vpnintranetapplication_binding.tf_bind tf_vpnintranetapplication
```
