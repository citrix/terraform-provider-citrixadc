---
subcategory: "Authentication"
---

# Resource: authenticationvserver_authenticationloginschemapolicy_binding

The authenticationvserver_authenticationloginschemapolicy_binding resource is used to bind authenticationloginschemapolicy to authenticationvserver.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name    = "tf_loginschemapolicy"
  rule    = "true"
  action  = citrixadc_authenticationloginschema.tf_loginschema.name
  comment = "samplenew_testing"
}
resource "citrixadc_authenticationvserver_authenticationloginschemapolicy_binding" "tf_binding" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy.name
  priority  = 77
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

* `id` - The id of the authenticationvserver_authenticationloginschemapolicy_binding. It is the concatenation of both `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_authenticationloginschemapolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_authenticationloginschemapolicy_binding.tf_binding tf_authenticationvserver,tf_loginschemapolicy
```
