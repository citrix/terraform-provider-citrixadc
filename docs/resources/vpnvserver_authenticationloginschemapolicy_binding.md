---
subcategory: "Vpn"
---

# Resource: vpnvserver_authenticationloginschemapolicy_binding

The vpnvserver_authenticationloginschemapolicy_binding resource is used to bind authenticationloginschemapolicy to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver_example"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
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
resource "citrixadc_vpnvserver_authenticationloginschemapolicy_binding" "tf_bind" {
  name            = citrixadc_vpnvserver.tf_vpnvserver.name
  policy          = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy.name
  priority        = 80
  secondary       = "false"
  groupextraction = "false"
  bindpoint       = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `policy` - (Required) The name of the policy, if any, bound to the VPN virtual server.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
* `gotopriorityexpression` - (Optional) Next priority expression.
* `groupextraction` - (Optional) Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_authenticationloginschemapolicy_binding. It is the concatenation of `name` and `policy` attributes seperated by comma.


## Import

A vpnvserver_authenticationloginschemapolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_authenticationloginschemapolicy_binding.tf_bind tf_vserver_example,tf_loginschemapolicy
```
