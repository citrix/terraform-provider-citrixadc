---
subcategory: "Content Switching"
---

# Resource: csvserver_appqoepolicy_binding

The csvserver_appqoepolicy_binding resource is used to bind an AppQoE policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver_appqoepolicy_binding" "tf_csvserver_appqoepolicy_binding" {
  name       = citrixadc_csvserver.tf_csvserver.name
  policyname = citrixadc_appqoepolicy.tf_appqoepolicy.name
  bindpoint  = "REQUEST"
  priority   = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
  name        = "tf_csvserver"
  ipv46       = "10.202.11.11"
  port        = 8080
  servicetype = "HTTP"
}

resource "citrixadc_appqoeaction" "tf_appqoeaction" {
  name      = "tf_appqoeaction"
  priority  = "HIGH"
  respondwith = "NS"
}

resource "citrixadc_appqoepolicy" "tf_appqoepolicy" {
  name   = "tf_appqoepolicy"
  rule   = "true"
  action = citrixadc_appqoeaction.tf_appqoeaction.name
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Required) Priority for the policy.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_appqoepolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_appqoepolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_appqoepolicy_binding.tf_csvserver_appqoepolicy_binding tf_csvserver,tf_appqoepolicy
```
