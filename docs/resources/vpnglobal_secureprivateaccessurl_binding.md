---
subcategory: "VPN"
---

# Resource: vpnglobal_secureprivateaccessurl_binding

This resource is used to bind a Secure Private Access URL to the global VPN bind point.


## Example usage

```hcl
resource "citrixadc_vpnglobal_secureprivateaccessurl_binding" "tf_bind" {
  secureprivateaccessurl = "https://app.example.com/"
}
```


## Argument Reference

* `secureprivateaccessurl` - (Required) Configured Secure Private Access URL. This is a literal URL string (maximum length 255) that serves as the binding key; it is not a reference to another Citrix ADC resource. Changing this value forces a new binding to be created.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this value forces a new binding to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_secureprivateaccessurl_binding. It has the same value as the `secureprivateaccessurl` attribute.


## Import

A vpnglobal_secureprivateaccessurl_binding can be imported using its secureprivateaccessurl value (which is also the resource id), e.g.

```shell
terraform import citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_bind "https://app.example.com/"
```
