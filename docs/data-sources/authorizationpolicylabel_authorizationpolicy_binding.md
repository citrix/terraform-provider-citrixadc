---
subcategory: "AUTHORIZATION"
---

# Data Source: authorizationpolicylabel_authorizationpolicy_binding

The authorizationpolicylabel_authorizationpolicy_binding data source allows you to retrieve information about a binding between an authorization policy label and an authorization policy.

## Example Usage

```terraform
data "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "tf_bind" {
  labelname  = "trans_http_url"
  policyname = "tp-authorize-1"
}

output "labelname" {
  value = data.citrixadc_authorizationpolicylabel_authorizationpolicy_binding.tf_bind.labelname
}

output "priority" {
  value = data.citrixadc_authorizationpolicylabel_authorizationpolicy_binding.tf_bind.priority
}
```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name of the authorization policy label to which to bind the policy.
* `policyname` - (Required) Name of the authorization policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authorizationpolicylabel_authorizationpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Specifies the priority of the policy.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then either forward the request or response to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - Type of invocation. Available settings function as follows: * reqvserver - Send the request to the specified request virtual server. * resvserver - Send the response to the specified response virtual server. * policylabel - Invoke the specified policy label.
