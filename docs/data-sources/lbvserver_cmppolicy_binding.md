---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_cmppolicy_binding

The lbvserver_cmppolicy_binding data source allows you to retrieve information about the binding between a load balancing virtual server and a compression policy.

## Example Usage

```terraform
data "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
  name       = "tf_lbvserver"
  policyname = "tf_cmppolicy"
  bindpoint  = "RESPONSE"
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_cmppolicy_binding.tf_bind.gotopriorityexpression
}

output "labeltype" {
  value = data.citrixadc_lbvserver_cmppolicy_binding.tf_bind.labeltype
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `bindpoint` - (Required) The bindpoint to which the policy is bound.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_cmppolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `priority` - Priority.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
