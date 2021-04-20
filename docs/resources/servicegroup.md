---
subcategory: "Basic"
---

# Resource: servicegroup

The servicegroup resource is used to create servicegroups.


## Example usage

```hcl
resource "citrixadc_servicegroup" "tf_servicegroup" {

  servicegroupname = "tf_servicegroup"
  servicetype = "HTTP"
  servicegroupmembers = ["172.20.0.9:80:10", "172.20.0.10:80:10", "172.20.0.11:8080:20"]
  lbvservers = [ citrixadc_lbvserver.tf_lbvserver1.name, citrixadc_lbvserver.tf_lbvserver2.name ]
}
```


## Argument Reference

* `servicegroupname` - (Optional) Name of the service group.
* `servicetype` - (Optional) Protocol used to exchange data with the service. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]
* `cachetype` - (Optional) Cache type supported by the cache server. Possible values: [ TRANSPARENT, REVERSE, FORWARD ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `maxclient` - (Optional) 
* `maxreq` - (Optional) Note: Connection requests beyond this value are rejected.
* `cacheable` - (Optional) Use the transparent cache redirection virtual server to forward the request to the cache server. Note: Do not set this parameter if you set the Cache Type. Possible values: [ YES, NO ]
* `cip` - (Optional) Insert the Client IP header in requests forwarded to the service. Possible values: [ ENABLED, DISABLED ]
* `cipheader` - (Optional) Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.
* `usip` - (Optional) Use client's IP address as the source IP address when initiating connection to the server. With the NO setting, which is the default, a mapped IP (MIP) address or subnet IP (SNIP) address is used as the source IP address to initiate server side connections. Possible values: [ YES, NO ]
* `pathmonitor` - (Optional) Path monitoring for clustering. Possible values: [ YES, NO ]
* `pathmonitorindv` - (Optional) Individual Path monitoring decisions. Possible values: [ YES, NO ]
* `useproxyport` - (Optional) Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES. Possible values: [ YES, NO ]
* `healthmonitor` - (Optional) Monitor the health of this service.  Available settings function as follows: YES - Send probes to check the health of the service. NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times. Possible values: [ YES, NO ]
* `sc` - (Optional) State of the SureConnect feature for the service group. Possible values: [ on, off ]
* `sp` - (Optional) Enable surge protection for the service group. Possible values: [ on, off ]
* `rtspsessionidremap` - (Optional) Enable RTSP session ID mapping for the service group. Possible values: [ on, off ]
* `clttimeout` - (Optional) Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - (Optional) Time, in seconds, after which to terminate an idle server connection.
* `cka` - (Optional) Enable client keep-alive for the service group. Possible values: [ YES, NO ]
* `tcpb` - (Optional) Enable TCP buffering for the service group. Possible values: [ YES, NO ]
* `cmp` - (Optional) Enable compression for the specified service. Possible values: [ YES, NO ]
* `maxbandwidth` - (Optional) 
* `monthreshold` - (Optional) 
* `state` - (Optional) Initial state of the service group. Possible values: [ ENABLED, DISABLED ]
* `downstateflush` - (Optional) Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]
* `tcpprofilename` - (Optional) Name of the TCP profile that contains TCP configuration settings for the service group.
* `httpprofilename` - (Optional) Name of the HTTP profile that contains HTTP configuration settings for the service group.
* `comment` - (Optional) Any information about the service group.
* `appflowlog` - (Optional) Enable logging of AppFlow information for the specified service group. Possible values: [ ENABLED, DISABLED ]
* `netprofile` - (Optional) Network profile for the service group.
* `autoscale` - (Optional) Auto scale option for a servicegroup. Possible values: [ DISABLED, DNS, POLICY, CLOUD, API ]
* `memberport` - (Optional) member port.
* `autodisablegraceful` - (Optional) Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service. Possible values: [ YES, NO ]
* `autodisabledelay` - (Optional) The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.
* `monconnectionclose` - (Optional) Close monitoring connections by sending the service a connection termination message with the specified bit set. Possible values: [ RESET, FIN ]
* `servername` - (Optional) Name of the server to which to bind the service group.
* `port` - (Optional) Server port number. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `weight` - (Optional) Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
* `customserverid` - (Optional) The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
* `serverid` - (Optional) The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `hashid` - (Optional) The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
* `nameserver` - (Optional) Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver.
* `dbsttl` - (Optional) Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `monitornamesvc` - (Optional) Name of the monitor bound to the service group. Used to assign a weight to the monitor.
* `dupweight` - (Optional) weight of the monitor that is bound to servicegroup.
* `delay` - (Optional) Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
* `graceful` - (Optional) Wait for all existing connections to the service to terminate before shutting down the service. Possible values: [ YES, NO ]
* `includemembers` - (Optional) Display the members of the listed service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.
* `lbvservers` - (Optional) List of lbvservers to bind the servicegroup to.
* `servicegroupmembers_by_servername` - (Optional) list of service members bindings by service name. e.g. `["service1:80:1", "service2:80:1"]`
* `servicegroupmembers` - (Optional) list of members bindings by server ip address. e.g.`["172.20.0.9:80:10", "172.20.0.10:80:10"]
* `lbmonitor` - (Optional) lbmonitor to bind the servicegroup to.
* `riseapbrstatsmsgcode` - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup. It has the same value as the `servicegroupname` attribute.


## Import

A servicegroup can be imported using its name, e.g.

```shell
terraform import citrixadc_servicegroup.tf_servicegroup tf_servicegroup
```
