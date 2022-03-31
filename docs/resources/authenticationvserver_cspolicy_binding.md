---
subcategory: "Authentication"
---

# Resource: authenticationvserver_cspolicy_binding

The authenticationvserver_cspolicy_binding resource is used to bind the authenticationvserver to cspolicy resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_policy_lb"
  servicetype = "HTTP"
  ipv46       = "192.122.3.3"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_csaction" "tf_csaction" {
  name            = "tf_csaction"
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
}
resource "citrixadc_cspolicy" "foo_cspolicy" {
  policyname = "test_policy"
  rule       = "TRUE"
  action     = citrixadc_csaction.tf_csaction.name
}
resource "citrixadc_authenticationvserver_cspolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_cspolicy.foo_cspolicy.policyname
  priority  = 90
  bindpoint = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - (Optional) Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - (Optional) The priority, if any, of the vpn vserver policy.
* `secondary` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_cspolicy_binding. It is the concatenation of `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_cspolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_cspolicy_binding.tf_bind tf_authenticationvserver,test_policy
```
