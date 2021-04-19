---
subcategory: "Content Switching"
---

# Resource: csvserver\_transformpolicy\_binding

The csvserver\_transformpolicy\_binding resource is used to bind content swtching virtual servers with transform policies.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_csvserver_transformpolicy_binding" "tf_binding" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_transformpolicy.tf_trans_policy.name
    priority = 100
    bindpoint = "REQUEST"
    gotopriorityexpression = "END"
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

* `id` - The id of the csvserver\_transformpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver\_transformpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_transformpolicy_binding.tf_binding tf_csvserver,tf_trans_policy
```
