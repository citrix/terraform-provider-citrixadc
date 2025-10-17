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

package basic

/**
* Configuration for service resource.
*/
type Service struct {
	/**
	* Name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the service has been created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP to assign to the service.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* Name of the server that hosts the service.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Protocol in which data is exchanged with the service.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* Port number of the service.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.
	*/
	Cleartextport *int `json:"cleartextport,omitempty"`
	/**
	* Cache type supported by the cache server.
	*/
	Cachetype string `json:"cachetype,omitempty"`
	/**
	* Maximum number of simultaneous open connections to the service.
	*/
	Maxclient *int `json:"maxclient,omitempty"`
	/**
	* Monitor the health of this service. Available settings function as follows:
		YES - Send probes to check the health of the service.
		NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.
	*/
	Healthmonitor string `json:"healthmonitor,omitempty"`
	/**
	* Maximum number of requests that can be sent on a persistent connection to the service.
		Note: Connection requests beyond this value are rejected.
	*/
	Maxreq *int `json:"maxreq,omitempty"`
	/**
	* Use the transparent cache redirection virtual server to forward requests to the cache server.
		Note: Do not specify this parameter if you set the Cache Type parameter.
	*/
	Cacheable string `json:"cacheable,omitempty"`
	/**
	* Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name "client-ip."
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service.
	*/
	Usip string `json:"usip,omitempty"`
	/**
	* Path monitoring for clustering
	*/
	Pathmonitor string `json:"pathmonitor,omitempty"`
	/**
	* Individual Path monitoring decisions
	*/
	Pathmonitorindv string `json:"pathmonitorindv,omitempty"`
	/**
	* Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.
		Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.
	*/
	Useproxyport string `json:"useproxyport,omitempty"`
	/**
	* Enable surge protection for the service.
	*/
	Sp string `json:"sp,omitempty"`
	/**
	* Enable RTSP session ID mapping for the service.
	*/
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle client connection.
	*/
	Clttimeout *int `json:"clttimeout,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle server connection.
	*/
	Svrtimeout *int `json:"svrtimeout,omitempty"`
	/**
	* Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.
	*/
	Customserverid string `json:"customserverid,omitempty"`
	/**
	* The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
	*/
	Serverid *int `json:"serverid,omitempty"`
	/**
	* Enable client keep-alive for the service.
	*/
	Cka string `json:"cka,omitempty"`
	/**
	* Enable TCP buffering for the service.
	*/
	Tcpb string `json:"tcpb,omitempty"`
	/**
	* Enable compression for the service.
	*/
	Cmp string `json:"cmp,omitempty"`
	/**
	* Maximum bandwidth, in Kbps, allocated to the service.
	*/
	Maxbandwidth *int `json:"maxbandwidth,omitempty"`
	/**
	* Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.
	*/
	Accessdown string `json:"accessdown,omitempty"`
	/**
	* Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.
	*/
	Monthreshold *int `json:"monthreshold,omitempty"`
	/**
	* Initial state of the service.
	*/
	State string `json:"state,omitempty"`
	/**
	* Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Name of the TCP profile that contains TCP configuration settings for the service.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the HTTP profile that contains HTTP configuration settings for the service.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service
	*/
	Contentinspectionprofilename string `json:"contentinspectionprofilename,omitempty"`
	/**
	* Name of QUIC profile which will be attached to the service.
	*/
	Quicprofilename string `json:"quicprofilename,omitempty"`
	/**
	* A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.
	*/
	Hashid *int `json:"hashid,omitempty"`
	/**
	* Any information about the service.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable logging of AppFlow information.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Network profile to use for the service.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.
	*/
	Processlocal string `json:"processlocal,omitempty"`
	/**
	* Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS, ADNS-TCP and ADNS-DOT services.
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`
	/**
	* Close monitoring connections by sending the service a connection termination message with the specified bit set.
	*/
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	/**
	* The new IP address of the service.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* Name of the monitor bound to the specified service.
	*/
	Monitornamesvc string `json:"monitor_name_svc,omitempty"`
	/**
	* Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
	*/
	Delay *int `json:"delay,omitempty"`
	/**
	* Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Display both user-configured and dynamically learned services.
	*/
	All bool `json:"all,omitempty"`
	/**
	* Display only dynamically learned services.
	*/
	Internal bool `json:"Internal,omitempty"`
	/**
	* New name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numofconnections string `json:"numofconnections,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Serviceconftype string `json:"serviceconftype,omitempty"`
	Serviceconftype2 string `json:"serviceconftype2,omitempty"`
	Value string `json:"value,omitempty"`
	Gslb string `json:"gslb,omitempty"`
	Dupstate string `json:"dup_state,omitempty"`
	Publicip string `json:"publicip,omitempty"`
	Publicport string `json:"publicport,omitempty"`
	Svrstate string `json:"svrstate,omitempty"`
	Monitorstate string `json:"monitor_state,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Lastresponse string `json:"lastresponse,omitempty"`
	Responsetime string `json:"responsetime,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Stateupdatereason string `json:"stateupdatereason,omitempty"`
	Clmonowner string `json:"clmonowner,omitempty"`
	Clmonview string `json:"clmonview,omitempty"`
	Serviceipstr string `json:"serviceipstr,omitempty"`
	Oracleserverversion string `json:"oracleserverversion,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Monuserstatusmesg string `json:"monuserstatusmesg,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
