---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_policymap_binding

The crvserver_policymap_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and a policy map.

## Example Usage

```terraform
data "citrixadc_crvserver_policymap_binding" "crvserver_policymap_binding" {
    name       = "my_vserver"
    policyname = "my_policymap"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_policymap_binding.crvserver_policymap_binding.name
}

output "policy_name" {
  value = data.citrixadc_crvserver_policymap_binding.crvserver_policymap_binding.policyname
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Name of the policy map to be bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_policymap_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `bindpoint` - For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - Name of the label to be invoked.
* `labeltype` - Type of label to be invoked.
* `priority` - An unsigned integer that determines the priority of the policy relative to other policies bound to this cache redirection virtual server. The lower the value, higher the priority.
* `targetvserver` - The CSW target server names.
