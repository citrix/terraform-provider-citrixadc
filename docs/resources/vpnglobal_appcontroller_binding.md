---
subcategory: "VPN"
---

# Resource: vpnglobal_appcontroller_binding

The vpnglobal_appcontroller_binding resource is used to bind an App Controller server to the global configuration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_appcontroller_binding" "tf_vpnglobal_appcontroller_binding" {
	appcontroller = "http://www.citrix.com"
}
```


## Argument Reference

* `appcontroller` - (Optional) Configured App Controller server.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_appcontroller_binding. It has the same value as the `appcontroller` attribute.


## Import

A vpnglobal_appcontroller_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnglobal_appcontroller_binding.tf_vpnglobal_appcontroller_binding http://www.citrix.com
```
