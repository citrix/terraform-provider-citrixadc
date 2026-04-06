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

* `id` - The id of the crvserver_policymap_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or a policy label.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `priority` - Priority for the policy.
* `targetvserver` - Name of the virtual server to which content is forwarded.
