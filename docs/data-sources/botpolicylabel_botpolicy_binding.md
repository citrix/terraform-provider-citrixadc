---
subcategory: "Bot"
---

# Data Source: botpolicylabel_botpolicy_binding

The botpolicylabel_botpolicy_binding data source allows you to retrieve information about a binding between a bot policy label and a bot policy.

## Example Usage

```terraform
data "citrixadc_botpolicylabel_botpolicy_binding" "tf_binding" {
  labelname  = "tf_botpolicylabel"
  policyname = "tf_botpolicy"
}

output "labelname" {
  value = data.citrixadc_botpolicylabel_botpolicy_binding.tf_binding.labelname
}

output "priority" {
  value = data.citrixadc_botpolicylabel_botpolicy_binding.tf_binding.priority
}
```

## Argument Reference

The following arguments are required:

* `labelname` - (Required) Name of the bot policy label to which to bind the policy.
* `policyname` - (Required) Name of the bot policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botpolicylabel_botpolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `invoke_labelname` - * If labelType is policylabel, name of the policy label to invoke. 
* If labelType is vserver, name of the virtual server.
* `labeltype` - Type of policy label to invoke. Available settings function as follows:
    * vserver - Invoke an unnamed policy label associated with a virtual server.
    * policylabel - Invoke a user-defined policy label.
