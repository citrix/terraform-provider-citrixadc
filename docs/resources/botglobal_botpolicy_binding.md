---
subcategory: "Bot"
---

# Resource: botglobal_botpolicy_binding

The botglobal_botpolicy_binding resource is used to bind botpolicy to botglobal.


## Example usage

```hcl
resource "citrixadc_botpolicy" "tf_botpolicy" {
  name        = "tf_botpolicy"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}
resource "citrixadc_botglobal_botpolicy_binding" "tf_binding" {
  policyname = citrixadc_botpolicy.tf_botpolicy.name
  priority   = 90
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the bot policy.
* `priority` - (Required) Specifies the priority of the policy.
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server.
* `labelname` - (Optional) Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
* `labeltype` - (Optional) Type of invocation, Available settings function as follows: * vserver - Forward the request to the specified virtual server. * policylabel - Invoke the specified policy label.
* `type` - (Optional) Specifies the bind point whose policies you want to display. Available settings function as follows: * REQ_OVERRIDE - Request override. Binds the policy to the priority request queue. * REQ_DEFAULT - Binds the policy to the default request queue.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botglobal_botpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A botglobal_botpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_botglobal_botpolicy_binding.tf_binding tf_botpolicy
```
