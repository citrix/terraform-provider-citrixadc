---
subcategory: "Content Switching"
---

# Resource: csvserver_appflowpolicy_binding

The csvserver_appflowpolicy_binding resource is used to bind an AppFlow policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver_appflowpolicy_binding" "tf_csvserver_appflowpolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_appflowpolicy"
  priority   = 100
  bindpoint  = "REQUEST"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `priority` - (Optional) Priority for the policy.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1. Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_appflowpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_appflowpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_appflowpolicy_binding.tf_csvserver_appflowpolicy_binding tf_csvserver,tf_appflowpolicy
```
