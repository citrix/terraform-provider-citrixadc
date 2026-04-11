---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_dnspolicy64_binding

The lbvserver_dnspolicy64_binding data source allows you to retrieve information about a DNS policy64 binding to a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_lbvserver_dnspolicy64_binding" "tf_lbvserver_dnspolicy64_binding" {
  name       = "tf_lbvserver"
  policyname = "tf_dnspolicy64"
}

output "bindpoint" {
  value = data.citrixadc_lbvserver_dnspolicy64_binding.tf_lbvserver_dnspolicy64_binding.bindpoint
}

output "priority" {
  value = data.citrixadc_lbvserver_dnspolicy64_binding.tf_lbvserver_dnspolicy64_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_dnspolicy64_binding.tf_lbvserver_dnspolicy64_binding.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the load balancing virtual server.
* `policyname` - (Required) Name of the DNS policy64 bound to the LB vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `priority` - Priority.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. policylabel - invoke the request or response against the specified user-defined policy label.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
* `id` - The id of the lbvserver_dnspolicy64_binding. It is a system-generated identifier.
