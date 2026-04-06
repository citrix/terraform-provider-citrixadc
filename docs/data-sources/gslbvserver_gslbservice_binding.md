---
subcategory: "GSLB"
---

# Data Source: gslbvserver_gslbservice_binding

The gslbvserver_gslbservice_binding data source allows you to retrieve information about a GSLB virtual server to GSLB service binding.

## Example Usage

```terraform
data "citrixadc_gslbvserver_gslbservice_binding" "tf_gslbvserver_gslbservice_binding" {
  name        = "gslb_vserver"
  servicename = "gslb1vservice"
}

output "name" {
  value = data.citrixadc_gslbvserver_gslbservice_binding.tf_gslbvserver_gslbservice_binding.name
}

output "servicename" {
  value = data.citrixadc_gslbvserver_gslbservice_binding.tf_gslbvserver_gslbservice_binding.servicename
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `servicename` - (Required) Name of the GSLB service for which to change the weight.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_gslbservice_binding. It is a composite identifier in the format "name,servicename".
* `domainname` - Domain name for which to change the time to live (TTL) and/or backup service IP address.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `weight` - Weight for the service.
