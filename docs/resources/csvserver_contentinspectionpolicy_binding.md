---
subcategory: "Content Switching"
---

# Resource: csvserver_contentinspectionpolicy_binding

The csvserver_contentinspectionpolicy_binding resource is used to bind a contentinspection policy to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_contentinspectionpolicy_binding" "tf_csvserver_contentinspectionpolicy_binding" {
	name = "tf_csvserver"
	policyname = "tf_contentinspectionpolicy"
	bindpoint = "REQUEST"
	gotopriorityexpression = "END"
	priority = 1    
}
```


## Argument Reference

* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_contentinspectionpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_contentinspectionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_contentinspectionpolicy_binding.tf_csvserver_contentinspectionpolicy_binding tf_csvserver,tf_contentinspectionpolicy
```
