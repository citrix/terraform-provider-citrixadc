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
* `internal` - Display only dynamically learned services.
* `accessdown` - Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.
* `aigwprofilename` - Name of the AIGW Profile that contains AIGW Endpoint setting for the service.
* `all` - Display both user-configured and dynamically learned services.
* `appflowlog` - Enable logging of AppFlow information.
* `cachetype` - Cache type supported by the cache server.
* `cipheader` - Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name "client-ip."
* `cka` - Enable client keep-alive for the service.
* `cleartextport` - Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.
* `cmp` - Enable compression for the service.
* `contentinspectionprofilename` - Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service
* `customserverid` - Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.
* `delay` - Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
* `dnsprofilename` - Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS, ADNS-TCP and ADNS-DOT services.
* `downstateflush` - Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
* `graceful` - Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.
* `hashid` - A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.
* `healthmonitor` - Monitor the health of this service. Available settings function as follows: YES - Send probes to check the health of the service. NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.
* `httpprofilename` - Name of the HTTP profile that contains HTTP configuration settings for the service.
* `maxbandwidth` - Maximum bandwidth, in Kbps, allocated to the service.
* `mcpprofilename` - Name of MCP profile which will be attached to the service.
* `monconnectionclose` - Close monitoring connections by sending the service a connection termination message with the specified bit set.
* `monitornamesvc` - Name of the monitor bound to the specified service.
* `monthreshold` - Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.
* `netprofile` - Network profile to use for the service.
* `riseapbrstatsmsgcode` - The code indicating the rise apbr status.
* `pathmonitor` - Path monitoring for clustering
* `pathmonitorindv` - Individual Path monitoring decisions
* `processlocal` - By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.
* `quicprofilename` - Name of QUIC profile which will be attached to the service.
* `rtspsessionidremap` - Enable RTSP session ID mapping for the service.
* `serverid` - The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `tcpb` - Enable TCP buffering for the service.
* `tcpprofilename` - Name of the TCP profile that contains TCP configuration settings for the service.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `wasmmodule` - Name of the WASM module to bind to this service.
* `weight` - Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.
* `snienable` - State of the Server Name Indication (SNI) feature on the service (SSL services only).
* `commonname` - Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL service.
