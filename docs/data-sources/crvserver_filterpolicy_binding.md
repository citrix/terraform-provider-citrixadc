---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_filterpolicy_binding

The crvserver_filterpolicy_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and a filter policy.

## Example Usage

```terraform
data "citrixadc_crvserver_filterpolicy_binding" "crvserver_filterpolicy_binding" {
  name       = "my_vserver"
  policyname = "my_filter_policy"
  bindpoint  = "REQUEST"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_filterpolicy_binding.crvserver_filterpolicy_binding.name
}

output "policy_name" {
  value = data.citrixadc_crvserver_filterpolicy_binding.crvserver_filterpolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_crvserver_filterpolicy_binding.crvserver_filterpolicy_binding.priority
}

output "bindpoint" {
  value = data.citrixadc_crvserver_filterpolicy_binding.crvserver_filterpolicy_binding.bindpoint
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Required) The bindpoint to which the policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_filterpolicy_binding. It is the concatenation of the `name`, `policyname` and `bindpoint` attributes separated by a comma.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `priority` - The priority for the policy.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetvserver` - Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.
