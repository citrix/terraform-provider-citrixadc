---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_cachepolicy_binding

The crvserver_cachepolicy_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and a cache policy.

## Example Usage

```terraform
data "citrixadc_crvserver_cachepolicy_binding" "crvserver_cachepolicy_binding" {
  name       = "my_vserver"
  policyname = "my_cache_policy"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_cachepolicy_binding.crvserver_cachepolicy_binding.name
}

output "policy_name" {
  value = data.citrixadc_crvserver_cachepolicy_binding.crvserver_cachepolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_crvserver_cachepolicy_binding.crvserver_cachepolicy_binding.priority
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_cachepolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `priority` - The priority for the policy.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `targetvserver` - Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.
