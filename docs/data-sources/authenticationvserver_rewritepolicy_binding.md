---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_rewritepolicy_binding

The authenticationvserver_rewritepolicy_binding data source allows you to retrieve information about the binding between an authentication virtual server and a rewrite policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_rewritepolicy_binding" "tf_binding" {
  name      = "tf_authenticationvserver"
  policy    = "tf_rewrite_policy"
  bindpoint = "RESPONSE"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_rewritepolicy_binding.tf_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_authenticationvserver_rewritepolicy_binding.tf_binding.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.
* `bindpoint` - (Required) Bindpoint to which the policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor

## Attribute Reference

* `id` - The id of the authenticationvserver_rewritepolicy_binding. It is a system-generated identifier.
