---
subcategory: "Cache Redirection"
---

# Data Source: crvserver_lbvserver_binding

The crvserver_lbvserver_binding data source allows you to retrieve information about a specific binding between a cache redirection virtual server and a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_crvserver_lbvserver_binding" "crvserver_lbvserver_binding" {
  name      = "my_vserver"
  lbvserver = "my_lb_vserver"
}

output "vserver_name" {
  value = data.citrixadc_crvserver_lbvserver_binding.crvserver_lbvserver_binding.name
}

output "lb_vserver" {
  value = data.citrixadc_crvserver_lbvserver_binding.crvserver_lbvserver_binding.lbvserver
}
```

## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `lbvserver` - (Required) The Default target server name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_lbvserver_binding. It is a system-generated identifier.
