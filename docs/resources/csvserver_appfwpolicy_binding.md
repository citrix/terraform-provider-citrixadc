---
subcategory: "Content Switching"
---

# Resource: csvserver_appfwpolicy_binding

The csvserver_appfwpolicy_binding resource is used to bind an application firewall policy to a content switching virtual server.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_appfwprofile" "tf_appfwprofile" {
  name = "tf_appfwprofile"
  type = ["HTML"]
}

resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
  name        = "tf_appfwpolicy"
  profilename = citrixadc_appfwprofile.tf_appfwprofile.name
  rule        = "true"
}

resource "citrixadc_csvserver_appfwpolicy_binding" "tf_bind" {
  name                   = citrixadc_csvserver.tf_csvserver.name
  policyname             = citrixadc_appfwpolicy.tf_appfwpolicy.name
  priority               = 100
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_appfwpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver_appfwpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_appfwpolicy_binding.tf_bind tf_csvserver,tf_appfwpolicy
```
