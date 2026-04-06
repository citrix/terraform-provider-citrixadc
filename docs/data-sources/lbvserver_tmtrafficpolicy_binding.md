---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_tmtrafficpolicy_binding

The lbvserver_tmtrafficpolicy_binding data source allows you to retrieve information about an existing tmtraffic policy binding to a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_lbvserver_tmtrafficpolicy_binding" "tf_binding" {
  name       = "my_lbvserver"
  policyname = "my_tmtrafficpolicy"
}

output "binding_priority" {
  value = data.citrixadc_lbvserver_tmtrafficpolicy_binding.tf_binding.priority
}

output "binding_bindpoint" {
  value = data.citrixadc_lbvserver_tmtrafficpolicy_binding.tf_binding.bindpoint
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.
* `policyname` - (Required) Name of the policy bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_tmtrafficpolicy_binding. It is a system-generated identifier.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. policylabel - invoke the request or response against the specified user-defined policy label.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
* `priority` - Priority.
