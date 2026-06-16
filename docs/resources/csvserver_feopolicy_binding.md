---
subcategory: "Content Switching"
---

# Resource: csvserver_feopolicy_binding

The csvserver_feopolicy_binding resource is used to create bindings between a content switching virtual server and a Front End Optimization (FEO) policy.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_feopolicy" "tf_feopolicy" {
  name   = "tf_feopolicy"
  rule   = "true"
  action = "tf_feoaction"
}

resource "citrixadc_csvserver_feopolicy_binding" "tf_bind" {
  name                   = citrixadc_csvserver.tf_csvserver.name
  policyname             = citrixadc_feopolicy.tf_feopolicy.name
  priority               = 100
  bindpoint              = "REQUEST"
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - (Optional) Name of the label to be invoked.
* `labeltype` - (Optional) Type of label to be invoked.
* `priority` - (Optional) Priority for the policy.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_feopolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_feopolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_feopolicy_binding.tf_bind tf_csvserver,tf_feopolicy
```
