---
subcategory: "Basic"
---

# Data Source: citrixadc_servicegroup

The servicegroup data source allows you to retrieve information about a service group configuration.

## Example Usage

```terraform
data "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "test_servicegroup"
}

output "servicetype" {
  value = data.citrixadc_servicegroup.tf_servicegroup.servicetype
}

output "state" {
  value = data.citrixadc_servicegroup.tf_servicegroup.state
}
```

## Argument Reference

* `servicegroupname` - (Required) Name of the service group.

## Attribute Reference

The following attributes are available:

* `id` - The id of the servicegroup. It is a system-generated identifier.
* `servicegroupname` - Name of the service group.
* `servicetype` - Protocol used to exchange data with the service. Example: `HTTP`, `SSL`, `TCP`, `UDP`, `DNS`.
* `state` - Initial state of the service group. Possible values: `ENABLED`, `DISABLED`.
* `cacheable` - Use the transparent cache redirection virtual server to forward requests to the cache server. Possible values: `YES`, `NO`.
* `cip` - Insert the Client IP header in requests forwarded to the service.
* `usip` - Use client's IP address as the source IP address when initiating connection to the server. Possible values: `YES`, `NO`.
* `useproxyport` - Use the proxy port as the source port when initiating connections with the server. Possible values: `YES`, `NO`.
* `sp` - Enable surge protection for the service group. Possible values: `ON`, `OFF`.
* `clttimeout` - Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - Time, in seconds, after which to terminate an idle server connection.
* `maxclient` - Maximum number of simultaneous open connections for the service group.
* `maxreq` - Maximum number of requests that can be sent on a persistent connection to the service group.
* `comment` - Any information about the service group.
* `autoscale` - Auto scale option for a servicegroup. Possible values: `DISABLED`, `DNS`, `POLICY`.
* `graceful` - Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service. Possible values: `YES`, `NO`.
