---
subcategory: "Application Firewall"
---

# Data Source: appfwglobal_appfwpolicy_binding

The appfwglobal_appfwpolicy_binding data source allows you to retrieve information about a specific appfwpolicy binding to appfwglobal configuration.

## Example Usage

```terraform
data "citrixadc_appfwglobal_appfwpolicy_binding" "tf_binding" {
  policyname = "tf_appfwpolicy"
  type       = "REQ_DEFAULT"
}

output "policyname" {
  value = data.citrixadc_appfwglobal_appfwpolicy_binding.tf_binding.policyname
}

output "globalbindtype" {
  value = data.citrixadc_appfwglobal_appfwpolicy_binding.tf_binding.globalbindtype
}

output "priority" {
  value = data.citrixadc_appfwglobal_appfwpolicy_binding.tf_binding.priority
}
```

## Argument Reference

* `policyname` - (Required) Name of the policy.
* `type` - (Required) Bind point to which to policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the appfwglobal_appfwpolicy_binding. It is a system-generated identifier.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - Type of policy label invocation.
* `priority` - The priority of the policy.
* `state` - Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.
