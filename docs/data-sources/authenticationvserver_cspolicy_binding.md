---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_cspolicy_binding

The authenticationvserver_cspolicy_binding data source allows you to retrieve information about the binding between an authentication virtual server and a content switching policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
  name      = "tf_authenticationvserver"
  policy    = "test_policy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_cspolicy_binding.tf_bind.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_authenticationvserver_cspolicy_binding.tf_bind.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor

## Attribute Reference

* `id` - The id of the authenticationvserver_cspolicy_binding. It is a system-generated identifier.
