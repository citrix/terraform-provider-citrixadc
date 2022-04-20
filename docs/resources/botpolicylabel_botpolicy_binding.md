---
subcategory: "Bot"
---

# Resource: botpolicylabel_botpolicy_binding

The botpolicylabel_botpolicy_binding resource is used to bind botpolicy to botpolicylabel.


## Example usage

```hcl
resource "citrixadc_botpolicylabel" "tf_botpolicylabel" {
  labelname = "tf_botpolicylabel"
}
resource "citrixadc_botpolicy" "tf_botpolicy" {
  name        = "tf_botpolicy"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}
resource "citrixadc_botpolicylabel_botpolicy_binding" "tf_binding" {
  labelname  = citrixadc_botpolicylabel.tf_botpolicylabel.labelname
  policyname = citrixadc_botpolicy.tf_botpolicy.name
  priority   = 50
}
```


## Argument Reference

* `labelname` - (Required) Name of the bot policy label to which to bind the policy.
* `policyname` - (Required) Name of the bot policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.
* `invoke_labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke.  * If labelType is vserver, name of the virtual server.
* `labeltype` - (Optional) Type of policy label to invoke. Available settings function as follows: * vserver - Invoke an unnamed policy label associated with a virtual server. * policylabel - Invoke a user-defined policy label.
* `priority` - (Optional) Specifies the priority of the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botpolicylabel_botpolicy_binding. It is the concatenation of `labelname` and `policyname`  attributes seperated by comma.


## Import

A botpolicylabel_botpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botpolicylabel_botpolicy_binding.tf_binding labelname,policyname
```
