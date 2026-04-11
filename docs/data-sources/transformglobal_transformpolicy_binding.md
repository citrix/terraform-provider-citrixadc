---
subcategory: "Transform"
---

# Data Source: transformglobal_transformpolicy_binding

The transformglobal_transformpolicy_binding data source allows you to retrieve information about a transform global transform policy binding.

## Example usage

```terraform
data "citrixadc_transformglobal_transformpolicy_binding" "tf_bind" {
  policyname = "tf_trans_policy"
  type       = "REQ_DEFAULT"
}

output "policyname" {
  value = data.citrixadc_transformglobal_transformpolicy_binding.tf_bind.policyname
}

output "priority" {
  value = data.citrixadc_transformglobal_transformpolicy_binding.tf_bind.priority
}

output "type" {
  value = data.citrixadc_transformglobal_transformpolicy_binding.tf_bind.type
}
```

## Argument Reference

* `policyname` - (Required) Name of the transform policy.
* `type` - (Required) Specifies the bind point to which to bind the policy. Available settings function as follows:
  * REQ_OVERRIDE. Request override. Binds the policy to the priority request queue.
  * REQ_DEFAULT. Binds the policy to the default request queue.
  * HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override request queue.
  * HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default request queue.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forwards the request or response to the specified virtual server or evaluates the specified policy label.
* `labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and the label type is Policy Label.
* `labeltype` - Type of invocation. Available settings function as follows:
  * reqvserver - Send the request to the specified request virtual server.
  * resvserver - Send the response to the specified response virtual server.
  * policylabel - Invoke the specified policy label.
* `priority` - Specifies the priority of the policy.
* `id` - The id of the transformglobal_transformpolicy_binding. It is a system-generated identifier.
