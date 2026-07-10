---
subcategory: "Load Balancing"
---

# Resource: lbvserver_feopolicy_binding

The lbvserver_feopolicy_binding resource is used to bind a front end optimization policy to a load balancing virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_feopolicy" "tf_feopolicy" {
  name   = "tf_feopolicy"
  rule   = "true"
  action = "pageextendcache"
}

resource "citrixadc_lbvserver_feopolicy_binding" "tf_bind" {
  name                   = citrixadc_lbvserver.tf_lbvserver.name
  policyname             = citrixadc_feopolicy.tf_feopolicy.name
  priority               = 100
  gotopriorityexpression = "END"
  bindpoint              = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows:
  * reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server.
  * resvserver - Evaluate the response against the response-based policies bound to the specified virtual server.
  * policylabel - invoke the request or response against the specified user-defined policy label.
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_feopolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_feopolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_feopolicy_binding.tf_bind tf_lbvserver,tf_feopolicy
```
