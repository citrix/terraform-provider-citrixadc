---
subcategory: "Cache Redirection"
---

# Resource: crvserver_spilloverpolicy_binding

The crvserver_spilloverpolicy_binding resource is used to bind a spillover policy to a cache redirection virtual server.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}

resource "citrixadc_crvserver_spilloverpolicy_binding" "crvserver_spilloverpolicy_binding" {
  name                   = citrixadc_crvserver.crvserver.name
  policyname             = "tf_spilloverpolicy"
  bindpoint              = "REQUEST"
  gotopriorityexpression = "END"
  invoke                 = false
  priority               = 1
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - (Optional) Name of the label to be invoked.
* `labeltype` - (Optional) Type of label to be invoked.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_spilloverpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A crvserver_spilloverpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_spilloverpolicy_binding.crvserver_spilloverpolicy_binding my_vserver,tf_spilloverpolicy
```
