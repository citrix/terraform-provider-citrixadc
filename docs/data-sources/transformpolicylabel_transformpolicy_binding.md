---
subcategory: "Transform"
---

# Data Source: transformpolicylabel_transformpolicy_binding

The transformpolicylabel_transformpolicy_binding data source allows you to retrieve information about a transform policylabel transform policy binding.

## Example usage

```terraform
data "citrixadc_transformpolicylabel_transformpolicy_binding" "tf_bind" {
  labelname  = "label_1"
  policyname = "tf_trans_policy"
}

output "labelname" {
  value = data.citrixadc_transformpolicylabel_transformpolicy_binding.tf_bind.labelname
}

output "policyname" {
  value = data.citrixadc_transformpolicylabel_transformpolicy_binding.tf_bind.policyname
}

output "priority" {
  value = data.citrixadc_transformpolicylabel_transformpolicy_binding.tf_bind.priority
}
```

## Argument Reference

* `labelname` - (Required) Name of the URL Transformation policy label to which to bind the policy.
* `policyname` - (Required) Name of the URL Transformation policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - Name of the policy label.
* `labeltype` - Type of invocation. Available settings function as follows:
  * reqvserver - Forward the request to the specified request virtual server.
  * policylabel - Invoke the specified policy label.
* `priority` - Specifies the priority of the policy.
* `id` - The id of the transformpolicylabel_transformpolicy_binding. It is a system-generated identifier.
