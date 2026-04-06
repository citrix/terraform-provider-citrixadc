---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_authorizationpolicy_binding

The lbvserver_authorizationpolicy_binding data source allows you to retrieve information about the binding between a load balancing virtual server and an authorization policy.

## Example Usage

```terraform
data "citrixadc_lbvserver_authorizationpolicy_binding" "tf_lbvserver_authorizationpolicy_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_authorizationpolicy"
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_authorizationpolicy_binding.tf_lbvserver_authorizationpolicy_binding.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_lbvserver_authorizationpolicy_binding.tf_lbvserver_authorizationpolicy_binding.invoke
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_authorizationpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Priority.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
