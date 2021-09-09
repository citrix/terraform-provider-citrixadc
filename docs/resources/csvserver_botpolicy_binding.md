---
subcategory: "Content Switching"
---

# Resource: csvserver_botpolicy_binding

The csvserver_botpolicy_binding resource is used to bind a bot policy to csvserver.


## Example usage

```hcl
resource "citrixadc_csvserver_botpolicy_binding" "tf_csvserver_botpolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = citrixadc_botpolicy.tf_botpolicy.name
	priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}

resource citrixadc_botpolicy tf_botpolicy {
	name = "tf_botpolicy"  
	profilename = "BOT_BYPASS"
	rule  = "true"
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

* `id` - The id of the csvserver_botpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_botpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_csvserver_botpolicy_binding.tf_csvserver_botpolicy_binding tf_csvserver,tf_botpolicy
```
