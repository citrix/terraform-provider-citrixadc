---
subcategory: "Responder"
---

# Resource: responderpolicylabel_responderpolicy_binding

The responderpolicylabel_responderpolicy_binding resource is used to bind a responder policy to responder policy label.


## Example usage

```hcl
resource "citrixadc_responderpolicylabel_responderpolicy_binding" "tf_responderpolicylabel_responderpolicy_binding" {
	labelname = citrixadc_responderpolicylabel.tf_responderpolicylabel.labelname
	policyname = citrixadc_responderpolicy.tf_responderpolicy.name
	priority = 5  
	gotopriorityexpression = "END"
	invoke = "false"
}

resource "citrixadc_responderpolicylabel" "tf_responderpolicylabel" {
	labelname = "tf_responderpolicylabel"
	policylabeltype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
```


## Argument Reference

* `policyname` - (Required) Name of the responder policy.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Available settings function as follows: * vserver - Invoke an unnamed policy label associated with a virtual server. * policylabel - Invoke a user-defined policy label. Possible values: [ vserver, policylabel ]
* `invoke_labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke.  * If labelType is reqvserver or resvserver, name of the virtual server.
* `labelname` - (Required) Name of the responder policy label to which to bind the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderpolicylabel_responderpolicy_binding. It is the concatenation of the `labelname` and `policyname` attributes separated by a comma.


## Import

A responderpolicylabel_responderpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding tf_responderpolicylabel,tf_responderpolicy
```
