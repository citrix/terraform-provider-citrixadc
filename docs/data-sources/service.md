---
subcategory: "Basic"
---

# Data Source: citrixadc_service

The service data source allows you to retrieve information about a service configuration.

## Example Usage

```terraform
data "citrixadc_service" "tf_service" {
  name = "test_service"
}

output "servicetype" {
  value = data.citrixadc_service.tf_service.servicetype
}

output "port" {
  value = data.citrixadc_service.tf_service.port
}
```

## Argument Reference

* `name` - (Required) Name for the service.

## Attribute Reference

The following attributes are available:

* `id` - The id of the service. It is a system-generated identifier.
* `name` - Name for the service.
* `servername` - Name of the server that hosts the service.
* `servicetype` - Protocol in which data is exchanged with the service. Example: `HTTP`, `SSL`, `TCP`, `UDP`, `DNS`.
* `port` - Port number of the service.
* `ip` - IP address of the service.
* `ipaddress` - IP address of the service.
* `state` - Initial state of the service. Possible values: `ENABLED`, `DISABLED`.
* `maxclient` - Maximum number of simultaneous open connections to the service.
* `maxreq` - Maximum number of requests that can be sent on a persistent connection to the service.
* `cacheable` - Use the transparent cache redirection virtual server to forward requests to the cache server. Possible values: `YES`, `NO`.
* `cip` - Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value.
* `usip` - Use the client's IP address as the source IP address when initiating a connection to the server. Possible values: `YES`, `NO`.
* `useproxyport` - Use the proxy port as the source port when initiating connections with the server. Possible values: `YES`, `NO`.
* `sp` - Enable surge protection for the service. Possible values: `ON`, `OFF`.
* `clttimeout` - Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - Time, in seconds, after which to terminate an idle server connection.
* `comment` - Any information about the service.
