---
subcategory: "Vpn"
---

# Resource: vpnglobal_intranetip_binding

The vpnglobal_intranetip_binding resource is used to bind intranetip to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_intranetip_binding" "tf_bind" {
  intranetip = "2.3.4.5"
  netmask    = "255.255.255.0"
}
```


## Argument Reference

* `intranetip` - (Required) The intranet ip address or range.
* `netmask` - (Required) The intranet ip address or range's netmask.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_intranetip_binding. It has the same value as the `intranetip` attribute.


## Import

A vpnglobal_intranetip_binding can be imported using its intranetip, e.g.

```shell
terraform import citrixadc_vpnglobal_intranetip_binding.tf_bind 2.3.4.5
```
