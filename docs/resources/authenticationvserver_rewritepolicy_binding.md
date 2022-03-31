---
subcategory: "Authentication"
---

# Resource: authenticationvserver_rewritepolicy_binding

The authenticationvserver_rewritepolicy_binding resource is used to bind the rewritepolicy to authenticationvserver resource.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
resource "citrixadc_authenticationvserver_rewritepolicy_binding" "tf_bind" {
  name                   = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy                 = citrixadc_rewritepolicy.tf_rewrite_policy.name
  priority               = 90
  bindpoint              = "RESPONSE"
  gotopriorityexpression = "END"
  groupextraction        = "false"
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.
* `bindpoint` - (Optional) Bindpoint to which the policy is bound.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - (Optional) Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - (Optional) The priority, if any, of the vpn vserver policy.
* `secondary` - (Optional) Applicable only while bindind classic authentication policy as advance authentication policy use nFactor


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_rewritepolicy_binding. It is the concatenation of  `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_rewritepolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_rewritepolicy_binding.tf_bind tf_authenticationvserver,tf_rewrite_policy
```
