---
subcategory: "Authentication"
---

# Resource: authenticationvserver_responderpolicy_binding

The authenticationvserver_responderpolicy_binding resource is used to bind responderpolicy with the authenticationvserver.


## Example usage

```hcl
resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
  name           = "tf_authenticationvserver"
  servicetype    = "SSL"
  comment        = "new"
  authentication = "ON"
  state          = "DISABLED"
}
resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_authenticationvserver_responderpolicy_binding" "tf_bind" {
  name      = citrixadc_authenticationvserver.tf_authenticationvserver.name
  policy    = citrixadc_responderpolicy.tf_responder_policy.name
  priority  = 200
  bindpoint = "REQUEST"
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

* `id` - The id of the authenticationvserver_responderpolicy_binding. It is the concatenation of both `name` and `policy` attributes seperated by comma.


## Import

A authenticationvserver_responderpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_authenticationvserver_responderpolicy_binding.tf_bind tf_authenticationvserver,tf_responder_policy
```
