---
subcategory: "Load Balancing"
---

# Resource: lbvserver\_responderpolicy\_binding

The lbvserver\_responderpolicy\_binding resource is used to bind load balancing virtual servers to responder policies.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name   = "tf_responder_policy"
  action = "NOOP"
  rule   = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}

resource "citrixadc_lbvserver_responderpolicy_binding" "tf_bind" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_responderpolicy.tf_responder_policy.name
  priority = 120
  bindpoint = "REQUEST"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver\_responderpolicy\_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver\_responderpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_responderpolicy_binding.tf_bind tf_lbvserver,tf_responder_policy
```
