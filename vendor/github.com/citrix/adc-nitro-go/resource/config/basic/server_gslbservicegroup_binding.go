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
* Binding class showing the gslbservicegroup that can be bound to server.
*/
type Servergslbservicegroupbinding struct {
	/**
	* servicegroups bind to this server
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* The type of bound service
	*/
	Svctype string `json:"svctype,omitempty"`
	/**
	* The IP address of the bound service
	*/
	Serviceipaddress string `json:"serviceipaddress,omitempty"`
	/**
	* The port number to be used for the bound service.
	*/
	Port int `json:"port,omitempty"`
	/**
	* A positive integer to identify the service. Used when the persistency type is set to Custom Server ID.
	*/
	Customserverid string `json:"customserverid,omitempty"`
	/**
	* The state of the bound service
	*/
	Svrstate string `json:"svrstate,omitempty"`
	/**
	* service type of the service.
	*/
	Dupsvctype string `json:"dup_svctype,omitempty"`
	/**
	* port of the service.
	*/
	Dupport int `json:"dup_port,omitempty"`
	/**
	* service flags to denote its a db enabled.
	*/
	Svrcfgflags int `json:"svrcfgflags,omitempty"`
	/**
	* This field has been intorduced to show the dbs services ip
	*/
	Serviceipstr string `json:"serviceipstr,omitempty"`
	/**
	* Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.
	*/
	Monthreshold int `json:"monthreshold,omitempty"`
	/**
	* Maximum number of simultaneous open connections for the service group.
	*/
	Maxclient int `json:"maxclient,omitempty"`
	/**
	* Maximum number of requests that can be sent on a persistent connection to the service group. 
		Note: Connection requests beyond this value are rejected.
	*/
	Maxreq int `json:"maxreq,omitempty"`
	/**
	* Maximum bandwidth, in Kbps, allocated for all the services in the service group.
	*/
	Maxbandwidth int `json:"maxbandwidth,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle client connection.
	*/
	Clttimeout int `json:"clttimeout,omitempty"`
	/**
	* Time, in seconds, after which to terminate an idle server connection.
	*/
	Svrtimeout int `json:"svrtimeout,omitempty"`
	/**
	* Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.
	*/
	Cipheader string `json:"cipheader,omitempty"`
	/**
	* Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Perform delayed clean-up of connections to all services in the service group.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* Enable logging of AppFlow information for the specified service group.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Boundtd int `json:"boundtd,omitempty"`
	/**
	* Name of the server for which to display parameters.
	*/
	Name string `json:"name,omitempty"`


}