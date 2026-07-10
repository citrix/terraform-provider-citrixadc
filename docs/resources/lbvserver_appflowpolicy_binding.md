---
subcategory: "Load Balancing"
---

# Resource: lbvserver_appflowpolicy_binding

The lbvserver_appflowpolicy_binding resource is used to bind an AppFlow policy to a load balancing virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver_appflowpolicy_binding" "tf_binding" {
  name                   = citrixadc_lbvserver.demo_lb.name
  policyname             = citrixadc_appflowpolicy.demo_appflowpolicy.name
  priority               = 100
  bindpoint              = "REQUEST"
  gotopriorityexpression = "END"
}

resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  ipv46       = "1.1.1.1"
  port        = "80"
  servicetype = "HTTP"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority. If not specified, the ADC assigns a value.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_appflowpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver\_appflowpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_appflowpolicy_binding.tf_binding tf_lbvserver,tf_appflowpolicy
```
