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
* Configuration for service group resource.
*/
type Servicegroup struct {
	/**
	* Name of the service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Protocol used to exchange data with the service.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* Cache type supported by the cache server.
	*/
	Cachetype string `json:"cachetype,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Maximum number of simultaneous open connections for the service group.
	*/
	Maxclient uint32 `json:"maxclient,omitempty"`
	/**
	* Maximum number of requests that can be sent on a persistent connection to the service group. 
		Note: Connection requests beyond this value are rejected.
	*/
	Maxreq uint32 `json:"maxreq,omitempty"`
	/**
	* Use the transparent cache redirection virtual server to forward the request to the cache server.
		Note: Do not set this parameter if you set the Cache Type.
	*/
	Cacheable string `json:"cacheable,omitempty"`
	/**
	* Insert the Client IP header in requests forwarded to the service.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Use client's IP address as the source IP address when initiating connection to the server. With the NO setting, which is the default, a mapped IP (MIP) address or subnet IP (SNIP) address is used as the source IP address to initiate server side connections.
	*/
	Usip string `json:"usip,omitempty"`
	/**
	* Path monitoring for clustering
	*/
	Pathmonitor string `json:"pathmonitor,omitempty"`
	/**
	* Individual Path monitoring decisions.
	*/
	Pathmonitorindv string `json:"pathmonitorindv,omitempty"`
	/**
	* Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection. 
		Note: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.
	*/
	Useproxyport string `json:"useproxyport,omitempty"`
	/**
	* Monitor the health of this service.  Available settings function as follows:
		YES - Send probes to check the health of the service.
		NO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.
	*/
	Healthmonitor string `json:"healthmonitor,omitempty"`
	/**
	* State of the SureConnect feature for the service group.
	*/
	Sc string `json:"sc,omitempty"`
	/**
	* Enable surge protection for the service group.
	*/
	Sp string `json:"sp,omitempty"`
	/**
	* Enable RTSP session ID mapping for the service group.
	*/
	Rtspsessionidremap string `json:"rtspsessionidremap,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle client connection.
	*/
	Clttimeout uint64 `json:"clttimeout,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle server connection.
	*/
	Svrtimeout uint64 `json:"svrtimeout,omitempty"`
	/**
	* Enable client keep-alive for the service group.
	*/
	Cka string `json:"cka,omitempty"`
	/**
	* Enable TCP buffering for the service group.
	*/
	Tcpb string `json:"tcpb,omitempty"`
	/**
	* Enable compression for the specified service.
	*/
	Cmp string `json:"cmp,omitempty"`
	/**
	* Maximum bandwidth, in Kbps, allocated for all the services in the service group.
	*/
	Maxbandwidth uint32 `json:"maxbandwidth,omitempty"`
	/**
	* Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.
	*/
	Monthreshold uint32 `json:"monthreshold,omitempty"`
	/**
	* Initial state of the service group.
	*/
	State string `json:"state,omitempty"`
	/**
	* Flush all active transactions associated with all the services in the service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Name of the TCP profile that contains TCP configuration settings for the service group.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the HTTP profile that contains HTTP configuration settings for the service group.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Any information about the service group.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable logging of AppFlow information for the specified service group.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Network profile for the service group.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Auto scale option for a servicegroup
	*/
	Autoscale string `json:"autoscale,omitempty"`
	/**
	* member port
	*/
	Memberport int32 `json:"memberport,omitempty"`
	/**
	* Indicates graceful shutdown of the service. System will wait for all outstanding connections to this service to be closed before disabling the service.
	*/
	Autodisablegraceful string `json:"autodisablegraceful,omitempty"`
	/**
	* The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.
	*/
	Autodisabledelay uint64 `json:"autodisabledelay,omitempty"`
	/**
	* Close monitoring connections by sending the service a connection termination message with the specified bit set.
	*/
	Monconnectionclose string `json:"monconnectionclose,omitempty"`
	/**
	* Name of the server to which to bind the service group.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Server port number.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
	*/
	Weight uint32 `json:"weight,omitempty"`
	/**
	* The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
	*/
	Customserverid string `json:"customserverid,omitempty"`
	/**
	* The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
	*/
	Serverid uint32 `json:"serverid,omitempty"`
	/**
	* The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
	*/
	Hashid uint32 `json:"hashid,omitempty"`
	/**
	* Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver
	*/
	Nameserver string `json:"nameserver,omitempty"`
	/**
	* Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors
	*/
	Dbsttl uint64 `json:"dbsttl,omitempty"`
	/**
	* Name of the monitor bound to the service group. Used to assign a weight to the monitor.
	*/
	Monitornamesvc string `json:"monitor_name_svc,omitempty"`
	/**
	* weight of the monitor that is bound to servicegroup.
	*/
	Dupweight uint32 `json:"dup_weight,omitempty"`
	/**
	* Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
	*/
	Delay uint64 `json:"delay,omitempty"`
	/**
	* Wait for all existing connections to the service to terminate before shutting down the service.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Display the members of the listed service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.
	*/
	Includemembers bool `json:"includemembers,omitempty"`
	/**
	* New name for the service group.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numofconnections string `json:"numofconnections,omitempty"`
	Serviceconftype string `json:"serviceconftype,omitempty"`
	Value string `json:"value,omitempty"`
	Svrstate string `json:"svrstate,omitempty"`
	Ip string `json:"ip,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`
	Statechangetimemsec string `json:"statechangetimemsec,omitempty"`
	Stateupdatereason string `json:"stateupdatereason,omitempty"`
	Clmonowner string `json:"clmonowner,omitempty"`
	Clmonview string `json:"clmonview,omitempty"`
	Groupcount string `json:"groupcount,omitempty"`
	Serviceipstr string `json:"serviceipstr,omitempty"`
	Servicegroupeffectivestate string `json:"servicegroupeffectivestate,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Svcitmactsvcs string `json:"svcitmactsvcs,omitempty"`
	Svcitmboundsvcs string `json:"svcitmboundsvcs,omitempty"`
	Monuserstatusmesg string `json:"monuserstatusmesg,omitempty"`

}
