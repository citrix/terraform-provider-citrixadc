---
subcategory: "Content Switching"
---

# Resource: csvserver\_responderpolicy\_binding

The csvserver\_responderpolicy\_binding resource is used to bind content switching virtual servers with responder policies


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}

resource "citrixadc_csvserver_responderpolicy_binding" "tf_bind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_responderpolicy.tf_responder_policy.name
    priority = 100
    bindpoint = "REQUEST"
}
```


## Argument Reference

* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver\_responderpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver\_responderpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_responderpolicy_binding.tf_bind tf_csvserver,tf_responder_policy
```
