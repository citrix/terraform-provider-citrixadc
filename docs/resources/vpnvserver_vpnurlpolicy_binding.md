---
subcategory: "Vpn"
---

# Resource: vpnvserver_vpnurlpolicy_binding

The vpnvserver_vpnurlpolicy_binding resource is used to bind vpnurlpolicy to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_tfvpnvserver" {
  name        = "tf_example_vserver"
  servicetype = "SSL"
  ipv46       = "9.6.77.8"
  port        = 443
}
resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
  name             = "tf_vpnurlaction"
  linkname         = "new_link"
  actualurl        = "www.citrix.com"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  ssotype          = "unifiedgateway"
  vservername      = "vserver1"
}
resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
  name   = "new_policy"
  rule   = "true"
  action = citrixadc_vpnurlaction.tf_vpnurlaction.name
}
resource "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_tfvpnvserver.name
  policy                 = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
  priority               = 20
  gotopriorityexpression = "next"
  bindpoint              = "REQUEST"
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

* `id` - The id of the vpnvserver_vpnurlpolicy_binding. It is concatenation of `name` and `policy` attributes seperated by comma.


## Import

A vpnvserver_vpnurlpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_vpnurlpolicy_binding.tf_bind tf_example_vserver,new_policy
```
