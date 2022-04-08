---
subcategory: "Vpn"
---

# Resource: vpnglobal_intranetip6_binding

The vpnglobal_intranetip6_binding resource is used to bind intranetip6 to vpnglobal congiguration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_intranetip6_binding" "tf_bind" {
  intranetip6 = "2.3.4.5"
  numaddr     = "45"
}
```


## Argument Reference

* `intranetip6` - (Required) The intranet ip address or range.
* `numaddr` - (Optional) The intranet ip address or range's netmask.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_intranetip6_binding. It has the same value as the `intranetip6` attribute.


## Import

A vpnglobal_intranetip6_binding can be imported using its intranetip6, e.g.

```shell
terraform import citrixadc_vpnglobal_intranetip6_binding.tf_bind 2.3.4.5
```
