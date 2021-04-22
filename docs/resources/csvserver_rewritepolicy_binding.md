---
subcategory: "Content Switching"
---

# Resource: csvserver\_rewritepolicy\_binding

The csvserver\_rewritepolicy\_binding resource is used to bind content switching virtual servers to rewrite policies.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
  name   = "tf_rewrite_policy"
  action = "DROP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_csvserver_rewritepolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
    priority = 100
    bindpoint = "REQUEST"
}
```


## Argument Reference

* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver\_rewritepolicy\_binding. It is the concatenation of `name` and `policyname` separated by a comma.


## Import

A csvserver\_rewritepolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_rewritepolicy_binding.tf_bind tf_csvserver,tf_rewrite_policy
```
