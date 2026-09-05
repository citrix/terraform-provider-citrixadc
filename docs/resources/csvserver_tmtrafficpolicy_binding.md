---
subcategory: "Content Switching"
---

# Resource: csvserver_tmtrafficpolicy_binding

The csvserver_tmtrafficpolicy_binding resource is used to bind a Traffic Management (TM) traffic policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver"
  ipv46       = "10.10.10.22"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_csvserver_tmtrafficpolicy_binding" "tf_binding" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = "tf_tmtrafficpolicy"
  priority   = 100
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `bindpoint` - (Optional) Bind point at which the policy needs to be bound. Note: Content switching policies are evaluated only at request time.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labeltype` - (Optional) Type of label to be invoked.
* `labelname` - (Optional) Name of the label to be invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if the policy rule is evaluated to be TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_tmtrafficpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_tmtrafficpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_tmtrafficpolicy_binding.tf_binding tf_csvserver,tf_tmtrafficpolicy
```
