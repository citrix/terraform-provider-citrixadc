---
subcategory: "Load Balancing"
---

# Resource: lbvserver_appqoepolicy_binding

The lbvserver_appqoepolicy_binding resource is used to bind load balancing virtual servers to AppQoE policies.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_appqoepolicy" "tf_appqoe_policy" {
  name   = "tf_appqoe_policy"
  rule   = "true"
  action = "DROP"
}

resource "citrixadc_lbvserver_appqoepolicy_binding" "tf_bind" {
  name       = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_appqoepolicy.tf_appqoe_policy.name
  priority   = 56
  bindpoint  = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. policylabel - invoke the request or response against the specified user-defined policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_appqoepolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_appqoepolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_appqoepolicy_binding.tf_bind tf_lbvserver,tf_appqoe_policy
```
