---
subcategory: "Load Balancing"
---

# Resource: lbvserver\_transformpolicy\_binding

The lbvserver\_transformpolicy\_binding resource is used to bind load balancing virtual servers to transform policies.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
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

resource "citrixadc_lbvserver_transformpolicy_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_transformpolicy.tf_trans_policy.name
  priority = 100
  bindpoint = "REQUEST"
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_transformpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver\_transformpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_transformpolicy_binding.tf_binding tf_lbvserver,tf_trans_policy
```
