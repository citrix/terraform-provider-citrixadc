---
subcategory: "Load Balancing"
---

# Resource: lbvserver_appfwpolicy_binding

The lbvserver_appfwpolicy_binding resource is used to bind a load balancing virtual server to an application firewall (AppFw) policy.


## Example usage

```hcl
resource "citrixadc_lbvserver" "demo_lb" {
  name        = "demo_lb"
  ipv46       = "1.1.1.1"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_appfwprofile" "demo_appfwprofile" {
  name = "demo_appfwprofile"
  type = ["HTML"]
}

resource "citrixadc_appfwpolicy" "demo_appfwpolicy" {
  name        = "demo_appfwpolicy"
  profilename = citrixadc_appfwprofile.demo_appfwprofile.name
  rule        = "true"
}

resource "citrixadc_lbvserver_appfwpolicy_binding" "demo_binding" {
  name       = citrixadc_lbvserver.demo_lb.name
  policyname = citrixadc_appfwpolicy.demo_appfwpolicy.name
  priority   = 100
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
* `labeltype` - (Optional) The invocation type.
* `labelname` - (Optional) Name of the label invoked.
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_appfwpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_appfwpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_appfwpolicy_binding.demo_binding demo_lb,demo_appfwpolicy
```
