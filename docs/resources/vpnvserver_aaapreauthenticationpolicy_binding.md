---
subcategory: "VPN"
---

# Resource: vpnvserver_aaapreauthenticationpolicy_binding

The vpnvserver_aaapreauthenticationpolicy_binding resource is used to to bind aaapreauthenticationpolicy to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserverexample"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnvserver_aaapreauthenticationpolicy_binding" "tf_binding" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = "tf_aaapolicy"
  priority  = 40
  secondary = "false"
  bindpoint = "OTHERTCP_REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `policy` - (Required) The name of the policy, if any, bound to the VPN virtual server.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: *  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
* `groupextraction` - (Optional) Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_aaapreauthenticationpolicy_binding. It is the concatenation of both `name` and `policy` attributes seperated by comma.


## Import

A vpnvserver_aaapreauthenticationpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_aaapreauthenticationpolicy_binding.tf_binding tf_vpnvserverexample,tf_aaapolicy
```
