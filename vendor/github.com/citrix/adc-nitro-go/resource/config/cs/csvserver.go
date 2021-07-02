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

package cs

/**
* Configuration for CS virtual server resource.
*/
type Csvserver struct {
	/**
	* Name for the content switching virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. 
		Cannot be changed after the CS virtual server is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my server or my server).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Protocol used by the virtual server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IP address of the content switching virtual server.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* Virtual server target type.
	*/
	Targettype string `json:"targettype,omitempty"`
	Dnsrecordtype string `json:"dnsrecordtype,omitempty"`
	Persistenceid uint32 `json:"persistenceid,omitempty"`
	/**
	* IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter. 
		For example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).
		If a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if the virtual servers, vs1 and vs2, have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.
	*/
	Ippattern string `json:"ippattern,omitempty"`
	/**
	* IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.
	*/
	Ipmask string `json:"ipmask,omitempty"`
	/**
	* Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.
	*/
	Range uint32 `json:"range,omitempty"`
	/**
	* Port number for content switching virtual server.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver
	*/
	Ipset string `json:"ipset,omitempty"`
	/**
	* Initial state of the load balancing virtual server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Enable state updates for a specific content switching virtual server. By default, the Content Switching virtual server is always UP, regardless of the state of the Load Balancing virtual servers bound to it. This parameter interacts with the global setting as follows:
		Global Level | Vserver Level | Result
		ENABLED      ENABLED        ENABLED
		ENABLED      DISABLED       ENABLED
		DISABLED     ENABLED        ENABLED
		DISABLED     DISABLED       DISABLED
		If you want to enable state updates for only some content switching virtual servers, be sure to disable the state update parameter.
	*/
	Stateupdate string `json:"stateupdate,omitempty"`
	/**
	* Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers.
	*/
	Cacheable string `json:"cacheable,omitempty"`
	/**
	* URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL.
		Caution: Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.
	*/
	Redirecturl string `json:"redirecturl,omitempty"`
	/**
	* Idle time, in seconds, after which the client connection is terminated. The default values are:
		180 seconds for HTTP/SSL-based services.
		9000 seconds for other TCP-based services.
		120 seconds for DNS-based services.
		120 seconds for other UDP-based services.
	*/
	Clttimeout uint64 `json:"clttimeout,omitempty"`
	/**
	* Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server. With the default (RULE) setting, incoming requests are evaluated against the rule-based content switching policies. If none of the rules match, the URL in the request is evaluated against the URL-based content switching policies.
	*/
	Precedence string `json:"precedence,omitempty"`
	/**
	* Consider case in URLs (for policies that use URLs instead of RULES). For example, with the ON setting, the URLs /a/1.html and /A/1.HTML are treated differently and can have different targets (set by content switching policies). With the OFF setting, /a/1.html and /A/1.HTML are switched to the same target.
	*/
	Casesensitive string `json:"casesensitive,omitempty"`
	/**
	* Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold. Connection spillover is based on the number of connections. Bandwidth spillover is based on the total Kbps of incoming and outgoing traffic.
	*/
	Somethod string `json:"somethod,omitempty"`
	/**
	* Maintain source-IP based persistence on primary and backup virtual servers.
	*/
	Sopersistence string `json:"sopersistence,omitempty"`
	/**
	* Time-out value, in minutes, for spillover persistence.
	*/
	Sopersistencetimeout uint32 `json:"sopersistencetimeout,omitempty"`
	/**
	* Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.
	*/
	Sothreshold uint32 `json:"sothreshold,omitempty"`
	/**
	* Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists
	*/
	Sobackupaction string `json:"sobackupaction,omitempty"`
	/**
	* State of port rewrite while performing HTTP redirect.
	*/
	Redirectportrewrite string `json:"redirectportrewrite,omitempty"`
	/**
	* Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Name of the backup virtual server that you are configuring. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the backup virtual server is created. You can assign a different backup virtual server or rename the existing virtual server.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks.
	*/
	Backupvserver string `json:"backupvserver,omitempty"`
	/**
	* Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state.
	*/
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	/**
	* Insert the virtual server's VIP address and port number in the request header. Available values function as follows:
		VIPADDR - Header contains the vserver's IP address and port number without any translation.
		OFF     - The virtual IP and port header insertion option is disabled.
		V6TOV4MAPPING - Header contains the mapped IPv4 address corresponding to the IPv6 address of the vserver and the port number. An IPv6 address can be mapped to a user-specified IPv4 address using the set ns ip6 command.
	*/
	Insertvserveripport string `json:"insertvserveripport,omitempty"`
	/**
	* Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.
	*/
	Vipheader string `json:"vipheader,omitempty"`
	/**
	* Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections.
	*/
	Rtspnat string `json:"rtspnat,omitempty"`
	/**
	* FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.
	*/
	Authenticationhost string `json:"authenticationhost,omitempty"`
	/**
	* Authenticate users who request a connection to the content switching virtual server.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.
	*/
	Listenpolicy string `json:"listenpolicy,omitempty"`
	/**
	* Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
	*/
	Listenpriority uint32 `json:"listenpriority,omitempty"`
	/**
	* Enable HTTP 401-response based authentication.
	*/
	Authn401 string `json:"authn401,omitempty"`
	/**
	* Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server. 
	*/
	Authnvsname string `json:"authnvsname,omitempty"`
	/**
	* Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL.
	*/
	Push string `json:"push,omitempty"`
	/**
	* Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.
	*/
	Pushvserver string `json:"pushvserver,omitempty"`
	/**
	* Expression for extracting the label from the response received from server. This string can be either an existing rule name or an inline expression. The service type of the virtual server should be either HTTP or SSL.
	*/
	Pushlabel string `json:"pushlabel,omitempty"`
	/**
	* Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.
	*/
	Pushmulticlients string `json:"pushmulticlients,omitempty"`
	/**
	* Name of the TCP profile containing TCP configuration settings for the virtual server.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the HTTP profile containing HTTP configuration settings for the virtual server. The service type of the virtual server should be either HTTP or SSL.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Name of the DB profile.
	*/
	Dbprofilename string `json:"dbprofilename,omitempty"`
	/**
	* Oracle server version
	*/
	Oracleserverversion string `json:"oracleserverversion,omitempty"`
	/**
	* Information about this virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* The version of the MSSQL server
	*/
	Mssqlserverversion string `json:"mssqlserverversion,omitempty"`
	/**
	* Use L2 Parameters to identify a connection
	*/
	L2conn string `json:"l2conn,omitempty"`
	/**
	* The protocol version returned by the mysql vserver.
	*/
	Mysqlprotocolversion uint32 `json:"mysqlprotocolversion,omitempty"`
	/**
	* The server version string returned by the mysql vserver.
	*/
	Mysqlserverversion string `json:"mysqlserverversion,omitempty"`
	/**
	* The character set returned by the mysql vserver.
	*/
	Mysqlcharacterset uint32 `json:"mysqlcharacterset,omitempty"`
	/**
	* The server capabilities returned by the mysql vserver.
	*/
	Mysqlservercapabilities uint32 `json:"mysqlservercapabilities,omitempty"`
	/**
	* Enable logging appflow flow information
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* The name of the network profile.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Can be active or passive
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
	* Name of the authentication profile to be used when authentication is turned on.
	*/
	Authnprofile string `json:"authnprofile,omitempty"`
	/**
	* Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`
	/**
	* This option starts/stops the dtls service on the vserver
	*/
	Dtls string `json:"dtls,omitempty"`
	/**
	*  Type of persistence for the virtual server. Available settings function as follows:
		* SOURCEIP - Connections from the same client IP address belong to the same persistence session.
		* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session.
		* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.
	*/
	Persistencetype string `json:"persistencetype,omitempty"`
	/**
	* Persistence mask for IP based persistence types, for IPv4 virtual servers.
	*/
	Persistmask string `json:"persistmask,omitempty"`
	/**
	* Persistence mask for IP based persistence types, for IPv6 virtual servers.
	*/
	V6persistmasklen uint32 `json:"v6persistmasklen,omitempty"`
	/**
	* Time period for which a persistence session is in effect.
	*/
	Timeout uint32 `json:"timeout,omitempty"`
	/**
	* Use this parameter to  specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.
	*/
	Cookiename string `json:"cookiename,omitempty"`
	/**
	* Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.
	*/
	Persistencebackup string `json:"persistencebackup,omitempty"`
	/**
	* Time period for which backup persistence is in effect.
	*/
	Backuppersistencetimeout uint32 `json:"backuppersistencetimeout,omitempty"`
	/**
	* Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.
	*/
	Tcpprobeport int32 `json:"tcpprobeport,omitempty"`
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
	* Name of QUIC profile which will be attached to the Content Switching VServer.
	*/
	Quicprofilename string `json:"quicprofilename,omitempty"`
	/**
	* Domain name for which to change the time to live (TTL) and/or backup service IP address.
	*/
	Domainname string `json:"domainname,omitempty"`
	Ttl uint64 `json:"ttl,omitempty"`
	Backupip string `json:"backupip,omitempty"`
	Cookiedomain string `json:"cookiedomain,omitempty"`
	Cookietimeout uint32 `json:"cookietimeout,omitempty"`
	Sitedomainttl uint64 `json:"sitedomainttl,omitempty"`
	/**
	* New name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my name" or 'my name').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Value string `json:"value,omitempty"`
	Ngname string `json:"ngname,omitempty"`
	Type string `json:"type,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Sc string `json:"sc,omitempty"`
	Status string `json:"status,omitempty"`
	Cachetype string `json:"cachetype,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Dnsvservername string `json:"dnsvservername,omitempty"`
	Domain string `json:"domain,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight string `json:"weight,omitempty"`
	Cachevserver string `json:"cachevserver,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
	Url string `json:"url,omitempty"`
	Bindpoint string `json:"bindpoint,omitempty"`
	Gt2gb string `json:"gt2gb,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Ruletype string `json:"ruletype,omitempty"`
	Lbvserver string `json:"lbvserver,omitempty"`
	Targetlbvserver string `json:"targetlbvserver,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Version string `json:"version,omitempty"`

}
