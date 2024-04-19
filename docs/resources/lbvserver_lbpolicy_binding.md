---
subcategory: "Load Balancing"
---

# Resource: lbvserver_lbpolicy_binding

The lbvserver_lbpolicy_binding resource is used to bind lbpolicy to lb vserver.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}
resource "citrixadc_lbpolicy" "tf_pol" {
  name   = "tf_pol"
  rule   = "true"
  action = "NOLBACTION"
}
resource "citrixadc_lbvserver_lbpolicy_binding" "tf_bind" {
  name       = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_lbpolicy.tf_pol.name
  priority   = 10
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values = REQUEST, RESPONSE, MQTT_JUMBO_REQ
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type. Possible values = reqvserver, resvserver, policylabel
* `order` - (Optional) Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
* `priority` - (Optional) Priority.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_lbpolicy_binding is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_lbpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_lbpolicy_binding.tf_bind tf_lbvserver,tf_pol
```
