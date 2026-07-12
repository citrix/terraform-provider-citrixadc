---
subcategory: "Cache Redirection"
---

# Resource: crvserver_filterpolicy_binding

The crvserver_filterpolicy_binding resource is used to create CRvserver Filterpolicy Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
  name      = "tf_filterpolicy"
  reqaction = "DROP"
  rule      = "REQ.HTTP.URL CONTAINS http://abcd.com"
}
resource "citrixadc_crvserver_filterpolicy_binding" "crvserver_filterpolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = citrixadc_filterpolicy.tf_filterpolicy.name
  priority   = 10
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_filterpolicy_binding. It is the concatenation of the `name`, `policyname` and `bindpoint` attributes separated by a comma.


## Import

A crvserver_filterpolicy_binding can be imported using the concatenation of the `name`, `policyname` and `bindpoint` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_crvserver_filterpolicy_binding.crvserver_filterpolicy_binding my_vserver,tf_filterpolicy,REQUEST
```
