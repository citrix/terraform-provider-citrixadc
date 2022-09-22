---
subcategory: "VPN"
---

# Resource: vpnvserver_authenticationsamlidppolicy_binding

The vpnvserver_authenticationsamlidppolicy_binding resource is used to bind authenticationsamlidppolicy to vpnvserver resource.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_vserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey = "tf_sslcertkey"
  cert    = "/var/tmp/certificate1.crt"
  key     = "/var/tmp/key1.pem"
}
resource "citrixadc_authenticationsamlidpprofile" "tf_samlidpprofile" {
  name                        = "tf_samlidpprofile"
  samlspcertname              = citrixadc_sslcertkey.tf_sslcertkey.certkey
  assertionconsumerserviceurl = "http://www.example.com"
  sendpassword                = "OFF"
  samlissuername              = "new_user"
  rejectunsignedrequests      = "ON"
  signaturealg                = "RSA-SHA1"
  digestmethod                = "SHA1"
  nameidformat                = "Unspecified"
}
resource "citrixadc_authenticationsamlidppolicy" "tf_samlidppolicy" {
  name    = "tf_samlidppolicy"
  rule    = "false"
  action  = citrixadc_authenticationsamlidpprofile.tf_samlidpprofile.name
  comment = "aSimpleTesting"
}
resource "citrixadc_vpnvserver_authenticationsamlidppolicy_binding" "tf_binding" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_authenticationsamlidppolicy.tf_samlidppolicy.name
  priority  = 9
  bindpoint = "REQUEST"
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

* `id` - The id of the vpnvserver_authenticationsamlidppolicy_binding. It is the concatenation of both `name` and `policy` attributes seperated by comma.


## Import

A vpnvserver_authenticationsamlidppolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_authenticationsamlidppolicy_binding.tf_binding tf_vserver,tf_samlidppolicy
```
