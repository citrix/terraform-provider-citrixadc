---
subcategory: "Load Balancing"
---

# Resource: csvserver_appflowpolicy_binding

The csvserver_appflowpolicy_binding resource is used to add AppFlow policies to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_appflowpolicy_binding" "tf_csvserver_appflowpolicy_binding" {
	name = "tf_csvserver"
	policyname = "tf_appflowpolicy"
	labelname = citrixadc_csvserver.demo_csvserver.name
	gotopriorityexpression = "END"
	invoke = true
	labeltype = "reqvserver"
	priority = 1
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the CS vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated ifl the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver\_appflowpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver\_appflowpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_appflowpolicy_binding.tf_csvserver_appflowpolicy_binding tf_csvserver,tf_appflowpolicy
```
