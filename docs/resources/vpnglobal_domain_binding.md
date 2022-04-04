---
subcategory: "Vpn"
---

# Resource: vpnglobal_domain_binding

The vpnglobal_domain_binding resource is used to domain to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_vpnglobal_domain_binding" "tf_bind" {
  intranetdomain = "http://www.example.com/"
}
```


## Argument Reference

* `intranetdomain` - (Required) The conflicting intranet domain name.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_domain_binding. It has the same value as the `intranetdomain` attribute.


## Import

A vpnglobal_domain_binding can be imported using its intranetdomain, e.g.

```shell
terraform import citrixadc_vpnglobal_domain_binding.tf_bind http://www.example.com/
```
