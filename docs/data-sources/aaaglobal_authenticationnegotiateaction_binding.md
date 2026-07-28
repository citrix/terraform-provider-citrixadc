---
subcategory: "AAA"
---

# Data Source: aaaglobal_authenticationnegotiateaction_binding

The aaaglobal_authenticationnegotiateaction_binding data source allows you to retrieve information about a negotiate (Kerberos/NTLM) authentication action that is bound to the global AAA configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_aaaglobal_authenticationnegotiateaction_binding" "example" {
  windowsprofile = "negotiate_action1"
}

output "bound_negotiate_profile" {
  value = data.citrixadc_aaaglobal_authenticationnegotiateaction_binding.example.windowsprofile
}
```


## Argument Reference

* `windowsprofile` - (Required) Name of the negotiate profile to look up in the global AAA binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaglobal_authenticationnegotiateaction_binding. It is the plain value of the `windowsprofile` attribute.
* `windowsprofile` - Name of the negotiate profile bound globally.
