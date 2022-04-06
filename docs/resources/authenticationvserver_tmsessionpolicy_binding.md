---
subcategory: "Authentication"
---

# Resource: authenticationvserver_tmsessionpolicy_binding

The authenticationvserver_tmsessionpolicy_binding resource is used to bind tmsessionpolicy to authenticationvserver resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_authenticationvserver_tmsessionpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = "tf_tmsesspolicy"
  priority  = 90
  bindpoint = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - (Optional) Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - (Optional) The priority, if any, of the vpn vserver policy.
* `secondary` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_tmsessionpolicy_binding. It is the concatenation of `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_tmsessionpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_tmsessionpolicy_binding.tf_bind tf_authenticationvserver,tf_tmsesspolicy
```
