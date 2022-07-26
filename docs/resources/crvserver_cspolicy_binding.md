---
subcategory: "Cache Redirection"
---

# Resource: crvserver_cspolicy_binding

The crvserver_cspolicy_binding resource is used to create CRvserver CSpolicy Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_policy_lbv"
  servicetype = "HTTP"
  ipv46       = "192.122.3.30"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_csaction" "tf_csaction" {
  name            = "test_csaction"
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
}
resource "citrixadc_cspolicy" "foo_cspolicy" {
  policyname = "test_cspolicy"
  rule       = "TRUE"
  action     = citrixadc_csaction.tf_csaction.name
}
resource "citrixadc_service" "tf_service" {
  lbvserver = citrixadc_lbvserver.foo_lbvserver.name
  name = "tf_service1"
  port = 8080
  ip = "10.202.22.111"
  servicetype = "HTTP"
  cachetype = "TRANSPARENT"
}
resource "citrixadc_crvserver_cspolicy_binding" "crvserver_cspolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_cspolicy.foo_cspolicy.policyname
  priority   = 90
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) The CSW target server names.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_cspolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_cspolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_cspolicy_binding.crvserver_cspolicy_binding my_vserver,test_cspolicy
```
