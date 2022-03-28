---
subcategory: "Vpn"
---

# Resource: vpnglobal_authenticationsamlpolicy_binding

The vpnglobal_authenticationsamlpolicy_binding resource is used ti bind authenticationsamlpolicy to vpnglobal configuration.


## Example usage

```hcl
resource "citrixadc_authenticationsamlaction" "tf_samlaction" {
  name                    = "tf_samlaction"
  metadataurl             = "http://www.example.com"
  samltwofactor           = "OFF"
  requestedauthncontext   = "minimum"
  digestmethod            = "SHA1"
  signaturealg            = "RSA-SHA256"
  metadatarefreshinterval = 1
}
resource "citrixadc_authenticationsamlpolicy" "tf_samlpolicy" {
  name      = "tf_samlpolicy"
  rule      = "NS_TRUE"
  reqaction = citrixadc_authenticationsamlaction.tf_samlaction.name
}
resource "citrixadc_vpnglobal_authenticationsamlpolicy_binding" "tf_bind" {
  policyname      = citrixadc_authenticationsamlpolicy.tf_samlpolicy.name
  priority        = 80
  groupextraction = false
  secondary       = false
}
```


## Argument Reference

* `policyname` - (Required) The name of the policy.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Bind the Authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_authenticationsamlpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A vpnglobal_authenticationsamlpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_vpnglobal_authenticationsamlpolicy_binding.tf_bind tf_samlpolicy
```
