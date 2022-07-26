---
subcategory: "AppFlow"
---

# Resource: appflowpolicylabel_appflowpolicy_binding

The appflowpolicylabel_appflowpolicy_binding resource is used to create appflowpolicylabel_appflowpolicy_binding.


## Example usage

```hcl
resource "citrixadc_appflowpolicylabel_appflowpolicy_binding" "tf_appflowpolicylabel_appflowpolicy_binding" {
  labelname  = "tf_policylabel"
  policyname = "test_policy"
  priority   = 30
}
# -------------------- ADC CLI ----------------------------

#add appflow collector tf_collector -IPAddress 192.168.2.2
#add appflowaction test_action -collectors tf_collector
#add appflowpolicy test_policy client.TCP.DSTPORT.EQ(22) test_action
#add appflowpolicylabel tf_policylabel -policylabeltype OTHERTCP


# ---------------- NOT YET IMPLEMENTED -------------------
# resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
#   labelname       = "tf_policylabel"
#   policylabeltype = "OTHERTCP"
# }

# resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
#   name      = "test_policy"
#   action    = citrixadc_appflowaction.tf_appflowaction.name
#   rule      = "client.TCP.DSTPORT.EQ(22)"
# }
# resource "citrixadc_appflowaction" "tf_appflowaction" {
#   name = "test_action"
#   collectors     = [citrixadc_appflowcollector.tf_appflowcollector.name]
#   securityinsight = "ENABLED"
#   botinsight      = "ENABLED"
#   videoanalytics  = "ENABLED"
# }
# resource "citrixadc_appflowcollector" "tf_appflowcollector" {
#   name      = "tf_collector"
#   ipaddress = "192.168.2.2"
#   port      = 80
# }
```


## Argument Reference

* `labelname` - (Required) Name of the policy label to which to bind the policy.
* `policyname` - (Required) Name of the AppFlow policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label to be invoked.
* `priority` - (Optional) Specifies the priority of the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowpolicylabel_appflowpolicy_binding is the concatenation of `labelname` and `policyname` attributes separated by comma.


## Import

A appflowpolicylabel_appflowpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowpolicylabel_appflowpolicy_binding.tf_appflowpolicylabel_appflowpolicy_binding tf_policylabel,test_policy