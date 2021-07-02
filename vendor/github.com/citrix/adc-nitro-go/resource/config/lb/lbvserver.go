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

package lb

/**
* Configuration for Load Balancing Virtual Server resource.
*/
type Lbvserver struct {
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol used by the service (also called the service type).
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IPv4 or IPv6 address to assign to the virtual server.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern.  Mutually exclusive with the IP Address parameter. 
		For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254.  You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).
		If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if virtual servers vs1 and vs2 have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.
	*/
	Ippattern string `json:"ippattern,omitempty"`
	/**
	* IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.
	*/
	Ipmask string `json:"ipmask,omitempty"`
	/**
	* Port number for the virtual server.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current lb vserver
	*/
	Ipset string `json:"ipset,omitempty"`
	/**
	* Number of IP addresses that the appliance must generate and assign to the virtual server. The virtual server then functions as a network virtual server, accepting traffic on any of the generated IP addresses. The IP addresses are generated automatically, as follows: 
		* For a range of n, the last octet of the address specified by the IP Address parameter increments n-1 times. 
		* If the last octet exceeds 255, it rolls over to 0 and the third octet increments by 1.
		Note: The Range parameter assigns multiple IP addresses to one virtual server. To generate an array of virtual servers, each of which owns only one IP address, use brackets in the IP Address and Name parameters to specify the range. For example:
		add lb vserver my_vserver[1-3] HTTP 192.0.2.[1-3] 80
	*/
	Range uint32 `json:"range,omitempty"`
	/**
	* Type of persistence for the virtual server. Available settings function as follows:
		* SOURCEIP - Connections from the same client IP address belong to the same persistence session.
		* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session. 
		* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.
		* CUSTOMSERVERID - Connections with the same server ID form part of the same session. For this persistence type, set the Server ID (CustomServerID) parameter for each service and configure the Rule parameter to identify the server ID in a request.
		* RULE - All connections that match a user defined rule belong to the same persistence session. 
		* URLPASSIVE - Requests that have the same server ID in the URL query belong to the same persistence session. The server ID is the hexadecimal representation of the IP address and port of the service to which the request must be forwarded. This persistence type requires a rule to identify the server ID in the request. 
		* DESTIP - Connections to the same destination IP address belong to the same persistence session.
		* SRCIPDESTIP - Connections that have the same source IP address and destination IP address belong to the same persistence session.
		* CALLID - Connections that have the same CALL-ID SIP header belong to the same persistence session.
		* RTSPSID - Connections that have the same RTSP Session ID belong to the same persistence session.
		* FIXSESSION - Connections that have the same SenderCompID and TargetCompID values belong to the same persistence session.
		* USERSESSION - Persistence session is created based on the persistence parameter value provided from an extension.
	*/
	Persistencetype string `json:"persistencetype,omitempty"`
	/**
	* Time period for which a persistence session is in effect.
	*/
	Timeout uint32 `json:"timeout,omitempty"`
	/**
	* Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.
	*/
	Persistencebackup string `json:"persistencebackup,omitempty"`
	/**
	* Time period for which backup persistence is in effect.
	*/
	Backuppersistencetimeout uint32 `json:"backuppersistencetimeout,omitempty"`
	/**
	* Load balancing method.  The available settings function as follows:
		* ROUNDROBIN - Distribute requests in rotation, regardless of the load. Weights can be assigned to services to enforce weighted round robin distribution.
		* LEASTCONNECTION (default) - Select the service with the fewest connections. 
		* LEASTRESPONSETIME - Select the service with the lowest average response time. 
		* LEASTBANDWIDTH - Select the service currently handling the least traffic.
		* LEASTPACKETS - Select the service currently serving the lowest number of packets per second.
		* CUSTOMLOAD - Base service selection on the SNMP metrics obtained by custom load monitors.
		* LRTM - Select the service with the lowest response time. Response times are learned through monitoring probes. This method also takes the number of active connections into account.
		Also available are a number of hashing methods, in which the appliance extracts a predetermined portion of the request, creates a hash of the portion, and then checks whether any previous requests had the same hash value. If it finds a match, it forwards the request to the service that served those previous requests. Following are the hashing methods: 
		* URLHASH - Create a hash of the request URL (or part of the URL).
		* DOMAINHASH - Create a hash of the domain name in the request (or part of the domain name). The domain name is taken from either the URL or the Host header. If the domain name appears in both locations, the URL is preferred. If the request does not contain a domain name, the load balancing method defaults to LEASTCONNECTION.
		* DESTINATIONIPHASH - Create a hash of the destination IP address in the IP header. 
		* SOURCEIPHASH - Create a hash of the source IP address in the IP header.  
		* TOKEN - Extract a token from the request, create a hash of the token, and then select the service to which any previous requests with the same token hash value were sent. 
		* SRCIPDESTIPHASH - Create a hash of the string obtained by concatenating the source IP address and destination IP address in the IP header.  
		* SRCIPSRCPORTHASH - Create a hash of the source IP address and source port in the IP header.  
		* CALLIDHASH - Create a hash of the SIP Call-ID header.
		* USER_TOKEN - Same as TOKEN LB method but token needs to be provided from an extension.
	*/
	Lbmethod string `json:"lbmethod,omitempty"`
	/**
	* Number of bytes to consider for the hash value used in the URLHASH and DOMAINHASH load balancing methods.
	*/
	Hashlength uint32 `json:"hashlength,omitempty"`
	/**
	* IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.
	*/
	V6netmasklen uint32 `json:"v6netmasklen,omitempty"`
	/**
	* Backup load balancing method. Becomes operational if the primary load balancing me
		thod fails or cannot be used.
		Valid only if the primary method is based on static proximity.
	*/
	Backuplbmethod string `json:"backuplbmethod,omitempty"`
	/**
	* Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
	*/
	Cookiename string `json:"cookiename,omitempty"`
	/**
	* Expression, or name of a named expression, against which traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character. 
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Expression identifying traffic accepted by the virtual server. Can be either an expression (for example, CLIENT.IP.DST.IN_SUBNET(192.0.2.0/24) or the name of a named expression. In the above example, the virtual server accepts all requests whose destination IP address is in the 192.0.2.0/24 subnet.
	*/
	Listenpolicy string `json:"listenpolicy,omitempty"`
	/**
	* Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
	*/
	Listenpriority uint32 `json:"listenpriority,omitempty"`
	/**
	* Expression specifying which part of a server's response to use for creating rule based persistence sessions (persistence type RULE). Can be either an expression or the name of a named expression.
		Example:
		HTTP.RES.HEADER("setcookie").VALUE(0).TYPECAST_NVLIST_T('=',';').VALUE("server1").
	*/
	Resrule string `json:"resrule,omitempty"`
	/**
	* Persistence mask for IP based persistence types, for IPv4 virtual servers.
	*/
	Persistmask string `json:"persistmask,omitempty"`
	/**
	* Persistence mask for IP based persistence types, for IPv6 virtual servers.
	*/
	V6persistmasklen uint32 `json:"v6persistmasklen,omitempty"`
	/**
	* Use priority queuing on the virtual server. based persistence types, for IPv6 virtual servers.
	*/
	Pq string `json:"pq,omitempty"`
	/**
	* Use SureConnect on the virtual server.
	*/
	Sc string `json:"sc,omitempty"`
	/**
	* Use network address translation (NAT) for RTSP data connections.
	*/
	Rtspnat string `json:"rtspnat,omitempty"`
	/**
	* Redirection mode for load balancing. Available settings function as follows:
		* IP - Before forwarding a request to a server, change the destination IP address to the server's IP address. 
		* MAC - Before forwarding a request to a server, change the destination MAC address to the server's MAC address.  The destination IP address is not changed. MAC-based redirection mode is used mostly in firewall load balancing deployments. 
		* IPTUNNEL - Perform IP-in-IP encapsulation for client IP packets. In the outer IP headers, set the destination IP address to the IP address of the server and the source IP address to the subnet IP (SNIP). The client IP packets are not modified. Applicable to both IPv4 and IPv6 packets. 
		* TOS - Encode the virtual server's TOS ID in the TOS field of the IP header. 
		You can use either the IPTUNNEL or the TOS option to implement Direct Server Return (DSR).
	*/
	M string `json:"m,omitempty"`
	/**
	* TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.
	*/
	Tosid uint32 `json:"tosid,omitempty"`
	/**
	* Length of the token to be extracted from the data segment of an incoming packet, for use in the token method of load balancing. The length of the token, specified in bytes, must not be greater than 24 KB. Applicable to virtual servers of type TCP.
	*/
	Datalength uint32 `json:"datalength,omitempty"`
	/**
	* Offset to be considered when extracting a token from the TCP payload. Applicable to virtual servers, of type TCP, using the token method of load balancing. Must be within the first 24 KB of the TCP payload.
	*/
	Dataoffset uint32 `json:"dataoffset,omitempty"`
	/**
	* Perform load balancing on a per-packet basis, without establishing sessions. Recommended for load balancing of intrusion detection system (IDS) servers and scenarios involving direct server return (DSR), where session information is unnecessary.
	*/
	Sessionless string `json:"sessionless,omitempty"`
	/**
	* When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored.
	*/
	Trofspersistence string `json:"trofspersistence,omitempty"`
	/**
	* State of the load balancing virtual server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Mode in which the connection failover feature must operate for the virtual server. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary appliance. Clients remain connected to the same servers. Available settings function as follows:
		* STATEFUL - The primary appliance shares state information with the secondary appliance, in real time, resulting in some runtime processing overhead. 
		* STATELESS - State information is not shared, and the new primary appliance tries to re-create the packet flow on the basis of the information contained in the packets it receives. 
		* DISABLED - Connection failover does not occur.
	*/
	Connfailover string `json:"connfailover,omitempty"`
	/**
	* URL to which to redirect traffic if the virtual server becomes unavailable. 
		WARNING! Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.
	*/
	Redirurl string `json:"redirurl,omitempty"`
	/**
	* Route cacheable requests to a cache redirection virtual server. The load balancing virtual server can forward requests only to a transparent cache redirection virtual server that has an IP address and port combination of *:80, so such a cache redirection virtual server must be configured on the appliance.
	*/
	Cacheable string `json:"cacheable,omitempty"`
	/**
	* Idle time, in seconds, after which a client connection is terminated.
	*/
	Clttimeout uint64 `json:"clttimeout,omitempty"`
	/**
	* Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:
		* CONNECTION - Spillover occurs when the number of client connections exceeds the threshold.
		* DYNAMICCONNECTION - Spillover occurs when the number of client connections at the virtual server exceeds the sum of the maximum client (Max Clients) settings for bound services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of bound services.
		* BANDWIDTH - Spillover occurs when the bandwidth consumed by the virtual server's incoming and outgoing traffic exceeds the threshold. 
		* HEALTH - Spillover occurs when the percentage of weights of the services that are UP drops below the threshold. For example, if services svc1, svc2, and svc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if svc1 and svc3 or svc2 and svc3 transition to DOWN. 
		* NONE - Spillover does not occur.
	*/
	Somethod string `json:"somethod,omitempty"`
	/**
	* If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers.
	*/
	Sopersistence string `json:"sopersistence,omitempty"`
	/**
	* Timeout for spillover persistence, in minutes.
	*/
	Sopersistencetimeout uint32 `json:"sopersistencetimeout,omitempty"`
	/**
	* Threshold in percent of active services below which vserver state is made down. If this threshold is 0, vserver state will be up even if one bound service is up.
	*/
	Healththreshold uint32 `json:"healththreshold,omitempty"`
	/**
	* Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).
	*/
	Sothreshold uint32 `json:"sothreshold,omitempty"`
	/**
	* Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists
	*/
	Sobackupaction string `json:"sobackupaction,omitempty"`
	/**
	* Rewrite the port and change the protocol to ensure successful HTTP redirects from services.
	*/
	Redirectportrewrite string `json:"redirectportrewrite,omitempty"`
	/**
	* Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Name of the backup virtual server to which to forward requests if the primary virtual server goes DOWN or reaches its spillover threshold.
	*/
	Backupvserver string `json:"backupvserver,omitempty"`
	/**
	* If the primary virtual server goes down, do not allow it to return to primary status until manually enabled.
	*/
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	/**
	* Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server. The format of the header is <vipHeader>: <virtual server IP address>_<port number >, where vipHeader is the name that you specify for the header. If the virtual server has an IPv6 address, the address in the header is enclosed in brackets ([ and ]) to separate it from the port number. If you have mapped an IPv4 address to a virtual server's IPv6 address, the value of this parameter determines which IP address is inserted in the header, as follows:
		* VIPADDR - Insert the IP address of the virtual server in the HTTP header regardless of whether the virtual server has an IPv4 address or an IPv6 address. A mapped IPv4 address, if configured, is ignored.
		* V6TOV4MAPPING - Insert the IPv4 address that is mapped to the virtual server's IPv6 address. If a mapped IPv4 address is not configured, insert the IPv6 address.
		* OFF - Disable header insertion.
	*/
	Insertvserveripport string `json:"insertvserveripport,omitempty"`
	/**
	* Name for the inserted header. The default name is vip-header.
	*/
	Vipheader string `json:"vipheader,omitempty"`
	/**
	* Fully qualified domain name (FQDN) of the authentication virtual server to which the user must be redirected for authentication. Make sure that the Authentication parameter is set to ENABLED.
	*/
	Authenticationhost string `json:"authenticationhost,omitempty"`
	/**
	* Enable or disable user authentication.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* Enable or disable user authentication with HTTP 401 responses.
	*/
	Authn401 string `json:"authn401,omitempty"`
	/**
	* Name of an authentication virtual server with which to authenticate users.
	*/
	Authnvsname string `json:"authnvsname,omitempty"`
	/**
	* Process traffic with the push virtual server that is bound to this load balancing virtual server.
	*/
	Push string `json:"push,omitempty"`
	/**
	* Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the load balancing virtual server that you are configuring.
	*/
	Pushvserver string `json:"pushvserver,omitempty"`
	/**
	* Expression for extracting a label from the server's response. Can be either an expression or the name of a named expression.
	*/
	Pushlabel string `json:"pushlabel,omitempty"`
	/**
	* Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.
	*/
	Pushmulticlients string `json:"pushmulticlients,omitempty"`
	/**
	* Name of the TCP profile whose settings are to be applied to the virtual server.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the HTTP profile whose settings are to be applied to the virtual server.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Name of the DB profile whose settings are to be applied to the virtual server.
	*/
	Dbprofilename string `json:"dbprofilename,omitempty"`
	/**
	* Any comments that you might want to associate with the virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to co-exist on the Citrix ADC.
	*/
	L2conn string `json:"l2conn,omitempty"`
	/**
	* Oracle server version
	*/
	Oracleserverversion string `json:"oracleserverversion,omitempty"`
	/**
	* For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version. Set this parameter if you expect some clients to run a version different from the version of the database. This setting provides compatibility between the client-side and server-side connections by ensuring that all communication conforms to the server's version.
	*/
	Mssqlserverversion string `json:"mssqlserverversion,omitempty"`
	/**
	* MySQL protocol version that the virtual server advertises to clients.
	*/
	Mysqlprotocolversion uint32 `json:"mysqlprotocolversion,omitempty"`
	/**
	* MySQL server version string that the virtual server advertises to clients.
	*/
	Mysqlserverversion string `json:"mysqlserverversion,omitempty"`
	/**
	* Character set that the virtual server advertises to clients.
	*/
	Mysqlcharacterset uint32 `json:"mysqlcharacterset,omitempty"`
	/**
	* Server capabilities that the virtual server advertises to clients.
	*/
	Mysqlservercapabilities uint32 `json:"mysqlservercapabilities,omitempty"`
	/**
	* Apply AppFlow logging to the virtual server.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Name of the network profile to associate with the virtual server. If you set this parameter, the virtual server uses only the IP addresses in the network profile as source IP addresses when initiating connections with servers.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* How the Citrix ADC responds to ping requests received for an IP address that is common to one or more virtual servers. Available settings function as follows:
		* If set to PASSIVE on all the virtual servers that share the IP address, the appliance always responds to the ping requests.
		* If set to ACTIVE on all the virtual servers that share the IP address, the appliance responds to the ping requests if at least one of the virtual servers is UP. Otherwise, the appliance does not respond.
		* If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance responds if at least one virtual server with the ACTIVE setting is UP. Otherwise, the appliance does not respond.
		Note: This parameter is available at the virtual server level. A similar parameter, ICMP Response, is available at the IP address level, for IPv4 addresses of type VIP. To set that parameter, use the add ip command in the CLI or the Create IP dialog box in the GUI.
	*/
	Icmpvsrresponse string `json:"icmpvsrresponse,omitempty"`
	/**
	* Route Health Injection (RHI) functionality of the NetSaler appliance for advertising the route of the VIP address associated with the virtual server. When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:
		* If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.
		* If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.
		* If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.
	*/
	Rhistate string `json:"rhistate,omitempty"`
	/**
	* Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode. A non-zero value indicates that slow-start is applicable. A zero value indicates that the global RR startup parameter is applied. Changing the value to zero will cause services currently in slow start to take the full traffic as determined by the LB method. Subsequently, any new services added will use the global RR factor.
	*/
	Newservicerequest uint32 `json:"newservicerequest,omitempty"`
	/**
	* Units in which to increment load at each interval in slow-start mode.
	*/
	Newservicerequestunit string `json:"newservicerequestunit,omitempty"`
	/**
	* Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP. A value of 0 (zero) specifies manual slow start.
	*/
	Newservicerequestincrementinterval uint32 `json:"newservicerequestincrementinterval,omitempty"`
	/**
	* Minimum number of members expected to be present when vserver is used in Autoscale.
	*/
	Minautoscalemembers uint32 `json:"minautoscalemembers,omitempty"`
	/**
	* Maximum number of members expected to be present when vserver is used in Autoscale.
	*/
	Maxautoscalemembers uint32 `json:"maxautoscalemembers,omitempty"`
	/**
	* Persist AVP number for Diameter Persistency. 
		In case this AVP is not defined in Base RFC 3588 and it is nested inside a Grouped AVP, 
		define a sequence of AVP numbers (max 3) in order of parent to child. So say persist AVP number X 
		is nested inside AVP Y which is nested in Z, then define the list as  Z Y X
	*/
	Persistavpno []uint32 `json:"persistavpno,omitempty"`
	/**
	* This argument decides the behavior incase the service which is selected from an existing persistence session has reached threshold.
	*/
	Skippersistency string `json:"skippersistency,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Name of the authentication profile to be used when authentication is turned on.
	*/
	Authnprofile string `json:"authnprofile,omitempty"`
	/**
	* This option is used to retain vlan information of incoming packet when macmode is enabled
	*/
	Macmoderetainvlan string `json:"macmoderetainvlan,omitempty"`
	/**
	* Enable database specific load balancing for MySQL and MSSQL service types.
	*/
	Dbslb string `json:"dbslb,omitempty"`
	/**
	* This argument is for enabling/disabling the dns64 on lbvserver
	*/
	Dns64 string `json:"dns64,omitempty"`
	/**
	* If this option is enabled while resolving DNS64 query AAAA queries are not sent to back end dns server
	*/
	Bypassaaaa string `json:"bypassaaaa,omitempty"`
	/**
	* When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on. Typically one would set this option to YES, when the vserver is load balancing a set of DNS servers thatsupport recursive queries.
	*/
	Recursionavailable string `json:"recursionavailable,omitempty"`
	/**
	* By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.
	*/
	Processlocal string `json:"processlocal,omitempty"`
	/**
	* Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`
	/**
	* Name of the LB profile which is associated to the vserver
	*/
	Lbprofilename string `json:"lbprofilename,omitempty"`
	/**
	* Port number for the virtual server, from which we absorb the traffic for http redirect
	*/
	Redirectfromport int32 `json:"redirectfromport,omitempty"`
	/**
	* URL to which all HTTP traffic received on the port specified in the -redirectFromPort parameter is redirected.
	*/
	Httpsredirecturl string `json:"httpsredirecturl,omitempty"`
	/**
	* This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.
	*/
	Retainconnectionsoncluster string `json:"retainconnectionsoncluster,omitempty"`
	/**
	* Name of the adfsProxy profile to be used to support ADFSPIP protocol for ADFS servers.
	*/
	Adfsproxyprofile string `json:"adfsproxyprofile,omitempty"`
	/**
	* Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.
	*/
	Tcpprobeport int32 `json:"tcpprobeport,omitempty"`
	/**
	* Name of QUIC profile which will be attached to the VServer.
	*/
	Quicprofilename string `json:"quicprofilename,omitempty"`
	/**
	* Name of the QUIC Bridge profile whose settings are to be applied to the virtual server.
	*/
	Quicbridgeprofilename string `json:"quicbridgeprofilename,omitempty"`
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
	Probeport int32 `json:"probeport,omitempty"`
	/**
	* Weight to assign to the specified service.
	*/
	Weight uint32 `json:"weight,omitempty"`
	/**
	* Service to bind to the virtual server.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* The redirect URL to be unset.
	*/
	Redirurlflags bool `json:"redirurlflags,omitempty"`
	/**
	* New name for the virtual server.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Value string `json:"value,omitempty"`
	Ipmapping string `json:"ipmapping,omitempty"`
	Ngname string `json:"ngname,omitempty"`
	Type string `json:"type,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Effectivestate string `json:"effectivestate,omitempty"`
	Status string `json:"status,omitempty"`
	Lbrrreason string `json:"lbrrreason,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Precedence string `json:"precedence,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Dnsvservername string `json:"dnsvservername,omitempty"`
	Domain string `json:"domain,omitempty"`
	Cachevserver string `json:"cachevserver,omitempty"`
	Health string `json:"health,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	Groupname string `json:"groupname,omitempty"`
	Cookiedomain string `json:"cookiedomain,omitempty"`
	Map string `json:"map,omitempty"`
	Gt2gb string `json:"gt2gb,omitempty"`
	Consolidatedlconn string `json:"consolidatedlconn,omitempty"`
	Consolidatedlconngbl string `json:"consolidatedlconngbl,omitempty"`
	Thresholdvalue string `json:"thresholdvalue,omitempty"`
	Bindpoint string `json:"bindpoint,omitempty"`
	Version string `json:"version,omitempty"`
	Totalservices string `json:"totalservices,omitempty"`
	Activeservices string `json:"activeservices,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Statechangetimeseconds string `json:"statechangetimeseconds,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Isgslb string `json:"isgslb,omitempty"`
	Vsvrdynconnsothreshold string `json:"vsvrdynconnsothreshold,omitempty"`
	Backupvserverstatus string `json:"backupvserverstatus,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`

}
