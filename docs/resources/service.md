---
subcategory: "Basic"
---

# Resource: service

The service resource is used to create services.


## Example usage

```hcl
resource "citrixadc_service" "tf_service" {

  lbvserver = "tf_lbvserver"
  name = "tf_service"
  port = 80
  ip = "10.202.22.12"
  servicetype = "HTTP"

  depends_on = ["citrixadc_lbvserver.tf_lbvserver"]

}
```


## Argument Reference

* `name` - (Optional) Name for the service.
* `ip` - (Optional) IP to assign to the service.
* `servername` - (Optional) Name of the server that hosts the service.
* `servicetype` - (Optional) Protocol in which data is exchanged with the service. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]
* `port` - (Optional) Port number of the service. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `cleartextport` - (Optional) Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.
* `cachetype` - (Optional) Cache type supported by the cache server. Possible values: [ TRANSPARENT, REVERSE, FORWARD ]
* `maxclient` - (Optional) 
* `healthmonitor` - (Optional) Monitor the health of this service. Available settings function as follows: YES - Send probes to check the health of the service. NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times. Possible values: [ YES, NO ]
* `maxreq` - (Optional) Note: Connection requests beyond this value are rejected.
* `cacheable` - (Optional) Use the transparent cache redirection virtual server to forward requests to the cache server. Note: Do not specify this parameter if you set the Cache Type parameter. Possible values: [ YES, NO ]
* `cip` - (Optional) Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option. Possible values: [ ENABLED, DISABLED ]
* `cipheader` - (Optional) Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name "client-ip.".
* `usip` - (Optional) Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service. Possible values: [ YES, NO ]
* `pathmonitor` - (Optional) Path monitoring for clustering. Possible values: [ YES, NO ]
* `pathmonitorindv` - (Optional) Individual Path monitoring decisions. Possible values: [ YES, NO ]
* `useproxyport` - (Optional) Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES. Possible values: [ YES, NO ]
* `sc` - (Optional) State of SureConnect for the service. Possible values: [ on, off ]
* `sp` - (Optional) Enable surge protection for the service. Possible values: [ on, off ]
* `rtspsessionidremap` - (Optional) Enable RTSP session ID mapping for the service. Possible values: [ on, off ]
* `clttimeout` - (Optional) Time, in seconds, after which to terminate an idle client connection.
* `svrtimeout` - (Optional) Time, in seconds, after which to terminate an idle server connection.
* `customserverid` - (Optional) Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.
* `serverid` - (Optional) The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
* `cka` - (Optional) Enable client keep-alive for the service. Possible values: [ YES, NO ]
* `tcpb` - (Optional) Enable TCP buffering for the service. Possible values: [ YES, NO ]
* `cmp` - (Optional) Enable compression for the service. Possible values: [ YES, NO ]
* `maxbandwidth` - (Optional) 
* `monthreshold` - (Optional) 
* `state` - (Optional) Initial state of the service. Possible values: [ ENABLED, DISABLED ]
* `downstateflush` - (Optional) Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions. Possible values: [ ENABLED, DISABLED ]
* `tcpprofilename` - (Optional) Name of the TCP profile that contains TCP configuration settings for the service.
* `httpprofilename` - (Optional) Name of the HTTP profile that contains HTTP configuration settings for the service.
* `contentinspectionprofilename` - (Optional) Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service.
* `hashid` - (Optional) A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.
* `comment` - (Optional) Any information about the service.
* `netprofile` - (Optional) Network profile to use for the service.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `processlocal` - (Optional) By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution. Possible values: [ ENABLED, DISABLED ]
* `dnsprofilename` - (Optional) Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS and ADNS-TCP services.
* `monconnectionclose` - (Optional) Close monitoring connections by sending the service a connection termination message with the specified bit set. Possible values: [ RESET, FIN ]
* `ipaddress` - (Optional) The new IP address of the service.
* `weight` - (Optional) Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.
* `monitornamesvc` - (Optional) Name of the monitor bound to the specified service.
* `delay` - (Optional) Time, in seconds, allocated to the Citrix ADC for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
* `graceful` - (Optional) Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed. Possible values: [ YES, NO ]
* `internal` - (Optional) Display only dynamically learned services.
* `lbvserver` - (Optional) The lb vserver to attach the service to.
* `lbmonitor` - (Optional) The lb monitor to attach the service to.

* `snienable` - (Optional) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, \*.sports.net can be used to secure domains such as login.sports.net and help.sports.net. Possible values: [ ENABLED, DISABLED ]
* `commonname` - (Optional) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL service.
* `all` - (Optional)
* `riseapbrstatsmsgcode` - (Optional)
* `accessdown` - (Optional) 
* `appflowlog` - (Optional)
* `wait_until_disabled` - (Optional) Boolean flag to signify if the resource will wait for the service to be in a disabled state after the disable operation has been issued.
* `disabled_timeout` - (Optional) Time period to wait for the service to be in a disabled state after the disable operation.
* `disabled_poll_delay` - (Optional) Time period to wait before the first poll for the disabled state read.
* `disabled_poll_interval` - (Optional) Time period for disabled state read poll interval between tries.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the service. It has the same value as the `name` attribute.


## Import

A service can be imported using its name, e.g.

```shell
terraform import citrixadc_service.tf_service tf_service
```
