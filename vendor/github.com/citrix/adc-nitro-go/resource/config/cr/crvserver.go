/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package cr

/**
* Configuration for CR virtual server resource.
*/
type Crvserver struct {
	/**
	* Name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the cache redirection virtual server is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my server" or 'my server').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* Protocol (type of service) handled by the virtual server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IPv4 or IPv6 address of the cache redirection virtual server. Usually a public IP address. Clients send connection requests to this IP address.
		Note: For a transparent cache redirection virtual server, use an asterisk (*) to specify a wildcard virtual server address.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* Port number of the virtual server.
	*/
	Port int `json:"port,omitempty"`
	/**
	* The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cr vserver
	*/
	Ipset string `json:"ipset,omitempty"`
	/**
	* Number of consecutive IP addresses, starting with the address specified by the IPAddress parameter, to include in a range of addresses assigned to this virtual server.
	*/
	Range int `json:"range,omitempty"`
	/**
	* Mode of operation for the cache redirection virtual server. Available settings function as follows:
		* TRANSPARENT - Intercept all traffic flowing to the appliance and apply cache redirection policies to determine whether content should be served from the cache or from the origin server.
		* FORWARD - Resolve the hostname of the incoming request, by using a DNS server, and forward requests for non-cacheable content to the resolved origin servers. Cacheable requests are sent to the configured cache servers.
		* REVERSE - Configure reverse proxy caches for specific origin servers. Incoming traffic directed to the reverse proxy can either be served from a cache server or be sent to the origin server with or without modification to the URL.
		The default value for cache type is TRANSPARENT if service is HTTP or SSL whereas the default cache type is FORWARD if the service is HDX.
	*/
	Cachetype string `json:"cachetype,omitempty"`
	/**
	* Type of cache server to which to redirect HTTP requests. Available settings function as follows:
		* CACHE - Direct all requests to the cache.
		* POLICY - Apply the cache redirection policy to determine whether the request should be directed to the cache or to the origin.
		* ORIGIN - Direct all requests to the origin server.
	*/
	Redirect string `json:"redirect,omitempty"`
	/**
	* Redirect requests that match the policy to either the cache or the origin server, as specified.
		Note: For this option to work, you must set the cache redirection type to POLICY.
	*/
	Onpolicymatch string `json:"onpolicymatch,omitempty"`
	/**
	* URL of the server to which to redirect traffic if the cache redirection virtual server configured on the Citrix ADC becomes unavailable.
	*/
	Redirecturl string `json:"redirecturl,omitempty"`
	/**
	* Time-out value, in seconds, after which to terminate an idle client connection.
	*/
	Clttimeout int `json:"clttimeout,omitempty"`
	/**
	* Type of policy (URL or RULE) that takes precedence on the cache redirection virtual server. Applies only to cache redirection virtual servers that have both URL and RULE based policies. If you specify URL, URL based policies are applied first, in the following order:
		1.   Domain and exact URL
		2.   Domain, prefix and suffix
		3.   Domain and suffix
		4.   Domain and prefix
		5.   Domain only
		6.   Exact URL
		7.   Prefix and suffix
		8.   Suffix only
		9.   Prefix only
		10.  Default
		If you specify RULE, the rule based policies are applied before URL based policies are applied.
	*/
	Precedence string `json:"precedence,omitempty"`
	/**
	* Use ARP to determine the destination MAC address.
	*/
	Arp string `json:"arp,omitempty"`
	Ghost string `json:"ghost,omitempty"`
	/**
	* Obsolete.
	*/
	Map string `json:"map,omitempty"`
	Format string `json:"format,omitempty"`
	/**
	* Insert a via header in each HTTP request. In the case of a cache miss, the request is redirected from the cache server to the origin server. This header indicates whether the request is being sent from a cache server.
	*/
	Via string `json:"via,omitempty"`
	/**
	* Name of the default cache virtual server to which to redirect requests (the default target of the cache redirection virtual server).
	*/
	Cachevserver string `json:"cachevserver,omitempty"`
	/**
	* Name of the DNS virtual server that resolves domain names arriving at the forward proxy virtual server.
		Note: This parameter applies only to forward proxy virtual servers, not reverse or transparent.
	*/
	Dnsvservername string `json:"dnsvservername,omitempty"`
	/**
	* Destination virtual server for a transparent or forward proxy cache redirection virtual server.
	*/
	Destinationvserver string `json:"destinationvserver,omitempty"`
	/**
	* Default domain for reverse proxies. Domains are configured to direct an incoming request from a specified source domain to a specified target domain. There can be several configured pairs of source and target domains. You can select one pair to be the default. If the host header or URL of an incoming request does not include a source domain, this option sends the request to the specified target domain.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Time-out, in minutes, for spillover persistence.
	*/
	Sopersistencetimeout int `json:"sopersistencetimeout,omitempty"`
	/**
	* For CONNECTION (or) DYNAMICCONNECTION spillover, the number of connections above which the virtual server enters spillover mode. For BANDWIDTH spillover, the amount of incoming and outgoing traffic (in Kbps) before spillover. For HEALTH spillover, the percentage of active services (by weight) below which spillover occurs.
	*/
	Sothreshold int `json:"sothreshold,omitempty"`
	/**
	* Reuse TCP connections to the origin server across client connections. Do not set this parameter unless the Service Type parameter is set to HTTP. If you set this parameter to OFF, the possible settings of the Redirect parameter function as follows:
		* CACHE - TCP connections to the cache servers are not reused.
		* ORIGIN - TCP connections to the origin servers are not reused.
		* POLICY - TCP connections to the origin servers are not reused.
		If you set the Reuse parameter to ON, connections to origin servers and connections to cache servers are reused.
	*/
	Reuse string `json:"reuse,omitempty"`
	/**
	* Initial state of the cache redirection virtual server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Perform delayed cleanup of connections to this virtual server.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Name of the backup virtual server to which traffic is forwarded if the active server becomes unavailable.
	*/
	Backupvserver string `json:"backupvserver,omitempty"`
	/**
	* Continue sending traffic to a backup virtual server even after the primary virtual server comes UP from the DOWN state.
	*/
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	/**
	* Use L2 parameters, such as MAC, VLAN, and channel to identify a connection.
	*/
	L2conn string `json:"l2conn,omitempty"`
	/**
	* Decides whether the backend connection made by Citrix ADC to the origin server will be HTTP or SSL. Applicable only for SSL type CR Forward proxy vserver.
	*/
	Backendssl string `json:"backendssl,omitempty"`
	/**
	* String specifying the listen policy for the cache redirection virtual server. Can be either an in-line expression or the name of a named expression.
	*/
	Listenpolicy string `json:"listenpolicy,omitempty"`
	/**
	* Priority of the listen policy specified by the Listen Policy parameter. The lower the number, higher the priority.
	*/
	Listenpriority int `json:"listenpriority,omitempty"`
	/**
	* Name of the profile containing TCP configuration information for the cache redirection virtual server.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the profile containing HTTP configuration information for cache redirection virtual server.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Comments associated with this virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Expression used to extract the source IP addresses from the requests originating from the cache. Can be either an in-line expression or the name of a named expression.
	*/
	Srcipexpr string `json:"srcipexpr,omitempty"`
	/**
	* Use the client's IP address as the source IP address in requests sent to the origin server.
		Note: You can enable this parameter to implement fully transparent CR deployment.
	*/
	Originusip string `json:"originusip,omitempty"`
	/**
	* Use a port number from the port range (set by using the set ns param command, or in the Create Virtual Server (Cache Redirection) dialog box) as the source port in the requests sent to the origin server.
	*/
	Useportrange string `json:"useportrange,omitempty"`
	/**
	* Enable logging of AppFlow information.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Name of the network profile containing network configurations for the cache redirection virtual server.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Criterion for responding to PING requests sent to this virtual server. If ACTIVE, respond only if the virtual server is available. If PASSIVE, respond even if the virtual server is not available.
	*/
	Icmpvsrresponse string `json:"icmpvsrresponse,omitempty"`
	/**
	* A host route is injected according to the setting on the virtual servers
		* If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.
		* If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.
		* If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.
	*/
	Rhistate string `json:"rhistate,omitempty"`
	/**
	* Use origin ip/port while forwarding request to the cache. Change the destination IP, destination port of the request came to CR vserver to Origin IP and Origin Port and forward it to Cache
	*/
	Useoriginipportforcache string `json:"useoriginipportforcache,omitempty"`
	/**
	* Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.
	*/
	Tcpprobeport int `json:"tcpprobeport,omitempty"`
	/**
	* Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck
	*/
	Probeprotocol string `json:"probeprotocol,omitempty"`
	/**
	* HTTP code to return in SUCCESS case.
	*/
	Probesuccessresponsecode string `json:"probesuccessresponsecode,omitempty"`
	/**
	* Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitring
	*/
	Probeport int `json:"probeport,omitempty"`
	/**
	* This is effective when a FORWARD type cr vserver is added. By default, this parameter is DISABLED. When it is ENABLED, backend services cannot be accessed through a FORWARD type cr vserver.
	*/
	Disallowserviceaccess string `json:"disallowserviceaccess,omitempty"`
	/**
	* New name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my name" or 'my name').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Value string `json:"value,omitempty"`
	Ngname string `json:"ngname,omitempty"`
	Type string `json:"type,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Status string `json:"status,omitempty"`
	Authentication string `json:"authentication,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Rule string `json:"rule,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Pipolicyhits string `json:"pipolicyhits,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight string `json:"weight,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
	Priority string `json:"priority,omitempty"`
	Somethod string `json:"somethod,omitempty"`
	Sopersistence string `json:"sopersistence,omitempty"`
	Lbvserver string `json:"lbvserver,omitempty"`
	Bindpoint string `json:"bindpoint,omitempty"`
	Invoke string `json:"invoke,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
