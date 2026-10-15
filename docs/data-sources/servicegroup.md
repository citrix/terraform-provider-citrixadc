---
subcategory: "Basic"
---

# Data Source: servicegroup

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
* `aigwprofilename` - Name of the backend AIGW Profile which will be attached to the servicegroup. This parameter enables the servicegroup to process the LLM request/response based on the profile config. Any service item bound to the servicegroup will inherit the backend AIGW Profile bound at the servicegroup level, if it does not have an explicit AIGW Profile given at bind time.
* `appflowlog` - Enable logging of AppFlow information for the specified service group.
* `autodelayedtrofs` - Indicates graceful movement of IP-Port binding/s to TROFS when IP addresses are removed from DNS response. System will wait for monitor response timeout period before moving to TROFS .
* `autodisabledelay` - The time allowed (in seconds) for a graceful shutdown.
* `autodisablegraceful` - Indicates graceful shutdown of the service.
* `bootstrap` - Flag to check if kafka broker servicegroup is of type bootstrap or not.
* `cachetype` - Cache type supported by the cache server.
* `cipheader` - Name of the HTTP header whose value must be set to the IP address of the client.
* `cka` - Enable client keep-alive for the service group.
* `cmp` - Enable compression for the specified service.
* `customserverid` - The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
* `dbsttl` - Specify the TTL for DNS record for domain based service.
* `delay` - Time, in seconds, allocated for a shutdown of the services in the service group.
* `downstateflush` - Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN.
* `dupweight` - weight of the monitor that is bound to servicegroup.
* `hashid` - The hash identifier for the service.
* `healthmonitor` - Monitor the health of this service.
* `httpprofilename` - Name of the HTTP profile that contains HTTP configuration settings for the service group.
* `includemembers` - Display the members of the listed service groups in addition to their settings.
* `maxbandwidth` - Maximum bandwidth, in Kbps, allocated for all the services in the service group.
* `mcpprofilename` - Name of MCP profile which will be attached to the servicegroup.
* `memberport` - member port
* `monconnectionclose` - Close monitoring connections by sending the service a connection termination message with the specified bit set.
* `monitornamesvc` - Name of the monitor bound to the service group. Used to assign a weight to the monitor.
* `monthreshold` - Minimum sum of weights of the monitors that are bound to this service.
* `nameserver` - Specify the nameserver to which the query for bound domain needs to be sent.
* `netprofile` - Network profile for the service group.
* `pathmonitor` - Path monitoring for clustering
* `pathmonitorindv` - Individual Path monitoring decisions.
* `quicprofilename` - Name of QUIC profile which will be attached to the service group.
* `riseapbrstatsmsgcode` - The code indicating the rise apbr status.
* `rtspsessionidremap` - Enable RTSP session ID mapping for the service group.
* `serverid` - The identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `servername` - Name of the server to which to bind the service group.
* `tcpb` - Enable TCP buffering for the service group.
* `tcpprofilename` - Name of the TCP profile that contains TCP configuration settings for the service group.
* `td` - Integer value that uniquely identifies the traffic domain.
* `topicname` - Name of the Kafka topic.
* `wasmmodule` - Name of the WASM module to bind to this service.
* `weight` - Weight to assign to the servers in the service group.

### Read-only servicegroup metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_servicegroup` resource). They are GET-only / Computed. Any attribute the appliance does not return is `null`.

* `numofconnections` - The number of client side connections still open.
* `serviceconftype` - The configuration type of the service group.
* `value` - SSL Status. Possible values = Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled.
* `svrstate` - The state of the service. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED.
* `ip` - IP Address.
* `monstatcode` - The code indicating the monitor response.
* `monstatparam1` - First parameter for use with message code.
* `monstatparam2` - Second parameter for use with message code.
* `monstatparam3` - Third parameter for use with message code.
* `statechangetimemsec` - Time when last state change occurred. Milliseconds part.
* `stateupdatereason` - Checks state update reason on the secondary node.
* `clmonowner` - Tells the mon owner of the service.
* `clmonview` - Tells the view id of the monitoring owner.
* `groupcount` - Servicegroup Count.
* `serviceipstr` - This field shows the dbs services ip.
* `servicegroupeffectivestate` - Indicates the effective servicegroup state based on the state of the bound service items. Possible values = UP, DOWN, OUT OF SERVICE, PARTIAL-UP, PARTIAL-DOWN.
* `nodefaultbindings` - To determine if the configuration is from stylebooks. Possible values = YES, NO.
* `svcitmactsvcs` - The total active service items for an FQDN for SRV type server binding.
* `svcitmboundsvcs` - The total bound items for an FQDN for SRV type server binding.
* `monuserstatusmesg` - User monitor failure reasons.
