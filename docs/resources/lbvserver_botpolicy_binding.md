---
subcategory: "Load Balancing"
---

# Resource: lbvserver_botpolicy_binding

The lbvserver_botpolicy_binding resource is used to bind a bot policy to a load balancing virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver_botpolicy_binding" "demo_binding" {
  name                   = citrixadc_lbvserver.demo_lb.name
  policyname             = citrixadc_botpolicy.demo_botpolicy.name
  priority               = 100
  bindpoint              = "REQUEST"
  labeltype              = "reqvserver"
  labelname              = citrixadc_lbvserver.demo_lb.name
  gotopriorityexpression = "END"
  invoke                 = true
}

resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  servicetype = "HTTP"
}

resource "citrixadc_botpolicy" "demo_botpolicy" {
  name        = "demo_botpolicy"
  profilename = citrixadc_botprofile.demo_botprofile.name
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}

resource "citrixadc_botprofile" "demo_botprofile" {
  name    = "demo_botprofile"
  comment = "demo_botprofile comment"
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_botpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_botpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_botpolicy_binding.tf_lbvserver_botpolicy_binding demo_lb,demo_botpolicy
```
