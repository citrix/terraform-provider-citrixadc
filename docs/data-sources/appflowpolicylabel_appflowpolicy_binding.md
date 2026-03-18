---
subcategory: "AppFlow"
---

# Data Source: appflowpolicylabel_appflowpolicy_binding

The appflowpolicylabel_appflowpolicy_binding data source allows you to retrieve information about a specific binding between an AppFlow policy label and an AppFlow policy.

## Example Usage

```terraform
data "citrixadc_appflowpolicylabel_appflowpolicy_binding" "example" {
  labelname  = "tf_policylabel"
  policyname = "test_policy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_appflowpolicylabel_appflowpolicy_binding.example.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_appflowpolicylabel_appflowpolicy_binding.example.invoke
}
```

## Argument Reference

* `labelname` - (Required) Name of the policy label to which to bind the policy.
* `policyname` - (Required) Name of the AppFlow policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the appflowpolicylabel_appflowpolicy_binding. It is a system-generated identifier.
* `invoke` - Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `invoke_labelname` - Name of the label to invoke if the current policy evaluates to TRUE.
* `labeltype` - Type of policy label to be invoked.
* `priority` - Specifies the priority of the policy.
