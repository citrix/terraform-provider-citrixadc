---
subcategory: "VPN"
---

# Resource: vpnglobal_auditnslogpolicy_binding

Binds an audit nslog policy to the global VPN bind point so that the policy is evaluated for traffic handled by the VPN (NetScaler Gateway) virtual servers globally. Use this resource to enable nslog auditing across all VPN sessions without binding the policy to individual vpnvservers.

The VPN global configuration is a singleton bind point on the Citrix ADC; this resource attaches a named audit nslog policy to it.


## Example usage

```hcl
resource "citrixadc_auditnslogaction" "tf_action" {
  name       = "tf_nslogaction"
  serverip   = "10.10.10.10"
  serverport = 514
  loglevel   = ["INFORMATIONAL"]
}

resource "citrixadc_auditnslogpolicy" "tf_policy" {
  name   = "tf_nslogpolicy"
  rule   = "true"
  action = citrixadc_auditnslogaction.tf_action.name
}

resource "citrixadc_vpnglobal_auditnslogpolicy_binding" "tf_bind" {
  policyname             = citrixadc_auditnslogpolicy.tf_policy.name
  priority               = 100
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `policyname` - (Required) The name of the audit nslog policy to bind to the global VPN bind point. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.
* `groupextraction` - (Optional) Bind the Authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded. Changing this forces a new resource to be created.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the priority number, the higher the policy's priority. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Changing this forces a new resource to be created.
* `secondary` - (Optional) Bind the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only to a primary authentication server but also to a secondary authentication server. User groups are aggregated across both authentication servers. The user name must be exactly the same on both authentication servers, but the authentication servers can require different passwords. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_auditnslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A vpnglobal_auditnslogpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_vpnglobal_auditnslogpolicy_binding.tf_bind tf_nslogpolicy
```
