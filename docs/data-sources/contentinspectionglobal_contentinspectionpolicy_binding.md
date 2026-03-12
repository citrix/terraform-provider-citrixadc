---
subcategory: "Content Inspection"
---

# Data Source: contentinspectionglobal_contentinspectionpolicy_binding

The contentinspectionglobal_contentinspectionpolicy_binding data source allows you to retrieve information about a specific contentinspectionpolicy binding to contentinspectionglobal configuration.

## Example Usage

```terraform
data "citrixadc_contentinspectionglobal_contentinspectionpolicy_binding" "tf_contentinspectionglobal_contentinspectionpolicy_binding" {
  policyname = "tf_contentinspectionpolicy_ds"
  type       = "REQ_OVERRIDE"
}

output "policyname" {
  value = data.citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_contentinspectionglobal_contentinspectionpolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_contentinspectionglobal_contentinspectionpolicy_binding.priority
}

output "type" {
  value = data.citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_contentinspectionglobal_contentinspectionpolicy_binding.type
}
```

## Argument Reference

* `policyname` - (Required) Name of the contentInspection policy.
* `type` - (Required) The bindpoint to which to policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the contentinspectionglobal_contentinspectionpolicy_binding. It is a system-generated identifier.
* `invoke` - Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
* `labeltype` - Type of invocation. Available settings function as follows: reqvserver - Forward the request to the specified request virtual server. resvserver - Forward the response to the specified response virtual server. policylabel - Invoke the specified policy label.
