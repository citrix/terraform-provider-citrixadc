---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_servicegroup_binding

The lbvserver_servicegroup_binding data source allows you to retrieve information about a service group binding to a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_lbvserver_servicegroup_binding" "tf_binding" {
    name = "tf_lbvserver"
    servicegroupname = "tf_servicegroup"
}

output "name" {
  value = data.citrixadc_lbvserver_servicegroup_binding.tf_binding.name
}

output "servicegroupname" {
  value = data.citrixadc_lbvserver_servicegroup_binding.tf_binding.servicegroupname
}

output "weight" {
  value = data.citrixadc_lbvserver_servicegroup_binding.tf_binding.weight
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `servicegroupname` - (Required) The service group name bound to the selected load balancing virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_servicegroup_binding. It is a system-generated identifier.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `weight` - Integer specifying the weight of the service. A larger number specifies a greater weight. Defines the capacity of the service relative to the other services in the load balancing configuration. Determines the priority given to the service in load balancing decisions.
