---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_service_binding

The lbvserver_service_binding data source allows you to retrieve information about an existing binding between a load balancing virtual server and a service.


## Example Usage

```terraform
data "citrixadc_lbvserver_service_binding" "tf_binding" {
  name        = "tf_lbvserver"
  servicename = "tf_service"
}

output "weight" {
  value = data.citrixadc_lbvserver_service_binding.tf_binding.weight
}

output "order" {
  value = data.citrixadc_lbvserver_service_binding.tf_binding.order
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
* `servicename` - (Optional) Service to bind to the virtual server.
* `servicegroupname` - (Optional) Name of the service group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_service_binding. It is the concatenation of the `name` and `servicename` attributes separated by a comma.
* `order` - Order number to be assigned to the service when it is bound to the lb vserver.
* `servicegroupname` - Name of the service group.
* `weight` - Weight to assign to the specified service.

### Read-only lbvserver_service_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_lbvserver_service_binding` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `vserverid` - Vserver Id.
* `vsvrbindsvcip` - Used for showing the ip of bound entities.
* `preferredlocation` - Used for displaying the location of bound services.
* `servicetype` - Protocol used by the service (also called the service type).
* `dynamicweight` - Dynamic weight.
* `orderstr` - Order in string form assigned to the service when it is bound to the lb vserver.
* `curstate` - Current LB vserver state.
* `port` - Port number for the virtual server.
* `cookieipport` - Encryped Ip address and port of the service that is inserted into the set-cookie http header.
* `cookiename` - Cookie name for COOKIE persistence type.
* `vsvrbindsvcport` - Used for showing ports of bound entities.
* `ipv46` - IPv4 or IPv6 address to assign to the virtual server.
