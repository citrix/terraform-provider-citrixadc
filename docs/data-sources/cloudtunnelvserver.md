---
subcategory: "Cloud"
---

# Data Source: cloudtunnelvserver

The cloudtunnelvserver data source allows you to retrieve information about a Cloud Tunnel virtual server on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name = "ct_vserver1"
}

output "cloudtunnelvserver_servicetype" {
  value = data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver.servicetype
}

output "cloudtunnelvserver_listenpolicy" {
  value = data.citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver.listenpolicy
}
```


## Argument Reference

* `name` - (Required) Name of the Cloud Tunnel virtual server to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `servicetype` - Service type of the listener using which traffic is tunneled through the cloud tunnel server. Possible values: [ TCP, UDP ].
* `listenpolicy` - String specifying the listen policy for the Cloud Tunnel virtual server. Can be either a named expression or an expression. The Cloud Tunnel virtual server processes only the traffic for which the expression evaluates to true.
* `listenpriority` - Integer specifying the priority of the listen policy. A higher number specifies a lower priority.
* `id` - The id of the cloudtunnelvserver. It has the same value as the `name` attribute.

### Read-only cloudtunnelvserver metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_cloudtunnelvserver` resource). Any attribute the appliance does not return is `null`.

* `state` - The current state of the virtual server, as UP, DOWN, BUSY, and so on.
* `effectivestate` - Effective state of the virtual server.
* `type` - The type of virtual server; for example, CONTENT based or ADDRESS based.
* `ip` - The Virtual IP address of the Cloud Tunnel virtual server.
* `ipv46` - IPv4 or IPv6 address of the Cloud Tunnel virtual server.
* `ippattern` - The IP pattern of the virtual server.
* `port` - Port on which the virtual server listens.
* `range` - Number of IP addresses that the appliance must generate and assign to the virtual server.
* `cachetype` - Virtual server cache type. The options are: TRANSPARENT, REVERSE, and FORWARD.
