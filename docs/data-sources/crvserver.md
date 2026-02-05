---
subcategory: "Cache Redirection"
---

# Data Source `crvserver`

The crvserver data source allows you to retrieve information about cache redirection virtual servers.


## Example usage

```terraform
data "citrixadc_crvserver" "tf_crvserver" {
  name = "my_vserver"
}

output "servicetype" {
  value = data.citrixadc_crvserver.tf_crvserver.servicetype
}

output "arp" {
  value = data.citrixadc_crvserver.tf_crvserver.arp
}
```


## Argument Reference

* `name` - (Required) Name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the cache redirection virtual server is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `appflowlog` - Enable logging of AppFlow information.
* `arp` - Use ARP to determine the destination MAC address.
* `backendssl` - Decides whether the backend connection made by Citrix ADC to the origin server will be HTTP or SSL. Applicable only for SSL type CR Forward proxy vserver.
* `backupvserver` - Name of the backup virtual server to which traffic is forwarded if the active server becomes unavailable.
* `cachetype` - Mode of operation for the cache redirection virtual server. Available settings function as follows: * TRANSPARENT - Intercept all traffic flowing to the appliance and apply cache redirection policies to determine whether content should be served from the cache or from the origin server. * FORWARD - Resolve the hostname of the incoming request, by using a DNS server, and forward requests for non-cacheable content to the resolved origin servers. Cacheable requests are sent to the configured cache servers. * REVERSE - Configure reverse proxy caches for specific origin servers. Incoming traffic directed to the reverse proxy can either be served from a cache server or be sent to the origin server with or without modification to the URL.
* `cachevserver` - Name of the default cache virtual server to which to redirect requests (the default target of the cache redirection virtual server).
* `clttimeout` - Time-out value, in seconds, after which to terminate an idle client connection.
* `comment` - Comments associated with this virtual server.
* `destinationvserver` - Destination virtual server for a transparent or forward proxy cache redirection virtual server.
* `disableprimaryondown` - Continue sending traffic to a backup virtual server even after the primary virtual server comes UP from the DOWN state.
* `disallowserviceaccess` - This is effective when a FORWARD type cr vserver is added. By default, this parameter is DISABLED. When it is ENABLED, backend services cannot be accessed through a FORWARD type cr vserver.
* `dnsvservername` - Name of the DNS virtual server that resolves domain names arriving at the forward proxy virtual server. Note: This parameter applies only to forward proxy virtual servers, not reverse or transparent.
* `domain` - Default domain for reverse proxies. Domains are configured to direct an incoming request from a specified source domain to a specified target domain. There can be several configured pairs of source and target domains. You can select one pair to be the default.
* `downstateflush` - Perform delayed cleanup of connections to this virtual server.
* `format` - Format parameter.
* `ghost` - Ghost parameter.
* `httpprofilename` - Name of the profile containing HTTP configuration information for cache redirection virtual server.
* `icmpvsrresponse` - Criterion for responding to PING requests sent to this virtual server. If ACTIVE, respond only if the virtual server is available. If PASSIVE, respond even if the virtual server is not available.
* `ipset` - The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cr vserver.
* `ipv46` - IPv4 or IPv6 address of the cache redirection virtual server. Usually a public IP address. Clients send connection requests to this IP address.
* `l2conn` - Use L2 parameters, such as MAC, VLAN, and channel to identify a connection.
* `listenpolicy` - String specifying the listen policy for the cache redirection virtual server. Can be either an in-line expression or the name of a named expression.
* `listenpriority` - Priority of the listen policy specified by the Listen Policy parameter. The lower the number, higher the priority.
* `map` - Map parameter (obsolete).
* `netprofile` - Name of the network profile containing network configurations for the cache redirection virtual server.
* `newname` - New name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.
* `onpolicymatch` - Redirect requests that match the policy to either the cache or the origin server, as specified. Note: For this option to work, you must set the cache redirection type to POLICY.
* `originusip` - Use the client's IP address as the source IP address in requests sent to the origin server. Note: You can enable this parameter to implement fully transparent CR deployment.
* `port` - Port number of the virtual server.
* `precedence` - Type of policy (URL or RULE) that takes precedence on the cache redirection virtual server.
* `probeport` - Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitoring.
* `probeprotocol` - Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck.
* `probesuccessresponsecode` - HTTP code to return in SUCCESS case.
* `range` - Number of consecutive IP addresses, starting with the address specified by the IPAddress parameter, to include in a range of addresses assigned to this virtual server.
* `redirect` - Type of cache server to which to redirect HTTP requests. Available settings function as follows: * CACHE - Direct all requests to the cache. * POLICY - Apply the cache redirection policy to determine whether the request should be directed to the cache or to the origin. * ORIGIN - Direct all requests to the origin server.
* `redirecturl` - URL of the server to which to redirect traffic if the cache redirection virtual server configured on the Citrix ADC becomes unavailable.
* `reuse` - Reuse TCP connections to the origin server across client connections.
* `rhistate` - A host route is injected according to the setting on the virtual servers.
* `servicetype` - Protocol (type of service) handled by the virtual server.
* `sopersistencetimeout` - Time-out, in minutes, for spillover persistence.
* `sothreshold` - For CONNECTION (or) DYNAMICCONNECTION spillover, the number of connections above which the virtual server enters spillover mode.
* `srcipexpr` - Expression used to extract the source IP addresses from the requests originating from the cache.
* `state` - Initial state of the cache redirection virtual server.
* `tcpprobeport` - Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port.
* `tcpprofilename` - Name of the profile containing TCP configuration information for the cache redirection virtual server.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity.
* `useoriginipportforcache` - Use origin ip/port while forwarding request to the cache.
* `useportrange` - Use a port number from the port range as the source port in the requests sent to the origin server.
* `via` - Insert a via header in each HTTP request. In the case of a cache miss, the request is redirected from the cache server to the origin server.

## Attribute Reference

* `id` - The id of the crvserver. It has the same value as the `name` attribute.


## Import

A crvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver.tf_crvserver my_vserver
```
