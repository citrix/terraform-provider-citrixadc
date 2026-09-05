---
subcategory: "Basic"
---

# Data Source: service

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

### Read-only service metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_service` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `numofconnections` - The number of client side connections that are still open.
* `policyname` - The name of the policy for which this service is bound.
* `serviceconftype` - The configuration type of the service.
* `serviceconftype2` - The configuration type of the service (`Internal`/`Dynamic`/`Configured`).
* `value` - SSL status of the service.
* `gslb` - The GSLB option for the corresponding virtual server (`REMOTE`, `LOCAL`).
* `dup_state` - State value from table (`ENABLED`, `DISABLED`).
* `publicip` - Public IP of the service.
* `publicport` - Public port of the service.
* `svrstate` - The state of the service (for example `UP`, `DOWN`, `OUT OF SERVICE`).
* `monitor_state` - The running state of the monitor on this service.
* `monstatcode` - The code indicating the monitor response.
* `lastresponse` - The string form of `monstatcode`.
* `responsetime` - Response time of this monitor.
* `monstatparam1` - First parameter for use with the message code.
* `monstatparam2` - Second parameter for use with the message code.
* `monstatparam3` - Third parameter for use with the message code.
* `statechangetimesec` - Time when the last state change happened (seconds part).
* `statechangetimemsec` - Time at which the last state change happened (milliseconds part).
* `tickssincelaststatechange` - Time in 10 millisecond ticks since the last state change.
* `stateupdatereason` - State update reason on the secondary node.
* `clmonowner` - The monitoring owner of the service.
* `clmonview` - The view id of the monitoring owner.
* `serviceipstr` - The DBS services IP.
* `oracleserverversion` - Oracle server version (`10G`, `11G`).
* `nodefaultbindings` - Whether the configuration will have default SSL CIPHER and ECC curve bindings (`YES`, `NO`).
* `monuserstatusmesg` - User monitor failure reasons.
* `builtin` - Whether the service is built-in. A list of strings (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`).
* `feature` - The feature to be checked while applying this configuration.
