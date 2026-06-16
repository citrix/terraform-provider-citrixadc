---
subcategory: "Content Switching"
---

# Data Source: csvserver_auditnslogpolicy_binding

The csvserver_auditnslogpolicy_binding data source allows you to retrieve information about an audit nslog policy bound to a content switching virtual server.


## Example usage

```terraform
data "citrixadc_csvserver_auditnslogpolicy_binding" "tf_csvserver_auditnslogpolicy_binding" {
  name       = "tf_csvserver"
  policyname = "tf_auditnslogpolicy"
}

output "priority" {
  value = data.citrixadc_csvserver_auditnslogpolicy_binding.tf_csvserver_auditnslogpolicy_binding.priority
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `policyname` - (Required) Policies bound to this vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_auditnslogpolicy_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.
* `priority` - Priority for the policy.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `bindpoint` - Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time.
* `invoke` - Invoke a policy label if this policy's rule evaluates to TRUE.
* `labeltype` - Type of label to be invoked.
* `labelname` - Name of the label to be invoked.
* `targetlbvserver` - Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
