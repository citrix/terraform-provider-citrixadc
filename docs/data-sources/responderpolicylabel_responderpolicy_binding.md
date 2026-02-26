---
subcategory: "Responder"
---

# Data Source: responderpolicylabel_responderpolicy_binding

The responderpolicylabel_responderpolicy_binding data source allows you to retrieve information about a responder policy binding to a responder policy label.

## Example Usage

```terraform
data "citrixadc_responderpolicylabel_responderpolicy_binding" "tf_responderpolicylabel_responderpolicy_binding" {
  labelname  = "tf_responderpolicylabel"
  policyname = "tf_responderpolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding.invoke
}

output "labeltype" {
  value = data.citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding.labeltype
}
```

## Argument Reference

* `labelname` - (Required) Name of the responder policy label to which to bind the policy.
* `policyname` - (Required) Name of the responder policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderpolicylabel_responderpolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server.
* `labeltype` - Type of policy label to invoke. Available settings function as follows: vserver - Invoke an unnamed policy label associated with a virtual server. policylabel - Invoke a user-defined policy label. Possible values: [ vserver, policylabel ]
