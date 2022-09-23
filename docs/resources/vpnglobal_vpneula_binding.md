---
subcategory: "VPN"
---

# Resource: vpnglobal_vpneula_binding

The vpnglobal_vpneula_binding resource is used to bind vpneula to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpneula" "tf_vpneula" {
	name = "tf_vpneula"	
}
resource "citrixadc_vpnglobal_vpneula_binding" "tf_bind" {
  eula = citrixadc_vpneula.tf_vpneula.name
}
```


## Argument Reference

* `eula` - (Required) Name of the EULA bound to vpnglobal
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_vpneula_binding. It has the same value as the `eula` attribute.


## Import

A vpnglobal_vpneula_binding can be imported using its eula, e.g.

```shell
terraform import citrixadc_vpnglobal_vpneula_binding.tf_bind tf_vpneula
```
