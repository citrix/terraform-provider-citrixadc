---
subcategory: "Cache Redirection"
---

# Resource: crvserver

The crvserver resource is used to create Cache Redirection Vserver.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
```


## Argument Reference

* `name` - (Required) Name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the cache redirection virtual server is created. The following requirement applies only to the Citrix ADC CLI:   If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my server" or 'my server').
* `servicetype` - (Required) Protocol (type of service) handled by the virtual server.
* `appflowlog` - (Optional) Enable logging of AppFlow information.
* `arp` - (Optional) Use ARP to determine the destination MAC address.
* `backendssl` - (Optional) Decides whether the backend connection made by Citrix ADC to the origin server will be HTTP or SSL. Applicable only for SSL type CR Forward proxy vserver.
* `backupvserver` - (Optional) Name of the backup virtual server to which traffic is forwarded if the active server becomes unavailable.
* `cachetype` - (Optional) Mode of operation for the cache redirection virtual server. Available settings function as follows: * TRANSPARENT - Intercept all traffic flowing to the appliance and apply cache redirection policies to determine whether content should be served from the cache or from the origin server. * FORWARD - Resolve the hostname of the incoming request, by using a DNS server, and forward requests for non-cacheable content to the resolved origin servers. Cacheable requests are sent to the configured cache servers. * REVERSE - Configure reverse proxy caches for specific origin servers. Incoming traffic directed to the reverse proxy can either be served from a cache server or be sent to the origin server with or without modification to the URL. The default value for cache type is TRANSPARENT if service is HTTP or SSL whereas the default cache type is FORWARD if the service is HDX.
* `cachevserver` - (Optional) Name of the default cache virtual server to which to redirect requests (the default target of the cache redirection virtual server).
* `clttimeout` - (Optional) Time-out value, in seconds, after which to terminate an idle client connection.
* `comment` - (Optional) Comments associated with this virtual server.
* `destinationvserver` - (Optional) Destination virtual server for a transparent or forward proxy cache redirection virtual server.
* `disableprimaryondown` - (Optional) Continue sending traffic to a backup virtual server even after the primary virtual server comes UP from the DOWN state.
* `dnsvservername` - (Optional) Name of the DNS virtual server that resolves domain names arriving at the forward proxy virtual server. Note: This parameter applies only to forward proxy virtual servers, not reverse or transparent.
* `domain` - (Optional) Default domain for reverse proxies. Domains are configured to direct an incoming request from a specified source domain to a specified target domain. There can be several configured pairs of source and target domains. You can select one pair to be the default. If the host header or URL of an incoming request does not include a source domain, this option sends the request to the specified target domain.
* `downstateflush` - (Optional) Perform delayed cleanup of connections to this virtual server.
* `format` - (Optional) 0
* `ghost` - (Optional) 0
* `httpprofilename` - (Optional) Name of the profile containing HTTP configuration information for cache redirection virtual server.
* `icmpvsrresponse` - (Optional) Criterion for responding to PING requests sent to this virtual server. If ACTIVE, respond only if the virtual server is available. If PASSIVE, respond even if the virtual server is not available.
* `ipset` - (Optional) The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cr vserver
* `ipv46` - (Optional) IPv4 or IPv6 address of the cache redirection virtual server. Usually a public IP address. Clients send connection requests to this IP address. Note: For a transparent cache redirection virtual server, use an asterisk (*) to specify a wildcard virtual server address.
* `l2conn` - (Optional) Use L2 parameters, such as MAC, VLAN, and channel to identify a connection.
* `listenpolicy` - (Optional) String specifying the listen policy for the cache redirection virtual server. Can be either an in-line expression or the name of a named expression.
* `listenpriority` - (Optional) Priority of the listen policy specified by the Listen Policy parameter. The lower the number, higher the priority.
* `map` - (Optional) Obsolete.
* `netprofile` - (Optional) Name of the network profile containing network configurations for the cache redirection virtual server.
* `onpolicymatch` - (Optional) Redirect requests that match the policy to either the cache or the origin server, as specified. Note: For this option to work, you must set the cache redirection type to POLICY.
* `originusip` - (Optional) Use the client's IP address as the source IP address in requests sent to the origin server.   Note: You can enable this parameter to implement fully transparent CR deployment.
* `port` - (Optional) Port number of the virtual server.
* `precedence` - (Optional) Type of policy (URL or RULE) that takes precedence on the cache redirection virtual server. Applies only to cache redirection virtual servers that have both URL and RULE based policies. If you specify URL, URL based policies are applied first, in the following order: 1.   Domain and exact URL 2.   Domain, prefix and suffix 3.   Domain and suffix 4.   Domain and prefix 5.   Domain only 6.   Exact URL 7.   Prefix and suffix 8.   Suffix only 9.   Prefix only 10.  Default If you specify RULE, the rule based policies are applied before URL based policies are applied.
* `probeport` - (Optional) Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitring
* `probeprotocol` - (Optional) Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck
* `probesuccessresponsecode` - (Optional) HTTP code to return in SUCCESS case.
* `range` - (Optional) Number of consecutive IP addresses, starting with the address specified by the IPAddress parameter, to include in a range of addresses assigned to this virtual server.
* `redirect` - (Optional) Type of cache server to which to redirect HTTP requests. Available settings function as follows: * CACHE - Direct all requests to the cache. * POLICY - Apply the cache redirection policy to determine whether the request should be directed to the cache or to the origin. * ORIGIN - Direct all requests to the origin server.
* `redirecturl` - (Optional) URL of the server to which to redirect traffic if the cache redirection virtual server configured on the Citrix ADC becomes unavailable.
* `reuse` - (Optional) Reuse TCP connections to the origin server across client connections. Do not set this parameter unless the Service Type parameter is set to HTTP. If you set this parameter to OFF, the possible settings of the Redirect parameter function as follows: * CACHE - TCP connections to the cache servers are not reused. * ORIGIN - TCP connections to the origin servers are not reused.  * POLICY - TCP connections to the origin servers are not reused. If you set the Reuse parameter to ON, connections to origin servers and connections to cache servers are reused.
* `rhistate` - (Optional) A host route is injected according to the setting on the virtual servers             * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.             * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.             * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.
* `sopersistencetimeout` - (Optional) Time-out, in minutes, for spillover persistence.
* `sothreshold` - (Optional) For CONNECTION (or) DYNAMICCONNECTION spillover, the number of connections above which the virtual server enters spillover mode. For BANDWIDTH spillover, the amount of incoming and outgoing traffic (in Kbps) before spillover. For HEALTH spillover, the percentage of active services (by weight) below which spillover occurs.
* `srcipexpr` - (Optional) Expression used to extract the source IP addresses from the requests originating from the cache. Can be either an in-line expression or the name of a named expression.
* `state` - (Optional) Initial state of the cache redirection virtual server.
* `tcpprobeport` - (Optional) Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.
* `tcpprofilename` - (Optional) Name of the profile containing TCP configuration information for the cache redirection virtual server.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `useoriginipportforcache` - (Optional) Use origin ip/port while forwarding request to the cache. Change the destination IP, destination port of the request came to CR vserver to Origin IP and Origin Port and forward it to Cache
* `useportrange` - (Optional) Use a port number from the port range (set by using the set ns param command, or in the Create Virtual Server (Cache Redirection) dialog box) as the source port in the requests sent to the origin server.
* `via` - (Optional) Insert a via header in each HTTP request. In the case of a cache miss, the request is redirected from the cache server to the origin server. This header indicates whether the request is being sent from a cache server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver. It has the same value as the `name` attribute.


## Import

A crvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver.crvserver my_vserver
```
