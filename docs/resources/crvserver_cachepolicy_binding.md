---
subcategory: "Cache Redirection"
---

# Resource: crvserver_cachepolicy_binding

The crvserver_cachepolicy_binding resource is used to create CRvserver Cachepolicy Binding.


## Example usage

```hcl
# Since the cachepolicy resource is not yet available on Terraform,
# the tf_cachepolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add cache policy tf_cachepolicy -rule "http.req.url.query.contains(\"IssuePage\")" -action CACHE
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}
resource "citrixadc_crvserver_cachepolicy_binding" "crvserver_cachepolicy_binding" {
  name       = citrixadc_crvserver.crvserver.name
  policyname = "tf_cachepolicy"
  priority   = 10
  bindpoint =  "REQUEST"

}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_cachepolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_cachepolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_cachepolicy_binding.crvserver_cachepolicy_binding my_vserver,tf_cachepolicy
```
