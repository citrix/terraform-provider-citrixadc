---
subcategory: "VPN"
---

# Resource: vpnglobal_authenticationpolicy_binding

Binds an advanced authentication policy to the global VPN bind point, so that the policy is evaluated for every VPN (NetScaler Gateway) connection rather than being scoped to a single VPN virtual server. Use this to enforce a common authentication chain — including primary, secondary (two-factor), and group-extraction steps — across all gateway logins.


## Example usage

```hcl
resource "citrixadc_vpnglobal_authenticationpolicy_binding" "tf_bind" {
  policyname = "advauth_pol1"
  priority   = 100
}
```


## Argument Reference

* `policyname` - (Required) The name of the policy.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Bind the Authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called it primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_authenticationpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A vpnglobal_authenticationpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_vpnglobal_authenticationpolicy_binding.tf_bind advauth_pol1
```
