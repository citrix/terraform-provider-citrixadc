---
subcategory: "Load Balancing"
---

# Resource: lbvserver_tmtrafficpolicy_binding

The lbvserver_tmtrafficpolicy_binding resource is used to bind a traffic management (tm) traffic policy to a load balancing virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver_tmtrafficpolicy_binding" "tf_binding" {
  name       = citrixadc_lbvserver.tf_lbvserver.name
  policyname = "tf_tmtrafficpolicy"
  priority   = 100
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT (evaluate the policy with the next higher priority number), END (end policy evaluation), USE_INVOCATION_RESULT (applicable if this policy invokes another policy label), or an expression that evaluates to a number.
* `bindpoint` - (Optional) Bind point to which to bind the policy.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver (evaluate the request against the request-based policies bound to the specified virtual server), resvserver (evaluate the response against the response-based policies bound to the specified virtual server), policylabel (invoke the request or response against the specified user-defined policy label).
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing virtual server's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_tmtrafficpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_tmtrafficpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lbvserver_tmtrafficpolicy_binding.tf_lbvserver_tmtrafficpolicy_binding tf_lbvserver,tf_tmttrafficpolicy
```
