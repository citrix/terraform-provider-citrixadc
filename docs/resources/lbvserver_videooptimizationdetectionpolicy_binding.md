---
subcategory: "Load Balancing"
---

# Resource: lbvserver_videooptimizationdetectionpolicy_binding

The lbvserver_videooptimizationdetectionpolicy_binding resource is used to bind a video optimization detection policy to a load balancing virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "1.1.1.1"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver_videooptimizationdetectionpolicy_binding" "tf_vopolicy_binding" {
  name                   = citrixadc_lbvserver.tf_lbvserver.name
  policyname             = "tf_vop"
  bindpoint              = "REQUEST"
  priority               = 1
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_videooptimizationdetectionpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.

## Import

A lbvserver_videooptimizationdetectionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lbvserver_videooptimizationdetectionpolicy_binding.tf_vopolicy_binding tf_lbvserver,tf_vop
```