---
subcategory: "AppFlow"
---

# Resource: appflowglobal_appflowpolicy_binding

The appflowglobal_appflowpolicy_binding resource is used to create appflowglobal_appflowpolicy_binding.


## Example usage

```hcl
resource "citrixadc_appflowglobal_appflowpolicy_binding" "tf_appflowglobal_appflowpolicy_binding" {
  policyname     = citrixadc_appflowpolicy.tf_appflowpolicy.name
  globalbindtype = "SYSTEM_GLOBAL"
  type           = "REQ_OVERRIDE"
  priority       = 55
}

resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
  name   = "test_policy"
  action = citrixadc_appflowaction.tf_appflowaction.name
  rule   = "client.TCP.DSTPORT.EQ(22)"
}
resource "citrixadc_appflowaction" "tf_appflowaction" {
  name            = "test_action"
  collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name]
  securityinsight = "ENABLED"
  botinsight      = "ENABLED"
  videoanalytics  = "ENABLED"
}
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
  name      = "tf_collector"
  ipaddress = "192.168.2.2"
  port      = 80
}
```


## Argument Reference

* `policyname` - (Required) Name of the AppFlow policy.
* `globalbindtype` - (Optional) 0
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `labelname` - (Optional) Name of the label to invoke if the current policy evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label to invoke. Specify vserver for a policy label associated with a virtual server, or policylabel for a user-defined policy label.
* `priority` - (Optional) Specifies the priority of the policy.
* `type` - (Optional) Global bind point for which to show detailed information about the policies bound to the bind point.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowglobal_appflowpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A appflowglobal_appflowpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowglobal_appflowpolicy_binding.tf_appflowglobal_appflowpolicy_binding test_policy
```
