---
subcategory: "AppFlow"
---

# Data Source: appflowglobal_appflowpolicy_binding

The appflowglobal_appflowpolicy_binding data source allows you to retrieve information about appflowglobal_appflowpolicy_binding.

## Example Usage

```terraform
data "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
  policyname = "my_appflowpolicy"
  type       = "REQ_OVERRIDE"
}

output "policyname" {
  value = data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding.policyname
}

output "globalbindtype" {
  value = data.citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding.globalbindtype
}
```

## Argument Reference

* `policyname` - (Required) Name of the AppFlow policy.
* `type` - (Required) Global bind point for which to show detailed information about the policies bound to the bind point. Possible values: `REQ_OVERRIDE`, `REQ_DEFAULT`, `OTHERTCP_REQ_OVERRIDE`, `OTHERTCP_REQ_DEFAULT`, `ICA_REQ_OVERRIDE`, `ICA_REQ_DEFAULT`, `HTTPQUIC_REQ_OVERRIDE`, `HTTPQUIC_REQ_DEFAULT`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Indicates whether binding is local or global.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `labelname` - Name of the label to invoke if the current policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Specify vserver for a policy label associated with a virtual server, or policylabel for a user-defined policy label.
* `id` - The id of the appflowglobal_appflowpolicy_binding. It has the same value as the `policyname` attribute.

### Read-only appflowglobal_appflowpolicy_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appflowglobal_appflowpolicy_binding` resource). They are Computed / GET-only, and any attribute the appliance does not return is `null`.

* `flowtype` - Flow type of the bound AppFlow policy.
* `numpol` - The number of policies bound to the bindpoint.
