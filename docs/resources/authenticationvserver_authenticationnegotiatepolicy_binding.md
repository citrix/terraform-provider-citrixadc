---
subcategory: "Authentication"
---

# Resource: authenticationvserver_authenticationnegotiatepolicy_binding

The authenticationvserver_authenticationnegotiatepolicy_binding resource is used to bind authenticationnegotiatepolicy to authenticationvserver resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
  name                       = "tf_negotiateaction"
  domain                     = "DomainName"
  domainuser                 = "usersame"
  domainuserpasswd           = "password"
  ntlmpath                   = "http://www.example.com/"
  defaultauthenticationgroup = "new_grpname"
}
resource "citrixadc_authenticationnegotiatepolicy" "tf_negotiatepolicy" {
  name      = "tf_negotiatepolicy"
  rule      = "ns_true"
  reqaction = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
}
resource "citrixadc_authenticationvserver_authenticationnegotiatepolicy_binding" "tf_binding" {
  name            = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy          = citrixadc_authenticationnegotiatepolicy.tf_negotiatepolicy.name
  priority        = 9
  groupextraction = "false"
  bindpoint       = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
* `gotopriorityexpression` - (Optional) Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: * If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `groupextraction` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - (Optional) Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - (Optional) The priority, if any, of the vpn vserver policy.
* `secondary` - (Optional) Bind the authentication policy to the secondary chain. Provides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method. Because user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_authenticationnegotiatepolicy_binding. It is concatenation of `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_authenticationnegotiatepolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_authenticationnegotiatepolicy_binding.tf_binding tf_authenticationvserver,tf_negotiatepolicy
```
