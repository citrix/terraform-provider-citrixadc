---
subcategory: "VPN"
---

# Data Source: vpnvserver_auditnslogpolicy_binding

The vpnvserver_auditnslogpolicy_binding data source allows you to retrieve information about a specific vpnvserver_auditnslogpolicy_binding.

## Example Usage

```terraform
data "citrixadc_vpnvserver_auditnslogpolicy_binding" "tf_bind" {
  name   = "tf_vpnvserver"
  policy = "tf_auditnslogpolicy"
}

output "priority" {
  value = data.citrixadc_vpnvserver_auditnslogpolicy_binding.tf_bind.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_vpnvserver_auditnslogpolicy_binding.tf_bind.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `policy` - (Required) The name of the policy, if any, bound to the VPN virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_auditnslogpolicy_binding. It has the value of the parent vpnvserver `name` and binding `policy` attributes separated by a comma. Example: `tf_vpnvserver,tf_auditnslogpolicy`.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `groupextraction` - Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `priority` - Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.
