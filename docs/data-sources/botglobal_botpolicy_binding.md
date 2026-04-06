---
subcategory: "Bot"
---

# Data Source: botglobal_botpolicy_binding

The botglobal_botpolicy_binding data source allows you to retrieve information about a bot policy binding to the bot global configuration.

## Example Usage

```terraform
data "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
  policyname = "tf_botpolicy"
  type       = "REQ_OVERRIDE"
}

output "policyname" {
  value = data.citrixadc_botglobal_botpolicy_binding.tf_binding.policyname
}

output "priority" {
  value = data.citrixadc_botglobal_botpolicy_binding.tf_binding.priority
}

output "type" {
  value = data.citrixadc_botglobal_botpolicy_binding.tf_binding.type
}
```

## Argument Reference

* `policyname` - (Required) Name of the bot policy.
* `type` - (Required) Specifies the bind point whose policies you want to display. Available settings function as follows:
  * REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.
  * REQ_DEFAULT - Binds the policy to the default request queue.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - The global bind type of the binding. Default value: "SYSTEM_GLOBAL"
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server.
* `labelname` - Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
* `labeltype` - Type of invocation. Available settings function as follows:
  * vserver - Forward the request to the specified virtual server.
  * policylabel - Invoke the specified policy label.
* `id` - The id of the botglobal_botpolicy_binding. It is a system-generated identifier.
