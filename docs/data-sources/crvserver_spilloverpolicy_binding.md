---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_spilloverpolicy_binding

This data source is used to retrieve information about a specific spillover policy binding to a cache redirection virtual server.

## Example Usage

```terraform
data "citrixadc_crvserver_spilloverpolicy_binding" "example" {
  name       = "my_vserver"
  policyname = "my_spilloverpolicy"
}

output "policy_name" {
  value = data.citrixadc_crvserver_spilloverpolicy_binding.example.policyname
}

output "target_vserver" {
  value = data.citrixadc_crvserver_spilloverpolicy_binding.example.targetvserver
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the binding.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `priority` - The priority for the policy.
* `labelname` - Name of the label to be invoked.
* `labeltype` - Type of label to be invoked.
* `targetvserver` - Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.
