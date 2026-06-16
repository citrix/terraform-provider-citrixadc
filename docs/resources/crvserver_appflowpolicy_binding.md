---
subcategory: "Cache Redirection"
---

# Resource: crvserver_appflowpolicy_binding

The crvserver_appflowpolicy_binding resource is used to create CRvserver Appflowpolicy Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  arp         = "OFF"
}

resource "citrixadc_appflowpolicy" "demo_appflowpolicy" {
  name       = "demo_appflowpolicy"
  rule       = "true"
  action     = "AF_ACTION_DEFAULT"
}

resource "citrixadc_crvserver_appflowpolicy_binding" "crvserver_appflowpolicy_binding" {
  name                   = citrixadc_crvserver.crvserver.name
  policyname             = citrixadc_appflowpolicy.demo_appflowpolicy.name
  priority               = 1
  gotopriorityexpression = "END"
}
```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label invoked.
* `labeltype` - (Optional) The invocation type.
* `policyname` - (Required) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_appflowpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A crvserver_appflowpolicy_binding can be imported using the concatenation of the `name` and `policyname` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_crvserver_appflowpolicy_binding.crvserver_appflowpolicy_binding my_vserver,demo_appflowpolicy
```
