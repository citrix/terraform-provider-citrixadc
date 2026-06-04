---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationsmartaccesspolicy_binding

The authenticationvserver_authenticationsmartaccesspolicy_binding data source allows you to retrieve information about a SmartAccess authentication policy bound to an authentication virtual server.


## Example usage

```terraform
data "citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding" "example" {
  name   = "authvserver1"
  policy = "smartaccesspolicy1"
}

output "binding_priority" {
  value = data.citrixadc_authenticationvserver_authenticationsmartaccesspolicy_binding.example.priority
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which the policy is bound.
* `policy` - (Required) The name of the policy bound to the authentication vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding, in the format `name:<name>,policy:<policy>` (the values are URL-encoded).
* `priority` - The priority of the policy within the authentication vserver's policy chain.
* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE (NEXT, END, USE_INVOCATION_RESULT, or an expression that evaluates to a number).
* `secondary` - Applicable only while binding a classic authentication policy; advance authentication policy uses nFactor.
* `groupextraction` - Applicable only while binding a classic authentication policy; advance authentication policy uses nFactor.
* `nextfactor` - Applicable only while binding an advance authentication policy; classic authentication policy does not support nFactor.
