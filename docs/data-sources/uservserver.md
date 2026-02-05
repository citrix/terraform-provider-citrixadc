---
subcategory: "User"
---

# Data Source: citrixadc_uservserver

The citrixadc_uservserver data source is used to retrieve information about an existing user virtual server.

## Example Usage

```terraform
data "citrixadc_uservserver" "example" {
  name = "my_user_vserver"
}

output "uservserver_ipaddress" {
  value = data.citrixadc_uservserver.example.ipaddress
}

output "uservserver_port" {
  value = data.citrixadc_uservserver.example.port
}

output "uservserver_protocol" {
  value = data.citrixadc_uservserver.example.userprotocol
}
```

## Argument Reference

* `name` - (Required) Name for the virtual server.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The id of the uservserver. It has the same value as the `name` attribute.
* `comment` - Comments associated with the virtual server.
* `defaultlb` - Name of the default Load Balancing virtual server used for load balancing of services.
* `ipaddress` - IPv4 or IPv6 address assigned to the virtual server.
* `params` - Comments associated with the protocol.
* `port` - Port number for the virtual server.
* `state` - Initial state of the user vserver. Possible values: ENABLED, DISABLED.
* `userprotocol` - User protocol used by the service.
