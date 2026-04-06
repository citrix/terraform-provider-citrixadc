---
subcategory: "Content Inspection"
---

# Data Source: contentinspectionpolicylabel_contentinspectionpolicy_binding

The contentinspectionpolicylabel_contentinspectionpolicy_binding data source allows you to retrieve information about a specific contentinspectionpolicy binding to contentinspectionpolicylabel configuration.

## Example Usage

```terraform
data "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_contentinspectionpolicylabel_contentinspectionpolicy_binding" {
    labelname  = "tf_contentinspectionpolicylabel_ds"
    policyname = "tf_contentinspectionpolicy_ds"
}

output "labelname" {
  value = data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding.labelname
}

output "policyname" {
  value = data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_contentinspectionpolicylabel_contentinspectionpolicy_binding.priority
}
```

## Argument Reference

* `labelname` - (Required) Name of the contentInspection policy label to which to bind the policy.
* `policyname` - (Required) Name of the contentInspection policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Specifies the priority of the policy.
* `id` - The id of the contentinspectionpolicylabel_contentinspectionpolicy_binding. It is a system-generated identifier.
* `invoke` - Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `labeltype` - Type of invocation. Available settings function as follows: reqvserver - Forward the request to the specified request virtual server. resvserver - Forward the response to the specified response virtual server. policylabel - Invoke the specified policy label.
