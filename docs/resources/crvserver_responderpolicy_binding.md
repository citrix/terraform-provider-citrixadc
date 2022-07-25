---
subcategory: "Cache Redirection"
---

# Resource: crvserver_responderpolicy_binding

The crvserver_responderpolicy_binding resource is used to create CRvserver Responderpolicy Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy1"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
resource "citrixadc_crvserver_responderpolicy_binding" "crvserver_responderpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_responderpolicy.tf_responderpolicy.name
    priority = 10
  
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_responderpolicy_binding. It has the same value as the `name` attribute.


## Import

A crvserver_responderpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_responderpolicy_binding.crvserver_responderpolicy_binding my_vserver,tf_responderpolicy1
```
