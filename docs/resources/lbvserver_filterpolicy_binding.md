---
subcategory: "Load Balancing"
---

# Resource: lbvserver_filterpolicy_binding

The lbvserver_filterpolicy_binding resource is used to bind a load balancing virtual server to a filter policy.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_filterpolicy" "tf_filterpolicy" {
  name      = "tf_filterpolicy"
  reqaction = "DROP"
  rule      = "REQ.HTTP.URL CONTAINS http://abc.com"
}

resource "citrixadc_lbvserver_filterpolicy_binding" "tf_bind" {
  name       = citrixadc_lbvserver.tf_lbvserver.name
  policyname = citrixadc_filterpolicy.tf_filterpolicy.name
  priority   = 100
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type.
* `labelname` - (Optional) Name of the label invoked.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_filterpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_filterpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_filterpolicy_binding.tf_bind tf_lbvserver,tf_filterpolicy
```
