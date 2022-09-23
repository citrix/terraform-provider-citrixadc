---
subcategory: "VPN"
---

# Resource: vpnvserver_responderpolicy_binding

The vpnvserver_responderpolicy_binding resource is used to bind responderpolicy to vpnvserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_vpnvserver_responderpolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_responderpolicy.tf_responder_policy.name
  priority  = 200
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `policy` - (Required) The name of the policy, if any, bound to the VPN virtual server.
* `bindpoint` - (Optional) Bindpoint to which the policy is bound.
* `gotopriorityexpression` - (Optional) Next priority expression.
* `groupextraction` - (Optional) Binds the authentication policy to a tertiary chain which will be used only for group extraction.  The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000.
* `secondary` - (Optional) Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_responderpolicy_binding. It is the concatenation of the `name` and `policy` attributes separated by a comma.


## Import

A vpnvserver_responderpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vpnvserver_responderpolicy_binding.tf_bind tf_examplevserver,tf_responder_policy
```
