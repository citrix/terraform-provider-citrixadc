---
subcategory: "VPN"
---

# Resource: vpnvserver_authenticationpolicy_binding

Binds an authentication policy to a specific VPN (NetScaler Gateway) virtual server so that the policy governs how users authenticate when connecting through that virtual server. Use this resource to attach primary, secondary (two-factor), or group-extraction authentication policies to an individual VPN vserver, controlling the evaluation order via priority.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "gatewayvserver1"
  servicetype = "SSL"
  ipv46       = "10.222.74.150"
  port        = 443
}

resource "citrixadc_authenticationpolicy" "tf_policy" {
  name   = "ldapauthpolicy1"
  rule   = "true"
  action = "ldapaction1"
}

resource "citrixadc_vpnvserver_authenticationpolicy_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_vpnvserver.name
  policy                 = citrixadc_authenticationpolicy.tf_policy.name
  bindpoint              = "REQUEST"
  priority               = 100
  secondary              = false
  groupextraction        = false
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `name` - (Required) Name of the VPN virtual server to which the authentication policy is bound. Changing this forces a new resource to be created.
* `policy` - (Required) The name of the authentication policy to bind to the VPN virtual server. Changing this forces a new resource to be created.
* `bindpoint` - (Optional) Bind point to which to bind the policy. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression. This value is part of the resource's composite ID. Changing this forces a new resource to be created. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST, AAA_REQUEST, AAA_RESPONSE ]
* `gotopriorityexpression` - (Optional) Applicable only to advanced VPN session policies. An expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE. Specify `NEXT` to evaluate the policy with the next higher priority number, `END` to end policy evaluation, or an expression that evaluates to a priority number. Changing this forces a new resource to be created.
* `groupextraction` - (Optional) Binds the authentication policy to a tertiary chain which will be used only for group extraction. The user will not authenticate against this server, and this will only be called if primary and/or secondary authentication has succeeded. Changing this forces a new resource to be created.
* `priority` - (Optional) Integer specifying the policy's priority. The lower the number, the higher the priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Changing this forces a new resource to be created.
* `secondary` - (Optional) Binds the authentication policy as the secondary policy to use in a two-factor configuration. A user must then authenticate not only via a primary authentication method but also via a secondary authentication method. User groups are aggregated across both. The user name must be exactly the same for both authentication methods, but they can require different passwords. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_authenticationpolicy_binding. It is a comma-separated list of `key:value` pairs in the form `bindpoint:<bindpoint>,name:<name>,policy:<policy>`, where each value is URL-encoded.


## Import

A vpnvserver_authenticationpolicy_binding can be imported using its id, which is a comma-separated list of `key:value` pairs in the form `bindpoint:<bindpoint>,name:<name>,policy:<policy>`, e.g.

```shell
terraform import citrixadc_vpnvserver_authenticationpolicy_binding.tf_bind bindpoint:REQUEST,name:gatewayvserver1,policy:ldapauthpolicy1
```
