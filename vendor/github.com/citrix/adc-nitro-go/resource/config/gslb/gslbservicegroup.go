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

package gslb

/**
* Configuration for GSLB service group resource.
*/
type Gslbservicegroup struct {
	/**
	* Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Protocol used to exchange data with the GSLB service.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* Maximum number of simultaneous open connections for the GSLB service group.
	*/
	Maxclient int `json:"maxclient,omitempty"`
	/**
	* Insert the Client IP header in requests forwarded to the GSLB service.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Monitor the health of this GSLB service.Available settings function are as follows:
		YES - Send probes to check the health of the GSLB service.
		NO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.
	*/
	Healthmonitor string `json:"healthmonitor,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle client connection.
	*/
	Clttimeout int `json:"clttimeout,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle server connection.
	*/
	Svrtimeout int `json:"svrtimeout,omitempty"`
	/**
	* Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.
	*/
	Maxbandwidth int `json:"maxbandwidth,omitempty"`
	/**
	* Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.
	*/
	Monthreshold int `json:"monthreshold,omitempty"`
	/**
	* Initial state of the GSLB service group.
	*/
	State string `json:"state,omitempty"`
	/**
	* Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Any information about the GSLB service group.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Enable logging of AppFlow information for the specified GSLB service group.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Auto scale option for a GSLB servicegroup
	*/
	Autoscale string `json:"autoscale,omitempty"`
	/**
	* Name of the GSLB site to which the service group belongs.
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.
	*/
	Sitepersistence string `json:"sitepersistence,omitempty"`
	/**
	* Name of the server to which to bind the service group.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Server port number.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
	*/
	Hashid int `json:"hashid,omitempty"`
	/**
	* The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
	*/
	Publicport int `json:"publicport,omitempty"`
	/**
	* The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`
	/**
	* Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.
	*/
	Monitornamesvc string `json:"monitor_name_svc,omitempty"`
	/**
	* weight of the monitor that is bound to GSLB servicegroup.
	*/
	Dupweight int `json:"dup_weight,omitempty"`
	/**
	* The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.
	*/
	Delay int `json:"delay,omitempty"`
	/**
	* Wait for all existing connections to the service to terminate before shutting down the service.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.
	*/
	Includemembers bool `json:"includemembers,omitempty"`
	/**
	* New name for the GSLB service group.
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
	Gslb string `json:"gslb,omitempty"`
	Svreffgslbstate string `json:"svreffgslbstate,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`

}
