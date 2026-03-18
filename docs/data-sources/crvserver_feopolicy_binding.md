---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_feopolicy_binding

The crvserver_feopolicy_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and a front end optimization policy.

## Example Usage

```terraform
data "citrixadc_crvserver_feopolicy_binding" "crvserver_feopolicy_binding" {
  name       = "my_vserver"
  policyname = "my_feo_policy"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_feopolicy_binding.crvserver_feopolicy_binding.name
}

output "policy_name" {
  value = data.citrixadc_crvserver_feopolicy_binding.crvserver_feopolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_crvserver_feopolicy_binding.crvserver_feopolicy_binding.priority
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_feopolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - Name of the label to be invoked.
* `priority` - The priority for the policy.
* `labeltype` - Type of label to be invoked.
* `targetvserver` - Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.
