---
subcategory: "VPN"
---

# Resource: vpnvserver_authenticationoauthidppolicy_binding

The vpnvserver_authenticationoauthidppolicy_binding resource is used to bind authenticationoauthidppolicy to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vpnvserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_authenticationoauthidpprofile" "tf_idpprofile" {
  name         = "tf_idpprofile"
  clientid     = "cliId"
  clientsecret = "secret"
  redirecturl  = "http://www.example.com/1/"
}
resource "citrixadc_authenticationoauthidppolicy" "tf_idppolicy" {
  name    = "tf_idppolicy"
  rule    = "true"
  action  = citrixadc_authenticationoauthidpprofile.tf_idpprofile.name
  comment = "aboutpolicy"
}
resource "citrixadc_vpnvserver_authenticationoauthidppolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationoauthidppolicy.tf_idppolicy.name
  priority  = 70
  bindpoint = "REQUEST"
  secondary = "false"
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

* `id` - The id of the vpnvserver_authenticationoauthidppolicy_binding. It is the concatenation of `name` and `policy` attributes seperated by comma.


## Import

A vpnvserver_authenticationoauthidppolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_authenticationoauthidppolicy_binding.tf_bind tf_vpnvserver,tf_idppolicy
```
