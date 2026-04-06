---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationradiuspolicy_binding

The authenticationvserver_authenticationradiuspolicy_binding data source allows you to retrieve information about an authenticationvserver to authenticationradiuspolicy binding.


## Example Usage

```terraform
data "citrixadc_authenticationvserver_authenticationradiuspolicy_binding" "tf_bind" {
  name   = "tf_authenticationvserver"
  policy = "tf_radiuspolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_authenticationradiuspolicy_binding.tf_bind.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_authenticationvserver_authenticationradiuspolicy_binding.tf_bind.gotopriorityexpression
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `id` - The id of the authenticationvserver_authenticationradiuspolicy_binding. It is the concatenation of the `name` and `policy` attributes separated by a comma.
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Bind the authentication policy to the secondary chain. Provides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method. Because user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.
