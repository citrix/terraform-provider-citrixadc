---
subcategory: "AAA"
---

# Resource: aaaglobal_authenticationnegotiateaction_binding

Binds a negotiate (Kerberos/NTLM) authentication action to the global AAA configuration on the Citrix ADC. Use this resource to enable Windows Integrated Authentication for traffic that is not matched by a more specific authentication virtual server or policy, so that users are authenticated against the negotiate profile system-wide.


## Example usage

```hcl
resource "citrixadc_authenticationnegotiateaction" "example" {
  name                = "negotiate_action1"
  domain              = "example.com"
  domainuser          = "svc_negotiate"
  domainuserpasswd    = "secret"
  ntlmpath            = "https://ntlm.example.com/"
}

resource "citrixadc_aaaglobal_authenticationnegotiateaction_binding" "example" {
  windowsprofile = citrixadc_authenticationnegotiateaction.example.name
}
```


## Argument Reference

* `windowsprofile` - (Required) Name of the negotiate profile (authenticationnegotiateaction) to be bound globally. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaaglobal_authenticationnegotiateaction_binding. Because `aaaglobal` is a singleton with no parent name, the id is the plain value of the `windowsprofile` attribute (the bound negotiate profile name), not a composite key.


## Import

A aaaglobal_authenticationnegotiateaction_binding can be imported using the bound `windowsprofile` value, e.g.

```shell
terraform import citrixadc_aaaglobal_authenticationnegotiateaction_binding.example negotiate_action1
```
