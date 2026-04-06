---
subcategory: "GSLB"
---

# Data Source: gslbvserver_gslbservicegroup_binding

The gslbvserver_gslbservicegroup_binding data source allows you to retrieve information about a GSLB virtual server's service group binding.

## Example Usage

```terraform
data "citrixadc_gslbvserver_gslbservicegroup_binding" "tf_gslbvserver_gslbservicegroup_binding" {
  name             = "Gslbv_server"
  servicegroupname = "tf_gslbvservicegroup"
}

output "name" {
  value = data.citrixadc_gslbvserver_gslbservicegroup_binding.tf_gslbvserver_gslbservicegroup_binding.name
}

output "servicegroupname" {
  value = data.citrixadc_gslbvserver_gslbservicegroup_binding.tf_gslbvserver_gslbservicegroup_binding.servicegroupname
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `servicegroupname` - (Required) The GSLB service group name bound to the selected GSLB virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_gslbservicegroup_binding. It is a system-generated identifier.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
