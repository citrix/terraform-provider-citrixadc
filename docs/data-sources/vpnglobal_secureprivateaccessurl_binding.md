---
subcategory: "VPN"
---

# Data Source: vpnglobal_secureprivateaccessurl_binding

The vpnglobal_secureprivateaccessurl_binding data source allows you to retrieve information about a Secure Private Access URL bound to the global VPN bind point on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_vpnglobal_secureprivateaccessurl_binding" "tf_bind" {
  secureprivateaccessurl = "https://app.example.com/"
}

output "secureprivateaccessurl" {
  value = data.citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_bind.secureprivateaccessurl
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `secureprivateaccessurl` - (Required) Configured Secure Private Access URL. This is the literal URL string used to look up the binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_secureprivateaccessurl_binding. It has the same value as the `secureprivateaccessurl` attribute.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
