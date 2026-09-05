---
subcategory: "Content Switching"
---

# Resource: csvserver_spilloverpolicy_binding

The csvserver_spilloverpolicy_binding resource is used to bind a spillover policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
  name   = "tf_spilloverpolicy"
  rule   = "SYS.VSERVER(\"tf_csvserver\").RESPTIME.GT(100)"
  action = "SPILLOVER"
}

resource "citrixadc_csvserver_spilloverpolicy_binding" "tf_csvserver_spilloverpolicy_binding" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_spilloverpolicy.tf_spilloverpolicy.name
  priority   = 1
  bindpoint  = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labeltype` - (Optional) Type of label to be invoked.
* `labelname` - (Optional) Name of the label to be invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1. Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_spilloverpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_spilloverpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_spilloverpolicy_binding.tf_csvserver_spilloverpolicy_binding tf_csvserver,tf_spilloverpolicy
```
