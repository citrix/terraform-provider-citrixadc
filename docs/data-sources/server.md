---
subcategory: "Basic"
---

# Data Source: server

The server data source allows you to retrieve information about a server configuration.

## Example Usage

```terraform
data "citrixadc_server" "tf_server" {
  name = "test_server"
}

output "ipaddress" {
  value = data.citrixadc_server.tf_server.ipaddress
}

output "state" {
  value = data.citrixadc_server.tf_server.state
}
```

## Argument Reference

* `name` - (Required) Name of the server.

## Attribute Reference

The following attributes are available:

* `id` - The id of the server. It is a system-generated identifier.
* `name` - Name for the server.
* `ipaddress` - IPv4 address of the server.
* `ipv6address` - IPv6 address of the server.
* `domain` - Domain name of the server.
* `state` - Initial state of the server. Possible values: `ENABLED`, `DISABLED`.
* `comment` - Any information about the server.
* `td` - Traffic Domain ID.
* `translationip` - IP address used to transform the server's IP address.
* `translationmask` - The netmask of the translation IP.

### Read-only server metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_server` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `statechangetimesec` - Time when last state change happened. Seconds part.
* `tickssincelaststatechange` - Time in 10 millisecond ticks since the last state change.
* `autoscale` - Auto scale option for a servicegroup. Possible values: `DISABLED`, `DNS`, `POLICY`, `CLOUD`, `API`.
* `usip` - Whether the client's IP address is used as the source IP address when initiating a connection to the server. Possible values: `YES`, `NO`.
* `cka` - Whether client keep-alive is enabled for the service group. Possible values: `YES`, `NO`.
* `tcpb` - Whether TCP buffering is enabled for the service group. Possible values: `YES`, `NO`.
* `cmp` - Whether compression is enabled for the specified service. Possible values: `YES`, `NO`.
* `cacheable` - Whether the transparent cache redirection virtual server is used to forward the request to the cache server. Possible values: `YES`, `NO`.
* `sp` - Whether surge protection is enabled for the service group. Possible values: `ON`, `OFF`.
