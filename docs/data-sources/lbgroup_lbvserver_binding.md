---
subcategory: "Load Balancing"
---

# Data Source: lbgroup_lbvserver_binding

The lbgroup_lbvserver_binding data source allows you to retrieve information about a specific binding between a load balancing group and a virtual server.

## Example Usage

```terraform
data "citrixadc_lbgroup_lbvserver_binding" "tf_lbvserverbinding" {
  name        = "tf_lbgroup"
  vservername = "tf_lbvserver"
}

output "lbgroup_name" {
  value = data.citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding.name
}

output "vserver_name" {
  value = data.citrixadc_lbgroup_lbvserver_binding.tf_lbvserverbinding.vservername
}
```

## Argument Reference

* `name` - (Required) Name for the load balancing virtual server group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `vservername` - (Required) Virtual server name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbgroup_lbvserver_binding. It is a system-generated identifier.
