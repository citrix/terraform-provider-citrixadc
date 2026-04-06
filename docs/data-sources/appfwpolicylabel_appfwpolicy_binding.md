---
subcategory: "Application Firewall"
---

# Data Source: appfwpolicylabel_appfwpolicy_binding

The appfwpolicylabel_appfwpolicy_binding data source allows you to retrieve information about a binding between an application firewall policy label and an application firewall policy.

## Example Usage

```terraform
data "citrixadc_appfwpolicylabel_appfwpolicy_binding" "tf_binding" {
  labelname  = "tf_appfwpolicylabel"
  policyname = "tf_appfwpolicy1"
}

output "labelname" {
  value = data.citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding.labelname
}

output "priority" {
  value = data.citrixadc_appfwpolicylabel_appfwpolicy_binding.tf_binding.priority
}
```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name of the application firewall policy label.
* `policyname` - (Required) Name of the application firewall policy to bind to the policy label.
## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwpolicylabel_appfwpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Positive integer specifying the priority of the policy. A lower number specifies a higher priority. Must be unique within a group of policies that are bound to the same bind point or label. Policies are evaluated in the order of their priority numbers.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows:
    * reqvserver. Invoke the unnamed policy label associated with the specified request virtual server.
    * policylabel. Invoke the specified user-defined policy label.
