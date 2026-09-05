---
subcategory: "Load Balancing"
---

# Resource: lbvserver_spilloverpolicy_binding

The lbvserver_spilloverpolicy_binding resource is used to bind a load balancing virtual server to a spillover policy.


## Example usage

```hcl
resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  ipv46       = "1.1.1.1"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_spilloverpolicy" "demo_spilloverpolicy" {
  name   = "demo_spilloverpolicy"
  rule   = "SYS.VSERVER(\"demo_lb\").RESPTIME.GT(50)"
  action = "SPILLOVER"
}

resource "citrixadc_lbvserver_spilloverpolicy_binding" "demo_binding" {
  name       = citrixadc_lbvserver.demo_lb.name
  policyname = citrixadc_spilloverpolicy.demo_spilloverpolicy.name
  priority   = 100
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies.
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_spilloverpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_spilloverpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_spilloverpolicy_binding.demo_binding demo_lb,demo_spilloverpolicy
```
